package common


type Database struct{
    dbName  string
    user    string
    pass    string
    ip      string
    port    string
    OpenDb  *sql.DB
}

