package settings

import (
	"giligili/cache"
	"giligili/model"
	"os"

	"github.com/joho/godotenv"
)

// Init 初始化配置项
func Init() {
	// 从本地读取环境变量
	godotenv.Load()

	// 连接数据库
	model.Database(os.Getenv("MYSQL_DSN"))
	cache.Redis()
}
