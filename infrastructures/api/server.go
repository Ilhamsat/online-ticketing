package api

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go.elastic.co/apm/module/apmgin/v2"
	"net/http"
	"online-ticketing/app/config"
	"online-ticketing/app/global"
	"online-ticketing/app/services"
	"online-ticketing/app/validator"
	"online-ticketing/database/mysql"
	"online-ticketing/infrastructures/middleware"
	"online-ticketing/infrastructures/persistence"
	"online-ticketing/interfaces/controllers"
	"online-ticketing/interfaces/routers"
)

func Serve() {
	dbConn := mysql.NewConnection()

	r := gin.Default()
	r.Use(apmgin.Middleware(r))
	r.Use(middleware.CORSMiddleware())
	//r.Use(middleware.RecoveryMiddleware())

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"host":    c.Request.Host,
			"message": config.App.AppName,
		})
	})

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    http.StatusNotFound,
			"host":    c.Request.Host,
			"path":    c.Request.URL.Path,
			"message": "Page not found",
		})
	})

	apiV1 := r.Group("/api/v1")

	customValidator := validator.NewCustomValidator()

	sitePersistence := persistence.NewSitePersistence(dbConn.Db)

	siteService := services.NewSiteService(sitePersistence)

	siteController := controllers.NewSiteController(*siteService, *customValidator)

	routers.SetupSiteRouter(apiV1, siteController)

	err := r.Run(":" + config.App.AppPort)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"app_version": global.BuildVersion,
			"app_port":    config.App.AppPort,
			"error":       err.Error(),
		})
	}
}
