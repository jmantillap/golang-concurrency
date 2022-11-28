package databooks

import "fmt"

type Book struct {
	ID       int
	Title    string
	Finished bool
}

var books = []*Book{
	{1, "libro titulo 1", false},
	{2, "libro titulo 2", false},
	{3, "libro titulo 3", false},
	{4, "libro titulo 4", false},
	{5, "libro titulo 5", false},
	{6, "libro titulo 6", false},
	{7, "libro titulo 7", false},
	{8, "libro titulo 8", false},
	{9, "libro titulo 9", false},
	{10, "libro titulo 10", false},
}

func findBook(id int) (int, *Book) {
	index := -1
	var book *Book
	for i, b := range books {
		if b.ID == id {
			index = i
			book = b
		}
	}

	return index, book
}

func FinishBook(id int) {
	i, book := findBook(id)
	if i < 0 {
		return
	}
	book.Finished = true
	books[id] = book
	fmt.Printf("Finished Book: %s\n", book.Title)
}
