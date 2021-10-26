package main

import (
	"fmt"
	"os"

	"github.com/iamgrewal7/masv/util"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("No arguments provided")
		return
	}

	args := os.Args[1:]
	err := util.ValidateInput(args)

	if err != nil {
		fmt.Printf("Error: %s", err.Error())
		return
	}

}
