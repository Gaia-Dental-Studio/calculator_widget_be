package helper

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"
)

func UploadFile(file multipart.File, header *multipart.FileHeader, destinationDir string) (string, error) {
	err := os.MkdirAll(destinationDir, os.ModePerm)
	if err != nil {
		return "", fmt.Errorf("failed to create directory: %v", err)
	}

	ext := filepath.Ext(header.Filename)
	uniqueName := fmt.Sprintf("%s%s", time.Now().Format("20060102150405"), ext)
	filePath := filepath.Join(destinationDir,uniqueName)
	
    outFile, err := os.Create(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to create file: %v", err)
	}
	defer outFile.Close()

	_, err = io.Copy(outFile, file)
	if err != nil {
		return "", fmt.Errorf("failed to save file: %v", err)
	}

	return filePath, nil
}

// Function to safely remove a file if it exists
func RemoveFileIfExists(filePath string) error {
    if filePath == "" {
        // File path is empty, no need to remove anything
        return nil
    }
    
    // Attempt to remove the file
    err := os.Remove(filePath)
    if err != nil {
        // If the error is because the file does not exist, ignore it
        if os.IsNotExist(err) {
            fmt.Printf("File does not exist: %s\n", filePath)
            return nil
        }
        // Return any other error
        return err
    }

    fmt.Printf("Successfully deleted file: %s\n", filePath)
    return nil
}
