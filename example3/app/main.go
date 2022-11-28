package main

import (
	"log"
	"net/http"

	limiter "github.com/jmantillap/limiter/services"
)

func main() {

	limiter := limiter.NewConcurrencyLimiter(10)

	httpGoogle := int(0)
	httpApple := int(0)
	for i := 0; i < 200; i++ {

		//httpGoogle := int(0)
		limiter.Execute(func() {
			resp, _ := http.Get("https://www.google.com/")
			defer resp.Body.Close()
			httpGoogle = resp.StatusCode
			log.Println("httpGoogle: ", i, httpGoogle)
		})
		//httpApple := int(0)

		limiter.Execute(func() {
			resp, _ := http.Get("https://www.apple.com/")
			//Expect(err).To(BeNil())
			defer resp.Body.Close()
			httpApple = resp.StatusCode
			log.Println("httpApple:", i, httpApple)
		})
	}
	limiter.WaitAndClose()

}
