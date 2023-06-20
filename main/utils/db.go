package utils

// db.go 是数据库初始化配置，和连接池配置
import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MysqlServer struct {
	UserName string `toml:"username"`
	Password string `toml:"password"`
	Protocol string `toml:"protocol"`
	Address  string `toml:"address"`
	DBName   string `toml:"dbname"`
}

// 返回连接 Mysql 的 dsn
func (m MysqlServer) dsn() string {
	return fmt.Sprintf("%s:%s@%s(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		m.UserName, m.Password, m.Protocol, m.Address, m.DBName)
}

// 打开数据库连接
func (m MysqlServer) Open() (db *gorm.DB, err error) {
	db, err = gorm.Open(mysql.Open(m.dsn()), &gorm.Config{})
	return
}
