package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Sub struct{
	Username string
	Data string
}

var tpl *template.Template

func main() {
	tpl = template.Must(tpl.ParseGlob("templates/*.html"))
	http.HandleFunc("/getform", getFormHandler)

}

func getFormHandler(w http.ResponseWriter, r *http.Request){
	tpl.ExecuteTemplate(w, "getform.html", nil)
}

func processGetHandler(w http.ResponseWriter, r *http.Request){
	fmt.Println("processGetHandler is running")
	var s Sub
	s.Username = r.URL.Query().Get("usernameName")
	s.Data = r.URL.Query().Get("dataName")
	fmt.Println("Username:", s.Username, "Sensitive Data:", s.Data)
	tpl.ExecuteTemplate(w, "thanks.html", s)
}