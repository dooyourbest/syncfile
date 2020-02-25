package lib

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"github.com/dooyourbest/syncfile/syncfile/cli"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)
type SyncServer struct{
	ca *Ca
	config *cli.Config
}
var local string
var dev string

func rmfile(w http.ResponseWriter, r *http.Request) {
	fmt.Println("rm")
	filename := r.PostFormValue(FILE_NAME_KEY)
	fmt.Println(filename)
	devPath := getDevPath(filename)
	err := os.Remove(devPath)
	if err!=nil{
		fmt.Println(err)
	}
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
	rootPath := dev
	file, e := os.Create(dev + "/" + LIST_FILE_NAME)
	if e != nil {
		fmt.Println(e)
		return
	}

	os.Open(rootPath+"/")
	err := filepath.Walk(rootPath, func(path string, f os.FileInfo, err error) error {
		if ( f == nil ) {return err}
		if f.IsDir() {
			path+=","+ PATH_IS_DIR
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

func Listen(server *SyncServer) {
	http.HandleFunc("/"+OPR_CREATE, uploadServer)
	http.HandleFunc("/"+OPR_REMOVE, rmfile)
	http.HandleFunc("/"+OPR_DIR_ADD, adddir)
	http.HandleFunc("/"+OPR_DOWNLOAD, download)
	http.HandleFunc("/"+OPR_LIST, listFile,)
	http.ListenAndServe(":"+server.config.Port, nil)
}
func ListenHttps(server *SyncServer){
	server.ca= NewCa(server.config.CaPath)
	pool := x509.NewCertPool()
	caCrt, err := ioutil.ReadFile(server.ca.ca_crt)
	if err != nil {
		fmt.Println("ReadFile err:", err)
		return
	}
	pool.AppendCertsFromPEM(caCrt)
	s := &http.Server{
		Addr:    ":"+server.config.Host,
		TLSConfig: &tls.Config{
			ClientCAs:  pool,
			ClientAuth: tls.RequireAndVerifyClientCert,
		},
	}
	http.HandleFunc("/"+OPR_CREATE, uploadServer)
	http.HandleFunc("/"+OPR_REMOVE, rmfile)
	http.HandleFunc("/"+OPR_DIR_ADD, adddir)
	http.HandleFunc("/"+OPR_DOWNLOAD, download)
	http.HandleFunc("/"+OPR_LIST, listFile)
	err=s.ListenAndServeTLS(server.ca.server_crt,server.ca.server_key)
	if err != nil {
		fmt.Println("ListenAndServeTLS err:", err)
	}
}

func RunListen(c *cli.Context){
	s:= SyncServer{config:c.Config}
	local=cli.Conf.Local
	dev=cli.Conf.Dev
	if c.Config.UseHttps==true{
		ListenHttps(&s)
	}else{
		Listen(&s)
	}
}