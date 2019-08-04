package conf

import (
	"lcb-go/cache"
	"fmt"
	"github.com/joho/godotenv"
	"lcb-go/excel"
)

func Init() {
	// 本地读取env
	godotenv.Load()
	// 创建Redis连接池
	cache.Redis()
	// 执行导入excel数据
	excel.Import()

	fmt.Println("init success")
}
