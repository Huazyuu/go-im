package main

import (
	"flag"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	chat "server/yu_chat/models"
	group "server/yu_group/models"
	user "server/yu_user/models"
	"time"
)

type Options struct {
	DB bool
}

func main() {
	var opt Options
	flag.BoolVar(&opt.DB, "db", false, "db")
	flag.Parse()

	if opt.DB {
		db := initMysql()
		err := db.AutoMigrate(
			user.UserModel{},
			user.UserConfModel{},
			user.FriendVerifyModel{},
			user.FriendsModel{},

			chat.ChatModel{},

			group.GroupModel{},
			group.GroupMsgModel{},
			group.GroupVerifyModel{},
			group.ChatMemberModel{},
		)
		if err != nil {
			fmt.Println("表结构生成失败", err)
			return
		}
		fmt.Println("表结构生成成功！")

	}
}
func initMysql() *gorm.DB {
	dsn := "root:123456@tcp(127.0.0.1:3307)/fim_server_db?charset=utf8mb4&parseTime=True&loc=Local"
	var mysqlLogger logger.Interface
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: mysqlLogger,
	})
	if err != nil {
		log.Fatalf(fmt.Sprintf("[%s] mysql连接失败", dsn))
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)               // 最大空闲连接数
	sqlDB.SetMaxOpenConns(100)              // 最多可容纳
	sqlDB.SetConnMaxLifetime(time.Hour * 4) // 连接最大复用时间，不能超过mysql的wait_timeout
	return db
}
