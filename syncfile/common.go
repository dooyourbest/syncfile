package syncfile

import "strings"

const (
	// syncfile local  即可监视本地文件目录
	RUN_LOCAL = "local"

	// syncfile dev  即可开启服务端(开发机端)
	RUN_DEV = "dev"

	// syncfile get 服务端文件绝对地址，手动将线上文件拉到本地对应目录
	GET_FILE = "get"

	// syncfile push 本地文件绝对地址，手动将本地文件推到线上对应目录
	PUSH_FILE = "push"
	LIST_FILE = "list"
	GET_ALL_DEV="getall"

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
	REMOTE_PORT    = "8081"                              //请求的port

	REMOTE_HOST    = "https://localhost"                  //server端host
	USE_HTTPS      = true //默认不使用https

	localDirPath   = "/Users/zhangziang_cd/go/src/github.com/dooyourbest/syncfile/a" //本地监听的目录 绝对路径
	//localDirPath   = "C:\\Users\\Administrator\\go\\src\\github.com\\dooyourbest\\syncfile\\syncfile\\b" //本地监听的目录 绝对路径
	developDirPath = "/Users/zhangziang_cd/go/src/github.com/dooyourbest/syncfile/b"  //远程代码目录 绝对路径
	//developDirPath = "C:\\Users\\Administrator\\go\\src\\github.com\\dooyourbest\\syncfile\\syncfile\\a"  //远程代码目录 绝对路径
	CA_ROOT = "/Users/zhangziang_cd/go/src/github.com/dooyourbest/syncfile/syncfile/ca/"
	CA_PATH = CA_ROOT+"ca.key"
	CA_CRT = CA_ROOT+"ca.crt"

	SERVER_KEY = CA_ROOT+"server.key"
	SERVRT_CRT = CA_ROOT+"server.crt"

	CLIENT_KEY = CA_ROOT+"client.key"
	CLIENT_CRT = CA_ROOT+"client.crt"


)

var DirIgnoreWord = []string{".git"}                         //监听中忽略目录
var FileIgnoreWord = []string{"~", "jb_tmp", "___jb_old___"} //监听中忽略文件

//根据本地地址返回开发机地址
func getDevPath(localPath string) string {
	path := strings.Replace(localPath, localDirPath, "", -1)
	devPath := developDirPath + "/" + path
	return devPath
}

//根据开发机地址返回本地地址
func getLocalPath(devPath string) string {
	path := strings.Replace(devPath, developDirPath, "", -1)
	localPath := localDirPath + "/" + path
	return localPath
}
