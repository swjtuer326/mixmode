package common

import (
	"sophgo.com/mixmode/config"
	"sophgo.com/mixmode/model"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

// 全局数据库对象
var DB *gorm.DB

// 初始化数据库
func InitDB() {
	switch config.Conf.Database.Driver {
	case "sqlite3":
		DB = ConnSqlite()
	}
	dbAutoMigrate()
}

// 自动迁移表结构
func dbAutoMigrate() {
	_ = DB.AutoMigrate(
		&model.Api{},
		&model.OperationLog{},
		&model.PowerLog{},
		&model.TemLog{},
	)
}

func ConnSqlite() *gorm.DB {
	db, err := gorm.Open(sqlite.Open(config.Conf.Database.Source), &gorm.Config{
		// 禁用外键(指定外键时不会在mysql创建真实的外键约束)
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		Log.Panicf("failed to connect sqlite3: %v", err)
	}
	dbObj, err := db.DB()
	if err != nil {
		Log.Panicf("failed to get sqlite3 obj: %v", err)
	}
	// 参见： https://github.com/glebarez/sqlite/issues/52
	dbObj.SetMaxOpenConns(1)
	return db
}
