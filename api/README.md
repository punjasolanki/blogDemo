# Go Demo App

GoLang Restful Api With Gin Gorm And Mysql 


## Installation of Gin

In the project directory, you can run:

### `go mod init gin.com/gin`

## To install Gin in your project, run the below command from your project directory folder.

### `go get -u github.com/gin-gonic/gin`


## Database setup


``` 
 
package models
import (
	"github.com/jinzhu/gorm"
	"github.com/jinzhu/gorm/dialects/mysql"
)
var DB *gorm.DB
func ConnectDatabase() {
	database, err := gorm.Open("mysql", "root:Tce#1234@tcp(127.0.0.1:3306)/blog_management?parseTime=true")
	if err != nil {
		panic("Failed to connect to database!")
	}
	database.AutoMigrate(&Article{})
	database.AutoMigrate(&Comment{}).AddForeignKey("article_id", "articles(id)", "RESTRICT", "RESTRICT")
	DB = database
}
```

## Run Go Project

Go to your project directory and run below command.

### `go run ./`


## Run TestCase

Go to your project directory and run below command.

### `go test`
