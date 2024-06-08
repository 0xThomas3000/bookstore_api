package main

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/0xThomas3000/bookstore_api/component/appcontext"
	"github.com/0xThomas3000/bookstore_api/features/upload/transport/ginupload"
	"github.com/0xThomas3000/bookstore_api/middleware"
	"github.com/0xThomas3000/bookstore_api/util"
	"github.com/gin-gonic/gin"
)

func main() {
	config, err := util.ReadEnvConfig(".")
	if err != nil {
		log.Fatal("unable to read config file, ", err)
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.DBUserName,
		config.DBPassword,
		config.DBHost,
		config.DBPort,
		config.DBName,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	db = db.Debug()

	appContext := appcontext.NewAppContext(db)

	request := gin.Default()
	request.Use(middleware.Recover(appContext))

	// Register a link for '/static' to display the image
	request.Static("/static", "./static")

	v1 := request.Group("/v1")
	v1.POST("/upload", ginupload.UploadImage(appContext))

	request.Run()
}
