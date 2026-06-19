package main

import (
	"fmt"
	"net/http"
)

func pingpong(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "pong")
}