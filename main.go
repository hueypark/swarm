package main

import (
	"log"

	"github.com/hueypark/swarm/benchmark"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	db, err := gorm.Open("postgres", "host=localhost port=26257 user=root dbname=cockroach sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}
	defer func() {
		err := db.Close()
		if err != nil {
			log.Println(err)
		}
	}()

	benchmark.Issue4234(db)
}
