package checkerr

import "fmt"

func CheckErr(err ...interface{}){
	var msg interface{}
	for _,v :=range err{
		switch v.(type) {
			case string:
				 msg = v
			case error:
				if v!=nil {
					fmt.Println("err:",msg,v)
					return
				}
		}
	}
	if err!=nil {
		fmt.Println("err:",err)
	}
}
