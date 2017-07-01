package main

import (
	"net/http"
	"fmt"
	"github.com/mmaia/go-playground/gowiki"
)

func handler(responseWriter http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(responseWriter, "<p>Hi there, you're at the root, unfortunately there's nothing here, <a href=/view/TestPage> try this link instead. </a></p>")
}

func viewHandler(responseWriter http.ResponseWriter, request *http.Request) {
	title := request.URL.Path[len("/view/"):]
	p, _ := gowiki.LoadPage(title)
	fmt.Fprintf(responseWriter, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/view/", viewHandler)
	http.ListenAndServe(":8080", nil)
}