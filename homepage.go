package main

import (
	"net/http"
	"html/template"
	"log"
	"os"
)

type PageParams struct {
    Title string
}

func handler(w http.ResponseWriter, r *http.Request) {
	//pp := &PageParams{"Шняга блеадь"}
    	t, _ := template.ParseFiles("index.html")
    	t.Execute(w, nil)
}

func main() {
	http.HandleFunc("/", handler)
	err := http.ListenAndServe(":80", http.FileServer(http.Dir("/home/zlodimir/gocode/src/github.com/zlodimir/homepage")))
	
    	if err != nil {
        	log.Println(err)
        	os.Exit(1)
	}
}



