package utils

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	UserName string
	jwt.StandardClaims
}

var jwtSecret = []byte("vrJ7oc1jMntcWYFe") // jwtSecret，这个建议复杂一点

func getToken() (string, error) { //生成Token，返回Encoded jwt
	expireTime := time.Now().Add(14 * 24 * time.Hour) // 设置14天的过期时间
	claims := &Claims{                                // jwt参数
		UserName: "root", //用户名
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(), //过期时间
			Issuer:    "blogo",           // 签名颁发者
		},
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(jwtSecret) //设置secret
	if err != nil {
		return "ERROR", err //错误：返回ERROR和错误信息
	}
	return token, nil //成功：返回Encoded jwt和nil
}

func verifyToken(tokenString string) (string, error) { //验证Token，接收一个Encoded jwt，返回UserName
	if tokenString == "" {
		return "ERROR", errors.New("tokenString is empty")
	}
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (i interface{}, err error) {
		return jwtSecret, nil
	})
	if err != nil || !token.Valid { // 验证此token时候过期
		return "ERROR", err
	}
	return claims.UserName, nil
}

/*
Use example
func yourfunc() {
	r := gin.Default()
	r.GET("/get", func(c *gin.Context) {
        token, err := getToken()
        if err != nil {
            c.JSON(401, gin.H{
                "Error": "获取Token失败",
                "Info": err,
            })
            return
        }
		c.JSON(200, gin.H{
			"JwtToken": token,
		})
	})
    r.GET("/ver", func(c *gin.Context) {
        tokenString := c.GetHeader("Authorization")
        userName, err := verifyToken(tokenString)
        if err != nil {
            fmt.Println(err)
            c.JSON(401, gin.H{
                "Error": "验证失败",
                "Info": err,
            })
            return
        }
        c.JSON(200, gin.H{
            "Info": "验证成功",
        })
        fmt.Println(userName)
    })
    r.Run(":8080")
}
*/
