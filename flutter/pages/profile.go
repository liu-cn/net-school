package pages

import (
	"fmt"
	"gin/db"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"strconv"
)

type Profile struct {
	Username  string `bson:"username" json:"username"`
	Age       int32  `bson:"age" json:"age"`
	Gender    int32  `bson:"gender" json:"gender"`
	Password  string `bson:"password" json:"password"`
	Introduce string `bson:"introduce" json:"introduce"`
	Nickname  string `bson:"nickname" json:"nickname"`
	Phone     string `bson:"phone" json:"phone"`
	Headimg   string `bson:"headimg" json:"headimg"`
	Uid       int32  `bson:"uid" json:"uid"`
}

func GetProfileData(ctx *gin.Context) {
	queru := ctx.Query("uid")
	fmt.Println("uid:=========", queru)
	uid, err := strconv.Atoi(queru)
	if err != nil {
		fmt.Println("字符串转换int失败！err", err)
		return
	}
	var profile Profile
	err = db.Mongdb.C("user").Find(bson.M{"uid": int32(uid)}).One(&profile)
	fmt.Println("查询结果：", profile)

	if err != nil {
		fmt.Println("查询失败！", err)
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"msg": "err",
		})
		return
	}
	fmt.Println(profile, profile.Uid)
	ctx.JSON(http.StatusOK, gin.H{
		"data": profile,
	})

}
