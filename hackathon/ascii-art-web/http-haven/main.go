package main

import "net/http"

func main(){
	mux := http.NewServeMux()
	mux.HandleFunc("/v2", Welcome)
	mux.HandleFunc("/legacy", Redirect)
	mux.HandleFunc("/dashboard", Secure)
	mux.HandleFunc("/agent", UserAgent)
	mux.HandleFunc("/calculate", BasicMath)
	mux.HandleFunc("/count", counter)
	mux.HandleFunc("/hello", hello)
	mux.HandleFunc("/ping", pingpong)
	http.ListenAndServe(":8080", mux)
}