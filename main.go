package main

import (
	"fmt"

	"github.com/kale-swapnil/azure-blob-storage-golang/blob"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Warning! .env file not found. Using default env vars.")
	}
	container := "golang"
	blobName := "demo2"
	filePath := "./hello.txt"
	destinationPath := "blob/hello.txt"

	//fmt.Println("Calling Blob function")

	var msg string

	msg, err = blob.UploadBlob(container, blobName, filePath)
	if err != nil {
		fmt.Println("Upload failed:", err)
		//os.Exit(1)
	} else {
		fmt.Println(msg)
	}

	blobs, err := blob.ListBlob(container)
	if err != nil {
		fmt.Println("Failed to list blobs:", err)
	}

	for _, blobItem := range blobs {
		fmt.Println(blobItem)

	}

	msg, err = blob.DownloadBlob(container, blobName, destinationPath)
	if err != nil {
		fmt.Println("Download failed:", err)
	} else {
		fmt.Println(msg)
	}

	msg, err = blob.DeleteBlob(container, blobName)
	if err != nil {
		fmt.Println("Delete failed:", err)
	} else {
		fmt.Println(msg)
	}

}
