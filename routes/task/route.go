package task

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

var Tasks []Task = []Task{}

func AddTask(w http.ResponseWriter, r *http.Request) {
	var newTask Task
	bodyBytes, _ := io.ReadAll(r.Body)
	err := json.Unmarshal(bodyBytes, &newTask)
	if err != nil {
		fmt.Println(err)
	}
	newTask.Id = uuid.New()
	Tasks = append(Tasks, newTask)
	w.WriteHeader(http.StatusOK)
}

func EditTask(w http.ResponseWriter, r *http.Request) {
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

func GetTasks(w http.ResponseWriter, r *http.Request) {
	b, err := json.Marshal(Tasks)
	if err != nil {
		fmt.Println(err)
	}
	w.Write(b)
}
func GetTask(w http.ResponseWriter, r *http.Request) {
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

func DeleteTask(w http.ResponseWriter, r *http.Request) {
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
