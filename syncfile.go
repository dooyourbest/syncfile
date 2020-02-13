package main

import (
	"fmt"
	"os"
	. "syncfile/syncfile"
)

func main() {
	param := os.Args[1]
	fmt.Println(param)
	switch param {
	case RUN_LOCAL:
		fmt.Println("watch")
		WatchDir()
	case RUN_DEV:
		fmt.Print("listen")
		Listen()
	case GET_FILE:
		fmt.Println("getfile")
		GetFile(os.Args[2])
	case PUSH_FILE:
		fmt.Println("pushfile")
		PushFile(os.Args[2])
	case LIST_FILE:
		fmt.Println("list file")
		GetList()
	case GET_ALL_DEV:
		fmt.Println("GET ALL")
		GetAllDev()
	}

}
