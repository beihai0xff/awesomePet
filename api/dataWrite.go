package api

import (
	"awesomePet/api/debug"
	. "awesomePet/models"
	"bytes"
	"fmt"
	"github.com/disintegration/imaging"
	"io"
	"io/ioutil"
	"mime/multipart"
	"os"
	"path"
	"strings"
)

func FileWrite(filePath string, file *multipart.FileHeader) error {
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

func MultipartFileWrite(uid string, form *multipart.Form) error {
	var builder strings.Builder
	// 向builder中写入字符/字符串
	builder.WriteString(OriginalFilePath)
	builder.WriteString(uid)
	builder.WriteString("/")
	tempPath := builder.String()
	fmt.Println(tempPath)
	var m Pet
	files := form.File["files"]
	for i, file := range files {
		// Source
		src, err := file.Open()
		if err != nil {
			return err
		}
		defer src.Close()
		filePath := tempPath + file.Filename
		// Destination
		dst, err := os.Create(filePath)
		if err != nil {
			return err
		}
		defer dst.Close()
		// Copy
		if _, err = io.Copy(dst, src); err != nil {
			return err
		}

		hash, err := DataHash(filePath)
		debug.PrintErr(err)
		ext := path.Ext(filePath)
		if err = os.Rename(filePath, tempPath+hash+ext); err != nil {
			err = os.Remove(tempPath)
			return err //file rename
		}
		m.Pic = append(m.Pic, Pic{ID: i, PetHash: hash, Ext: ext})
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
