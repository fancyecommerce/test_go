package model

import (    
    _ "fmt"
    "database/sql" 
    "github.com/gin-gonic/gin"
)
type UserType struct {
    Id      int    `form:"id" json:"id" `
    Name    string `form:"name" json:"name" binding:"required"`
    Age     int    `form:"age" json:"age" binding:"required"`
}
var User UserType;

func (user UserType) Coll(db *sql.DB) gin.H{
    body := make(gin.H)
    rows, err := db.Query("SELECT * From user")  
    if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
    var dbdata []gin.H
    for rows.Next() {  
        var userOne UserType;
        err = rows.Scan(&userOne.Id, &userOne.Name ,&userOne.Age) 
        if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
        dd := gin.H{
            "id": userOne.Id ,
            "name": userOne.Name,
            "age": userOne.Age,
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



func (user UserType) Save(db *sql.DB, json *UserType) gin.H{
    //return gin.H{"id":json.Id,"name":json.Name, "age": json.Age}
    
    body := make(gin.H)
    // id不存在，则插入
    if json.Id == 0 {
        stmtIns, err := db.Prepare("INSERT INTO user (`name`, `age`) VALUES( ?, ? )") // ? = placeholder
        if err != nil {
            panic(err.Error()) // proper error handling instead of panic in your app
        }
        defer stmtIns.Close()
        res, err := stmtIns.Exec(json.Name, json.Age) 
        id, err := res.LastInsertId()
        body["id"] = id
        if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
    } else {
        stmtIns, err := db.Prepare(" update user set `name` = ? , `age` = ? where `id` = ? ") // ? = placeholder
        if err != nil {
            panic(err.Error()) // proper error handling instead of panic in your app
        }
        defer stmtIns.Close()
        res, err := stmtIns.Exec(json.Name, json.Age, json.Id) 
        count, err := res.RowsAffected()
        body["updateCount"] = count
        if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		} 
    }
    
    body["status"] = "success"
    
    return body;
    
}



func (user UserType) Delete(db *sql.DB, userId int) gin.H{
    //return gin.H{"id":json.Id,"name":json.Name, "age": json.Age}
    body := make(gin.H)
    stmtIns, err := db.Prepare(" delete from user where `id` = ? ") // ? = placeholder
    if err != nil {
        panic(err.Error()) // proper error handling instead of panic in your app
    }
    defer stmtIns.Close()
    res, err := stmtIns.Exec(userId) 
    count, err := res.RowsAffected()
    body["deleteCount"] = count
    if err != nil {
        panic(err.Error()) // proper error handling instead of panic in your app
    } 
    body["status"] = "success"
    
    return body;
    
}

