import re

content = '''
/*
  PROTO_NUM:1
  RPC_NAME:Hello
  REQ_NAME:HelloReq
  RSP_NAME:HelloRsp
*/
message HelloReq  {
  string HelloMsg = 1;
}
message HelloRsp  {
  string ReplyMsg = 1;
}

/*
  PROTO_NUM:1
  RPC_NAME:Hello
  REQ_NAME:HelloReq
  RSP_NAME:HelloRsp
*/
message HelloReq  {
  string HelloMsg = 1;
}
message HelloRsp  {
  string ReplyMsg = 1;
}
'''

RPC_COMMENT_NUM = 4
PROTO_NUM = "PROTO_NUM:"
RPC_NAME = "RPC_NAME:"
REQ_NAME = "REQ_NAME:"
RSP_NAME = "RSP_NAME:"

rex_block = "\/\*[\w\W]*?\*\/"
rex_field = f"{PROTO_NUM}.*|{RPC_NAME}.*|{REQ_NAME}.*|{RSP_NAME}.*"

block = re.findall(rex_block, content)
for b in block:
    bf = re.findall(rex_field, b)
    if len(bf) != RPC_COMMENT_NUM:
        print("proto file comment err\n", b)
        exit()
    print(bf)
