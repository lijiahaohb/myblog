package routers

import (
	v1 "ginblog/api/v1"
	"ginblog/middleware"
	"ginblog/utils"

	"github.com/gin-gonic/gin"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(middleware.Logger())

	auth := r.Group("api/v1")
	auth.Use(middleware.JwtToken())
	{
		// 用户模块的路由接口
		auth.PUT("user/:id", v1.EditUser)

		auth.DELETE("user/:id", v1.DeleteUser)

		// 分类模块的路由接口
		auth.POST("category/add", v1.AddCategory)

		auth.PUT("category/:id", v1.EditCategory)

		auth.DELETE("category/:id", v1.DeleteCategory)

		// 文章模块的路由接口
		auth.POST("article/add", v1.AddArticle)

		auth.PUT("article/:id", v1.EditArticle)

		auth.DELETE("article/:id", v1.DeleteArticle)

		// 上传文件
		auth.POST("upload", v1.Upload)
	}

	routerV1 := r.Group("api/v1")
	{
		routerV1.GET("users", v1.GetUsers)
		routerV1.POST("user/add", v1.AddUser)

		routerV1.GET("categories", v1.GetCategory)

		/*查询单个文章*/
		routerV1.GET("article/:id", v1.GetArticleInfo)
		/*查询文章列表*/
		routerV1.GET("articles", v1.GetArticles)
		/*查询一个分类下的所有文章*/
		routerV1.GET("articles/:id", v1.GetCateArticles)

		// 登录验证
		routerV1.POST("login", v1.Login)
	}

	r.Run(utils.HttpPort)
}
