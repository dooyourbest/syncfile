package syncfile

func PushFile(localFilePath string) {
	c := Client{fileName: localFilePath}
	c.opreate = OPR_CREATE
	c.post()
}
