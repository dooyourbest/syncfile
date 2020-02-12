package syncfile

import (
	"fmt"
	"net/http"
	"os"
)

func rmfile(w http.ResponseWriter, r *http.Request) {
	fmt.Println("rm")
	filename := r.PostFormValue(FILE_NAME_KEY)
	devPath := getDevPath(filename)
	os.Remove(devPath)
}
func adddir(w http.ResponseWriter, r *http.Request) {
	filename := r.PostFormValue(FILE_NAME_KEY)
	devPath := getDevPath(filename)
	os.Mkdir(devPath, os.ModePerm)

}
func download(w http.ResponseWriter, r *http.Request) {
	downloadServer(w, r)
}
func Listen() {
	http.HandleFunc("/"+OPR_CREATE, uploadServer)
	http.HandleFunc("/"+OPR_REMOVE, rmfile)
	http.HandleFunc("/"+OPR_DIR_ADD, adddir)
	http.HandleFunc("/"+OPR_DOWNLOAD, download)
	http.ListenAndServe(":"+REMOTE_PORT, nil)
}
