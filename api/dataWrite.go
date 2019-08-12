package api

import (
	. "awesomePet/models"
	"bytes"
	"github.com/disintegration/imaging"
	"io"
	"io/ioutil"
	"mime/multipart"
	"os"
	"path"
	"strings"
)

func FileWrite(filePath string, file *multipart.FileHeader) (err error) {
	src, err := file.Open()
	if err != nil {
		return
	}
	defer src.Close()
	// Destination
	dst, err := os.Create(filePath)
	if err != nil {
		return
	}
	defer dst.Close()
	// Copy
	_, err = io.Copy(dst, src)
	return
}

func MultipartFileWrite(uid string, form *multipart.Form) (*Pet, error) {
	var builder strings.Builder
	// 向builder中写入字符/字符串
	builder.WriteString(OriginalFilePath)
	builder.WriteString(uid)
	builder.WriteString("/")
	tempPath := builder.String()
	err := os.MkdirAll(tempPath, os.ModePerm) // mkdir
	PrintErr(err)
	err = os.MkdirAll(ThumbnailFilePath+uid, os.ModePerm) // mkdir
	PanicErr(err)
	var m Pet
	var i uint
	files := form.File["files"]
	for _, file := range files {
		// Source
		src, err := file.Open()
		if err != nil {
			return &m, err
		}
		//defer src.Close()
		filePath := tempPath + file.Filename
		// Destination
		dst, err := os.Create(filePath)
		if err != nil {
			return &m, err
		}
		//defer dst.Close()
		// Copy
		if _, err = io.Copy(dst, src); err != nil {
			return &m, err
		}
		err = dst.Close()
		err = src.Close()
		hash, err := DataHash(filePath)
		PrintErr(err)
		ext := path.Ext(filePath)
		if err = os.Rename(filePath, tempPath+hash+ext); err != nil {
			_ = os.Remove(filePath)
			return &m, err //file rename
		}
		var builder strings.Builder
		builder.WriteString(hash)
		builder.WriteString(ext)
		filePath = builder.String()
		if err = Resize(&uid, &filePath); err != nil {
			return &m, err
		}
		m.Pic = append(m.Pic, Pic{OrderID: i, PetHash: hash, Ext: ext})
		i++
	}
	return &m, nil
}

func Resize(uid, filename *string) (err error) {
	imgData, _ := ioutil.ReadFile(OriginalFilePath + *uid + "/" + *filename)
	buf := bytes.NewBuffer(imgData)
	image, err := imaging.Decode(buf)
	if err != nil {
		return
	}
	//生成缩略图，传0表示等比例放缩
	image = imaging.Resize(image, 0, 480, imaging.Lanczos)
	return imaging.Save(image, ThumbnailFilePath+*uid+"/tn_"+*filename)
}

func ShowPP(filename string) error {
	imgData, _ := ioutil.ReadFile(OriginalPPPath + filename)
	buf := bytes.NewBuffer(imgData)
	image, err := imaging.Decode(buf)
	if err != nil {
		return err
	}
	//生成缩略图
	image = imaging.Resize(image, 72, 72, imaging.Lanczos)
	return imaging.Save(image, ThumbnailPPPath+"tn_"+filename)
}
