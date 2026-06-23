package main

import (
	"fmt"
	"net/http"
)

func FormDecode(w http.ResponseWriter, r *http.Request){
	if r.Method != "POST"{
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	err := r.ParseForm()
	if err != nil{
		http.Error(w, "Parsing Form Failed", http.StatusInternalServerError)
		return
	}

	username := r.FormValue("username")
	language := r.FormValue("language")

	if username == ""{
		http.Error(w, "Missing username → username is required", http.StatusBadRequest)
		return
	}

		if language == ""{
		http.Error(w, "Missing language → language is required", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "Hello %s, you are coding in %s!", username, language)
}