package main

import(
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/method-inspector", MethodInspector)
	http.ListenAndServe(":8080", mux)
}