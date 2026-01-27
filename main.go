package main

import (
	"flag"
	"fmt"
)

var name string

func init() {
	flag.StringVar(&name, "name", "", "Name to greet")
	flag.Parse()
}

func main() {
	if name == "" {
		fmt.Println("Please provide a name using the -name flag")
		return
	}
	fmt.Printf("Hello, %s!\n", name)
}

func add() {
	fmt.Println("Adding")
}

func list() {
	fmt.Println("Listing")
}

func complete() {
	fmt.Println("Completing")
}

func delete() {
	fmt.Println("Deleting")
}
