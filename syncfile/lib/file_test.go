package lib

import (
	"testing"
)

func TestIgnore(t *testing.T) {
	if !isIgnorePath("C:\\Users\\Administrator\\go\\src\\test1\\12.txt___jb_tmp___", FileIgnoreWord) {
		t.Error("t___jb_tmp___" + "==>false")
	}
}
