#!/usr/bin/env sh

alias terminal='/Applications/iTerm.app/Contents/MacOS/iTerm2'

curr_dir=$(pwd)
server_dir="$curr_dir/server"
client_dir="$curr_dir/client"

terminal "$server_dir" &
terminal "$client_dir" &
