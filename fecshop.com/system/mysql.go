package system

import (    
    _ "fmt"
)

var database    = "go_test"
var user        = "root"
var pass        = "Zhaoyong2017fdsfds3f3GDs3fgsd"
var ip          = "127.0.0.1"
var port        = "3306"

func MysqlDSN() string{
    return user + ":" + pass + "@tcp(" + ip + ":" + port + ")/" + database;
}










