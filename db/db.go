package db

import (
	"fmt"
	"log"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
    "github.com/furuTra/gin_train/entity"
)

var (
	db *gorm.DB
	err error
)

const (
    // データベース
    Dialect = "mysql"
    // ユーザー名
    DBUser = "docker"
    // パスワード
    DBPass = "Tribu242"
    // プロトコル
    DBProtocol = "tcp(mysql:3306)"
    // DB名
    DBName = "gin_test"
)

// MySQL接続の初期化
func Init() {
	connect := fmt.Sprintf("%s:%s@%s/%s", DBUser, DBPass, DBProtocol, DBName)
	db, err = gorm.Open(Dialect, connect)

    if err != nil {
        log.Println(err.Error())
        panic(err.Error())
    }
    autoMigration()
}

func GetDB() *gorm.DB {
	return db
}

func CloseDB() {
	err = db.Close()
	if err != nil {
		panic(err.Error())
	}
}

func autoMigration() {
    db.AutoMigrate(&entity.User{})
}