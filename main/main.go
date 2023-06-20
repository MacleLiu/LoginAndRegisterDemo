package main

import (
	"LoginAndRegisterDemo/main/middlewares"
	"LoginAndRegisterDemo/main/routers"
	"LoginAndRegisterDemo/main/utils"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	config, err := utils.UnmarshalConfig("./conf/config.toml")
	if err != nil {
		log.Fatalf("解析配置文件错误：%v", err)
	}

	// 创建数据库连接
	//db, err := mysql.Open()
	mysql, ok := config.MysqlConf("business")
	if !ok {
		log.Fatal("获取mysql配置失败")
	}
	db, err := mysql.Open()
	if err != nil {
		log.Fatalf("打开数据库连接错误：%v", err)
	}

	// 创建一个默认路由引擎，其中已经附加了日志记录器和恢复中间件。
	r := gin.Default()

	// 设置中间件
	r.Use(middlewares.Set("DB", db))

	// 加载 HTML 模板文件。
	r.LoadHTMLGlob("templates/*")

	// 默认路由
	routers.DefaultRouterInit(r)

	// 用户路由
	routers.UserRouterInit(r)

	// 监听并启动服务
	r.Run(config.ListenAddr())

}
