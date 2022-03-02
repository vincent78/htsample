package global

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// InitDB 连接数据库
func InitDB(host string, port int, user string, pwd string, db string) error {
	dsnStr := fmt.Sprintf("host=%s user=%s password=%s DB.name=%s port=%d sslmode=disable TimeZone=Asia/Shanghai", host, user, pwd, db, port)
	va, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsnStr, // data source name, refer https://github.com/jackc/pgx
		PreferSimpleProtocol: true,   // disables implicit prepared statement usage. By default pgx automatically uses the extended protocol
	}), &gorm.Config{})
	if err == nil {
		DB = va
	}
	return err
}
