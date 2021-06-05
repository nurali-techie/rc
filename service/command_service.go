package service

import (
	"context"

	"github.com/nurali-techie/rc/domain"
)

type commandService struct {
	repo domain.CommandRepository
}

func NewCommandService(repo domain.CommandRepository) domain.CommandService {
	commandService := new(commandService)
	commandService.repo = repo
	return commandService
}

func (s *commandService) Add(ctx context.Context, command *domain.Commmand) error {
	return s.repo.Add(ctx, command)
}

func (s *commandService) Get(ctx context.Context, key string) (*domain.Commmand, error) {
	return s.repo.Get(ctx, key)
}

func (s *commandService) List(ctx context.Context, query string) ([]string, error) {
	return s.repo.List(ctx, query)
}
