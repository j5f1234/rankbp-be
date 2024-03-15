package router

import (
	"rankbp-be/middleware"

	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	r.Use(middleware.Error)

	apiRouter := r.Group("/api")
	{
		// begin

		apiRouter.POST("/herosdata/base", ctr.HeroBaseController.UpdateHeroBaseData)
		apiRouter.POST("/herospool/change", ctr.HeroPoolController.ChangeHeropool)
		apiRouter.GET("/herospool/get", ctr.HeroPoolController.SendHeropool)
		// end
	}
}
