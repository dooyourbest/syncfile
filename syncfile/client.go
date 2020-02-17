package syncfile

import (
	"bytes"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

var MapFile = make(map[string]string)

type Watch struct {
	watch *fsnotify.Watcher
}

type Client struct {
	remoteUrl string
	fileName  string
	opreate   string
	contentType string
	bodyBuf *bytes.Buffer
}

func (c Client) post() (*http.Response, error) {
	fmt.Print("client->post")
	c.remoteUrl = REMOTE_HOST + ":" + REMOTE_PORT + "/" + c.opreate
	fmt.Println(c.remoteUrl)
	fmt.Println(c.opreate)
	var resp *http.Response
	var err error
	if c.opreate == OPR_CREATE {
		resp, err = uploadClient(&c)
	} else {
		resp, err = postForm(&c)
	}
	return resp, err
}


//是否含有过滤词文件夹
func isIgnorePath(dirName string, ignoreWord []string) bool {
	flag := false
	for i := 0; i < len(ignoreWord); i++ {
		if strings.Contains(dirName, ignoreWord[i]) {
			flag = true
		}
	}
	return flag
}

//监控目录
func (w *Watch) watchDir(dir string) {
	//通过Walk来遍历目录下的所有子目录
	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		//这里判断是否为目录，只需监控目录即可
		//目录下的文件也在监控范围内，不需要我们一个一个加
		if info.IsDir() {
			path, err := filepath.Abs(path)
			if err != nil {
				return err
			}
			if !isIgnorePath(path, DirIgnoreWord) {
				err = w.watch.Add(path)
				if err != nil {
					return err
				}
				fmt.Println("监控 : ", path)
			}
		}
		return nil
	})
	var c Client
	go func() {
		for {
			select {
			case ev := <-w.watch.Events:
				{
					if ev.Op&fsnotify.Create == fsnotify.Create {
						if !isIgnorePath(ev.Name, FileIgnoreWord) {
							log.Println("创建文件 : ", ev.Name)
							//这里获取新创建文件的信息，如果是目录，则加入监控中
							fi, err := os.Stat(ev.Name)
							if err == nil {
								c.opreate = OPR_CREATE
								c.fileName = ev.Name
								if fi.IsDir() {
									c.opreate = OPR_DIR_ADD
									w.watch.Add(ev.Name)
								}
								c.post()
							}
						}

					}
					if ev.Op&fsnotify.Write == fsnotify.Write {
						if !isIgnorePath(ev.Name, FileIgnoreWord) {
							c.fileName = ev.Name
							c.opreate = OPR_CREATE
							c.post()
							fmt.Println("写入文件 : ", ev.Name)
						}

					}

					if ev.Op&fsnotify.Remove == fsnotify.Remove {
						if !isIgnorePath(ev.Name, FileIgnoreWord) {
							fmt.Println("删除文件 : ", ev.Name)
							//如果删除文件是目录，则移除监控
							fi, err := os.Stat(ev.Name)
							c.fileName = ev.Name
							c.opreate = OPR_REMOVE
							c.post()
							if err == nil {
								if fi.IsDir() {
									w.watch.Remove(ev.Name)
									fmt.Println("删除监控 : ", ev.Name)
								}
							} else {
								fmt.Println(err)
							}
						}

					}
					if ev.Op&fsnotify.Rename == fsnotify.Rename {
						if !isIgnorePath(ev.Name, FileIgnoreWord) {
							c.fileName = ev.Name
							c.opreate = OPR_REMOVE
							c.post()
							fmt.Println("重命名 : ", ev.Name)
						}
						//如果重命名文件是目录，则移除监控
						//注意这里无法使用os.Stat来判断是否是目录了
						//因为重命名后，go已经无法找到原文件来获取信息了
						//所以这里就简单粗爆的直接remove好了
						w.watch.Remove(ev.Name)

					}
					if ev.Op&fsnotify.Chmod == fsnotify.Chmod {
						if isIgnorePath(ev.Name, FileIgnoreWord) {
							MapFile[ev.Name] = "change"
							fmt.Println("修改权限 : ", ev.Name)
						}

					}
				}
			case err := <-w.watch.Errors:
				{
					fmt.Println("error : ", err)
					return
				}
			}
		}
	}()
}

//入口函数
func WatchDir() {
	//监控本地文件变化
	watch, _ := fsnotify.NewWatcher()
	w := Watch{
		watch: watch,
	}

	w.watchDir(localDirPath)
	select {}
}


