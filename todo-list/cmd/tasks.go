package cmd

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func Add(task string) {
	fmt.Printf("Adding task '%s'\n", task)

	file, err := os.OpenFile("tasks.csv", os.O_APPEND|os.O_CREATE, 0660)
	if err != nil {
		fmt.Println("Error opening tasks.csv:", err)
		return
	}

	s := []string{task}
	w := csv.NewWriter(file)
	if err := w.Write(s); err != nil {
		log.Fatalln("error writing to tasks.csv:", err)
	}

	w.Flush()

	if err := w.Error(); err != nil {
		log.Fatal(err)
	}
}

func List() {
	fmt.Println("Listing")

	file, err := os.Open("tasks.csv")
	if err != nil {
		fmt.Println("Error opening tasks.csv:", err)
		return
	}

	r := csv.NewReader(file)
	records, err := r.ReadAll()
	if err != nil {
		fmt.Println("Error reading tasks.csv:", err)
		return
	}

	fmt.Println(records)
}

func Complete(id string) {
	fmt.Printf("Completing task with id '%s'\n", id)
}

func Delete(id string) {
	fmt.Printf("Deleting id '%s'\n", id)
}
