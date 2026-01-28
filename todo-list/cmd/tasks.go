package cmd

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"text/tabwriter"
)

type Task struct {
	id     string
	title  string
	status string
}

func Add(task string) {
	// open csv readFile if available, otherwise create it
	readFile, err := os.OpenFile("tasks.csv", os.O_RDONLY|os.O_CREATE, 0660)
	if err != nil {
		fmt.Println("Error opening tasks.csv:", err)
		return
	}

	// read the data from the csv readFile
	tasks := getTasksFromFile(readFile)
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
	if len(tasks) == 0 {
		header := Task{"id", "title", "status"}
		err := w.Write(taskToRecord(header))
		if err != nil {
			log.Fatalln("error writing header to Task.csv:", err)
		}
	}

	// if file is empty or only contains header, start with id=1
	if len(tasks) <= 1 {
		nextId = "1"
	} else { // otherwise set id to the next number
		lastId, err := strconv.Atoi(tasks[len(tasks)-1].id)
		if err != nil {
			fmt.Println("error evaluating last id:", err)
			return
		}
		nextId = strconv.Itoa(lastId + 1)
	}

	// create new data & write it to the readFile
	s := Task{nextId, task, "open"}
	if err := w.Write(taskToRecord(s)); err != nil {
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
	tasks := getTasksFromFile(file)

	var result []Task
	if listAll {
		// list all tasks
		result = tasks
	} else {
		// only list tasks with status "open"
		result = make([]Task, 0, len(tasks))

		for i, t := range tasks {
			if i == 0 || t.status == "open" {
				result = append(result, t)
			}
		}
	}

	tw := new(tabwriter.Writer)
	tw.Init(os.Stdout, 0, 8, 2, ' ', 0)

	for i, r := range result {
		fmt.Fprintf(tw, "%s\t%s\t%s\n", r.id, r.title, r.status)
		if i == 0 {
			fmt.Fprintf(tw, "--\t----\t------\n")
		}
	}
	tw.Flush()
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
	tasks := getTasksFromFile(file)

	// find index of the entry which should be modified
	idx := getIndexOfId(id, tasks)

	// update entry
	tasks[idx].status = "done"

	updateFile(file, tasks)
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
	records := getTasksFromFile(file)

	// find index of the entry which should be deleted
	idx := getIndexOfId(id, records)

	// update data
	isLastElement := len(records)-1 == idx
	if isLastElement {
		records = append(records[:idx])
	} else {
		records = append(records[:idx], records[idx+1:]...)
	}

	updateFile(file, records)
}
