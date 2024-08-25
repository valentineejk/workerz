package repository

import (
	"context"

	dbq "github.com/valentineejk/workerz/database/sqlc"
	"github.com/valentineejk/workerz/model"
)

type WorkerzRepo interface {
	InsertWorkerz(ctx context.Context, Worker model.Workerz) error
}

type DbWorkerzRepoImpl struct {
	db *dbq.Queries
}

func NewDbWorkerzRepo(db *dbq.Queries) *DbWorkerzRepoImpl {
	return &DbWorkerzRepoImpl{
		db: db,
	}
}

func (w *DbWorkerzRepoImpl) InsertWorkerz(ctx context.Context, Worker model.Workerz) error {
	return nil
}
