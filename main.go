package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type Task struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Id          uuid.UUID `json:"id"`
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
	var newTask Task
	bodyBytes, _ := io.ReadAll(r.Body)
	err := json.Unmarshal(bodyBytes, &newTask)
	if err != nil {
		fmt.Println(err)
	}
	taskCreate := NewTask(newTask.Name, newTask.Description)
	Tasks = append(Tasks, taskCreate)
	w.WriteHeader(http.StatusOK)
}

func editTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var currentTask Task
	bodyBytes, _ := io.ReadAll(r.Body)
	err := json.Unmarshal(bodyBytes, &currentTask)
	if err != nil {
		fmt.Println(err)
	}
	i := 0
	for i <= len(Tasks) {
		if (Tasks[i].Id.String()) == id {
			Tasks[i].Name = currentTask.Name
			Tasks[i].Description = currentTask.Description
			b, err := json.Marshal(Tasks[i])
			if err != nil {
				fmt.Println(err)
			}
			w.Write(b)
			return
		}
		i++
	}
}

func getTasks(w http.ResponseWriter, r *http.Request) {
	b, err := json.Marshal(Tasks)
	if err != nil {
		fmt.Println(err)
	}
	w.Write(b)
}
func getTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	i := 0
	for i < len(Tasks) {
		fmt.Println(Tasks[i].Id.String())
		if (Tasks[i].Id).String() == id {
			b, err := json.Marshal(Tasks[i])
			if err != nil {
				fmt.Println(err)
			}
			w.Write(b)
			return
		}
		i++
	}
}

func deleteTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	i := 0
	for i < len(Tasks) {
		fmt.Println(Tasks[i].Id.String())
		if (Tasks[i].Id).String() == id {
			Tasks = append(Tasks[:i], Tasks[i+1:]...)
			w.WriteHeader(http.StatusOK)
		}
		i++
	}
}

func main() {
	m := mux.NewRouter()
	m.HandleFunc("/addTask", addTask).Methods("POST")
	m.HandleFunc("/editTask/{id}", editTask).Methods("PATCH")
	m.HandleFunc("/getTasks", getTasks).Methods("GET")
	m.HandleFunc("/getTasks/{id}", getTask).Methods("GET")
	m.HandleFunc("/deleteTask/{id}", deleteTask).Methods("DELETE")
	err := http.ListenAndServe(":8080", m)
	if err != nil {
		fmt.Println(err)
	}

}
