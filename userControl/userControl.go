package userControl

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func Register(ctx *gin.Context){
	//
	name := ctx.PostForm("name")
	telephone := ctx.PostForm("telephone")
	password := ctx.PostForm("password")


	//
	if len(password) < 6 {
		ctx.JSON(http.StatusUnprocessableEntity,gin.H{"code":422,"msg":"password must be cover 6st"})
		return
	}

	//如果不存在名字 则随机生成
	if len(name) == 0{
		name = RandomString(10)
	}

	newUser := User{
		Name: name,
		Password: password,
	}
	db.Create(&newUser)
	log.Println(name,telephone,password)
	ctx.JSON(200,gin.H{
		"msg":"success register",
	})
}

func RandomString(n int) string{
	var letters = []byte("abcdefghijklmnobqrsduvwxyz")
	result := make([]byte,n)

	rand.Seed(time.Now().Unix())
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}

func InitDB() *gorm.DB{
	driverName := "mysql"
	host := "127.0.0.1"
	port := "3306"
	database := "gvs"
	username := "root"
	password := "123456"
	charset := "utf8"
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
		username,
		password,
		host,
		port,
		database,
		charset)

	db,err := gorm.Open(driverName,args)
	if err != nil  {
		panic("failed to connection database,err :" + err.Error())
	}
	db.AutoMigrate(&User{})
	return db
}