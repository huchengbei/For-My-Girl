package models

import (
	"errors"
	"fmt"
	"github.com/huchengbei/for-my-girl/backend/pkg/logging"
	"github.com/huchengbei/for-my-girl/backend/pkg/setting"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
)

var db *gorm.DB

func init()  {
	var (
		err error
		dbType, dbName, user, password, host, tablePrefix string
	)

	sec, err := setting.Cfg.GetSection("database")
	if err != nil {
		log.Fatal(2, "Fail to get section 'database': %v", err)
	}

	dbType = sec.Key("TYPE").String()
	dbName = sec.Key("NAME").String()
	user = sec.Key("USER").String()
	password = sec.Key("PASSWORD").String()
	host = sec.Key("HOST").String()
	tablePrefix = sec.Key("TABLE_PREFIX").String()

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user,
		password,
		host,
		dbName)
	if dbType == "mysql" {
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				TablePrefix: tablePrefix,
			},
		})
	} else {
		err = errors.New("dbType is not mysql")
	}


	if err != nil {
		logging.Warn(err)
	}

	//gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
	//	return tablePrefix + defaultTableName
	//}

	//db.SingularTable(true)
	//db.LogMode(true)
	//db.DB().SetMaxIdleConns(10)
	//db.DB().setMaxOpenConns(100)
}

