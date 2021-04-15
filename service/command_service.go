package service

import (
	"context"

	"github.com/nurali-techie/rc/domain"
	"github.com/nurali-techie/rc/store"
)

type CommandService interface {
	Add(ctx context.Context, command *domain.Commmand) error
	Get(ctx context.Context, key string) (*domain.Commmand, error)
}

type commandService struct {
	store store.CommandStore
}

func NewCommandService(store store.CommandStore) CommandService {
	commandService := new(commandService)
	commandService.store = store
	return commandService
}

func (s *commandService) Add(ctx context.Context, command *domain.Commmand) error {
	return s.store.Add(ctx, command)
}

func (s *commandService) Get(ctx context.Context, key string) (*domain.Commmand, error) {
	return s.store.Get(ctx, key)
}
