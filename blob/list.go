package blob

import (
	"context"
	"fmt"
	"os"

	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
)

func ListBlob(containerName string) ([]string, error) {
	connStr := os.Getenv("AZURE_STORAGE_CONNECTION_STRING")
	if connStr == "" {
		return nil, fmt.Errorf("connection string is empty")

	}

	serviceClient, err := azblob.NewClientFromConnectionString(connStr, nil)

	if err != nil {
		return nil, fmt.Errorf("failed to create service client: %w", err)
	}

	containerClient := serviceClient.ServiceClient().NewContainerClient(containerName)

	ctx := context.Background()
	pager := containerClient.NewListBlobsFlatPager(nil)

	var blobNames []string

	for pager.More() {
		pageResp, err := pager.NextPage(ctx)
		if err != nil {
			return nil, fmt.Errorf("error getting blobs list: %w", err)
		}

		for _, blobItem := range pageResp.Segment.BlobItems {
			blobNames = append(blobNames, *blobItem.Name)
		}
	}

	return blobNames, nil
}
