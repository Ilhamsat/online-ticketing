package routers

import (
	"github.com/gin-gonic/gin"
	"online-ticketing/interfaces/controllers"
)

func SetupSiteRouter(router *gin.RouterGroup, controller *controllers.SiteController) {
	r := router.Group("/site")
	r.POST("", controller.PostSite)
	r.GET("", controller.GetSites)
	// r.PATCH("/:id", controller.PatchSite)
	//r.DELETE("/:id", controller.DeleteSite)
}
