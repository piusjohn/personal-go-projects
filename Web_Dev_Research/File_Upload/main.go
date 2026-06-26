package main

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
)

var tpl *template.Template

func main() {
	tpl, _ = tpl.ParseGlob("templates/*.html")
	http.HandleFunc("/upload", uploadFile)
	//http.HandleFunc("/", homePageHandler)
	http.HandleFunc("/", Home)
	myDir := http.Dir("./uploads")
	filehandler := http.FileServer(myDir)
	handle := http.StripPrefix("/downloads", filehandler)

	http.Handle("/downloads/", handle)
	http.ListenAndServe(":8080", nil)
}


func Home(w http.ResponseWriter, r *http.Request){
	tpl.ExecuteTemplate(w, "upload.html", nil)
}


func uploadFile(w http.ResponseWriter, r *http.Request){
	file, header, err := r.FormFile("myFile")
	if err != nil {
		http.Error(w, "retrieving file failed", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	dst, err := os.Create("./uploads/" + header.Filename)
	if err != nil {
		http.Error(w, "creating file failed", http.StatusInternalServerError)
		return
	}

	defer dst.Close()

	io.Copy(dst, file)

	fmt.Println("file uploaded successfully")
	http.Redirect(w, r, "/downloads", http.StatusSeeOther)

}


