package main

import (
	"fmt"
	"net/http"
)

func createWebserver() {

	http.HandleFunc("/", processCall)
	err := http.ListenAndServe("localhost:8080", nil)

	if err != nil {
		fmt.Println("Error!!", err)
	}
}

func processCall(w http.ResponseWriter, r *http.Request) {

	headerMap := w.Header()

	headerMap.Add("Go-Test", "Test 1.0")

	switch r.Method {

	case "GET":

		w.Write([]byte("Prueba webserver"))

	case "POST":

		headerMap.Add("Unauthorized", "You need to have a valid token")
		w.WriteHeader(401)
		return
		
	}
}
