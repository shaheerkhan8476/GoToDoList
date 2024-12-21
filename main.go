package main

import (
	"fmt"
	"net/http"
	"todo/routes/task"

	"github.com/gorilla/mux"
)

func main() {
	m := mux.NewRouter()
	m.HandleFunc("/addTask", task.AddTask).Methods("POST")
	m.HandleFunc("/editTask/{id}", task.EditTask).Methods("PATCH")
	m.HandleFunc("/getTasks", task.GetTasks).Methods("GET")
	m.HandleFunc("/getTasks/{id}", task.GetTask).Methods("GET")
	m.HandleFunc("/deleteTask/{id}", task.DeleteTask).Methods("DELETE")
	err := http.ListenAndServe(":8080", m)
	if err != nil {
		fmt.Println(err)
	}

}
