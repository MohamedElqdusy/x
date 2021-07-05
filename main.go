package main

import (
	"flag"
	"myhttp/pool"
	"myhttp/url"
	"os"
)

const (
	defaultParallelRequests = 10
)

func main() {
	parallel := flag.Int("parallel", defaultParallelRequests, "an int")
	// starting from index 2 to skip the program name and its parallel flag
	urls := os.Args[2:]

	inputs := make(chan string, *parallel)
	results := make(chan url.Result, len(urls))

	// deliver the urls to the workers
	go pool.Publish(urls, inputs)

	// used to signal the result printing has been done
	done := make(chan bool)
	// open a goroutine to receive the results
	go pool.Result(done, results)
	// start the work of visit url
	pool.CreateWorkerPool(*parallel, url.VisitURL, results, inputs)

	// waiting untill the results prinitng is done
	<-done
}
