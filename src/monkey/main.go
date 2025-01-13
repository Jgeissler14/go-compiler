package main

import (
	"fmt"
	"monkey/repl"
	"os"
	"os/user"
)

func main() {
	if currentUser, err := user.Current(); err != nil {
		panic(err)
	} else {
		fmt.Printf("Hello %s! This is the Monkey programming language!\n", currentUser.Username)
	}

	fmt.Println("Feel free to type in commands")
	repl.Start(os.Stdin, os.Stdout)
}
