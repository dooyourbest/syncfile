package syncfile

import (
	"io/ioutil"
	"net/http"
)

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
