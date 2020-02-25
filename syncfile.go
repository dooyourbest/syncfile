package main

import (
	"fmt"
	"github.com/dooyourbest/syncfile/syncfile/cli"
	"github.com/dooyourbest/syncfile/syncfile/cmd"
	"os"
)

func main() {
	app,err:=cli.NewApp()
	//注册实践
	app.CmdList=[]cli.Command{
		cmd.RunDev,
		cmd.RunLocal,
	}
	if err != nil{
		fmt.Print(err)
	}
	app.Run(os.Args)
}



