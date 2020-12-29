package register

import (
	"fmt"
	"gin/db"
	"gin/token"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
	"math/rand"
	"net/http"
	"time"
)

type User struct {
	Username string `bson:"username" db:"username" json:"username" binding:"required"`
	Password string `bson:"password" db:"password" json:"password" binding:"required"`
	Phone    string `bson:"phone" db:"phone" json:"phone" binding:"required"`
	Uid      int32  `bson:"uid" db:"uid" json:"uid"`
}

type HasUser struct {
	Username string `bson:"username" db:"username" json:"username" binding:"required"`
	Phone    string `bson:"phone" db:"phone" json:"phone" binding:"required"`
}

func GetUid() (int32, bool) {
	data := make(map[string]int32, 1)
	err := db.Mongdb.C("app").Find(bson.M{}).One(&data)
	if err != nil {
		fmt.Println("err:", err)
		rand.Seed(time.Now().UnixNano())
		var i int32
		i = int32(rand.Intn(1000) + 10000)
		return i, false
	}
	fmt.Printf("第%v个用户", data["id"])

	err = db.Mongdb.C("app").Update(bson.M{}, bson.M{"$inc": bson.M{"id": 1}})

	if err != nil {
		fmt.Println(err)
		rand.Seed(time.Now().UnixNano())
		var any int32
		any = int32(rand.Intn(1000) + 10000)
		return any, false
	}
	//如果没出错，返回查询到的uid，并且给uid+1，否则返回一个10000-11000的随机数，不能和uid冲突。
	return data["id"], true
}

//数据库验证用户是否已经被注册
func (this *User) Register() (bool, string, *User) {
	fmt.Println("register user:", this)
	var hasUser HasUser
	err := db.Mongdb.C("user").Find(
		bson.M{"$or": []bson.M{
			bson.M{"username": this.Username},
			bson.M{"phone": this.Phone},
		}}).One(&hasUser)
	fmt.Println("结果", hasUser)
	if err != nil { //说明没查到，未注册！
		fmt.Println("find err or nil:", err)
		//获取当前用户uid
		uid, ok := GetUid()
		if ok {
			this.Uid = uid
			//user := User{
			//	Uid:      uid,
			//	Phone:    this.Phone,
			//	Password: this.Password,
			//	Username: this.Username,
			//}
			err := db.Mongdb.C("user").Insert(this)
			if err != nil {
				return false, "注册插入失败，字段值可能有误！", this
			}
			return true, "注册成功", this
		}
	}
	return false, "手机号或用户名已经被注册！", &User{}
}

//前端请求注册时返回的处理函数
func Register() func(*gin.Context) {

	return func(ctx *gin.Context) {
		var user User
		if err := ctx.ShouldBindJSON(&user); err != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"msg": "注册失败！",
			})
			fmt.Println(user)
			fmt.Println(err)
			return
		}
		fmt.Println("处理函数的user:", user)
		ok, msg, getuser := user.Register()
		if ok {
			token, err := token.GenerateToken(user.Username, user.Password)
			if err != nil {
				ctx.JSON(http.StatusBadGateway, gin.H{
					"msg": "生成token失败",
				})
				return
			}
			ctx.JSON(http.StatusOK, gin.H{
				"registerStatus": "ok",
				"msg":            msg,
				"token":          token,
				"username":       user.Username,
				"phone":          user.Phone,
				"uid":            getuser.Uid,
			})
		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"registerStatus": "no",
				"msg":            msg,
			})
		}

	}
}
