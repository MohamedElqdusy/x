package pool

import (
	"myhttp/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWorkers(t *testing.T) {
	urls := []string{"adjust.com", "google.com", "facebook.com", "yahoo.com", "yandex.com", "twitter.com", "reddit.com/r/funny"}

	inputs := make(chan string, 3)
	results := make(chan url.Result, len(urls))

	// deliver the urls to the workers
	go Publish(urls, inputs)

	// used to signal the result printing has been done
	done := make(chan bool)
	// open a goroutine to receive the results
	go resultMock(t, done, results, len(urls))
	// start the work of visit url
	CreateWorkerPool(3, visitURLMock, results, inputs)

	// waiting untill the results prinitng is done
	<-done
}

func visitURLMock(s string) url.Result {
	return url.Result{
		Address: s,
		Hash:    s,
	}
}

func resultMock(t *testing.T, done chan bool, results chan url.Result, ResultCounter int) {
	counter := 0
	for result := range results {
		counter++
		assert.Equal(t, result.Address, result.Hash)
	}
	assert.Equal(t, ResultCounter, counter)
	done <- true
}
