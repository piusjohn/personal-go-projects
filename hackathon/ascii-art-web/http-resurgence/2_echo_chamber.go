package main

import (
	//"fmt"
	"io"
	"net/http"
)

func EchoChamber(w http.ResponseWriter, r *http.Request){
	if r.Method != "POST"{
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	data, err := io.ReadAll(r.Body)
	if err != nil{
		http.Error(w, "error reading r.Body", http.StatusInternalServerError)
		return
	}

	defer r.Body.Close()
	if len(data) == 0{
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "text/plain")
	// fmt.Print(w, data)
	w.Write(data)
}