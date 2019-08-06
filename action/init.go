package action

import (
	"awesomePet/api/debug"
	. "awesomePet/models"
	"os"
)

func Init() {
	err := os.MkdirAll(OriginalPPPath, os.ModePerm) // mkdir
	debug.PanicErr(err)
	err = os.MkdirAll(ThumbnailPPPath, os.ModePerm) // mkdir
	debug.PanicErr(err)
}
