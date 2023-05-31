package config

import (
	//"github.com/jinzhu/gorm"
	"fmt"
	"sync"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var lock = &sync.Mutex{}

var (
	db *gorm.DB
)

func GetDBInstance() *gorm.DB {
	if db == nil {
		lock.Lock()
		defer lock.Unlock()
		if db == nil {
			fmt.Println("Creating single instance of db.")

			dbURL := "postgres://postgres:aaditya@localhost:5432/db2"

			d, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

			if err != nil {
				panic(err)
			}
			db = d

		} else {
			fmt.Println("Single instance of db already created.")
		}
	} else {
		fmt.Println("Single instance of db already created.")
	}

	return db

}

func Hello() {
	fmt.Println("Hello")
}
