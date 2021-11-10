package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/huchengbei/for-my-girl/middleware"
	"github.com/huchengbei/for-my-girl/middleware/jwt"
	"github.com/huchengbei/for-my-girl/pkg/setting"
	routersApi "github.com/huchengbei/for-my-girl/routers/api"
)

func InitRouter() *gin.Engine {
	gin.SetMode(setting.RunMode)

	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	r.Use(middleware.FrontendFileHandler())

	r.GET("/auth", routersApi.GetAuth)

	api := r.Group("/api")
	{
		api.GET("slides", routersApi.GetSlides)
		api.GET("moments", routersApi.GetMoments)
	}
	api.Use(jwt.JWT())
	{
		api.POST("slides", routersApi.AddSlide)
		api.PUT("slides", routersApi.EditSlide)
		api.DELETE("slides", routersApi.DeleteSlide)
		api.POST("moments", routersApi.AddMoment)
		api.PUT("moments", routersApi.EditMoment)
		api.DELETE("moments", routersApi.DeleteMoment)
	}

	return r
}
