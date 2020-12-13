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
	Phone    string `db:"phone" json:"phone" binding:"required"`
}

type HasUser struct {
	Username string `db:"username" json:"username" binding:"required"`
	Phone    string `db:"phone" json:"phone" binding:"required"`
}

//数据库验证用户是否已经被注册
func RegisterUser(user *User) (bool, string) {
	fmt.Println("register user:", user)
	var hasUser HasUser

	queryStr := `select username,phone from users where username=? or phone=?`
	if err := db.FlutterDb.Get(&hasUser, queryStr, user.Username, user.Phone); err != nil {
		fmt.Println("Select err:", err, hasUser)
		sqlstr := `insert into users(username,password,phone)values(?,?,?)`
		res := db.FlutterDb.MustExec(sqlstr, user.Username, user.Password, user.Phone)
		fmt.Println(res)

		return true, "注册成功"
	} else {
		fmt.Println(hasUser)
		if hasUser.Username == user.Username {
			return false, "账号已经被注册！"
		} else if hasUser.Phone == user.Phone {
			return false, "手机号已经被注册！"
		} else {
			return false, "服务器故障，注册失败！"
		}
	}
}

//前端请求注册时返回的处理函数
func Register() func(*gin.Context) {
	return func(c *gin.Context) {
		var user User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusOK, gin.H{
				"msg": "注册失败！",
			})
			fmt.Println(user)
			fmt.Println(err)
			return
		}
		fmt.Println("处理函数的user:", user)
		registerStatus, msg := RegisterUser(&user)
		if registerStatus {
			token, err := token.GenerateToken(user.Username, user.Password)
			if err != nil {
				c.JSON(http.StatusBadGateway, gin.H{
					"msg": "生成token失败",
				})
				return
			}
			c.JSON(http.StatusOK, gin.H{
				"registerStatus": "ok",
				"msg":            msg,
				"token":          token,
				"username":       user.Username,
				"phone":          user.Phone,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"registerStatus": "no",
				"msg":            msg,
			})
		}

	}
}
