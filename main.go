package main

import (
	"GoWeb/core"
	"GoWeb/global"
	"GoWeb/routers"
)

func main() {
	// 读取配置文件
	core.InitConf()
	//初始化日志
	global.Log = core.InitLogger()
	// 连接数据库
	global.DB = core.InitGorm()
	//连接网页路由
	router := routers.InitRouter()
	addr := global.Config.System.Addr()
	global.Log.Infoln("go_web 运行在 ", addr)
	_ = router.Run(addr)
}
