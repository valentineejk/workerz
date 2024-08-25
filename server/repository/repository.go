package repository

import (
	"context"

	dbq "github.com/valentineejk/workerz/database/sqlc"
)

type WorkerzRepo interface {
	InsertWorkerz(ctx context.Context, Worker dbq.InsertWorkerzParams) error
	ListWorkerz(ctx context.Context) ([]dbq.Workerz, error)
}

type DbWorkerzRepoImpl struct {
	db *dbq.Queries
}

func NewDbWorkerzRepo(db *dbq.Queries) *DbWorkerzRepoImpl {
	return &DbWorkerzRepoImpl{
		db: db,
	}
}

func (w *DbWorkerzRepoImpl) InsertWorkerz(ctx context.Context, Worker dbq.InsertWorkerzParams) error {

	_, err := w.db.InsertWorkerz(ctx, Worker)
	if err != nil {
		return err
	}

	return nil
}

func (w *DbWorkerzRepoImpl) ListWorkerz(ctx context.Context) ([]dbq.Workerz, error) {

	wrzListData, err := w.db.GetAllWorkerz(ctx)
	if err != nil {
		return nil, err
	}

	workerzList := append([]dbq.Workerz{}, wrzListData...)

	return workerzList, nil
}
