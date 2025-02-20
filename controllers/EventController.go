package controllers

import (
	"aheadPMP/services"
	"aheadPMP/utils"
	"github.com/gin-gonic/gin"
	"net/http"
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
	expensesResult, incomeResult := services.SearchForEventData(query)
	utils.Success(c, map[string]interface{}{"code": int(utils.ApiCode.SUCCESS), "expenses": expensesResult, "income": incomeResult})
	return
}

func (e EventController) DownloadExcel(c *gin.Context) {
	file, err := services.ExportExcel()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 设置响应头，以便用户可以下载 Excel 文件
	c.Header("Content-Disposition", "attachment; filename=PMP账户余额汇总.xlsx")
	c.Header("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")

	// 将 Excel 文件写入响应
	err = file.Write(c.Writer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}
