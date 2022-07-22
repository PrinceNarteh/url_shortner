package model

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type URL struct {
	ID       uint64 `json:"id" gorm:"primaryKey"`
	Redirect string `json:"redirect" gorm:"not null"`
	Link     string `json:"link" gorm:"unique;not null"`
	Clicked  uint64 `json:"clicked"`
	Random   bool   `json:"random"`
}

func Setup() {
	var err error
	dns := "host=172.17.0.2 user=postgres password=root dbname=shortener sslmode=disable"

	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&URL{})
	if err != nil {
		log.Println(err)
	}
}
