package main

import (
	"gin/api"
	"gin/api/wall"
	"gin/download"
	"gin/flutter/pages"
	"gin/middleware"
	"gin/qiniu"
	"gin/upload"
	"gin/video"
	"gin/views/home"
	"gin/views/login"
	"gin/views/profile"
	"gin/views/register"
	"github.com/gin-gonic/gin"
)

func main() {
	route := gin.Default()
	route.Use(middleware.Cors())
	//route.GET("/", middleware.JWTAuth,home.Home())
	route.GET("/", home.Home())

	//route.POST("/api/wall",middleware.JWTAuth, wall.GetContentWallList())
	route.POST("/api/wall", wall.GetContentWallList())
	route.POST("/api/commitwall", wall.CommitWallContent())
	route.POST("/GetUserInfo", profile.GetUserInfo())
	route.POST("/upload", upload.QiniuUpload())
	route.GET("/video", video.Video())
	route.POST("/getqiniutoken", qiniu.GetQiniuToken)
	route.GET("/download", download.Download)

	//aiyingyuan api
	route.GET("/profile", pages.GetProfileData)
	route.POST("/login", login.Login())
	route.POST("/register", register.Register())
	route.GET("/getblog", api.GetBlog)
	//test.TestMongo()
	//test.TestPostMongdbData()
	//register.GetUid()
	route.Run() // listen and serve on 0.0.0.0:8080

}
