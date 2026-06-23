package main

import (
	"fmt"
	"net/http"
)

func pingHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "pong")
}

func greetHandler(w http.ResponseWriter, r *http.Request){
	if r.Method != "GET"{
		http.Error(w, "Method Not Available", http.StatusMethodNotAllowed)
		return
	}
	text := r.URL.Query().Get("name")
	if text == ""{
		text = "Stranger!"
	} 
	fmt.Fprintf(w, "Greetings, %s!", text )
}

