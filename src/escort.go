package main

import (
	"client"
	//. "common"
	"flag"
	"fmt"
	"glog"
	"server"
)

var Usage = func() {
	fmt.Println("USAGE: escort -role [master|slave]")
	fmt.Println("\nThe commands are:\n\tmaster\tSet this server is master server.\n\tslave\tSet this server is slave server.")
}

func main() {

	role := flag.String("role", "master", "Escort role")
	flag.Set("alsologtostderr", "true")
	flag.Set("v", "5")
	flag.Parse()

	if role == nil || len(*role) == 0 {
		Usage()
		return
	}

	glog.Infoln(*role)

	switch *role {
	case "master":
		// 主服务
		server.Work()
	case "slave":
		// 备份服务
		client.Work()
	default:
		Usage()
	}
}
