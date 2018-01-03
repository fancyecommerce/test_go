package main

import(
    "github.com/gin-gonic/gin"
    "net/http"
    f_customer "fecshop.com/customer"
)
func main() {
	r := gin.Default()
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
    
    r.GET("/customer/account/login", f_customer.AccountLogin)
    r.GET("/customer/account/register", f_customer.AccountRegister)
    
	r.Run("120.24.37.249:3000") // listen and serve on 0.0.0.0:8080
}














