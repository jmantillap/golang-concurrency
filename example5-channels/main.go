package main

import (
	"fmt"
	"sync"

	"github.com/google/uuid"
)

//COMUNICACION EN GOROUTINES

//bidireccionales --> IdsChan := make(chan string)
//receive-only
//send-only

func main() {

	wg := &sync.WaitGroup{}

	IdsChan := make(chan string)
	FakeIdsChan := make(chan string)
	ClosedChans := make(chan int)

	wg.Add(3)

	go generateIDS(wg, IdsChan, ClosedChans)
	go generateFakeIds(wg, FakeIdsChan, ClosedChans)
	go logIds(wg, IdsChan, FakeIdsChan, ClosedChans)

	wg.Wait()
}

func generateFakeIds(wg *sync.WaitGroup, fakeIdsChan chan<- string, closedChans chan<- int) {
	for i := 0; i < 50; i++ {
		id := uuid.New()
		fakeIdsChan <- fmt.Sprintf("%d. %s", i+1, id.String()) // El valor del id se lo va a enviar al chanel
	}
	close(fakeIdsChan)
	closedChans <- 1
	wg.Done()
}

func generateIDS(wg *sync.WaitGroup, idsChan chan<- string, closedChans chan<- int) { //idsChan chan <- string : Channel de recibir datos --> flecha a la izquierda

	//id := uuid.New()
	//idsChan <- id.String() // El valor del id se lo va a enviar al chanel
	for i := 0; i < 100; i++ {
		id := uuid.New()
		//idsChan <- id.String() // El valor del id se lo va a enviar al chanel
		idsChan <- fmt.Sprintf("%d. %s", i+1, id.String()) // El valor del id se lo va a enviar al chanel
	}
	close(idsChan)
	closedChans <- 1
	wg.Done()
}

func logIds(wg *sync.WaitGroup, idsChan <-chan string, fakeIdsChan <-chan string, closedChans chan int) { // idsChan <-chan string : Chanel de enviar datos con la flechita a la derecha
	//id := <-idsChan // El valor que viene del channe se los pasamos a la variable id
	//fmt.Println(id)

	//for id := range idsChan {
	//	fmt.Println(id)
	//}
	closeCounter := 0

	for {
		select {
		case id, ok := <-idsChan:
			if ok {
				fmt.Println("IDS:", id)
			}
		case id, ok := <-fakeIdsChan:
			if ok {
				fmt.Println("Fake ID", id)
			}
		case count, ok := <-closedChans:
			if ok {
				closeCounter += count
			}
		}
		if closeCounter == 2 {
			close(closedChans)
			break
		}
	}

	wg.Done()

}
