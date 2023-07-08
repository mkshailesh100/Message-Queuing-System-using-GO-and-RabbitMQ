package api_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/mkshailesh100/message-queue-system/internal/api"
)

func BenchmarkCreateProduct(b *testing.B) {
	// Create a new Gin router
	r := gin.Default()
	r.POST("/products", api.CreateProduct)

	// Perform the request in a loop
	for i := 0; i < b.N; i++ {
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
			b.Fatal(err)
		}

		// Perform the request
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		// Check the response status code
		if w.Code != http.StatusOK {
			b.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
		}
	}
}

func BenchmarkCreateProductParallel(b *testing.B) {
	// Create a new Gin router
	r := gin.Default()
	r.POST("/products", api.CreateProduct)

	// Run the benchmark in parallel
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
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
				b.Fatal(err)
			}

			// Perform the request
			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)

			// Check the response status code
			if w.Code != http.StatusOK {
				b.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
			}
		}
	})
}
