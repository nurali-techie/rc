package service

import (
	"context"
	"strconv"

	"github.com/nurali-techie/rc/domain"
)

type commandService struct {
	repo          domain.CommandRepository
	memoryService domain.MemroyService
}

func NewCommandService(repo domain.CommandRepository, memoryService domain.MemroyService) domain.CommandService {
	commandService := new(commandService)
	commandService.repo = repo
	commandService.memoryService = memoryService
	return commandService
}

func (s *commandService) Add(ctx context.Context, command *domain.Commmand) error {
	return s.repo.Add(ctx, command)
}

func (s *commandService) Get(ctx context.Context, key string) (*domain.Commmand, error) {
	seq, err := strconv.Atoi(key)
	if err == nil {
		memory, err := s.memoryService.Get(ctx, seq)
		if err != nil {
			return nil, err
		}
		key = memory.Key
	}
	return s.repo.Get(ctx, key)
}

func (s *commandService) List(ctx context.Context, query string) ([]domain.Memory, error) {
	keys, err := s.repo.List(ctx, query)
	if err != nil {
		return nil, err
	}
	err = s.memoryService.Clear(ctx)
	if err != nil {
		return nil, err
	}
	return s.memoryService.AddAll(ctx, keys)
}

func (s *commandService) Move(ctx context.Context, oldKey, newKey string) error {
	return s.repo.Move(ctx, oldKey, newKey)
}

func (s *commandService) Remove(ctx context.Context, key string) error {
	return s.repo.Remove(ctx, key)
}
