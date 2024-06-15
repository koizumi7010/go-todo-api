package mysql

import (
	"fmt"
	"os"

	"github.com/go-sql-driver/mysql"
	gormmysql "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// NewDB is a function to connect to the database.
func NewDB() (*gorm.DB, error) {
	// DBの設定
	c := mysql.Config{
		DBName:               os.Getenv("DATABASE_NAME"),
		User:                 os.Getenv("DATABASE_USER"),
		Passwd:               os.Getenv("DATABASE_PASSWORD"),
		Addr:                 os.Getenv("DATABASE_HOST") + ":" + os.Getenv("DATABASE_PORT"),
		Net:                  "tcp",
		ParseTime:            true,
		Collation:            "utf8mb4_general_ci",
		AllowNativePasswords: true,
	}
	// DBに接続
	db, err := gorm.Open(gormmysql.Open(c.FormatDSN()), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	return db, nil
}
