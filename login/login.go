package login

import (
	"fmt"
	"gin/token"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserLogin struct {
	//结构体这里字段名必须大写，不认接收不到数据。
	Username string `form:"username" json:"username" uri:"username" binding:"required"`
	Password string `form:"password" json:"password" uri:"password" binding:"required"`
}


func Login()func(c *gin.Context){
	return func(c *gin.Context) {
		var login UserLogin
		if err:=c.ShouldBindJSON(&login); err!=nil{
			c.JSON(http.StatusBadRequest,gin.H{
				"msg":"参数解析错误！",
			})
			fmt.Println(err)
			return
		}
		fmt.Println(login)
		if login.Username!="admin" || login.Password!="123456" {
			c.JSON(http.StatusOK,gin.H{
				"LoginStatus":"failed",
				"msg":"账号或密码错误！",
			})
			return
		}

		token,err:=token.GenerateToken(login.Username,login.Password)
		if err != nil {
			c.JSON(http.StatusBadGateway,gin.H{
				"msg":"生成token失败",
			})
			return
		}
		c.JSON(http.StatusOK,gin.H{
			"username":login.Username,
			"password":login.Password,
			"token":token,
		})


	}
}