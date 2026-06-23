package main

import (
	"net/http"
	"html/template"
)

func TemplateRender(w http.ResponseWriter, r *http.Request){
	const tmplStr = `
<!DOCTYPE html>
<html>
<head><title>{{.Title}}</title></head>
<body>
  <h1>{{.Title}}</h1>
  <p>{{.Body}}</p>
</body>
</html>
`

type PageData struct {
	Title string
	Body  string
}

tmpl := template.Must(template.New("page").Parse(tmplStr))

title :=  r.URL.Query().Get("title")
body  :=  r.URL.Query().Get("body")
	if title == "" || body == "" {
		http.Error(w, "title and body are required", 400)
		return
	}
	w.Header().Set("Content-Type", "text/html")
	if err := tmpl.Execute(w, PageData{Title: title, Body: body}); err != nil{
		http.Error(w, "template execution failed", 500)
		return
	}

}