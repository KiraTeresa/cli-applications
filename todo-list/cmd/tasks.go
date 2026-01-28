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
	// open csv readFile if available, otherwise create it
	readFile, err := os.OpenFile("tasks.csv", os.O_RDONLY|os.O_CREATE, 0660)
	if err != nil {
		fmt.Println("Error opening tasks.csv:", err)
		return
	}

	// read the data from the csv readFile
	r := csv.NewReader(readFile)
	records, err := r.ReadAll()
	if err != nil {
		fmt.Println("Error reading tasks.csv:", err)
		return
	}
	readFile.Close()

	// open csv file with append flag
	appendFile, err := os.OpenFile("tasks.csv", os.O_APPEND, 0660)
	if err != nil {
		fmt.Println("Error opening tasks.csv:", err)
		return
	}
	defer appendFile.Close()

	var nextId string
	w := csv.NewWriter(appendFile)

	// write a header if the readFile was newly created
	if len(records) == 0 {
		header := []string{"id", "task", "status"}
		err := w.Write(header)
		if err != nil {
			log.Fatalln("error writing header to task.csv:", err)
		}
	}

	// if file is empty or only contains header, start with id=1
	if len(records) <= 1 {
		nextId = "1"
	} else { // otherwise set id to the next number
		lastId, err := strconv.Atoi(records[len(records)-1][0])
		if err != nil {
			fmt.Println("error evaluating last id:", err)
			return
		}
		nextId = strconv.Itoa(lastId + 1)
	}

	// create new data & write it to the readFile
	s := []string{nextId, task, "open"}
	if err := w.Write(s); err != nil {
		log.Fatalln("error writing to tasks.csv:", err)
	}

	// ensure all buffered data was written to the readFile
	w.Flush()

	// ensure no I/O errors occurred
	if err := w.Error(); err != nil {
		log.Fatal(err)
	}
}

func List(listAll bool) {
	// open the csv file
	file, err := os.Open("tasks.csv")
	if err != nil {
		fmt.Println("Error opening tasks.csv:", err)
		return
	}
	defer file.Close()

	// retrieve all data from the csv file
	r := csv.NewReader(file)
	records, err := r.ReadAll()
	if err != nil {
		fmt.Println("Error reading tasks.csv:", err)
		return
	}

	if listAll {
		// list all tasks
		fmt.Println(records)
	} else {
		// only list tasks with status "open"
		result := make([][]string, 0, len(records))

		for i, r := range records {
			if i == 0 || r[2] == "open" {
				result = append(result, r)
			}
		}

		fmt.Println(result)
	}
}

func Complete(id string) {
	// open csv file
	file, err := os.OpenFile("tasks.csv", os.O_RDWR, 0660)
	if err != nil {
		fmt.Println("Error opening tasks.csv:", err)
		return
	}
	defer file.Close()

	// read data from csv file
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading tasks.csv:", err)
		return
	}

	// find index of the entry which should be modified
	idx := slices.IndexFunc(records, func(r []string) bool {
		return len(r) > 0 && r[0] == id
	})
	if idx == -1 {
		log.Printf("Task with id %s not found\n", id)
		return
	}

	// update entry
	records[idx][2] = "done"

	// truncate & reset file before writing
	if err := file.Truncate(0); err != nil {
		log.Println("Error truncating file:", err)
		return
	}

	// ensure to start writing at the beginning
	if _, err := file.Seek(0, 0); err != nil {
		log.Println("Error seeking file:", err)
		return
	}

	// write updated data to csv file
	w := csv.NewWriter(file)
	if err := w.WriteAll(records); err != nil {
		log.Fatalln("error updating tasks.csv:", err)
	}

	// ensure all buffered data was written to the file
	w.Flush()

	// ensure no I/O errors occurred
	if err := w.Error(); err != nil {
		log.Fatal(err)
	}
}

func Delete(id string) {
	// open csv file
	file, err := os.OpenFile("tasks.csv", os.O_RDWR, 0660)
	if err != nil {
		fmt.Println("Error opening tasks.csv:", err)
		return
	}
	defer file.Close()

	// read data from csv file
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading tasks.csv:", err)
		return
	}

	// find index of the entry which should be deleted
	idx := slices.IndexFunc(records, func(r []string) bool {
		return len(r) > 0 && r[0] == id
	})
	if idx == -1 {
		log.Printf("Task with id %s not found\n", id)
		return
	}

	// update data
	isLastElement := len(records)-1 == idx
	if isLastElement {
		records = append(records[:idx])
	} else {
		records = append(records[:idx], records[idx+1:]...)
	}

	// truncate & reset file before writing
	if err := file.Truncate(0); err != nil {
		log.Println("Error truncating file:", err)
		return
	}

	// ensure to start writing at the beginning
	if _, err := file.Seek(0, 0); err != nil {
		log.Println("Error seeking file:", err)
		return
	}

	// write updated data to csv file
	w := csv.NewWriter(file)
	if err := w.WriteAll(records); err != nil {
		log.Fatalln("error updating tasks.csv:", err)
	}

	// ensure all buffered data was written to the file
	w.Flush()

	// ensure no I/O errors occurred
	if err := w.Error(); err != nil {
		log.Fatal(err)
	}
}
