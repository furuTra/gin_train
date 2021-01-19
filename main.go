package main

import (
	"github.com/furuTra/gin_train/db"
	"github.com/furuTra/gin_train/server"
)

func main() {
    // DB接続
    db.Init()
    // Router初期化
    server.Init()
    db.CloseDB()
}