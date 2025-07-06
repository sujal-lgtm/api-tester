package tester

import (
	"log"
	"sync"
	"time"
	"testapi/internal/models"
)

type LoadResult struct {
	URL           string `json:"url"`
	TotalRequests int    `json:"total_requests"`
	Concurrency   int    `json:"concurrency"`
	AverageTimeMS int64  `json:"average_time_ms"`
	SuccessCount  int    `json:"success_count"`
	FailureCount  int    `json:"failure_count"`
}

func RunLoadTest(request models.TestRequest) (LoadResult, error) {
	var wg sync.WaitGroup

	times := make(chan int64, request.Requests)
	results := make(chan bool, request.Requests)
	con := make(chan struct{}, request.Concurrency) // semaphore

	for i := 0; i < request.Requests; i++ {
		wg.Add(1)
		con <- struct{}{} // acquire slot

		go func(i int) {
			defer func() {
				<-con // release slot
				wg.Done()
			}()

			start := time.Now()
			resp, err := SendRequest(request.Method, request.URL, request.Body, request.Headers)
			duration := time.Since(start).Milliseconds()
			times <- duration

			if err == nil && resp.StatusCode < 400 {
				results <- true
			} else {
				log.Println("Failed:", err)
				results <- false
			}
			if resp != nil {
				resp.Body.Close()
			}
			log.Println("Completed", i)
		}(i)
	}

	wg.Wait()
	close(times)
	close(results)

	var totalTime int64
	var successCount, failureCount int

	for t := range times {
		totalTime += t
	}

	for r := range results {
		if r {
			successCount++
		} else {
			failureCount++
		}
	}

	avg := int64(0)
	if request.Requests > 0 {
		avg = totalTime / int64(request.Requests)
	}

	result := LoadResult{
		URL:           request.URL,
		TotalRequests: request.Requests,
		Concurrency:   request.Concurrency,
		AverageTimeMS: avg,
		SuccessCount:  successCount,
		FailureCount:  failureCount,
	}

	return result, nil
}
