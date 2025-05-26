package main

import (
	"fmt"
	"os"

	"github.com/kale-swapnil/azure-blob-storage-golang/upload/"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Warning! .env file not found. Using default env vars.")
	}
	container := "golang"
	blobName := "demo"
	filePath := "./hello.txt"

	fmt.Println("Calling Blob function")

	err = upload.UploadBlob(container, blobName, filePath)
	if err != nil {
		fmt.Println("Upload failed:", err)
		os.Exit(1)
	}

	fmt.Println("Upload Successful")

}
