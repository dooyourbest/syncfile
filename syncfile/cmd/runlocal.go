package cmd

import "github.com/dooyourbest/syncfile/syncfile/cli"
import (
	"github.com/dooyourbest/syncfile/syncfile/lib"
)

var RunLocal = cli.Command{
	Name:      "local",
	ShortName: "local",
	Action:    runlocal,
}
func runlocal(ctx *cli.Context){
	lib.WatchDir(ctx)
}
