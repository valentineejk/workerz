package router

import (
	"github.com/gin-gonic/gin"
	"github.com/valentineejk/workerz/controller"
)

func Router(g *gin.Engine, s *controller.WorkerzController) {

	r := g.Group("/v1")

	r.GET("/getin", s.GetData)
}
