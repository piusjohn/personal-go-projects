package main

import(
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	apiMux := http.NewServeMux()
	apiMux.HandleFunc("/v1/ping", pingHandler)
	apiMux.HandleFunc("/v1/greet", greetHandler)

	mainMux := http.NewServeMux()

	mainMux.Handle("/api/", http.StripPrefix("/api", apiMux))

	mux.HandleFunc("/status", StatusCodeFactory)
	mux.HandleFunc("/form", FormDecode)
	mux.HandleFunc("/headers", HeaderDetective)
	mux.HandleFunc("/echo", EchoChamber)
	mux.HandleFunc("/method-inspector", MethodInspector)

	http.ListenAndServe(":8080", mainMux)
	http.ListenAndServe(":8080", mux)
}