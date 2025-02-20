package routes

import (
	"aheadPMP/controllers"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()

	r.GET("/getEventData", controllers.NewEventControllerr().EventData)
	r.GET("/findEventData", controllers.NewEventControllerr().SearchForEvent)

	return r
}
