package action

import (
	"awesomePet/api"
	"awesomePet/gorm_mysql"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func UploadPetBlog(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	uidString := claims["uid"].(string)
	uid, _ := strconv.ParseUint(uidString, 10, 64)
	description := c.FormValue("description")
	form, err := c.MultipartForm()
	if err != nil {
		return err
	}
	m, err := api.MultipartFileWrite(uidString, form)
	if err != nil {
		return err
	} else {
		m.Uid = uid
		m.Description = description
		err := gorm_mysql.CreatePet(m)
		if err != nil {
			return err
		}
	}
	return c.JSON(http.StatusOK, m)
	//return c.JSON(http.StatusOK, models.ResultWithNote{Result: true, Note: "blog 发布成功"})
}
