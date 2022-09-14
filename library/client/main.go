package main

import (
	"fmt"

	"github.com/u-shylianok/go-workspaces/library"
)

func main() {
	fmt.Println("Library client up")
	library.Greetings("Boba")
	fmt.Println("Library client down")
}
