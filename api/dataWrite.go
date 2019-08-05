package api

import (
	"io"
	"mime/multipart"
	"os"
)

func DataWrite(filePath string, file *multipart.FileHeader) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()
	// Destination
	dst, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer dst.Close()
	// Copy
	if _, err := io.Copy(dst, src); err != nil {
		return err
	}
	return nil
}
