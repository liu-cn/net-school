package token

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

//密钥
var jwtSecret = "hello"

type Index struct {
	Token string `form:"token" json:"token" binding:"required"`
}

type Claims struct {
	Username string
	Password string
	jwt.StandardClaims
}


//生成token
func GenerateToken(username,password string)(string,error){
	nowtime:=time.Now()
	expireTime:= nowtime.Add(3 *time.Hour)
	claims:=Claims{
		username,
		password,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer: "mrliu",

		},
	}
	tokenClaims:=jwt.NewWithClaims(jwt.SigningMethodHS256,claims)
	token,err:=tokenClaims.SignedString([]byte(jwtSecret))

	return token,err

}

//解析token
func ParseToken(token string)(*Claims,error){
	tokenClaims,err:=jwt.ParseWithClaims(token,&Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecret),nil
	})
	if tokenClaims!=nil {
		if claims,ok:=tokenClaims.Claims.(*Claims);ok&&tokenClaims.Valid{
			return claims,err
		}
	}
	return nil, err
}