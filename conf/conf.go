package conf

import (
	"lcb-go/cache"
	"fmt"
	"lcb-go/excel"
)

func Init() {
	// 创建Redis连接池
	cache.Redis()
	// 执行导入excel数据
	excel.Import()

	fmt.Println("init success")
}
