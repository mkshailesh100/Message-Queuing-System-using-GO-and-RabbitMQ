package compression

import (
	"image/jpeg"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/nfnt/resize"
)
func DownloadAndCompressImages(imageURLs[] string) ([]string, error) {
    var compressedPaths []string

    for _, imageURL := range imageURLs {
        // Download the image
		log.Println("Image URL is ",imageURL)
        response, err := http.Get(imageURL)
        if err != nil {
            log.Println("Failed to download image:", err)
            continue
        }
        defer response.Body.Close()
        timestamp := time.Now().Format("20060102150405.000000")
        imageName := timestamp+".jpeg"
        // Create a file to store the downloaded image
        file, err := os.Create("images/"+imageName)
        if err != nil {
            log.Println("Failed to create file:", err)
            continue
        }
        defer file.Close()

        // Copy the image data to the file
        _, err = io.Copy(file, response.Body)
        if err != nil {
            log.Println("Failed to write image data to file:", err)
            continue
        }

        // Open the downloaded image file
        img, err := os.Open("images/"+imageName)
        if err != nil {
            log.Println("Failed to open image file:", err)
            continue
        }
        defer img.Close()

        // Decode the image
        imgData, err := jpeg.Decode(img)
        if err != nil {
            log.Println("Failed to decode image:", err)
            // continue
        }

        // Compress the image
        compressedData := resize.Resize(200, 0, imgData, resize.Lanczos3)

		// Create a new file for the compressed image with timestamp
		compressedPath := filepath.Join("compressedImages", "compressed_"+imageName)
        compressedFile, err := os.Create(compressedPath)
        if err != nil {
            log.Println("Failed to create compressed image file:", err)
            continue
        }
        defer compressedFile.Close()

        // Write the compressed image data to the file
        err = jpeg.Encode(compressedFile, compressedData, nil)
        if err != nil {
            log.Println("Failed to write compressed image data to file:", err)
            continue
        }

        // Append the compressed image path to the list
        compressedPaths = append(compressedPaths, compressedPath)
    }

    return compressedPaths, nil
}
