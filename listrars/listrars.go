package listrars

import (
	"io/ioutil"
	"path"
	"regexp"
	"strings"
)

// Path is the directory string.
var Path string

// File is an interface type to represent a file (only the filename is needed).
type File interface {
	Name() string
}

// Init initializes the package listrars with a path to work on.
func Init(path string) {
	Path = path
}

// GetRars extracts the rar files in the given directory.
// If multipart rar files are detected, only the first part will
// be included in the result.
// The test argument is only used for testing, so to use this
// function in a non-testing scenario, an empty slice should be
// passed.
func GetRars(test []interface{}) ([]string, error) {
	files := []string{}
	if len(test) > 0 {
		for _, elem := range test {
			files = append(files, elem.(File).Name())
		}
	} else {
		fs, err := ioutil.ReadDir(Path)
		if err != nil {
			return nil, err
		}
		for _, file := range fs {
			files = append(files, file.Name())
		}
	}
	rarNames := []string{}
	for _, file := range files {
		extension := path.Ext(file)
		if strings.ToLower(extension) == ".rar" {
			rarNames = append(rarNames, file)
		}
	}
	strippedRars := []string{}
	regexGeneralPart := "(?i)part[0-9]+\\.rar$"
	regexFirstPart := "(?i)part[0]*1\\.rar$"
	for _, rar := range rarNames {
		matchFirst, errorFirst := regexp.MatchString(regexFirstPart, rar)
		if errorFirst != nil {
			return nil, errorFirst
		}
		matchGeneral, errorGeneral := regexp.MatchString(regexGeneralPart, rar)
		if errorGeneral != nil {
			return nil, errorGeneral
		}
		if matchFirst {
			strippedRars = append(strippedRars, rar)
		} else if matchGeneral {
			continue
		} else {
			strippedRars = append(strippedRars, rar)
		}
	}
	return strippedRars, nil
}
