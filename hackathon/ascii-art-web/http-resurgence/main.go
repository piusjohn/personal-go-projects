package main

import(
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/headers", HeaderDetective)
	mux.HandleFunc("/echo", EchoChamber)
	mux.HandleFunc("/method-inspector", MethodInspector)
	http.ListenAndServe(":8080", mux)
}