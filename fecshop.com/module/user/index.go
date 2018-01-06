package user

import (    
    _ "fmt"
    "github.com/gin-gonic/gin"
    "net/http"
    "fecshop.com/common/cdb" 
    "fecshop.com/model" 
    "database/sql"
    "strconv"
)

func List(c *gin.Context) {
    // 数据库连接
    db, err := sql.Open("mysql", cdb.Mysql.MysqlDSN())
    if err != nil {  
        panic(err.Error())  
    } 
    defer db.Close()
    err = db.Ping()  
    if err != nil {  
        panic(err.Error())  
    } 
    // 进行数据库操作
    data := model.User.Coll(db);
    c.JSON(http.StatusOK, data)
}



func AddOne(c *gin.Context){
    // 数据库连接
    db, err := sql.Open("mysql", cdb.Mysql.MysqlDSN())
    if err != nil {  
        panic(err.Error())  
    } 
    defer db.Close()
    err = db.Ping()  
    if err != nil {  
        panic(err.Error())  
    } 
    // 保存
    var json model.UserType
    if err := c.ShouldBindJSON(&json); err == nil {
        // 进行数据库操作
        //c.JSON(http.StatusOK, gin.H{"id":json.Id,"name":json.Name, "age": json.Age})
        json.Id = 0
        data := model.User.Save(db,&json);
        c.JSON(http.StatusOK, data)
    } else {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    }
}



func UpdateById(c *gin.Context){
    userId, err := strconv.Atoi(c.Param("id"))
    if userId == 0 {
        panic("userId can not empty")  
    }
    // 数据库连接
    db, err := sql.Open("mysql", cdb.Mysql.MysqlDSN())
    if err != nil {  
        panic(err.Error())  
    } 
    defer db.Close()
    err = db.Ping()  
    if err != nil {  
        panic(err.Error())  
    } 
    // 保存
    var json model.UserType
    if err := c.ShouldBindJSON(&json); err == nil {
        // 进行数据库操作
        //c.JSON(http.StatusOK, gin.H{"id":json.Id,"name":json.Name, "age": json.Age})
        json.Id = userId;
        data := model.User.Save(db,&json);
        c.JSON(http.StatusOK, data)
    } else {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    }
}




func DeleteById(c *gin.Context){
    userId, err := strconv.Atoi(c.Param("id"))
    if userId == 0 {
        panic("userId can not empty")  
    }
    // 数据库连接
    db, err := sql.Open("mysql", cdb.Mysql.MysqlDSN())
    if err != nil {  
        panic(err.Error())  
    } 
    defer db.Close()
    err = db.Ping()  
    if err != nil {  
        panic(err.Error())  
    } 
    // 删除
    data := model.User.Delete(db,userId);
    c.JSON(http.StatusOK, data)
}