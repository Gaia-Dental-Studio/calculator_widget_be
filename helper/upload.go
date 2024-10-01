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
