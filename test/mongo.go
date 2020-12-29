package test

import (
	"fmt"
	"gin/db"
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	Name string `bson:"name"`
	Age  string `bson:"age"`
}

type TestUser struct {
	//Id_       bson.ObjectId `bson:_id`
	Nickname  string        `bson:nickname`
	Username  string        `bson:username`
	Password  string        `bson:password`
	Age       int32         `bson:age`
	Gender    int32         `bson:gender`
	Introduce string        `bson:introduce`
	Phone     int64         `bson:phone`
	HeadImg   string        `bson:headImg`
	Uid       int32         `bson:uid`
}

func TestPostMongdbData() {
	data := make(map[string]int, 1)
	err := db.Mongdb.C("app").Find(bson.M{}).One(&data)
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Println(data["id"])

	if err != nil {
		fmt.Println("err",err)
		return
	}
	err= db.Mongdb.C("app").Update(bson.M{}, bson.M{"$inc": bson.M{"id": 1}})

	if err != nil {
		fmt.Println(err)
		return
	}
}
