package cmd

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"slices"
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

func updateFile(file *os.File, data []Task) {
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
	if err := w.WriteAll(tasksToRecords(data)); err != nil {
		log.Fatalln("error updating tasks.csv:", err)
	}

	// ensure all buffered data was written to the file
	w.Flush()

	// ensure no I/O errors occurred
	if err := w.Error(); err != nil {
		log.Fatal(err)
	}
}

func getIndexOfId(id string, tasks []Task) int {
	// find index of the entry which should be modified
	idx := slices.IndexFunc(tasks, func(t Task) bool {
		return t.id == id
	})

	if idx <= 0 {
		log.Printf("Task with id %s not found\n", id)
		os.Exit(1)
	}
	return idx
}

func getTasksFromFile(file *os.File) []Task {
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading tasks.csv:", err)
		os.Exit(1)
	}
	return recordsToTask(records)
}

func recordsToTask(records [][]string) []Task {
	var tasks []Task
	for _, r := range records {
		tasks = append(tasks, Task{
			id:     r[0],
			title:  r[1],
			status: r[2],
		})
	}
	return tasks
}

func taskToRecord(t Task) []string {
	return []string{t.id, t.title, t.status}
}

func tasksToRecords(tasks []Task) [][]string {
	var records [][]string
	for _, t := range tasks {
		records = append(records, taskToRecord(t))
	}
	return records
}
