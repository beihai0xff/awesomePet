package action

import (
	"awesomePet/api"
	"awesomePet/gorm_mysql"
	"awesomePet/models"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/pbkdf2"
	"net/http"
	"os"
	"path"
	"strconv"
	"time"
)

func Register(c echo.Context) error {
	m := new(models.RequestUser)
	if err := c.Bind(m); err != nil {
		return err
	}
	fmt.Printf("uid为: %d 密码为: %s \n", m.Uid, m.Password)
	if gorm_mysql.Has(m.Uid) {
		return c.JSON(http.StatusOK, models.ResultWithNote{Result: false, Note: "该用户已存在"})
	} else {
		//pbkdf2加密
		salt := make([]byte, 32)
		_, err := rand.Read(salt)
		if err != nil {
			return err
		}
		key := pbkdf2.Key([]byte(m.Password), salt, 1323, 32, sha256.New)
		User := models.User{Uid: m.Uid, Salt: hex.EncodeToString(salt), Key: hex.EncodeToString(key)}
		if err = gorm_mysql.CreateUser(&User); err != nil {
			return err
		}
		UserInfo := models.UserInfo{
			Uid:         m.Uid,
			Nickname:    m.UserName,
			Sex:         m.Sex,
			Description: m.Description,
			Email:       m.Email,
			City:        m.City,
			Street:      m.Street,
		}
		if err = gorm_mysql.CreateUserInfo(&UserInfo); err != nil {
			return err
		}
		return c.JSON(http.StatusOK, models.ResultWithNote{Result: true, Note: "注册成功"})
	}
}

func Login(c echo.Context) error {
	m := new(models.RequestUser)
	if err := c.Bind(m); err != nil {
		return err
	}
	fmt.Printf("uid为: %d 密码为: %s \n", m.Uid, m.Password)
	userInfo, err := gorm_mysql.GetUserPassword(&m.Uid)
	if err != nil {
		return err
	}
	getSalt, err := hex.DecodeString(userInfo.Salt)
	if err != nil {
		return err
	}
	key := pbkdf2.Key([]byte(m.Password), getSalt, 1323, 32, sha256.New)
	if hex.EncodeToString(key) == userInfo.Key {
		// CreateUser token
		token := jwt.New(jwt.SigningMethodHS256)
		// Set claims
		claims := token.Claims.(jwt.MapClaims)
		claims["uid"] = strconv.FormatUint(userInfo.Uid, 10)
		claims["exp"] = time.Now().Add(time.Hour * 72).Unix() //有效期三天
		// Generate encoded token and send it as response.
		t, err := token.SignedString([]byte("yourSecret"))
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, models.Token{Result: true, Token: t})
	} else {
		return c.JSON(http.StatusOK, models.Token{Result: false, Token: "用户名或密码错误"})
	}
}

func Reset(c echo.Context) error {
	m := new(models.PasswordReset)
	if err := c.Bind(m); err != nil {
		return err
	}
	fmt.Printf("uid为: %d 密码为: %s \n", m.Uid, m.OldPassword)
	userInfo, err := gorm_mysql.GetUserPassword(&m.Uid)
	if err != nil {
		return err
	}
	getSalt, err := hex.DecodeString(userInfo.Salt)
	if err != nil {
		return err
	}
	key := pbkdf2.Key([]byte(m.OldPassword), getSalt, 1323, 32, sha256.New)
	if hex.EncodeToString(key) == userInfo.Key {
		//pbkdf2加密
		salt := make([]byte, 32)
		_, err = rand.Read(salt)
		if err != nil {
			return err
		}
		key := pbkdf2.Key([]byte(m.NewPassword), salt, 1323, 32, sha256.New)
		user := models.User{Salt: hex.EncodeToString(salt), Key: hex.EncodeToString(key)}
		err := gorm_mysql.UpdateUserPassword(&user)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, models.ResultWithNote{Result: true, Note: "密码更新成功"})
	} else {
		return c.JSON(http.StatusOK, models.ResultWithNote{Result: false, Note: "用户名或密码错误"})
	}
}

func ProfilePhoto(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	uidString := claims["uid"].(string)
	uid, _ := strconv.ParseUint(uidString, 10, 64)
	file, err := c.FormFile("profile")
	if err != nil {
		return err
	}
	tempPath := models.OriginalPPPath + file.Filename
	if err = api.DataWrite(tempPath, file); err != nil {
		fmt.Println(err)
		return err //data copy
	}
	ext := path.Ext(tempPath)
	filename := uidString + ext
	if err = os.Rename(tempPath, models.OriginalPPPath+filename); err != nil {
		err = os.Remove(tempPath)
		return err //file rename
	}
	//生成缩略图
	if err := api.Resize(filename); err != nil { //生成缩略图
		err = os.Remove(filename)
		return err
	}
	if err := gorm_mysql.UpdatePPExt(uid, ext); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, models.Ext{Result: true, Ext: ext})
}

func ThumbnailProfilePhoto(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	uidString := claims["uid"].(string)
	ext := c.QueryParam("ext")
	return c.Inline(models.ThumbnailPPPath+"tn_"+uidString+ext, "thumbnail"+ext)
}

func GetUserInfo(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	uidString := claims["uid"].(string)
	uid, _ := strconv.ParseUint(uidString, 10, 64)
	m, err := gorm_mysql.GetUserInfo(&uid)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, m)
}
