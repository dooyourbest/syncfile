## 取代samba进行代码同步
#### 代码分为两端 client server,需要在本地和开发机同时安装代码，才能进行通信
## 安装和配置
### 1 依赖于 github.com/fsnotify/fsnotify
   没有安装需要安装一下 go get  github.com/fsnotify/fsnotify
### 2 下载syncfile：
   go get github.com/dooyourbest/syncfile 
   这样会下载到你的gopath 目录
## 修改配置
### 1 找到gopath 下修改 go/src/github.com/dooyourbest/syncfile/syncfile/common.go ，
### 2 修改26 -29行即可 （本地和开发机此文件需要一致，修改好一份，粘贴过去替换即可）
 注意：host需要加上schema 即 localhost需要写成http:\\localhost
## 编译
go install syncfile
 本地开发机都需要编译

### 即可享受伪samba，
命令:
  本地运行     syncfile local 命令
  开发机运行  syncfile dev 命令
  两边同时开启即可，本地修改回同步到开发机
 其他
  手动push文件  syncfile push 本地文件绝对路径
  手动get文件  syncfile get 线上文件绝对路径
  需要完善：
  #### 功能点：
  
  1 全量拉取线上目录下文件
  2 增加权限控制
  #### 目前已知的缺点：
  手动push,get 需要不能递归创建文件，需要存在上级目录，才能同步成功
  比如本地是 ~/test/
  拉取线上  ~/test/a1/1.txt， ~/test/a1/a2/a.txt 会失败
  只能拉取 ~/test/层级下的文件 , 如~/test/c.txt，，可以手动创建文件加进行避免

