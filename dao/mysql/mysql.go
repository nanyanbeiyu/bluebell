package mysql

import (
	"bluebell/models"
	"bluebell/settings"
	"fmt"

	"go.uber.org/zap"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"database/sql"
)

var db *gorm.DB
var mysqldb *sql.DB

func Init(conf *settings.MysqlConfig) (err error) {
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		conf.User,
		conf.Password,
		conf.Host,
		conf.Port,
		conf.DB)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	mysqldb, _ = db.DB()
	if err != nil {
		zap.L().Error("init mysql failde,err:", zap.Error(err))
		return
	}
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	mysqldb.SetMaxIdleConns(conf.MaxIdleConns)
	// SetMaxOpenConns 设置打开数据库连接的最大数量
	mysqldb.SetMaxOpenConns(conf.MaxOpenConns)

	db.AutoMigrate(&models.User{}, &models.Community{}, &models.Post{})
	return
}

func Close() {
	_ = mysqldb.Close()
}
