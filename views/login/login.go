package login

import (
	"fmt"
	"gin/db"
	"gin/token"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserLogin struct {
	//结构体这里字段名必须大写，不认接收不到数据。
	//Username string `form:"username" json:"username" uri:"username" binding:"required"`
	Phone    string `form:"phone" json:"phone" uri:"phone" binding:"required"`
	Password string `form:"password" json:"password" uri:"password" binding:"required"`
}

//验证账号是否存在
func VerifyAccount(phone, password string) (bool, string) {
	var loginUser UserLogin
	sqlStr := `select phone,password from users where phone=? && password=?`
	if err := db.FlutterDb.Get(&loginUser, sqlStr, phone, password); err != nil {
		fmt.Println("查询结果：err：", err)
		return false, "账号或密码错误！"
	} else {
		return true, "登录成功！"
	}
}

func Login() func(c *gin.Context) {
	return func(c *gin.Context) {
		var login UserLogin
		if err := c.ShouldBindJSON(&login); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": "err",
				"msg":    "参数解析错误！",
			})
			fmt.Println(err)
			return
		}
		fmt.Println(login)

		ok, msg := VerifyAccount(login.Phone, login.Password)
		if !ok {
			c.JSON(http.StatusOK, gin.H{
				"status": "fail",
				"msg":    msg,
			})
			return
		} else {
			token, err := token.GenerateToken(login.Phone, login.Password)
			if err != nil {
				c.JSON(http.StatusBadGateway, gin.H{
					"status": "err",
					"msg":    "生成token失败",
				})
				return
			}
			c.JSON(http.StatusOK, gin.H{
				"status": "ok",
				"phone":  login.Phone,
				"token":  token,
				"msg":    msg,
			})
		}
	}
}
