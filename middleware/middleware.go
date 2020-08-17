package middleware

import (
	"fmt"
	"gin/token"
	"github.com/gin-gonic/gin"
	"net/http"
)

//中间件验证token
func JWTAuth(c *gin.Context)  {
	//token:=c.DefaultQuery("token","")
	var index token.Index
	if err:=c.ShouldBindJSON(&index); err!=nil{
		fmt.Println(1)
		c.JSON(http.StatusBadRequest,gin.H{
			"msg":"token参数解析失败！",
		})
		return
	}
	fmt.Println("index:",index)

	if index.Token==""{
		fmt.Println(2)
		c.JSON(http.StatusBadRequest,gin.H{
			"msg":"未携带token",
		})
	}
	claims,err:=token.ParseToken(index.Token)
	if err != nil {
		fmt.Println(3)
		c.JSON(http.StatusBadRequest,gin.H{
			"msg":err.Error(),
		})
		c.Abort()
		return
	}
	fmt.Println("claims:",claims)
	c.Set("claims",claims)
}

//允许跨域访问，
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")

		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		// 处理请求
		c.Next()
	}
}