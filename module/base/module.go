package base

import (
	"fmt"
	"reflect"
	"runtime"

	"Jottings/tiny_rpc/log"
	"Jottings/tiny_rpc/util"
)

var mm = newModuleMgr()

func ModuleMgrIns() *ModuleMgr {
	return mm
}

func SendWork(mo, me string, arg, reply interface{}) error {
	return ModuleMgrIns().sendWork(mo, me, arg, reply)
}

// ModuleMgr ----------------------------------------------------------------------------------------------------
type ModuleMgr struct {
	wg         *util.WGWrapper
	moduleMap  map[string]*Module
	moduleList []string
}

func newModuleMgr() *ModuleMgr {
	return &ModuleMgr{
		wg:         new(util.WGWrapper),
		moduleMap:  make(map[string]*Module, 128),
		moduleList: make([]string, 0, 128),
	}
}

func (r *ModuleMgr) RegModule(m interface{}, chanSize int) error {
	t := reflect.TypeOf(m)
	name := t.Elem().Name()
	if t.Kind() != reflect.Ptr {
		name = t.Name()
		return fmt.Errorf("reg module %s fail, module not ptr", name)
	}
	if r.moduleMap[name] != nil {
		return nil
	}

	r.moduleMap[name] = NewModule(m, chanSize)
	r.moduleList = append(r.moduleList, name)
	log.Info("reg module %s chan size %d", name, chanSize)
	return nil
}

func (r *ModuleMgr) Start() error {
	for i := 0; i < len(r.moduleList); i++ {
		name := r.moduleList[i]
		if r.moduleMap[name] == nil {
			return fmt.Errorf("module mgr start %s err", name)
		}
		r.wg.Wrap(r.moduleMap[name].Start)
	}
	log.Info("modules start, count %d", len(r.moduleMap))
	return nil
}

func (r *ModuleMgr) Stop() error {
	for i := len(r.moduleList) - 1; i >= 0; i-- {
		name := r.moduleList[i]
		if r.moduleMap[name] == nil {
			log.Error("module mgr stop %s err", name)
			continue
		}
		if err := r.moduleMap[name].Stop(); err != nil {
			log.Error("module %s stop err %v", name, err)
			continue
		}
	}
	log.Info("modules stop, count %d", len(r.moduleMap))
	return nil
}

func (r *ModuleMgr) sendWork(mo, me string, arg, reply interface{}) error {
	if r.moduleMap[mo] == nil {
		return fmt.Errorf("module %s nil", mo)
	}
	w := &ModuleWork{
		Method:  me,
		Arg:     arg,
		Reply:   reply,
		RetChan: make(chan struct{}),
	}
	r.moduleMap[mo].Receive(w)

	select {
	case <-w.RetChan:
		if w.Err != nil {
			return w.Err
		}
	}
	return nil
}

type ModuleWork struct {
	Method  string
	Arg     interface{}
	Reply   interface{}
	Err     error
	RetChan chan struct{}
}

type methodType struct {
	method    reflect.Method
	ArgType   reflect.Type
	ReplyType reflect.Type
}

type Module struct {
	name      string
	workChan  chan *ModuleWork
	closeChan chan interface{}
	rel       interface{}
	typ       reflect.Type
	rcvr      reflect.Value
	method    map[string]*methodType
}

func NewModule(m interface{}, chanSize int) *Module {
	return &Module{
		name:      reflect.TypeOf(m).Elem().Name(),
		workChan:  make(chan *ModuleWork, chanSize),
		closeChan: make(chan interface{}),
		rel:       m,
		typ:       reflect.TypeOf(m),
		rcvr:      reflect.ValueOf(m),
	}
}

func (r *Module) Start() {
	defer util.InfoPanic("module %v panic", r.name)
	if err := r.register(); err != nil {
		log.Error("module start err %v", err)
		return
	}
	log.Info("module %v start", r.name)

	for {
		select {
		case w := <-r.workChan:
			func(w *ModuleWork) {
				defer util.InfoPanic("module %v panic, work %s param %+v", r.name, w.Method, w.Arg)

				if err := r.dealWork(w); err != nil {
					log.Error("module %v deal work err %v", r.name, err)
				}
			}(w)
		case <-r.closeChan:
			return
		}
	}
}

func (r *Module) Stop() error {
	close(r.closeChan)
	return nil
}

func (r *Module) Receive(w *ModuleWork) {
	r.workChan <- w
}

func (r *Module) register() error {
	r.method = make(map[string]*methodType)
	for m := 0; m < r.typ.NumMethod(); m++ {
		method := r.typ.Method(m)
		mtype := method.Type
		mname := method.Name
		// Method must be exported.
		if method.PkgPath != "" {
			continue
		}
		// Method needs three ins: receiver, *args, *reply.
		if mtype.NumIn() != 3 {
			log.Debug("method %s has wrong number of ins: %d", mname, mtype.NumIn())
			continue
		}

		argType := mtype.In(1)
		replyType := mtype.In(2)

		if argType.Kind() != reflect.Ptr || replyType.Kind() != reflect.Ptr {
			return fmt.Errorf("arg or reply not ptr")
		}
		// Method needs one out.
		if mtype.NumOut() != 1 {
			log.Info("method %s has wrong number of outs: %d", mname, mtype.NumOut())
			continue
		}
		r.method[mname] = &methodType{method: method, ArgType: argType, ReplyType: replyType}
		log.Info("module %s reg method %s", r.name, mname)
	}
	return nil
}

func (r *Module) dealWork(w *ModuleWork) error {
	mtype := r.method[w.Method]
	if mtype == nil {
		return fmt.Errorf("%s not find", w.Method)
	}
	err := r.call(mtype, reflect.ValueOf(w.Arg), reflect.ValueOf(w.Reply))
	if w.RetChan != nil {
		if err != nil {
			w.Err = err
		}
		w.RetChan <- struct{}{}
	}
	return nil
}

func (r *Module) call(mtype *methodType, argv, replyv reflect.Value) (err error) {
	defer func() {
		if err := recover(); err != nil {
			buf := make([]byte, 4096)
			n := runtime.Stack(buf, false)
			buf = buf[:n]

			err = fmt.Errorf("[module internal error]: %v, method: %s, argv: %+v, stack: %s",
				err, mtype.method.Name, argv.Interface(), buf)
			log.Error("%v", err)
		}
	}()

	function := mtype.method.Func
	// Invoke the method, providing a new value for the reply.
	returnValues := function.Call([]reflect.Value{r.rcvr, argv, replyv})
	// The return value for the method is an error.
	errInter := returnValues[0].Interface()
	if errInter != nil {
		return errInter.(error)
	}
	return nil
}
