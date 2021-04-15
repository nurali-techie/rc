package store

import (
	"context"
	"database/sql"

	"github.com/nurali-techie/rc/domain"
)

type CommandStore interface {
	Add(ctx context.Context, command *domain.Commmand) error
	Get(ctx context.Context, key string) (*domain.Commmand, error)
}

type commandStore struct {
	db *sql.DB
}

func NewCommandStore(db *sql.DB) CommandStore {
	commandStore := new(commandStore)
	commandStore.db = db
	return commandStore
}

func (s *commandStore) Add(ctx context.Context, command *domain.Commmand) error {
	s.commands[command.Key] = command
	return nil
}

func (s *commandStore) Get(ctx context.Context, key string) (*domain.Commmand, error) {
	return s.commands[key], nil
}
