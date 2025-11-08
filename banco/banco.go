package banco

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Banco() *gorm.DB {
	db, err := gorm.Open(mysql.Open("root:root@tcp(127.0.0.1:3306)/watermelon_db"), &gorm.Config{})
	if err != nil {panic(err)}

	return db
}

