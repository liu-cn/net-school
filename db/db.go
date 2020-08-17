package db

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var Db *sqlx.DB

type User struct {
	Id int64 `db:"id"`
	Name string `db:"name"`
	Age int64 `db:"age"`
	Content string `db:"content"`
	Gender string `db:"gender"`
}

func init() {
	var err error
	Db,err=sqlx.Open("mysql","test:123456@tcp(127.0.0.1:3306)/user")
	if err != nil {
		panic(err)
	}
	var  user  []User
	err=Db.Select(&user,"select id,name,age,gender,content from list where id>?",0)
	if err != nil {
		panic(err)
	}
	fmt.Println(user)
}

func Get520List() (*[]User,error){
	var user []User
	err:=Db.Select(&user,"select id,name,age,gender,content from list where id>?",0)
	if err != nil {
		return &user, err
	}
	return &user, nil
}

