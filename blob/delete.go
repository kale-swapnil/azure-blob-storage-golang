package blob

import (
	"context"
	"fmt"
	"os"

	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
)

func DeleteBlob(containerName, blobName string) (string, error) {
	connStr := os.Getenv("AZURE_STORAGE_CONNECTION_STRING")
	if connStr == "" {
		return "", fmt.Errorf("missing Azure connection string")
	}

	serviceClient, err := azblob.NewClientFromConnectionString(connStr, nil)
	if err != nil {
		return "", fmt.Errorf("failed to create service client: %w", err)
	}

	containerClient := serviceClient.ServiceClient().NewContainerClient(containerName)

	blobClient := containerClient.NewBlobClient(blobName)

	blobClient.Delete(context.Background(), nil)

	return "Delete Successful", nil
}
