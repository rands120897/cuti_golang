package connection

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func OpenConnection() *gorm.DB {
	dsn := "host=localhost user=postgres password=postgres dbname=absent port=5432"
	db, errDb := gorm.Open(postgres.Open(dsn))
	if errDb != nil {
		fmt.Printf("errorr")
	}

	return db
}
