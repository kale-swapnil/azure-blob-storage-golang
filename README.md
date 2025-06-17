# Azure Blob Storage in Golang 🚀

This code demonstrates how to use **Azure Blob Storage** in a Go application — including uploading, listing, downloading, and deleting blobs from a container. It's built for developers who want a minimal, clean, and functional example of how to work with Azure Storage in Go.

---

## 📦 Features

- ✅ Upload a file to Azure Blob Storage
- 📄 List all blobs in a container
- 📥 Download a blob to your local machine
- 🗑️ Delete a blob from the container
---

## 🛠️ Prerequisites

- Go 1.18 or later
- Azure Storage account
- Blob container created (e.g. `golang`)
- Azure Storage connection string

---

## 🔐 Setup

1. Clone the repo:

```bash
git clone https://github.com/kale-swapnil/azure-blob-storage-golang.git
cd azure-blob-storage-golang
```
2. Update the .env file with the Azure Storage Connection string:
3. Make sure you have a blob container already created in Azure (or create one named golang).
4. Run Go from the terminal
```
run main.go
```

| File           | Description                                                 |
| -------------- | ----------------------------------------------------------- |
| `main.go`      | Entry point. Executes upload, list, download, and delete    |
| `upload.go`    | Contains `UploadBlob()` for uploading files to Azure        |
| `list.go`      | Contains `ListBlob()` to list all files in a blob container |
| `download.go`  | Contains `DownloadBlob()` to download a file from Azure     |
| `delete.go`    | Contains `DeleteBlob()` to remove a file from the container |
| `.env`         | File for environment variables                          |




