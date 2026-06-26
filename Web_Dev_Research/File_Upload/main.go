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
	http.HandleFunc("/", uploadFile)
	//http.HandleFunc("/", homePageHandler)
	myDir := http.Dir("./uploads")
	filehandler := http.FileServer(myDir)
	handle := http.StripPrefix("/downloads", filehandler)

	http.Handle("/downloads/", handle)
	http.ListenAndServe(":8080", nil)
}


func uploadFile(w http.ResponseWriter, r *http.Request){

	if r.Method == "GET"{
		tpl.ExecuteTemplate(w, "upload.html", nil)
		return
	}

	file, Fileheader, err := r.FormFile("myFile")
	if err != nil {
		http.Error(w, "retrieving file failed", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	contentType := Fileheader.Header["Content-Type"][0]
	fmt.Println("Content Type:", contentType)

	var osFile *os.File

	if contentType == "image/jpeg" {
		osFile, err = os.CreateTemp("images", "*.jpg")
	} else if contentType == "application/pdf"{
		osFile, err = os.CreateTemp("PDFs", "*pdf")
	} else if contentType == "text/javascript" {
		osFile, err = os.CreateTemp("js", "*.js")
	}
	fmt.Println("error:", err)
	defer osFile.Close()

	// dst, err := os.Create("./uploads/" + Fileheader.Filename)
	// if err != nil {
	// 	http.Error(w, "creating file failed", http.StatusInternalServerError)
	// 	return
	// }

	

	io.Copy(osFile, file)

	fmt.Println("file uploaded successfully")
	http.Redirect(w, r, "/downloads/", http.StatusSeeOther)

}


