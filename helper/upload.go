package helper

import (
    "io"
    "mime/multipart"
    "os"
    "path/filepath"
    "fmt"
)

// UploadFile is a helper function to upload a file and save it to the given directory
func UploadFile(file multipart.File, header *multipart.FileHeader, destinationDir string) (string, error) {
    // Ensure the directory exists, create it if not
    err := os.MkdirAll(destinationDir, os.ModePerm)
    if err != nil {
        return "", fmt.Errorf("failed to create directory: %v", err)
    }

    // Create a unique file name to avoid overwriting
    filePath := filepath.Join(destinationDir, header.Filename)

    // Create the destination file
    outFile, err := os.Create(filePath)
    if err != nil {
        return "", fmt.Errorf("failed to create file: %v", err)
    }
    defer outFile.Close()

    // Copy the file content to the destination file
    _, err = io.Copy(outFile, file)
    if err != nil {
        return "", fmt.Errorf("failed to save file: %v", err)
    }

    return filePath, nil
}
