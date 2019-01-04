package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/kcwebapply/kc-interpreter/repl"
)

func main() {
	user, _ := user.Current()

	fmt.Printf("Hello %s!", user.Username)

	repl.Start(os.Stdin, os.Stdout)
}
