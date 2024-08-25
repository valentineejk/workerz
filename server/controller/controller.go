package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/valentineejk/workerz/services"
)

type WorkerzController struct {
	service *services.WorkerzService
}

func NewWorkerzController(service *services.WorkerzService) *WorkerzController {
	return &WorkerzController{service: service}
}

func (cs *WorkerzController) GetData(c *gin.Context) {

	c.JSON(200, gin.H{
		"status": "mike",
	})

}
