package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Task struct {
	Name        string
	Description string
}

var Tasks []Task = []Task{}

func exampleEndpoint(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		var newTask Task
		bodyBytes, _ := io.ReadAll(r.Body)

		err := json.Unmarshal(bodyBytes, &newTask)
		if err != nil {
			fmt.Println(err)
		}
		Tasks = append(Tasks, newTask)
	case "GET":
		w.Write([]byte(fmt.Sprintf("%v", Tasks)))
	}

	w.WriteHeader(http.StatusOK)
}
func main() {
	m := http.NewServeMux()
	m.HandleFunc("/exampleEndpoint", exampleEndpoint)
	err := http.ListenAndServe(":8080", m)
	if err != nil {
		fmt.Println(err)
	}
}
