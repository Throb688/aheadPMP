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
