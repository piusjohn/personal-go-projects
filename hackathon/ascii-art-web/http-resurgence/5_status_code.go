package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func StatusCodeFactory(w http.ResponseWriter, r *http.Request){
	code := r.URL.Query().Get("code")
	if code == ""{
		http.Error(w, "code parameter is required", http.StatusBadRequest)
		return
	}
	n, err := strconv.Atoi(code)
	if err != nil{
		http.Error(w, "code must be a valid integer", http.StatusBadRequest)
		return
	}

	if n < 100 || n > 599{
		http.Error(w, "code must be a valid HTTP status code", http.StatusBadRequest)
		return
	} 
	w.WriteHeader(n)
	fmt.Fprintf(w, "Responding with status %s", code)
}