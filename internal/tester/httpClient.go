package tester

import (
    "net/http"
    "strings"
    "time"
)

func SendRequest(method, url, body string, headers map[string]string) (*http.Response, error) {

	if method == "" {
    	method = "GET"
	}

    var bodyReader *strings.Reader
    if body != "" {
        bodyReader = strings.NewReader(body)
    } else {
        bodyReader = strings.NewReader("")
    }

    req, err := http.NewRequest(strings.ToUpper(method), url, bodyReader)
    if err != nil {
        return nil, err
    }

    for k, v := range headers {
        req.Header.Set(k, v)
    }

    client := &http.Client{
        Timeout: 10 * time.Second,
    }

    return client.Do(req)
}
