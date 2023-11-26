package main

import (
	"github.com/gin-gonic/gin"
	"github.com/whywehere/smart-city-backend/config/db"
	"github.com/whywehere/smart-city-backend/web/handlers"
	"github.com/whywehere/smart-city-backend/web/middleware"
)

func main() {
	if err := db.InitMysql(); err != nil {
		panic(err.Error())
	}
	r := gin.Default()
	r.Use(middleware.CORSMiddleware())
	r.POST("/user/register", handlers.RegisterHandler)
	r.POST("/user/login", handlers.LoginHandler)

	v1 := r.Group("/cityDetail")
	v1.Use(middleware.JWTAuthMiddleware())
	{
		//v1.POST("/", handlers.)
		v1.POST("/transportation", handlers.TransportationHandler)
		v1.POST("/fireEvents", handlers.FireEventHandler)
		v1.POST("/crimeEvents", handlers.CrimeEventHandler)
		v1.POST("/transportation/week", handlers.WeekTransportationHandler)
		v1.POST("/people", handlers.PopulationHandler)
		v1.POST("/source", handlers.SourceHandler)
	}

	v2 := r.Group("/home")
	v2.Use(middleware.JWTAuthMiddleware())
	{
		//查询上海市近五年来四种问题 总数
		v2.GET("/selectAllProblems", handlers.SelectAllProblemsHandler)

		//按照年份、问题类型查询各年四种问题的总数
		v2.GET("/selectByYear", handlers.SelectProblemsNumByYearHandler)

		//按照区域查询问题的总数
		v2.GET("/selectAllByArea", handlers.SelectAllByAreaHandler)

		//查询空气质量
		v2.GET("/selectAirQ", handlers.SelectAirQualityHandler)

		//查询各个区域资源使用情况
		v2.GET("/selectSource", handlers.SelectSourceHandler)

		//查询上海市人口及增长率
		v2.GET("/selectPersons", handlers.SelectPersons)

	}

	if err := r.Run(":8080"); err != nil {
		panic(err.Error())
	}
}
