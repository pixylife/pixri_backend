package config

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

//"root:root@/pws?charset=utf8&parseTime=True&loc=Local"


type Env struct {
	DB *gorm.DB
}

func NewMySQLdb(dataSourceName string) (*gorm.DB, error) {

	const dbType = "mysql"

	db, err := gorm.Open(dbType, dataSourceName)

	if err != nil {
		fmt.Println(err)
	}

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	db.DB().SetMaxIdleConns(10)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	db.DB().SetMaxOpenConns(100)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	db.DB().SetConnMaxLifetime(time.Hour)

	return db, nil
}



