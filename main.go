package main

import (
	"flag"
	"fmt"
	"github.com/seedgo/seedgo/action"
)

var createType string
var createName string

func init() {
	flag.StringVar(&createType, "create", "project", "specify the create type: project")
	flag.StringVar(&createName, "name", "", "specify the name")
}

func main() {
	flag.Parse()
	var err error
	if createType == "project" {
		err = action.CreateProject(createName)
	}

	if err != nil {
		fmt.Println(err.Error())
	}

}
