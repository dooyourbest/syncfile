package lib

import (
	"github.com/dooyourbest/syncfile/syncfile/cli"
	"strings"
)

const (


	LIST_FILE_NAME = "index.xml"
	PATH_IS_DIR = "1"
	PATH_IS_FILE = "0"

	OPR_REMOVE   = "rm"       //删除操作对应的urlpath
	OPR_CREATE   = "create"   //创建文件操作对应的urlpath
	OPR_DIR_ADD  = "addDir"   //创建目录对应的urlpath
	OPR_DOWNLOAD = "download" //下载文件对应的urlpath
	OPR_LIST = "list" //下载文件对应的urlpath

	FILE_NAME_KEY = "filename"

	//需要进行修改的内容



)

var DirIgnoreWord = []string{".git"}                         //监听中忽略目录
var FileIgnoreWord = []string{"~", "jb_tmp", "___jb_old___"} //监听中忽略文件

//根据本地地址返回开发机地址
func getDevPath(localPath string) string {
	path := strings.Replace(localPath, cli.Conf.Local, "", -1)
	devPath := cli.Conf.Dev + "/" + path
	return devPath
}

//根据开发机地址返回本地地址
func getLocalPath(devPath string) string {
	path := strings.Replace(devPath, cli.Conf.Dev, "", -1)
	localPath := cli.Conf.Local + "/" + path
	return localPath
}

type conf struct {
	local string
	dev string
	port string
	host string
}

//func loadConfFile(*app APP,path string)(error){
//	msg,err := ioutil.ReadFile("file/test")
//	if err==nil{
//		return nil
//	}
//}
