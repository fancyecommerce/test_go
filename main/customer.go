package main

import(
    "github.com/gin-gonic/gin"
    "net/http"
    f_customer "fecshop.com/customer"
    "fecshop.com/module/user" 
    _ "fmt"  
    _ "github.com/go-sql-driver/mysql" 
)


func main() { 
	r := gin.Default()
    // 查询部分
    r.GET("/user/list", user.List)
    
    /*
    r.GET("/customer/account/login", f_customer.AccountLogin)
    r.GET("/customer/account/register", f_customer.AccountRegister)
	r.Run("120.24.37.249:3000") // listen and serve on 0.0.0.0:8080
    
    r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
    r.GET("/customer", func(c *gin.Context){
        c.JSON(http.StatusOK, gin.H{
            "customer": "terry",
        })
    })
    */
}






// http://blog.csdn.net/rambo_huang/article/details/60604924









