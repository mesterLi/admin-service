package utils

import (
	"admin-service/model"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"strconv"
	"time"
)

type JWTClaims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Uid int `json:"uid"`
	jwt.StandardClaims
}

const (
	ExpireHours = 12
	Secret = "脆弱时间到，我们一起来祷告"
)

func GetToken(user model.User) (string, error) {
	claims := JWTClaims{
		user.Username,
		user.Password,
		user.Uid,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(ExpireHours * time.Hour).Unix(),
			Issuer: "thankslyh",
		},
	}
	//fmt.Println(claims)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	siginedToken, err := token.SignedString([]byte(Secret))
	if err != nil {
		fmt.Println("err.....", err)
		return "", err
	}
	return siginedToken, nil
}

func ValidadteToken(token string) (bool, error) {
	t, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(Secret), nil
	})
	if err != nil {
		return false, err
	}
	if t.Valid == false {
		fmt.Println("token is not valid")
		return false, nil
	}
	return true, nil
}

func GetUidFromToken(token string, id *int) {
	t, _ := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(Secret), nil
	})
	claims := t.Claims.(jwt.MapClaims)
	for key, val := range claims {
		if key == "uid" {
			//fmt.Println(key, val)
			//fmt.Println(reflect.TypeOf(val))
			*id, _ = strconv.Atoi(fmt.Sprintf("%1.0f", val.(float64)))
			break
		}
	}
	
}