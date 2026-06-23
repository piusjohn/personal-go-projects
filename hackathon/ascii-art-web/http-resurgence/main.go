package main

import(
	"net/http"
)

func main() {
	apiMux := http.NewServeMux()
	apiMux.HandleFunc("/v1/ping", pingHandler)
	apiMux.HandleFunc("/v1/greet", greetHandler)

	mainMux := http.NewServeMux()

	mainMux.Handle("/api/", http.StripPrefix("/api", apiMux))

	mainMux.HandleFunc("/render", TemplateRender)
	mainMux.HandleFunc("/status", StatusCodeFactory)
	mainMux.HandleFunc("/form", FormDecode)
	mainMux.HandleFunc("/headers", HeaderDetective)
	mainMux.HandleFunc("/echo", EchoChamber)
	mainMux.HandleFunc("/method-inspector", MethodInspector)

	http.ListenAndServe(":8080", mainMux)
}