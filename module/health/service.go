package health

import (
	"context"

	logger "github.com/test_cache_CQRS/infrastructure/log"
)

type InterfaceService interface {
	CheckUpTime(ctx context.Context) (err error)
}

type Service struct {
	repository RepositoryInterface
}

func NewService(repository RepositoryInterface) InterfaceService {
	return &Service{
		repository: repository,
	}
}

func (u *Service) CheckUpTime(ctx context.Context) (err error) {
	ctxName := "CheckUpTimeDB"
	err = u.repository.CheckUpTimeDB(ctx)
	if err != nil {
		logger.Error(ctx, ctxName, "got error when %s : %v", ctxName, err)
		return err
	}

	return nil
}
