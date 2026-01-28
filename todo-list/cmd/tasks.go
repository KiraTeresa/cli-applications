package cmd

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
)

func Add(task string) {
	fmt.Printf("Adding task '%s'\n", task)

	// open csv file if available, otherwise create it
	file, err := os.OpenFile("tasks.csv", os.O_RDWR|os.O_CREATE, 0660)
	if err != nil {
		fmt.Println("Error opening tasks.csv:", err)
		return
	}
	defer file.Close()

	// read the data from the csv file
	r := csv.NewReader(file)
	records, err := r.ReadAll()
	if err != nil {
		fmt.Println("Error reading tasks.csv:", err)
		return
	}

	// ensure to start writing at the beginning
	if _, err := file.Seek(0, 0); err != nil {
		log.Println("Error seeking file:", err)
		return
	}

	w := csv.NewWriter(file)

	// write a header if the file was newly created
	if len(records) == 0 {
		header := []string{"id", "task", "status"}
		err := w.Write(header)
		if err != nil {
			log.Fatalln("error writing header to task.csv:", err)
		}
	}

	// create new data & write it to the file
	nextId := strconv.Itoa(len(records) + 1)
	s := []string{nextId, task, "open"}
	if err := w.Write(s); err != nil {
		log.Fatalln("error writing to tasks.csv:", err)
	}

	// ensure all buffered data was written to the file
	w.Flush()

	// ensure no I/O errors occurred
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
