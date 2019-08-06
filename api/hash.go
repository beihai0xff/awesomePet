package api

import (
	"crypto/sha256"
	"encoding/hex"
	"io"
	"os"
)

func DataHash(filePath string) (string, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer f.Close()
	h := sha256.New()
	_, err = io.Copy(h, f)
	if err != nil {
		return "", err
	}
	dataHash := hex.EncodeToString(h.Sum(nil))
	return dataHash, nil
}

func MsgHash(Msg string) string {
	h := sha256.New()
	h.Write([]byte(Msg))
	return hex.EncodeToString(h.Sum(nil))
}
