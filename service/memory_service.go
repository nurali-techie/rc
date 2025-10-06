package service

import (
	"context"

	"github.com/nurali-techie/rc/domain"
)

type memoryService struct {
	repo domain.MemoryRepository
}

func NewMemoryService(repo domain.MemoryRepository) domain.MemroyService {
	memoryService := new(memoryService)
	memoryService.repo = repo
	return memoryService
}

func (s *memoryService) AddAll(ctx context.Context, keys []string) ([]domain.Memory, error) {
	return s.repo.AddAll(ctx, keys)
}

func (s *memoryService) Get(ctx context.Context, seq int) (*domain.Memory, error) {
	return s.repo.Get(ctx, seq)
}

func (s *memoryService) Clear(ctx context.Context) error {
	return s.repo.Clear(ctx)
}
