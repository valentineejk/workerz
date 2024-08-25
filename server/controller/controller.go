package controller

import (
	"context"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
	dbq "github.com/valentineejk/workerz/database/sqlc"
	"github.com/valentineejk/workerz/model"
	"github.com/valentineejk/workerz/services"
)

type WorkerzController struct {
	service *services.WorkerzService
}

func NewWorkerzController(service *services.WorkerzService) *WorkerzController {
	return &WorkerzController{service: service}
}

func (cs *WorkerzController) AddWorkerz(c *gin.Context) {

	var ctx, cancel = context.WithTimeout(c, 10*time.Second)
	defer cancel()

	var w model.Workerz

	if err := c.ShouldBind(&w); err != nil {
		c.JSON(200, gin.H{
			"status":  false,
			"message": "invalid phone number",
			"error":   "invalid-body",
		})
		return
	}

	NID, _ := services.NanoIDS()

	addWorkersParams := dbq.InsertWorkerzParams{
		UserID:    NID,
		FullName:  w.FullName,
		Email:     w.Email,
		Salary:    w.Salary,
		Role:      w.Role,
		Country:   w.Country,
		CreatedAt: pgtype.Timestamp{Time: time.Now(), Valid: true},
		UpdatedAt: pgtype.Timestamp{Time: time.Now(), Valid: true},
	}
	err := cs.service.AddWorkerzService(ctx, addWorkersParams)

	if err != nil {
		log.Println("error adding workerz to postgres")
		return

	}
	c.JSON(200, gin.H{
		"status": true,
	})

}

func (cs *WorkerzController) GetAllWorkerz(c *gin.Context) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	resp := services.Response{}

	res, err := cs.service.GetAllWorkerz(ctx)
	if err != nil {
		resp.Status = false
		resp.Error = "Error listing workerz"
		resp.Raw = err.Error()
		c.JSON(401, resp)
		return
	}

	resp.Status = true
	resp.Data = res
	c.JSON(200, resp)

}
