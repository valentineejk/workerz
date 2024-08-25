package services

import (
	"context"

	nanoid "github.com/matoous/go-nanoid/v2"
	dbq "github.com/valentineejk/workerz/database/sqlc"
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

func (w WorkerzService) AddWorkerzService(ctx context.Context, Worker dbq.InsertWorkerzParams) error {

	return w.repo.InsertWorkerz(ctx, Worker)

}

func (w WorkerzService) GetAllWorkerz(ctx context.Context) ([]dbq.Workerz, error) {

	wrz, wrzErr := w.repo.ListWorkerz(ctx)
	if wrzErr != nil {
		return nil, wrzErr
	}

	return wrz, nil
}

const (
	alphabet = "0123456789abcdefghijklmnopqrstuvwxyz"
	length   = 12
)

// New generates a unique public ID.
func NanoIDS() (string, error) { return nanoid.Generate(alphabet, length) }
