package main

import (
	"fmt"
	"github.com/guneyin/bookstore/cmd"
	"github.com/guneyin/bookstore/cmd/app"
	"github.com/guneyin/bookstore/cmd/gen"
)

func main() {
	cmd.RootCmd.AddCommand(app.Cmd)
	cmd.RootCmd.AddCommand(gen.Cmd)

	err := cmd.RootCmd.Execute()
	if err != nil {
		fmt.Println(err)
	}
}
