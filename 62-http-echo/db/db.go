package db

import (
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	RETRY_TIMES    = 10
	RETRY_DURATION = time.Second * 5
)

func GetConnection(dsn string) (*gorm.DB, error) {
	c := 1
retry:
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		if c > RETRY_TIMES {
			return nil, err
		} else {
			log.Println("trying to connect to the databas--", c, "times")
			time.Sleep(RETRY_DURATION)
			c++
			goto retry
		}
	}
	return db, err
}

//dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
