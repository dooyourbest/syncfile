package cmd

import (
	"github.com/dooyourbest/syncfile/syncfile/cli"
	"github.com/dooyourbest/syncfile/syncfile/lib"
)

var RunDev = cli.Command{
	Name:      "dev",
	ShortName: "dev",
	Action:    rundev,
}
func rundev(ctx *cli.Context){
	lib.RunListen(ctx)
}