package routes

import (
	"cld/controller"
	_ "cld/docs"
	"cld/logger"
	"cld/middlewares"
	"cld/settings"
	"fmt"
	"io/ioutil"

	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var r *gin.Engine

func Setup(cfg *settings.AppConfig) {
	r = gin.Default()
	//翻译器初始化
	controller.InitTrans("zh")

	//日志写入中间件
	r.Use(logger.GinLogger(), logger.GinRecovery(false))
	//接口文档UI
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	baseapi := r.Group("/api/v1")
	baseapi.GET("/log", showLog)

	authen := baseapi.Group("/auth")
	{
		authen.POST("/login", controller.LoginHandler)
		authen.POST("/signup", controller.SignUpHandler)
		authen.POST("/sendemail", controller.SendEmailHandler)
		authen.POST("/resetpass", controller.ReSetPassHandler)
		authen.POST("/recoverpass", controller.ReCoverPassHandler)
	}

	edu := baseapi.Group("/edu")
	//加入中间件验证
	edu.Use(middlewares.JWTAuthMiddleware())
	{
		edu.POST("/bind", controller.BingEducationalHandler)

		edu.GET("/cookie", controller.CookieHandler)
		edu.GET("/courses", controller.CourseHandler)
		edu.GET("/grades", controller.GradesHandler)
	}

	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, cfg.Name)
	})
}

func StartServer(p int) {
	r.Run(fmt.Sprintf(":%s", strconv.Itoa(p)))
}

func showLog(c *gin.Context) {
	data, _ := ioutil.ReadFile("/home/hackerxiao/go_project/sylu_project/web_app.log")
	c.String(200, string(data))
}
