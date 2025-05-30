package blob

import (
	"context"
	"fmt"
	"os"

	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
)

func UploadBlob(containerName, blobName, filePath string) (string, error) {
	connStr := os.Getenv("AZURE_STORAGE_CONNECTION_STRING")

	if connStr == "" {
		return "", fmt.Errorf("missing Azure connection string")
	}

	serviceClient, err := azblob.NewClientFromConnectionString(connStr, nil)
	if err != nil {
		return "", fmt.Errorf("failed to create service client: %w", err)
	}

	containerClient := serviceClient.ServiceClient().NewContainerClient(containerName)

	file, err := os.Open(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	blobClient := containerClient.NewBlockBlobClient(blobName)

	_, err = blobClient.UploadStream(context.Background(), file, nil)
	if err != nil {
		return "", fmt.Errorf("failed to upload to blob: %w", err)
	}

	//fmt.Println("Upload successful")

	return "Upload successful", nil

}
