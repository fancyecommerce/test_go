package customer

import (  
    "fmt"
    "github.com/gin-gonic/gin"
    "net/http"
    
    f_account "fecshop.com/customer/account"
)

func AccountRegister(c *gin.Context) {
    //message := "customer account register success"
    body := f_account.Register();
    c.String(http.StatusOK, body)
    fmt.Println("customer register account success");
}

func AccountLogin(c *gin.Context) {
    email       := c.DefaultQuery("email", "Guest")
    password    := c.Query("password") // shortcut for c.Request.URL.Query().Get("lastname")
    body        := f_account.Login(email, password);
    c.String(http.StatusOK, body)
    fmt.Println("customer login account success");
}

func AccountIndex(c *gin.Context) {
    body := "welcome to account index page";
    c.String(http.StatusOK, body)
    fmt.Println("welcome to account index page");
}













