package pool

import (
	"fmt"
	"myhttp/url"
	"sync"
)

// worker reads the input from inputs chan, applies function f() on them then, sends the result to result chan
func worker(wg *sync.WaitGroup, f func(string) url.Result, results chan url.Result, inputs chan string) {
	for input := range inputs {
		results <- f(input)
	}
	wg.Done()
}

// CreateWorkerPool creates the workers pool
func CreateWorkerPool(workersNum int, f func(string) url.Result, results chan url.Result, inputs chan string) {
	var wg sync.WaitGroup
	for i := 0; i < workersNum; i++ {
		wg.Add(1)
		go worker(&wg, f, results, inputs)
	}
	wg.Wait()
	close(results)
}

// Publish send the urls to the workers pool
func Publish(urls []string, inputs chan string) {
	for _, url := range urls {
		inputs <- url
	}
	close(inputs)
}

// Result prints the url response
func Result(done chan bool, results chan url.Result) {
	for result := range results {
		if result.Err != nil {
			fmt.Printf("%s has an Error: %s \n", result.Address, result.Err.Error())
		} else {
			fmt.Printf("%s  %s \n", result.Address, result.Hash)
		}
	}
	done <- true
}
