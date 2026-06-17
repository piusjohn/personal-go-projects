package main

import (
	//"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("http://webcode.me")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	///, err := ioutil.ReadAll(resp.Body)
	//instead of ReadAll wwe can use io.copy
	_, err = io.Copy(os.Stdout, resp.Body)

	if err != nil {
		log.Fatal(err)
	}
//fmt.Println(string(body))
}
