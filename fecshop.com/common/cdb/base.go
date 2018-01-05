package cdb

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

