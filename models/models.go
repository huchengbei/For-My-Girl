package models

import (
	"errors"
	"fmt"
	"github.com/huchengbei/for-my-girl/pkg/logging"
	"github.com/huchengbei/for-my-girl/pkg/setting"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"os"
)

var db *gorm.DB

func mustString(value interface{}, defaultValue string) string {
	newV, ok := value.(string)
	if ok {
		return newV
	} else {
		return defaultValue
	}
}

func init() {
	var (
		err    error
		dbType string
	)

	database := setting.CfgMap["database"].(map[string]interface{})
	dbType = database["TYPE"].(string)
	if dbType == "mysql" {
		mysqlCfg := database["MYSQL"].(map[string]interface{})
		var (
			err                               error
			user, password, host, tablePrefix string
		)
		dbName, ok := mysqlCfg["NAME"].(string)
		if !ok {
			logging.Error("dbName of mysql is invalid in config")
		}
		user, ok = mysqlCfg["USER"].(string)
		if !ok {
			logging.Error("username of mysql is invalid in config")
		}
		password, ok = mysqlCfg["PASSWORD"].(string)
		if !ok {
			logging.Error("password of mysql is invalid in config")
		}
		host, ok = mysqlCfg["HOST"].(string)
		if !ok {
			logging.Error("host of mysql is invalid in config")
		}
		tablePrefix = mustString(mysqlCfg["TABLE_PREFIX"], "")

		dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			user,
			password,
			host,
			dbName)

		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				TablePrefix: tablePrefix,
			},
		})
		if err != nil {
			logging.Error("Cant Open database, Err: %v", err)
		}
	} else if dbType == "sqlite" {
		var (
			dbName, path, tablePrefix string
		)
		sqliteCfg := database["SQLITE"].(map[string]interface{})
		dbName, ok := sqliteCfg["NAME"].(string)
		if !ok {
			logging.Error("dbName of sqlite is invalid in config")
		}
		path, ok = sqliteCfg["PATH"].(string)
		if !ok {
			logging.Error("path of sqlite is invalid in config")
		}
		tablePrefix = mustString(sqliteCfg["TABLE_PREFIX"], "")

		_, err := os.Stat(path)
		if err != nil {
			logging.Error("db is not exist, Err: %v", err)
			fmt.Printf("file not exist")
		}

		db, err = gorm.Open(sqlite.Open(path+"?parseTime=True&loc=Local"), &gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				TablePrefix: tablePrefix,
			},
		})
		fmt.Printf("dbName: %s\n", dbName)
		if err != nil {
			logging.Error("Cant Open database, Err: %v", err)
		}
	} else {
		err = errors.New("dbType is not mysql or sqlite")
	}

	if err != nil {
		logging.Warn(err)
	}
	db.AutoMigrate(&Moment{})
	db.AutoMigrate(&Slide{})

}
