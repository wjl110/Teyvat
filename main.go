package main

import (
	"douying/dao"
	"douying/setting"
	"fmt"
	"os"
)

func main() {
	/**
	需要在运行参数中（Program arguments）添加 conf/config.ini ，可以根据不同参数执行不同环境下的配置文件
	 */
	if len(os.Args) < 2 {
		fmt.Println("Usage：./douyin conf/config.ini")
		return
	}
	// 加载配置文件
	if err := setting.Init(os.Args[1]);err != nil {
		fmt.Printf("load config from file failed, err:%v\n", err)
		return
	}
	//初始化数据源
	if  err := dao.InitDB(setting.Conf.MySQLConfig);err != nil {
		fmt.Printf("init mysql failed, err:%v\n", err)
		return
	}
	// 初始化redis
	dao.InitRedis(setting.Conf.RedisConfig)
	// 初始化路由
	router := routerInit()
	if err := router.Run(fmt.Sprintf(":%d", setting.Conf.Port)); err != nil {
		fmt.Printf("server startup failed, err:%v\n", err)
	}

}
