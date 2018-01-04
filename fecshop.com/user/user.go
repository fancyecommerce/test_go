package user

import (    
    _ "fmt"
    "github.com/gin-gonic/gin"
    "net/http"
    f_system "fecshop.com/system" 
    "database/sql" 
)

func List(c *gin.Context) {
    db, err := sql.Open("mysql", f_system.MysqlDSN());
    if err != nil {  
        panic(err.Error())  
    } 
    defer db.Close()
    err := db.Ping()  
    if err != nil {  
        panic(err.Error())  
    } 
    data := coll(db)
    c.JSON(http.StatusOK, data)
}

func coll(db *sql.DB) gin.H{
    body := make(gin.H)
    rows, err := db.Query("SELECT * From user")  
    defer rows.Close()  
    var dbdata []gin.H
    for rows.Next() {  
        var id int  
        var name string  
        var age string
        err = rows.Scan(&id, &name ,&age) 
        //body += ",value = " + name 
        dd := gin.H{
            "id": id ,
            "name": name,
            "age": age,
        }
        dbdata = append(dbdata, dd)
    }  
    body["data"] = dbdata;
    err = rows.Err()  
    if err != nil {  
        panic(err.Error())  
    }  
    return body;
}













