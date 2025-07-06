package models

type TestRequest struct {
    URL         string            `json:"url"`
    Type        string            `json:"type"`
    Method      string            `json:"method,omitempty"`
    Headers     map[string]string `json:"headers,omitempty"`
    Body        string            `json:"body,omitempty"`
    Requests    int               `json:"requests,omitempty"`
    Concurrency int               `json:"concurrency,omitempty"`
}
