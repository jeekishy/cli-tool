package repository

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestTodo_GetTodo(t *testing.T) {
	mockMessage := &Message{
		UserID:    2,
		ID:        1,
		Title:     "My Mock Tile",
		Completed: false,
	}
	mockResponse := getMockHttpResponse(mockMessage)
	mockHttpClient := http.Client{Transport: &mockRoundTripper{mockResponse}}

	mockTodo := &todo{
		client:   &mockHttpClient,
		basePath: "https://mock",
	}

	m, err := mockTodo.GetMessage(5)
	assert.NoError(t, err)
	assert.Equal(t, mockMessage, m, "unexpected response")
}

func TestTodo_GetTodoInValid(t *testing.T) {
	mockResponse := new(http.Response)
	mockHttpClient := http.Client{Transport: &mockInvalidRoundTripper{mockResponse}}

	mockTodo := &todo{
		client:   &mockHttpClient,
		basePath: "https://mock",
	}

	m, err := mockTodo.GetMessage(2)
	assert.Equal(t, "request failed with status code 400", err.Error(), "unexpected error message")
	assert.Nil(t, m, "unexpected response")
}

func TestTodo_getFullPath(t *testing.T) {
	mockBasePath := "https://mock.com"
	mockPath := 5
	td := &todo{basePath: mockBasePath}

	// get full path
	fp, err := td.getFullPath(mockPath)

	assert.NoError(t, err, "error has been received")
	assert.Equal(t, fmt.Sprintf("%s/%d", mockBasePath, mockPath), fp)

}

func TestTodo_getFullPathWithErr(t *testing.T) {
	mockBasePath := "a%20b%1"
	mockPath := 5
	td := &todo{basePath: mockBasePath}

	// get full path
	fp, err := td.getFullPath(mockPath)

	assert.Error(t, err, "error has been received")
	assert.Equal(t, "", fp)
}

func getMockHttpResponse(message *Message) *http.Response {
	mockResponse, err := json.Marshal(message)
	if err != nil {
		return nil
	}

	recorder := httptest.NewRecorder()
	recorder.Header().Add("Content-Type", "application/json")

	_, err = recorder.Write(mockResponse)
	if err != nil {
		return nil
	}

	return recorder.Result()
}

// ------------------------------
// Mock Round trippers
type mockRoundTripper struct {
	response *http.Response
}

func (mt *mockRoundTripper) RoundTrip(request *http.Request) (*http.Response, error) {
	r := mt.response
	r.Status = http.StatusText(http.StatusOK)
	r.StatusCode = http.StatusOK

	return r, nil
}

type mockInvalidRoundTripper struct {
	response *http.Response
}

func (mit *mockInvalidRoundTripper) RoundTrip(request *http.Request) (*http.Response, error) {
	r := mit.response
	r.Status = http.StatusText(http.StatusBadRequest)
	r.StatusCode = http.StatusBadRequest

	return r, nil
}
