package wall

import "gin/db"

type User struct {
	Id int64 `db:"id"`
	UserName string `db:"username"`
	Content string `db:"content"`
}


func SqlWallList() (*[]User,error){
	var user []User
	err:=db.Db.Select(&user,"select id,username,content from list where id>?",0)
	if err != nil {
		return &user, err
	}
	return &user, nil
}