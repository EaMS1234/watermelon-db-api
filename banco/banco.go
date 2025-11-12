package banco

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Banco() *gorm.DB {
	db, err := gorm.Open(mysql.Open("watermelon:watermelon@tcp(db:3306)/watermelon_db"), &gorm.Config{})
	if err != nil {panic(err)}

	return db
}

