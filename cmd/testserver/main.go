package main

import (
	"bcbtest/bcblog"
	"bcbtest/config"
	"bcbtest/dao"
	"bcbtest/parse"
	"bcbtest/webapi"
)

var (
	err error
)

// @title BCBTest API Docs
// @version  v1
// @description BCBTest API解释说明文档
func main() {
	err = config.LoadConfig("../conf/config.yaml")
	if err != nil {
		panic(err)
	}
	// 初始化数据库
	dao.Init()

	// 初始化日志
	bcblog.Init()

	// 初始化testKind对象
	parse.Init()

	// 获取TestKind对象
	testKinds := parse.GetTestKinds()

	// 将TestKind对象映射到数据库中
	dao.LoadTestKinds(*testKinds)

	// 启动路由
	router := webapi.InitRouter()
	err = router.Run(":8089")
	if err != nil {
		panic(err)
	}
}
