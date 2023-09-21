package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockBook struct {
	mock.Mock
}

func (m *MockBook) getAllBooks() []Book {
	args := m.Called()
	return args.Get(0).([]Book)
}

func TestGetAllBooks(t *testing.T) {

	mockBook := new(MockBook)

	mockBook.On("getAllBooks").Return([]Book{
		{ID: 1, Title: "Kitap 1", Author: "Yazar 1"},
		{ID: 2, Title: "Kitap 2", Author: "Yazar 2"},
	})

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		books := mockBook.getAllBooks()
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(books)
	})

	req := httptest.NewRequest("GET", "/getAll", nil)
	w := httptest.NewRecorder()

	handler.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response []Book
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.Nil(t, err)

	expectedResponse := []Book{
		{ID: 1, Title: "Kitap 1", Author: "Yazar 1"},
		{ID: 2, Title: "Kitap 2", Author: "Yazar 2"},
	}

	assert.Equal(t, expectedResponse, response)

	mockBook.AssertExpectations(t)
}
