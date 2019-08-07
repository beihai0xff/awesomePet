package action

import (
	"awesomePet/api"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"strconv"
)

func UploadPet(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	uidString := claims["uid"].(string)
	uid, _ := strconv.ParseUint(uidString, 10, 64)

	form, err := c.MultipartForm()
	if err != nil {
		return err
	}
	if err := api.MultipartFileWrite(uidString, form); err != nil {

	}

}
