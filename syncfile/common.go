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

	OPR_REMOVE   = "rm"       //删除操作对应的urlpath
	OPR_CREATE   = "create"   //创建文件操作对应的urlpath
	OPR_DIR_ADD  = "addDir"   //创建目录对应的urlpath
	OPR_DOWNLOAD = "download" //下载文件对应的urlpath

	FILE_NAME_KEY = "filename"

	//需要进行修改的内容
	REMOTE_PORT    = "8081"                              //请求的port

	REMOTE_HOST    = "http://localhost"                  //server端host
	USE_HTTPS      = false //默认不使用https

	localDirPath   = "/Users/zhangziang_cd/go/src/test1" //本地监听的目录 绝对路径
	developDirPath = "/Users/zhangziang_cd/go/src/test"  //远程代码目录 绝对路径



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
