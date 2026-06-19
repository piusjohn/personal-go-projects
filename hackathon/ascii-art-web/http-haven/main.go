package main

import "net/http"

func main(){
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", hello)
	mux.HandleFunc("/ping", pingpong)
	http.ListenAndServe(":8080", mux)
}