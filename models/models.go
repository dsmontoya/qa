package models

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func Connect() error {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_HOST"), os.Getenv("MYSQL_DATABASE"))
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{PrepareStmt: true})
	return err
}

func Close() error {
	s, err := db.DB()
	if err != nil {
		return err
	}
	return s.Close()
}

func Create(value interface{}) error {
	return db.Create(value).Error
}

func First(dest interface{}, conds ...interface{}) error {
	return db.First(dest, conds...).Error
}

func Migrate() error {
	return db.Migrator().AutoMigrate(&Post{})
}

func Where(query interface{}, args ...interface{}) *gorm.DB {
	return db.Where(query, args...)
}
