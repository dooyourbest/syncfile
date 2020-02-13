package syncfile

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
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

func listFile(w http.ResponseWriter, r *http.Request){
	fmt.Println(123)
	rootPath := developDirPath
	file, e := os.Create(developDirPath + "/" + LIST_FILE_NAME)
	if e != nil {
		fmt.Println(e)
		return
	}

	os.Open(rootPath+"/")
	err := filepath.Walk(rootPath, func(path string, f os.FileInfo, err error) error {
		if ( f == nil ) {return err}
		if f.IsDir() {
			path+=","+PATH_IS_DIR
		}else{
			path+=","+ PATH_IS_FILE
		}
		file.WriteString(path+"\n")
		return nil
	})
	if err != nil {
		fmt.Printf("filepath.Walk() returned %v\n", err)
	}
	downloadServer(w, r)
}

func Listen() {
	http.HandleFunc("/"+OPR_CREATE, uploadServer)
	http.HandleFunc("/"+OPR_REMOVE, rmfile)
	http.HandleFunc("/"+OPR_DIR_ADD, adddir)
	http.HandleFunc("/"+OPR_DOWNLOAD, download)
	http.HandleFunc("/"+OPR_LIST, listFile)
	http.ListenAndServe(":"+REMOTE_PORT, nil)
}
