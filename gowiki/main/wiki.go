package main

import (
	"fmt"
	"github.com/mmaia/go-playground/gowiki"
)

// create Page type instance save it to disk then loads it and print it's contents to console.
func main() {
	p1 := &gowiki.Page{Title: "TestPage", Body: []byte("This is a sample Page.")}
	p1.Save()
	p2, _ := gowiki.LoadPage("TestPage")
	fmt.Println(string(p2.Body))
}

