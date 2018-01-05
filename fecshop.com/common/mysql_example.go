package common

import (    
    _ "fmt"
    "database/sql" 
)

var Mysql Database

func init(){
    Mysql.dbName      = "go_test"
    Mysql.user        = "root"
    Mysql.pass        = "xxx"
    Mysql.ip          = "127.0.0.1"
    Mysql.port        = "3306"
}

func (db Database) MysqlDSN() string{
    return db.user + ":" + db.pass + "@tcp(" + db.ip + ":" + db.port + ")/" + db.dbName;
}


