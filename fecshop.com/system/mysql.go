package system

import (    
    _ "fmt"
    "database/sql" 
)

type Database struct{
    dbName  string
    user    string
    pass    string
    ip      string
    port    string
    OpenDb  *sql.DB
}

var I int

var Mysql Database

func init(){
    
    Mysql.dbName      = "go_test"
    Mysql.user        = "root"
    Mysql.pass        = ""
    Mysql.ip          = "127.0.0.1"
    Mysql.port        = "3306"
}

func (db Database) MysqlDSN() string{
    return db.user + ":" + db.pass + "@tcp(" + db.ip + ":" + db.port + ")/" + db.dbName;
}

// 打开db
func (db Database) Open() (*sql.DB){
    if db.OpenDb == nil {
        I = I + 1;
        dbHandle, err := sql.Open("mysql", db.MysqlDSN())
        if err != nil {  
            panic(err.Error())  
        } 
        //defer dbOpen.Close()
        err = dbHandle.Ping()  
        if err != nil {  
            panic(err.Error())  
        } 
        db.OpenDb = dbHandle;
    }
    return  db.OpenDb
}

// 关闭db
func (db Database) Close() {
    if db.OpenDb != nil {
        db.OpenDb.Close()
    }
}


