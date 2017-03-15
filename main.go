package main

import (
	"fmt"
	"os"

	docopt "github.com/docopt/docopt-go"
	"github.com/goggle/listrars/listrars"
)

const version = "0.5"

func main() {
	usage := `List Rars.

Usage:
  listrars [PATH]

Options:
  -h --help     Show this screen.
  --version     Show version.`

	arguments, _ := docopt.Parse(usage, nil, true, "listrars "+version, false)
	var path string
	p := arguments["PATH"]
	if p == nil {
		pp, err := os.Getwd()
		if err != nil {
			fmt.Printf("%v\n", err)
			os.Exit(1)
		}
		path = pp
	} else {
		path = p.(string)
	}

	listrars.Init(path)
	rarList, error := listrars.GetRars(nil)
	if error != nil {
		fmt.Printf("%v", error)
		os.Exit(1)
	}
	for _, rar := range rarList {
		fmt.Println(rar)
	}

}
