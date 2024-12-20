package main

import (
	"fmt"
	"net/http"
)

func exampleEndpoint(w http.ResponseWriter, r *http.Request) {
	fmt.Println("exampleEndpoint Hit")
	w.WriteHeader(http.StatusOK)
}
func main() {
	fmt.Println("Hello")
	m := http.NewServeMux()
	m.HandleFunc("/exampleEndpoint", exampleEndpoint)
	err := http.ListenAndServe(":8080", m)
	if err != nil {
		fmt.Println(err)
	}
}
