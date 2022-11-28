package main

import (
	"fmt"
	"math/rand"
	"time"
)

// se agrega la palabra go y un tiempo de espera para que se puedan ejecutar todas las goroutines
func main() {
	for i := 0; i < 10; i++ {
		go ShowGouroutine(i)
	}
	time.Sleep(time.Minute)
}

func ShowGouroutine(id int) {

	delay := rand.Intn(500)

	fmt.Printf("Gouroutine #%d with %dms  \n", id, delay)

	time.Sleep(time.Millisecond * time.Duration(delay))
}
