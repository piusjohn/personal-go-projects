package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET"{
		http.Error(w, "Method Not Available", http.StatusMethodNotAllowed)
		return
	}
	text := r.URL.Query().Get("name")
	if text == ""{
		text = "Guest!"
	} 
	fmt.Fprintf(w, "Hello, %s!", text )
}