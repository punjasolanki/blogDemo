package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open("mysql", "root:Tce#1234@tcp(192.168.1.68:3306)/blog_management?parseTime=true")

	if err != nil {
		panic("Failed to connect to database!")
	}

	//Migrate the table
	database.AutoMigrate(&Article{})
	database.AutoMigrate(&Comment{}).AddForeignKey("article_id", "articles(id)", "RESTRICT", "RESTRICT")

	// Seed the data
	seedData(database)

	DB = database
}

func seedData(database *gorm.DB) {
	if database.HasTable(&Article{}) {
		var article Article
		if err := database.First(&article).Error; err == gorm.ErrRecordNotFound {
			articles := []Article{
				{Title: "GO", Author: "Google", Content: "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industrys standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum."},
				{Title: "React", Author: "Facebook", Content: "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industrys standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum."},
			}

			for _, article := range articles {
				database.Create(&article)
			}
		}
	}
}
