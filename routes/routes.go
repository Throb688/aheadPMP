package routes

import (
	"aheadPMP/controllers"
	middleware "aheadPMP/middlewares"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.CORSMiddleware())

	r.GET("/getEventData", controllers.NewEventControllerr().EventData)
	r.GET("/findEventData", controllers.NewEventControllerr().SearchForEvent)
	r.GET("/download", controllers.NewEventControllerr().DownloadExcel)

	return r
}
