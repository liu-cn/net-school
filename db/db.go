package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var Db *sqlx.DB

//type User struct {
//	Id int64 `db:"id"`
//	UserName string `db:"username"`
//	Content string `db:"content"`
//}

func init() {
	var err error
	Db,err=sqlx.Open("mysql","test:123456@tcp(127.0.0.1:3306)/user")
	if err != nil {
		panic(err)
	}
	//var  user  []User
	//err=Db.Select(&user,"select id,username,content from list where id>?",0)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(user)
}

//func SqlWallList() (*[]User,error){
//	var user []User
//	err:=Db.Select(&user,"select id,username,content from list where id>?",0)
//	if err != nil {
//		return &user, err
//	}
//	return &user, nil
//}

