package listrars

import (
	"io/ioutil"
	"os"
	"path"
	"strings"
)

var Path string

func Init(path string) {
	Path = path
}

func getRars() ([]os.FileInfo, error) {
	files := ioutil.ReadDir(Path)
	rars := []os.FileInfo{}
	for _, file := range files {
		extension := path.Ext(file.Name())
		if strings.ToLower(extension) == ".rar" {
			rars = append(rars, file)
		}
	}
}
