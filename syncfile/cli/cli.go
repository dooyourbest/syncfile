package cli

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)
import "flag"

type App struct {
	name string
	version string
	author string
	CmdList []Command
}

type Command struct {
	Name string
	ShortName string
	Action func(context *Context)
}
type Context struct {
	App       *App
	Command   *Command
	Config    *Config
}
type Config struct {
	Local string
	Dev string
	Host string
	Port string
	UseHttps bool
	ConfigPath string
	CaPath string
}
var Conf=Config{
	Local:      "",
	Dev:        "",
	Host:       "",
	Port:       "",
	UseHttps:   false,
	ConfigPath: "",
	CaPath:     "",
}


func NewApp()(* App,error){
	app:= &App{
		name:"syncfile",
		version:"v1",
		author:"zhangziang",
	}
	return app,nil
}
func (a *App)Run(args []string)error{
	var filename = flag.String("f", "", "please input confile path")
	var cmd = flag.String("s", "", "please input cmd")
	flag.Parse()
	Conf.ConfigPath=*filename
	Conf.parseConfig()
	c := a.Command(*cmd)
	context:=&Context{App:a,Command:c,Config:&Conf}
	c.Action(context)
	return nil
}

func (a *App)Command(cmdName string)*Command  {
	for _, c := range a.CmdList {
		if c.HasName(cmdName) {
			return &c
		}
	}
	return nil
}
func (c *Command)HasName(name string)bool {
	return  c.Name==name || c.ShortName == name
}

func (f * Config)parseConfig(){
	if _, err := os.Stat(f.ConfigPath); os.IsNotExist(err){
		panic("error")
	}

	fileContent, err := ioutil.ReadFile(f.ConfigPath)
	if err != nil {
		fmt.Println(err)
	}
	confMsg := string(fileContent)
	confLine:=strings.Split(confMsg,"\n")
	confMap :=make(map[string]string)
	for _,value:=range confLine{
		kv:=strings.Split(value,":")
		if(len(kv)>1){
			confMap[kv[0]]=kv[1]
		}
	}
	f.Local=confMap["Local"]
	f.Dev=confMap["Dev"]
	f.Host=confMap["Host"]
	f.Port=confMap["Port"]
	if confMap["UseHttps"]=="true"{
		f.UseHttps=true
	}else{
		f.UseHttps=false
	}
	return nil
}

