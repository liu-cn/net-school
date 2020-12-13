package pages

import (
	"fmt"
	"gin/db"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Profile struct {
	Name      string `json:"name"`
	Age       int64  `json:"age"`
	Gender    string `json:"gender"`
	Password  string `json:"password"`
	Introduce string `json:"introduce"`
	Account   string `json:"account"`
	Phone     int64  `json:"phone"`
}

func GetProfileData(c *gin.Context) {
	var proflie Profile
	sqlStr := `select name,age,gender,password,introduce,account,phone from users where account=?`

	err := db.FlutterDb.Get(&proflie, sqlStr, "liubaorui")

	if err != nil {
		fmt.Println("查询失败！", err)
		return
	}
	fmt.Println(proflie)
	c.JSON(http.StatusOK, gin.H{
		"data": proflie,
		//"age":proflie,
		//"gender":proflie,
		//"password":proflie,
		//"introduce":proflie,
		//"account":proflie,
		//"phone" :proflie,
	})

}
