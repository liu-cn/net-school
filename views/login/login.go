package login

import (
	"fmt"
	"gin/db"
	"gin/token"
	"github.com/gin-gonic/gin"
	"net/http"
)

//结构体这里字段名必须大写，不然接收不到数据。
type UserLogin struct {
	Uid       int64  `json:"uid"`
	Phone     string `json:"phone" binding:"required"`
	Password  string `json:"password" binding:"required"`
	Age       int8   `json:"age"`
	Gender    int8   `json:"gender"`
	Introduce string `json:"introduce"`
	Username  string `json:"username"`
	Headimg   string `json:"headimg"`
	Nickname  string `json:"nickname"`
}

//验证账号是否存在
func VerifyAccount(phone, password string) (bool, string, UserLogin) {
	var loginUser UserLogin
	sqlStr := `select phone,password,uid,age,gender,introduce,username,headimg,nickname from users where phone=? && password=?`
	if err := db.LocalDb.Get(&loginUser, sqlStr, phone, password); err != nil {
		fmt.Println("查询结果：err：", err)
		return false, "账号或密码错误！", loginUser
	} else {
		return true, "登录成功！", loginUser
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

		ok, msg, loginInfo := VerifyAccount(login.Phone, login.Password)
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
				"data":   loginInfo,
				"token":  token,
				//"phone":     login.Phone,
				//"msg":       msg,
				//"uid":       loginInfo.Uid,
				//"password":  loginInfo.Password,
				//"age":       loginInfo.Age,
				//"gender":    loginInfo.Gender,
				//"introduce": loginInfo.Introduce,
				//"username":  loginInfo.Username,
				//"headimg":   loginInfo.Headimg,
				//"nickname":  loginInfo.Nickname,
			})
		}
	}
}
