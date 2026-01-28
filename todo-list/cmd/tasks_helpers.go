package cmd

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"slices"
)

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
