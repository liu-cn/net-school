package main

import (
	"fmt"
	"gin/db"
	"gin/login"
	"gin/middleware"
	"gin/register"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {


	route := gin.Default()
	route.Use(middleware.Cors())
	route.POST("/login", login.Login())
	route.GET("/", middleware.JWTAuth,func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello world",
		})

	})

	route.POST("/api/list520",middleware.JWTAuth, func(c *gin.Context) {
		user,err:=db.Get520List()
		if err != nil {
			c.JSON(http.StatusOK,gin.H{
				"msg":"获取数据失败！",
				"list520":"",
			})
			fmt.Println(err,user)
			return
		}
		c.JSON(http.StatusOK,gin.H{
			"msg":"获取数据成功！",
			"list520":*user,
		})

	})
	route.POST("/register", register.Register())

	route.Run() // listen and serve on 0.0.0.0:8080
}


