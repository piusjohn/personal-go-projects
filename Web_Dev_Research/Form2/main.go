package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)
var tpl *template.Template
func main(){
	tpl = template.Must(tpl.ParseGlob("templates/*.html"))
	http.HandleFunc("/getform", getFormHandler)
	http.HandleFunc("/processget", processGetHandler)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

type Sub struct {
	Username string
	Num int
	MyFloat float64
	Updates bool
}

func getFormHandler(w http.ResponseWriter, r *http.Request){
	tpl.ExecuteTemplate(w, "getform.html", nil)
}

func processGetHandler(w http.ResponseWriter, r *http.Request){
	fmt.Println("processGetHandler running")
	var s Sub
	s.Username = r.FormValue("username")
	num := r.FormValue("numberName")
	s.Num, _ = strconv.Atoi(num)
	s.Num = s.Num * 2
	var err error

	s.MyFloat, err = strconv.ParseFloat(r.FormValue("myFltName"), 64)
	if err != nil {
		log.Fatal("error parsing float64")
	}
	if r.FormValue("upName") == "true" {
		s.Updates = true
	}else if r.FormValue("upName") == "false" {
		s.Updates = false
	}
	tpl.ExecuteTemplate(w, "thanks.html", s)
}