package main

import (
	"fmt"
	"os"

	"github.com/rk295/go-freshservice/cmd/freshservice-cli/commands"
)

func main() {
	err := commands.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
