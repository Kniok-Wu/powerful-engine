package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// NewSQLiteConn 新建立一个sqlite3的数据库链接
func NewSQLiteConn(dbPath string) (db *gorm.DB, err error) {
	if db, err = gorm.Open(sqlite.Open(dbPath), &gorm.Config{}); err != nil {
		return nil, err
	}
	return db, nil
}

// NewMySQLConn 新建立一个sqlite3的数据库链接
func NewMySQLConn(dbPath string) (db *gorm.DB, err error) {
	if db, err = gorm.Open(sqlite.Open(dbPath), &gorm.Config{}); err != nil {
		return nil, err
	}
	return db, nil
}
