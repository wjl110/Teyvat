package main

import (
	"douying/controller"
	"douying/setting"
	"github.com/gin-gonic/gin"
)

func routerInit() *gin.Engine {

	if setting.Conf.Release {
		/**
			总共三种模式，默认debug模式
		 */
		//DebugMode = "debug"
		//ReleaseMode = "release"
		//TestMode = "test"
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()
	// 基础接口
	apiRouter := router.Group("/douyin")
	{
		// 基础
		apiRouter.GET("/feed", controller.Feed)
		apiRouter.GET("/user/", controller.UserInfo)
		apiRouter.POST("/user/register/", controller.Register)
		apiRouter.POST("/user/login/", controller.Login)
		apiRouter.POST("/publish/action/", controller.Publish)
		apiRouter.GET("/publish/list/", controller.PublishList)

		// 扩展1
		apiRouter.POST("/favorite/action/", controller.FavoriteAction)
		apiRouter.GET("/favorite/list/", controller.FavoriteList)
		apiRouter.POST("/comment/action/", controller.CommentAction)
		apiRouter.GET("/comment/list/", controller.CommentList)

		// 扩展2
		apiRouter.POST("/relation/action/", controller.RelationAction)
		apiRouter.GET("/relation/follow/list/", controller.FollowList)
		apiRouter.GET("/relation/follower/list/", controller.FollowerList)

	}
	return router
}
