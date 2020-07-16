package dao

import (
	"bcbtest/config"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

const DRIVER = "mysql"

var (
	db  *gorm.DB
	err error
)

func Init() {
	cfg := config.GetConfig()
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8&parseTime=True&loc=UTC", cfg.DbUser, cfg.DbPwd, cfg.DbIP, cfg.DbPort, cfg.DbName)
	db, err = gorm.Open(DRIVER, dsn)
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&TestCaseKind{})
	db.AutoMigrate(&TestSuite{})
	db.AutoMigrate(&TestCase{})
	db.AutoMigrate(&Report{})
	db.AutoMigrate(&ReportCase{})
	db.AutoMigrate(&ReportCaseHis{})
}

func GetClient() *gorm.DB {
	if db == nil {
		log.Print("must init db first")
	}
	return db
}
