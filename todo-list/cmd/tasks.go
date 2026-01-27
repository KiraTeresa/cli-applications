package cmd

import "fmt"

func Add(task string) {
	fmt.Printf("Adding task '%s'\n", task)
}

func List() {
	fmt.Println("Listing")
}

func Complete(id string) {
	fmt.Printf("Completing task with id '%s'\n", id)
}

func Delete(id string) {
	fmt.Printf("Deleting id '%s'\n", id)
}
