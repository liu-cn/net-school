package profile

import (
	"fmt"
	"gin/db"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserInfo struct {
	Nickname string `db:"nickname" json:"nickname"`
	Gender int8 `db:"gender" json:"gender"`
	UserName string `db:"username" json:"username"`
	Age int8 `db:"age" json:"age"`
	Profile string `db:"profile" json:"profile"`
	Phone int64 `db:"phone" json:"phone"`
	Email string `db:"email" json:"email"`
	Name string `db:"name" json:"name"`
}

type UserName struct {
	UserName string `db:"username" json:"username"`
}

func GetUserInfo() func(c *gin.Context) {
	return func(c *gin.Context) {
		var username UserName
		if err := c.ShouldBindJSON(&username);err!=nil{
			fmt.Println("解析参数失败！：err:",err)
			c.JSON(http.StatusOK,gin.H{
				"msg":"解析参数失败！",
			})
			return
		}
		fmt.Println("解析后数据:username",username)
		var userinfo  UserInfo
		sqlStr:=`select nickname,gender,username,age,profile,phone,email from userlist where username=?`
		if err := db.Db.Get(&userinfo, sqlStr, username.UserName);err!=nil{
			fmt.Println("找不到该用户：err",err)
			c.JSON(http.StatusOK,gin.H{
				"msg":"找不到用户信息",
			})
			return
		}else {
			c.JSON(http.StatusOK,gin.H{
				"msg":"获取用户数据成功!",
				"nickname":userinfo.Nickname,
				"gender":userinfo.Gender,
				"username":userinfo.UserName,
				"age":userinfo.Age,
				"profile":userinfo.Profile,
				"phone":userinfo.Phone,
				"email":userinfo.Email,
				"name":userinfo.Name,
			})
		}
	}
}