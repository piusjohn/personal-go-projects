package main

import (
	"html/template"
	"log"
	"net/http"
)

type Sub struct{
	Username string
	Data string
}

var tpl *template.Template

func main() {
	tpl, err := tpl.ParseGlob("templates/*.html")
	log.Fatal(err)
}

func getFormHandler(w http.ResponseWriter, r *http.Request){
	tpl.ExecuteTemplate(w, "getform.html", nil)
}

func processGetHandler(w http.ResponseWriter, r *http.Request){
	
}