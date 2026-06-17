package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

  mux := http.NewServeMux()
  mux.HandleFunc("/", handleRoot)
  mux.HandleFunc("/blogs", handleBlogs)
  mux.HandleFunc("/profile", handleProfile)
  
  err := http.ListenAndServe(":4001", mux)
  if err != nil{
	log.Fatal(err)
  }

}

func handleRoot(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "welcome to the root page")
}

func handleBlogs(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "welcome to the blog page")
}

func handleProfile(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "welcome to the profile page")
}

//https://webreference.com/go/web/servers/