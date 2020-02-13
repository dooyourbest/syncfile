package main

import (
	"fmt"
	"os"
	"github.com/dooyourbest/syncfile/syncfile"
)

func main() {
	param := os.Args[1]
	fmt.Println(param)
	switch param {
	case syncfile.RUN_LOCAL:
		fmt.Println("watch")
		syncfile.WatchDir()
	case syncfile.RUN_DEV:
		fmt.Print("listen")
		syncfile.Listen()
	case syncfile.GET_FILE:
		fmt.Println("getfile")
		syncfile.GetFile(os.Args[2])
	case syncfile.PUSH_FILE:
		fmt.Println("pushfile")
		syncfile.PushFile(os.Args[2])
	case syncfile.LIST_FILE:
		fmt.Println("list file")
		syncfile.GetList()
	case syncfile.GET_ALL_DEV:
		fmt.Println("GET ALL")
		syncfile.GetAllDev()
	}

}
