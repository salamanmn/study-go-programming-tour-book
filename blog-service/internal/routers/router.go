package routers

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-programming-tour-book/blog-service/docs"
	"github.com/go-programming-tour-book/blog-service/global"
	"github.com/go-programming-tour-book/blog-service/internal/middleware"
	"github.com/go-programming-tour-book/blog-service/internal/routers/api"
	"github.com/go-programming-tour-book/blog-service/internal/routers/api/v1"
	"github.com/go-programming-tour-book/blog-service/pkg/limiter"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
	"time"
)

var methodLimiters = limiter.NewMethodLimiter().AddBuckets(
	limiter.LimiterBucketRule{
		Key:          "/auth",
		FillInterval: time.Second,
		Capacity:     10,
		Quantum:      10,
	},
)

func NewRouter() *gin.Engine {
	router := gin.New()

	if global.ServerSetting.RunMode == "debug" {
		router.Use(gin.Logger())
		router.Use(gin.Recovery())
	} else {
		router.Use(middleware.AccessLog())
		router.Use(middleware.Recovery())
	}

	router.Use(middleware.RateLimiter(methodLimiters))
	router.Use(middleware.ContextTimeout(global.AppSetting.DefaultContextTimeout))
	router.Use(middleware.Translations())

	article := v1.NewArticle()
	tag := v1.NewTag()
	upload := api.NewUpload()
	auth := api.NewAuthority()
	router.POST("/auth", auth.GetAuth)
	router.POST("/upload/file", upload.UploadFile)
	router.StaticFS("/static", http.Dir(global.AppSetting.UploadSavePath))
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// 简单的路由组: v1
	apiv1 := router.Group("/api/v1")
	apiv1.Use(middleware.JWT())
	{
		apiv1.POST("/tags", tag.Create)
		apiv1.DELETE("/tags/:id", tag.Delete)
		apiv1.PUT("/tags/:id", tag.Update)
		apiv1.PATCH("/tags/:id/state", tag.Update)
		apiv1.GET("/tags", tag.List)

		apiv1.POST("/articles", article.Create)
		apiv1.DELETE("/articles/:id", article.Delete)
		apiv1.PUT("/articles/:id", article.Update)
		apiv1.PATCH("/articles/:id/state", article.Update)
		apiv1.GET("/articles", article.List)
		apiv1.GET("/articles/:id", article.Get)
	}
	return router
}
