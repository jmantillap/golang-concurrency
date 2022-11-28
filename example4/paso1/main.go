package main

import (
	"fmt"
	"math/rand"
	"time"
)

//secuencial sin gorutines
func main() {
	for i := 0; i < 10; i++ {
		ShowGouroutine(i)
	}

}

func ShowGouroutine(id int) {

	delay := rand.Intn(500)

	fmt.Printf("Gouroutine #%d with %dms  \n", id, delay)

	time.Sleep(time.Millisecond * time.Duration(delay))
}
