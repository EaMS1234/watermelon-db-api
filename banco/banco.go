package banco

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func Banco() *gorm.DB {
	db, err := gorm.Open(mysql.Open("watermelon:watermelon@tcp(db:3306)/watermelon_db"), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
        SingularTable: true,
				NoLowerCase: true,
    },
	})
	if err != nil {panic(err)}

	return db
}

