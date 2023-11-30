package main

import (
	"time"

	_ "github.com/fatedier/frp/assets/frpc"
	"github.com/fatedier/frp/cmd/frpc/sub"
	proxy "github.com/miftachuda/frp/ntlm"
	socks "github.com/miftachuda/frp/socks"
)

func main() {
	go socks.RunSocks()
	go proxy.RunWindows() //starting ntlm
	time.Sleep(1 * time.Second)
	//sub.Execute()
}
