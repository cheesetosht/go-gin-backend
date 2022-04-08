package models

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Book struct {
	ID     uint   `json:"id" gorm:"primary_key"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

var DB *gorm.DB

func ConnectDatabase() {
  dsn := "host=john.db.elephantsql.com user=gqpmmuvh password=G9qVWhK34iHyPdPydWgqRNPn5aTc7eWv dbname=gqpmmuvh port=5432 sslmode=disable TimeZone=Asia/Kolkata"
  db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})


  if err != nil {
    panic("Failed to connect to database!")
  }

  db.AutoMigrate(&Book{})

  DB = db
}