package repository

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
)

type TodoClient interface {
	GetMessage(number int) (*Message, error)
}

type Message struct {
	UserID    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

type todo struct {
	client   *http.Client
	basePath string
}

func New(client *http.Client, basePath string) TodoClient {
	return &todo{
		basePath: basePath,
		client:   client,
	}
}

func (t *todo) GetMessage(number int) (*Message, error) {
	fullPath, err := t.getFullPath(number)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest(http.MethodGet, fullPath, nil)
	if err != nil {
		return nil, err
	}

	response, err := t.client.Do(request)
	if err != nil {
		return nil, err
	}

	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			// ideally we want to log this instead
			fmt.Println("failed to close response body")
		}
	}(response.Body)

	// check we have a valid response
	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("request failed with status code %d", response.StatusCode)
	}

	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	// decode response message
	rMessage := new(Message)
	if err = json.Unmarshal(bodyBytes, rMessage); err != nil {
		return nil, err
	}

	return rMessage, nil
}

func (t *todo) getFullPath(number int) (string, error) {
	fullPath, err := url.JoinPath(t.basePath, strconv.Itoa(number))
	if err != nil {
		return "", err
	}

	return fullPath, nil
}
