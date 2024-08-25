package services

import (
	"context"
	"fmt"
	"time"

	"github.com/valentineejk/workerz/model"
	"github.com/valentineejk/workerz/repository"
)

type WorkerzService struct {
	repo repository.WorkerzRepo
}

func NewWorkerzService(wr repository.WorkerzRepo) *WorkerzService {
	return &WorkerzService{
		repo: wr,
	}
}

func (w WorkerzService) CreateWorker() error {
	fmt.Println("create worker service")
	return nil
}

func (w WorkerzService) GetAllWorkerz() ([]*model.Workerz, error) {

	_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	return nil, nil
}
