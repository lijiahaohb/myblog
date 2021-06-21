package main

import (
	"ginblog/model"
	"ginblog/routers"
)

func main() {
	// 连接数据库
	model.InitDb()

	routers.InitRouter()
}
