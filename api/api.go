package api

import (
	"fmt"
	"gin/db"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
	"net/http"
)

func GetBlog(ctx *gin.Context) {
	var data map[string]interface{}
	err := db.Mongdb.C("article").Find(bson.M{"id": 0}).One(&data)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data":   data,
		"status": "ok",
	})

}
