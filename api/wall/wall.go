package wall

import (
	"fmt"
	"gin/db"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CommitObj struct {
	Content string `db:"content" json:"content"`
	IsPublic string `db:"IsPublic" json:"IsPublic"`
	UserName string `db:"username" json:"username"`
}

func GetContentWallList() func(c *gin.Context) {
	return func(c *gin.Context) {
		user,err:=SqlWallList()
		if err != nil {
			c.JSON(http.StatusOK,gin.H{
				"msg":"获取数据失败！",
				"list520":"",
			})
			fmt.Println(err,user)
			return
		}
		c.JSON(http.StatusOK,gin.H{
			"msg":"获取数据成功！",
			"wallContentList":*user,
		})

	}
}



func CommitWallContent()func(c *gin.Context) {
	return func(c *gin.Context) {
		var commit CommitObj
		if err := c.ShouldBindJSON(&commit);err!=nil{
			c.JSON(http.StatusOK,gin.H{
				"msg":"解析参数失败！",
				"CommitStatus":"no",
			})
			return
		}
		fmt.Println(commit)
		sqlStr:="insert into list(username,content,IsPublic)values(?,?,?)"
		result := db.Db.MustExec(sqlStr,commit.UserName, commit.Content, commit.IsPublic)
		fmt.Println(result)
		c.JSON(http.StatusOK,gin.H{
			"msg":"表白成功！",
			"CommitStatus":"ok",
		})
	}
}
