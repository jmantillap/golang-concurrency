package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// se agrega waitgroup para esperar que termine mejor si en tiempo de espera
func main() {

	start := time.Now()

	wg := &sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go ShowGouroutine(i, wg)
	}

	wg.Wait()
	duration := time.Since(start).Milliseconds()
	fmt.Printf("Process Duration %dms\n", duration)
}

func ShowGouroutine(id int, wg *sync.WaitGroup) {

	delay := rand.Intn(500)
	fmt.Printf("Gouroutine #%d with %dms  \n", id, delay)
	time.Sleep(time.Millisecond * time.Duration(delay))
	wg.Done()
}
