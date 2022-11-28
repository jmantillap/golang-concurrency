package databooks

import (
	"fmt"
	"sync"
)

type Book struct {
	ID       int
	Title    string
	Finished bool
}

var books = []*Book{
	{0, "libro titulo 1", false},
	{1, "libro titulo 2", false},
	{2, "libro titulo 3", false},
	{3, "libro titulo 4", false},
	{4, "libro titulo 5", false},
	{5, "libro titulo 6", false},
	{6, "libro titulo 7", false},
	{7, "libro titulo 8", false},
	{8, "libro titulo 9", false},
	{9, "libro titulo 10", false},
}

func findBook(id int, m *sync.RWMutex) (int, *Book) {

	index := -1
	var book *Book
	m.RLock()
	for i, b := range books {
		if b.ID == id {
			index = i
			book = b
		}
	}
	m.RUnlock()
	return index, book
}

func FinishBook(id int, m *sync.RWMutex) {
	i, book := findBook(id, m)
	if i < 0 {
		return
	}
	m.Lock()
	book.Finished = true
	books[id] = book
	m.Unlock()
	fmt.Printf("Finished Book: %s\n", book.Title)

}
