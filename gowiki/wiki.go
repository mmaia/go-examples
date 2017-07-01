package main

import (
	"fmt"
	"io/ioutil"
)

type Page struct {
	Title string
	Body []byte
}


// Writes a page body to a file with the same name as the page title
// The method ioutil.WriteFile returns an error if one occurs and nil(null value) if
// all goes well, hence the error return type of this function.
// return nil, if all goes well and the body of the page is successfully written to disk
// return error, if something goes wrong
func (page *Page) save() error {
	filename := page.Title + ".txt"
	return ioutil.WriteFile(filename, page.Body, 0600)
}

