package register

import (
	"fmt"
	"gin/db"
	"gin/token"
	"github.com/gin-gonic/gin"
	"net/http"
)

type User struct {
	Username string `db:"username" json:"username" binding:"required"`
	Password string `db:"password" json:"password" binding:"required"`
	Phone string `db:"phone" json:"phone" binding:"required"`
}

//数据库查询用户是否注册
func RegisterUser(user *User) bool{
	fmt.Println("register user:",user)
	sqlstr:=`insert into userlist(username,password,phone)values(?,?,?)`
	res:=db.Db.MustExec(sqlstr,user.Username,user.Password,user.Phone)
	fmt.Println(res)

	return true
}

//前端请求注册时返回的处理函数
func Register()func(*gin.Context){
	return func(c *gin.Context){
		var user User
		if err := c.ShouldBindJSON(&user);err != nil {
			c.JSON(http.StatusOK,gin.H{
				"msg":"注册失败！",
			})
			fmt.Println(user)
			fmt.Println(err)
			return
		}
		fmt.Println("处理函数的user:",user)
		registerStatus := RegisterUser(&user)
		if registerStatus{
			token,err:=token.GenerateToken(user.Username,user.Password)
			if err != nil {
				c.JSON(http.StatusBadGateway,gin.H{
					"msg":"生成token失败",
				})
				return
			}
			c.JSON(http.StatusOK,gin.H{
				"registerStatus":"ok",
				"token":token,
			})
		}

	}
}