package main

import (
	"bufio"
	"fmt"
	"os/exec"
	"regexp"
	"strings"
	"sync"
)

func main() {
	NumEl := 100 // Number of times the external program is called
	NumCore := 4 // Number of available cores
	c := make(chan bool, NumCore-1)
	wg := new(sync.WaitGroup)
	wg.Add(NumEl)
	for i := 0; i < NumEl; i++ {
		go callProg(i, c, wg)
		c <- true
	}
	wg.Wait() // Wait for all the children to die
	close(c)

}

func callProg(i int, c chan bool, wg *sync.WaitGroup) {
	defer func() {
		<-c
		wg.Done() // Decrease the number of alive goroutines
	}()
	args := "-n 5 127.0.0.1"
	cmd := exec.Command("ping", strings.Split(args, " ")...)
	output, _ := cmd.StdoutPipe()
	cmd.Start()

	scanner := bufio.NewScanner(output)
	for scanner.Scan() {
		m := scanner.Text()

		matchPackets, _ := regexp.MatchString("Packets", m)
		matchMinimum, _ := regexp.MatchString("Minimum", m)
		fmt.Printf("-------------RESULTADOS %v-------------------\n", i)
		if matchPackets {
			fmt.Println("Ping statistics for 127.0.0.1 para la Solicitud N ", i, "-->", m)
			//fmt.Println(m)
		}

		if matchMinimum {
			fmt.Println("Approximate round trip times in milli-seconds para la Solicitud N ", i, "-seg->", m)
			fmt.Println(m)
		}
		fmt.Printf("-------------FIN RESULTADOS %v---------------\n", i)
	}
	cmd.Wait()
}
