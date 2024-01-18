package executor

import (
	"cron-request/internal/config"
	"github.com/google/logger"
	"io"
	_ "io"
	"net/http"
	"slices"
)

type IExecutor interface {
	Execute() (string, *error)
}

type Executor struct {
	config config.RequestConfig
}

func New(config config.RequestConfig) IExecutor {
	return &Executor{config: config}
}

// Execute executes the request and returns the response body and an error
func (e *Executor) Execute() (string, *error) {
	// create a new request
	req, err := http.NewRequest(e.config.Method, e.config.Url, nil)
	if err != nil {
		logger.Error("Error creating request: ", err)
		return "", &err
	}

	// add headers to the request
	for key, value := range e.config.Headers {
		req.Header.Add(key, value)
	}

	// execute the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		logger.Error("Error executing request: ", err)
		return "", &err
	}

	// read the response body
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.Error("Error reading response body: ", err)
		return "", &err
	}

	successStatusCodes := []int{http.StatusOK, http.StatusCreated}
	if !slices.Contains(successStatusCodes, resp.StatusCode) {
		logger.Error("Error executing request, Status: ", resp.StatusCode, " response: ", string(data))
		return "", &err
	}

	return string(data), nil
}
