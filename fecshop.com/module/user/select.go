package user

import (    
    _ "fmt"
    "github.com/gin-gonic/gin"
    "net/http"
    "fecshop.com/common/cdb" 
    "fecshop.com/model" 
    "database/sql"
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





