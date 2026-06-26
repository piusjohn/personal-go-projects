package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var tpl *template.Template
func main(){
	tpl, _ = tpl.ParseGlob("templates/*.html")
	myDir := http.Dir("./public")

	fmt.Printf("myDir: %T", myDir)
	myHandler := http.FileServer(myDir)
	handle := http.StripPrefix("/public/", myHandler)
	http.Handle("/public/", handle)

	http.HandleFunc("/hello", hellohandler)
	http.ListenAndServe(":8080", nil)
}

func hellohandler(w http.ResponseWriter, r *http.Request){
	tpl.ExecuteTemplate(w, "hello.html", nil)
}