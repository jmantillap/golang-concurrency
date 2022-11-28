package main

import (
	data "example4/paso4/databooks"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// se Ajusta para que ho existan bloqueos al leer y escribir propiedades de algunos objetos
//go run --race main.go ;  Advertencia entonces es de lectura  y ecritura
func main() {

	start := time.Now()
	wg := &sync.WaitGroup{}
	m := &sync.RWMutex{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go reedBook(i, wg, m)
	}

	wg.Wait()
	duration := time.Since(start).Milliseconds()
	fmt.Printf("Process Duration %dms\n", duration)
}

func reedBook(id int, wg *sync.WaitGroup, m *sync.RWMutex) {

	data.FinishBook(id, m)

	delay := rand.Intn(800)
	//fmt.Printf("Gouroutine #%d with %dms  \n", id, delay)
	time.Sleep(time.Millisecond * time.Duration(delay))

	wg.Done()
}
