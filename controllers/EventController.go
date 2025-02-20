package controllers

import (
	"aheadPMP/services"
	"aheadPMP/utils"
	"github.com/gin-gonic/gin"
)

type EventController struct{}

func NewEventControllerr() EventController {
	return EventController{}
}

func (e EventController) EventData(c *gin.Context) {
	results := services.GetEventData()
	utils.Success(c, map[string]interface{}{"code": int(utils.ApiCode.SUCCESS), "msg": results})
	return
}

func (e EventController) SearchForEvent(c *gin.Context) {
	query := c.Query("q")
	result := services.SearchForEventData(query)
	utils.Success(c, map[string]interface{}{"code": int(utils.ApiCode.SUCCESS), "msg": result})
	return
}
