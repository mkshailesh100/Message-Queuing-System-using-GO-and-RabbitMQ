package api_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/mkshailesh100/message-queue-system/internal/api"
)

func TestCreateProduct(t *testing.T) {
	// Create a new Gin router
	r := gin.Default()
	r.POST("/products", api.CreateProduct)

	// Create a test request
	payload := `{
		"user_id": 1,
		"product_name": "Sample Product",
		"product_description": "This is a sample product",
		"product_images": "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcSB-Sl-zoom2cnZQvm8yKk409Pg0ts0gZ7pxEAvOL38oQ&s,https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTNtCh17cCUl3OeiiqnqYb72OPfHLLRVte3sg5Lz5duGg&s",
		"product_price": 9.99
	}`
	req, err := http.NewRequest("POST", "/products", strings.NewReader(payload))
	if err != nil {
		t.Fatal(err)
	}

	// Perform the request
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Check the response status code
	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}

	// Check the response body
	expectedResponse := `{"message":"Product created successfully"}`
	if w.Body.String() != expectedResponse {
		t.Errorf("Expected response body %s, got %s", expectedResponse, w.Body.String())
	}
}
