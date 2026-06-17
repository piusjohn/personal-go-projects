package main

import (
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request ){
	if r.URL.Path != "/"{
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	// if r.Method == "GET"{
	// 	fmt.Fprintf(w, "got a GET request\n")
	// } else if r.Method == "POST"{
	// 	fmt.Fprintf(w, "got a POST request\n")
	// }
	w.Header().Set("content-type", "text/html")

	fmt.Fprintf(w, `<html>
            <body>
                <h1>Hello from Go</h1>
                <p>This is HTML</p>
            </body>
        </html>`)
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.ListenAndServe(":8080", nil)
}