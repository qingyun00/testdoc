// db/db.go
package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

// InitDB 初始化数据库连接
func InitDB() {
	var err error
	db, err = gorm.Open("mysql", "username:password@tcp(localhost:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic("无法连接数据库")
	}

	// 设置数据库连接池
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)

	// 自动迁移数据模型
	db.AutoMigrate(&User{}, &Todo{})
}

// CloseDB 关闭数据库连接
func CloseDB() {
	if db != nil {
		db.Close()
	}
}

// GetDB 获取数据库连接实例
func GetDB() *gorm.DB {
	return db
}
