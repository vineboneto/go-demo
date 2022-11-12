package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	requestId := make(chan int)
	concurrency := 400

	for i := 1; i <= concurrency; i++ {
		go worker(requestId, i)
	}

	for i := 0; i < 20*10000; i++ {
		requestId <- i
	}
}

func worker(requestId chan int, worker int) {

	for r := range requestId {
		res, err := http.Get("http://localhost:8585/products")

		if err != nil {
			log.Fatal("Error")
		}

		defer res.Body.Close()

		// content, _ := ioutil.ReadAll(res.Body)

		fmt.Printf("Worker %d. RequestId %d\n", worker, r)

		time.Sleep(time.Second * time.Duration(rand.Intn(5)))
	}
}
