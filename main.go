package main

import (
	"fmt"
	"monkey/repl"
	"os"
	"os/user"
)

func main() {
	u, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! Welcome to the Monkey programming language!\n", u.Username)
	fmt.Println("Begin typing commands")
	repl.Start(os.Stdin, os.Stdout)
}
