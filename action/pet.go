package action

import (
	"awesomePet/api"
	"awesomePet/gorm_mysql"
	"awesomePet/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func UploadBlog(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	uidString := claims["uid"].(string)
	uid, _ := strconv.ParseUint(uidString, 10, 64)
	description := c.FormValue("description")
	title := c.FormValue("title")
	tag := c.FormValue("tag")
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
		m.Title = title
		m.Tag = tag
		err := gorm_mysql.CreatePet(m)
		if err != nil {
			return err
		}
	}
	return c.JSON(http.StatusOK, m)
	//return c.JSON(http.StatusOK, models.ResultWithNote{Result: true, Note: "blog 发布成功"})
}

func UpdateBlogContext(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	uidString := claims["uid"].(string)
	uid, _ := strconv.ParseUint(uidString, 10, 64)
	m := new(models.Pet)
	if err := c.Bind(m); err != nil {
		return err
	}
	if err := gorm_mysql.UpdatePet(uid, m); err != nil {
		return err
	}
	//return c.JSON(http.StatusOK, m)
	return c.JSON(http.StatusOK, models.ResultWithNote{Result: true, Note: "blog 重编辑成功"})
}

func DeleteBlog(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	uidString := claims["uid"].(string)
	uid, _ := strconv.ParseUint(uidString, 10, 64)
	var id uint
	if err := api.StrToUint(c.Param("id"), &id); err != nil {
		return c.JSON(http.StatusOK, models.ResultWithNote{Result: true, Note: "blog id未找到"})
	}
	m := models.Pet{Uid: uid, ID: id}
	if err := gorm_mysql.DeleteBlog(&m); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, models.ResultWithNote{Result: true, Note: "blog 删除成功"})
}

func GetUserBlog(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	uidString := claims["uid"].(string)
	uid, _ := strconv.ParseUint(uidString, 10, 64)
	m, err := gorm_mysql.GetUserBlog(&uid)
	if err != nil {
		return err
	}
	return c.JSONPretty(http.StatusOK, m, " ")
}
