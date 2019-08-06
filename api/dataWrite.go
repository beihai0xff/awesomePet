package api

import (
	. "awesomePet/models"
	"bytes"
	"github.com/disintegration/imaging"
	"io"
	"io/ioutil"
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

func Resize(filename string) error {
	imgData, _ := ioutil.ReadFile(OriginalPPPath + filename)
	buf := bytes.NewBuffer(imgData)
	image, err := imaging.Decode(buf)
	if err != nil {
		return err
	}
	//生成缩略图，传0表示等比例放缩
	image = imaging.Resize(image, 72, 72, imaging.Lanczos)
	err = imaging.Save(image, ThumbnailPPPath+"tn_"+filename)
	if err != nil {
		return err
	}
	return nil
}
