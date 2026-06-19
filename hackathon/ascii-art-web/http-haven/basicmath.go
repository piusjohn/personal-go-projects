package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func BasicMath(w http.ResponseWriter, r *http.Request){
	A := r.URL.Query().Get("a")
	B := r.URL.Query().Get("b")
	opt :=  r.URL.Query().Get("op")
	fnum, err := strconv.Atoi(A)

	if err != nil{
		http.Error(w, "string conversion failed", 400)
		return
	}
	snum, err := strconv.Atoi(B)

	if err != nil{
		http.Error(w, "string conversion failed", 500)
		return
	}
	var result int
	if opt == "add"{
		result = fnum + snum
	} else if opt == "subtract"{
		result = fnum - snum
	} else if opt == "multiply"{
		result = fnum * snum
	} else {
		http.Error(w, "Unknown Operation", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "Result: %d", result)

}