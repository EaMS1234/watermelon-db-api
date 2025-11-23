package banco

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func Banco() *gorm.DB {
	db, err := gorm.Open(mysql.Open("watermelon:watermelon@tcp(127.0.0.1:3306)/watermelon_db"), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
        SingularTable: true,
				NoLowerCase: true,
    },
	})
	if err != nil {panic(err)}

	return db
}

