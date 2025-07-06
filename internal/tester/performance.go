package tester

import (
	"time"
	"testapi/internal/models"
)

type PerfResult struct {
	URL          string `json:"url"`
	ResponseTime int64  `json:"response_time_ms"`
	StatusCode   int    `json:"status_code"`
	Success      bool   `json:"success"`
}

func RunPerformanceTest(request models.TestRequest) (PerfResult, error) {
	start := time.Now()
	resp, err := SendRequest(request.Method, request.URL, request.Body, request.Headers)
	duration := time.Since(start)

	result := PerfResult{
		URL:          request.URL,
		ResponseTime: duration.Milliseconds(),
		Success:      err == nil && resp.StatusCode < 400,
	}

	if resp != nil {
		result.StatusCode = resp.StatusCode
		defer resp.Body.Close()
	}

	return result, nil
}
