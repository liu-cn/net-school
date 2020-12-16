package pages

import (
	"fmt"
	"gin/db"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Profile struct {
	Username  string `json:"username"`
	Age       int64  `json:"age"`
	Gender    string `json:"gender"`
	Password  string `json:"password"`
	Introduce string `json:"introduce"`
	Nickname  string `json:"nickname""`
	Phone     int64  `json:"phone"`
	Headimg   string `json:"headimg"`
	Uid       int64  `json:"uid"`
}

func GetProfileData(ctx *gin.Context) {
	query := ctx.Query("uid")
	var proflie Profile
	sqlStr := `select username,age,gender,password,introduce,nickname,phone,headimg,uid from users where uid=?`

	err := db.LocalDb.Get(&proflie, sqlStr, query)

	if err != nil {
		fmt.Println("查询失败！", err)
		return
	}
	fmt.Println(proflie, proflie.Uid)
	ctx.JSON(http.StatusOK, gin.H{
		"data": proflie,
	})

}
