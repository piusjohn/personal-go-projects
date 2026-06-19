package main

import (
	"fmt"
	"net/http"
	"io"
)

func counter(w http.ResponseWriter, r *http.Request){
	if r.Method == "GET"{
		fmt.Fprint(w, "Send a POST request with text to count words")
		return
	}
	text, err := io.ReadAll(r.Body)
		if err != nil{
			http.Error(w, "500 internal server error", 500)
			return
		}
	l := len(text)
	fmt.Fprint(w, l)
	
}