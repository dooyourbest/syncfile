package syncfile

import "strings"

const (
	RUN_LOCAL    = "local"
	RUN_DEV      = "dev"
	GET_FILE    = "get"
	PUSH_FILE    = "push"

	OPR_REMOVE   = "rm"
	OPR_CREATE   = "create"
	OPR_DIR_ADD  = "addDir"
	OPR_DOWNLOAD = "download"

	FILE_NAME_KEY = "filename"

	//一下需要修改，可以在本地进行测试
	REMOTE_PORT    = "8089"
	targetUrl      = "http://localhost:" + REMOTE_PORT   //server端监听接口
	localDirPath   = "/Users/zhangziang_cd/go/src/test1" //local监听目录
	developDirPath = "/Users/zhangziang_cd/go/src/test"  //远程代码目录
)
var DirIgnoreWord = []string{".git"}  //忽略目录
var FileIgnoreWord = []string{"~", "jb_tmp","___jb_old___"} //忽略文件

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
