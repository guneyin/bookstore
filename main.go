package main

import (
	"fmt"
	"github.com/guneyin/bookstore/cmd"
	"github.com/guneyin/bookstore/cmd/app"
	"github.com/guneyin/bookstore/cmd/gen"
)

// @title The Book Store API Doc
// @version 1.0
// @description A case study project to demonstrate an online bookstore based on Golang

// @contact.name Hüseyin Güney
// @contact.url https://github.com/guneyin
// @contact.email guneyin@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:3000
// @BasePath /api/
// @schemes http
func main() {
	cmd.RootCmd.AddCommand(app.Cmd)
	cmd.RootCmd.AddCommand(gen.Cmd)

	err := cmd.RootCmd.Execute()
	if err != nil {
		fmt.Println(err)
	}
}
