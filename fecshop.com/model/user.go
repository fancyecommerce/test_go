package model

import (    
    _ "fmt"
    "fecshop.com/system"
    "database/sql" 
    "github.com/gin-gonic/gin"
)
type UserTb struct {
    id      int
    name    string
    age     int
}
var User UserTb;

func (user UserTb) Coll(db *sql.DB) gin.H{
    data := list(db)
    return data
}

func list(db *sql.DB) gin.H{
    body := make(gin.H)
    rows, err := db.Query("SELECT * From user")  
    if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
    var dbdata []gin.H
    for rows.Next() {  
        // var id int  
        // var name string  
        // var age string
        var userOne UserTb;
        err = rows.Scan(&userOne.id, &userOne.name ,&userOne.age) 
        if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
        dd := gin.H{
            "id": userOne.id ,
            "name": userOne.name,
            "age": userOne.age,
        }
        dbdata = append(dbdata, dd)
    }  
    
    err = rows.Err()  
    if err != nil {  
        panic(err.Error())  
    }  
    body["data"] = dbdata
    body["i"] = system.I
    return body;
}


func (user UserTb) Coll2() gin.H{
    db := system.Mysql.Open()
    data := list(db)
    return data
}





