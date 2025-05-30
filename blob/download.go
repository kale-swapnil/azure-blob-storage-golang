package blob

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
)

func DownloadBlob(containerName, blobName, destinationPath string) (string, error) {
	strConn := os.Getenv("AZURE_STORAGE_CONNECTION_STRING")
	if strConn == "" {
		return "", fmt.Errorf("missing Azure Connection string")
	}

	serviceClient, err := azblob.NewClientFromConnectionString(strConn, nil)
	if err != nil {
		return "", fmt.Errorf("failed to create service client: %w", err)
	}

	containerClient := serviceClient.ServiceClient().NewContainerClient(containerName)

	blobClient := containerClient.NewBlobClient(blobName)

	downloadResp, err := blobClient.DownloadStream(context.Background(), nil)
	if err != nil {
		return "", fmt.Errorf("failed to get blob: %w", err)
	}

	localFile, err := os.Create(destinationPath)
	if err != nil {
		return "", fmt.Errorf("failed to get destination path: %w", err)
	}

	io.Copy(localFile, downloadResp.Body)

	defer downloadResp.Body.Close()
	defer localFile.Close()

	return "Download successful", nil

}
