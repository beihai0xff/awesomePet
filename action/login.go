package action

import (
	"awesomePet/api/debug"
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
			UserName:    m.UserName,
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
	userInfo := gorm_mysql.GetUserPassword(&m.Uid)
	getSalt, err := hex.DecodeString(userInfo.Salt)
	debug.PanicErr(err)
	key := pbkdf2.Key([]byte(m.Password), getSalt, 1323, 32, sha256.New)
	if hex.EncodeToString(key) == userInfo.Key {
		// CreateUser token
		token := jwt.New(jwt.SigningMethodHS256)
		// Set claims
		claims := token.Claims.(jwt.MapClaims)
		claims["uid"] = userInfo.Uid
		claims["exp"] = time.Now().Add(time.Hour * 72).Unix() //有效期三天
		// Generate encoded token and send it as response.
		t, err := token.SignedString([]byte("233333"))
		debug.PanicErr(err)
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
	userInfo := gorm_mysql.GetUserPassword(&m.Uid)
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
		err := gorm_mysql.UpdateUserPassword(user)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, models.ResultWithNote{Result: true, Note: "密码更新成功"})
	} else {
		return c.JSON(http.StatusOK, models.ResultWithNote{Result: false, Note: "用户名或密码错误"})
	}
}
