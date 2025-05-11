package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	// "github.com/joho/godotenv"
	"github.com/kollzey539/hash-store/handler"
	"github.com/kollzey539/hash-store/storage"

	"github.com/gorilla/mux"
)

func main() {
	// Load environment variables from .env file
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// }

	// Read environment variables
	accessKey := os.Getenv("AWS_ACCESS_KEY_ID")
	secretKey := os.Getenv("AWS_SECRET_ACCESS_KEY")
	region := os.Getenv("AWS_REGION")
	bucketName := os.Getenv("S3_BUCKET_NAME")

	// Initialize S3 storage
	s3Storage, err := storage.NewS3Storage(accessKey, secretKey, region, bucketName)
	if err != nil {
		log.Fatal("Error initializing S3 storage:", err)
	}

	r := mux.NewRouter()

	// Define the routes
	r.HandleFunc("/hash", handler.CreateHashHandler(s3Storage)).Methods("POST")
	r.HandleFunc("/hash/{hash}", handler.GetStringHandler(s3Storage)).Methods("GET")

	fmt.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
