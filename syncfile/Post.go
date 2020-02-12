package syncfile

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"strings"
)

//发送信息

func postForm(c *Client) (*http.Response, error) {
	var r http.Request
	r.ParseForm()
	r.Form.Add(FILE_NAME_KEY, c.fileName)
	bodystr := strings.TrimSpace(r.Form.Encode())
	request, err := http.NewRequest("POST", c.remoteUrl, strings.NewReader(bodystr))
	if err != nil {
	}
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := http.DefaultClient.Do(request)
	return resp, err
}

func downloadServer(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() //解析url传递的参数，对于POST则解析响应包的主体（request body）
	//注意:如果没有调用ParseForm方法，下面无法获取表单的数据
	fileName := r.Form["filename"] //filename  文件名
	file, err := os.Open(fileName[0])
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	content, err := ioutil.ReadAll(file)
	fileNames := url.QueryEscape(fileName[0]) // 防止中文乱码
	w.Header().Add("Content-Type", "application/octet-stream")
	w.Header().Add("Content-Disposition", "attachment; filename=\""+fileNames+"\"")

	if err != nil {
		fmt.Println("Read File Err:", err.Error())
	} else {
		w.Write(content)
	}
}

//SERVER 端对上传的处理， hundle中包含上传本地路径，可以映射到开发机地址
func uploadServer(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get file")
	r.ParseMultipartForm(32 << 20)
	file, handler, err := r.FormFile(FILE_NAME_KEY)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	fmt.Fprintf(w, "%v", handler.Header)
	fileName := handler.Filename
	devPath := getDevPath(fileName)
	os.Remove(devPath)
	f, err := os.OpenFile(devPath, os.O_WRONLY|os.O_CREATE, 0666) // 此处假设当前目录下已存在test目录
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	io.Copy(f, file)
}

//发送信息并上传文件
func uploadClient(client *Client) (*http.Response, error) {
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)

	//关键的一步操作
	fileWriter, err := bodyWriter.CreateFormFile(FILE_NAME_KEY, client.fileName)
	if err != nil {
		fmt.Println("error writing to buffer")
		return nil, err
	}

	//打开文件句柄操作
	fh, err := os.Open(client.fileName)
	if err != nil {
		fmt.Println("error opening file")
		return nil, err
	}
	defer fh.Close()

	//iocopy
	_, err = io.Copy(fileWriter, fh)
	if err != nil {
		fmt.Println("get err")
		return nil, err
	}

	contentType := bodyWriter.FormDataContentType()
	bodyWriter.Close()
	resp, err := http.Post(client.remoteUrl, contentType, bodyBuf)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	resp_body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	fmt.Println(resp.Status)
	fmt.Println(string(resp_body))
	return resp, err
}
