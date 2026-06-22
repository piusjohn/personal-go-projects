package main

import (
	"fmt"
	"net/http"
)

func HeaderDetective(w http.ResponseWriter, r *http.Request){
	
	//if there are no value associated with the given key, get returns "".

	token := r.Header.Get("X-Custom-Token")
	if len(token) == 0 {
		http.Error(w, "X-Custom-Token header is missing", http.StatusBadRequest)
		return
	}

	contenttype := r.Header.Get("Content-Type")
	if len(contenttype) == 0{
		token = token + "Content-Type not provided"
	}
	token = token + " " + contenttype
	fmt.Fprint(w, token)
}