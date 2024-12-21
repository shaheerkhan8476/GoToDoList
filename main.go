package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/google/uuid"
)

type Task struct {
	Name        string
	Description string
	Id          uuid.UUID
}

func NewTask(name string, description string) Task {
	return Task{
		Name:        name,
		Description: description,
		Id:          uuid.New(),
	}
}

var Tasks []Task = []Task{}

func addTask(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		var newTask Task
		bodyBytes, _ := io.ReadAll(r.Body)
		err := json.Unmarshal(bodyBytes, &newTask)
		if err != nil {
			fmt.Println(err)
		}
		taskCreate := NewTask(newTask.Name, newTask.Description)
		Tasks = append(Tasks, taskCreate)
	case "GET":
		b, err := json.Marshal(Tasks)
		if err != nil {
			fmt.Println(err)
		}
		w.Write(b)

	}
	w.WriteHeader(http.StatusOK)
}

func editTask(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		var currentTask Task
		bodyBytes, _ := io.ReadAll(r.Body)
		err := json.Unmarshal(bodyBytes, &currentTask)
		if err != nil {
			fmt.Println(err)
		}
		i := 1
		for i <= len(Tasks) {
			if (Tasks[i].Id) == currentTask.Id {
				Tasks[i].Name = currentTask.Name
				Tasks[i].Description = currentTask.Description
				w.WriteHeader(http.StatusOK)
				return
			}
		}
		w.WriteHeader(http.StatusNotModified)
	}

}

func main() {
	m := http.NewServeMux()
	m.HandleFunc("/addTask", addTask)
	m.HandleFunc("/editTask", editTask)
	err := http.ListenAndServe(":8080", m)
	if err != nil {
		fmt.Println(err)
	}

}
