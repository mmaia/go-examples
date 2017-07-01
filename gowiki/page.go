package gowiki

import (
	"io/ioutil"
	"os"
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
func (page *Page) Save() error {
	pwd, _ := os.Getwd()
	filename := page.Title + ".txt"
	return ioutil.WriteFile(pwd + "/build/" + filename, page.Body, 0600)
}

// Returns a file from disk using the title parameter as the file name and returning it's content as a Page type.
// Functions can return multiple values. The standard library function io.ReadFile returns []byte and error.
func LoadPage(title string) (*Page, error) {
	pwd, _ := os.Getwd()
	filename := pwd + "/build/" + title + ".txt"
	body, err := ioutil.ReadFile(filename) // ReadFile returns []byte and error
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}
