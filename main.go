package main

import (
	"GinVueS/userControl"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)


func main() {
	db := userControl.InitDB()
	defer db.Close()

	router := gin.Default()
			router.POST("/api/auth/register",userControl.Register)
	panic(router.Run("127.0.0.1:9999"))
}



