package main

import (
	"GoWeb/core"
	"GoWeb/global"
)

func main() {
	// 读取配置文件
	core.InitConf()
	//初始化日志
	global.Log = core.InitLogger()
	global.Log.Warnln("111")
	global.Log.Errorln("111")
	global.Log.Infoln("111")
	// 连接数据库
	global.DB = core.InitGorm()
}
