package model

import (    
    _ "fmt"
    "database/sql" 
    "github.com/gin-gonic/gin"
)
type BaseUser struct {
    id      int
    name    string
    age     int
}
var User BaseUser;

func (user BaseUser) Coll(db *sql.DB) gin.H{
    body := make(gin.H)
    rows, err := db.Query("SELECT * From user")  
    if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
    var dbdata []gin.H
    for rows.Next() {  
        var userOne BaseUser;
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
    return body;
}






