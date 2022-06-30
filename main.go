package main

import (
	"fmt"
	"net/http"
)

type Book struct {
	ID    int
	Title string
}

var book []Book

func getid(w http.ResponseWriter, r *http.Request) {
	for i := 0; i < len(book); i++ {
		fmt.Fprintf(w, "%d - %s\n", book[i].ID, book[i].Title)
	}
}

func Read_db() {
	book = append(book, Book{1, "Tom"})
	book = append(book, Book{2, "gfg"})
}

func main() {
	Read_db()

	http.HandleFunc("/", getid)
	http.ListenAndServe(":80", nil)
}
