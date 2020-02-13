package syncfile

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func PushFile(localFilePath string) {
	c := Client{fileName: localFilePath}
	c.opreate = OPR_CREATE
	c.post()
}
func GetFile(path string) {
	var c Client
	c.fileName = path
	c.opreate = OPR_DOWNLOAD
	var resp *http.Response
	resp, _ = c.post()
	message, _ := ioutil.ReadAll(resp.Body)
	localpath := getLocalPath(path)
	ioutil.WriteFile(localpath, message, 0644)
}

func GetList() {
	var c Client
	c.fileName = developDirPath+"/"+LIST_FILE_NAME
	c.opreate = OPR_LIST
	var resp *http.Response
	resp, _ = c.post()
	message, _ := ioutil.ReadAll(resp.Body)
	localpath := getLocalPath(c.fileName)
	ioutil.WriteFile(localpath, message, 0644)
}
func GetAllDev(){
	GetList()
	f,err :=os.Open(localDirPath+"/"+LIST_FILE_NAME)
	if err != nil {
		fmt.Println(err.Error())
	}
	//建立缓冲区，把文件内容放到缓冲区中
	buf := bufio.NewReader(f)
	for {
		//遇到\n结束读取
		b, errR := buf.ReadBytes('\n')
		if errR != nil {
			if errR == io.EOF {
				break
			}
			fmt.Println(errR.Error())
		}
		fileNameList := string(b)
		msg:=strings.Split(fileNameList,",")
		fileType := strings.Replace(msg[1], "\n", "", -1)
		fileName := msg[0]
		if(fileType == PATH_IS_DIR){
			dirName  := getLocalPath(fileName)
			os.Mkdir(dirName,os.ModePerm)
		}else{
			GetFile(fileName)
		}
	}
}

