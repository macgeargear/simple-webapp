package main

import (
	"fmt"
	"net/http"
	"os"
)

type Page struct {
	Title string
	Body  []byte // byte slice
}

// save
func (p *Page) save() error {
	filename := p.Title + ".txt"
	return os.WriteFile(filename, p.Body, 0600) // WriteFile return error value
}

// load
func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := os.ReadFile(filename) // ReadFile return []byte error
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	p, _ := loadPage(title)
	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
}

func main() {
	// p1 := &Page{Title: "TestPage", Body: []byte("This is a simple page.")}
	// p1.save()
	// p2, _ := loadPage("TestPage")
	// fmt.Println(string(p2.Body))
	http.HandleFunc("/view/", viewHandler)
	http.ListenAndServe(":8080", nil)
}
