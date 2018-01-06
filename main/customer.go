package main

import(
    "github.com/gin-gonic/gin"
    // "net/http"
    "fecshop.com/module/user" 
    _ "fmt"  
    _ "github.com/go-sql-driver/mysql" 
)


func main() { 
	r := gin.Default()
    v1 := r.Group("/v1")
    {
        // 查询部分
        v1.GET("/users", user.List)
        v1.POST("/users", user.AddOne)
        v1.PATCH("/users/:id", user.UpdateById)
        v1.DELETE("/users/:id", user.DeleteById)
    }
    /*
    //r.GET("/customer/account/login", f_customer.AccountLogin)
    //r.GET("/customer/account/register", f_customer.AccountRegister)
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
    r.Run("120.24.37.249:3000") // listen and serve on 0.0.0.0:8080
    
}






// http://blog.csdn.net/rambo_huang/article/details/60604924









