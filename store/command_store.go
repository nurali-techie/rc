package store

import (
	"context"
	"database/sql"

	"github.com/nurali-techie/rc/domain"
)

const (
	INSERT_COMMAND_SQL = "INSERT INTO commands (key, text) VALUES (?, ?)"
	SELECT_COMMAND_SQL = "SELECT key, text from commands WHERE key = ?"
	LIST_COMMAND_SQL   = "SELECT key from commands WHERE key like ?"
)

type CommandStore interface {
	Add(ctx context.Context, command *domain.Commmand) error
	Get(ctx context.Context, key string) (*domain.Commmand, error)
	List(ctx context.Context, query string) ([]string, error)
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
	stmt, err := s.db.Prepare(INSERT_COMMAND_SQL)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(command.Key, command.Text)
	return err
}

func (s *commandStore) Get(ctx context.Context, key string) (*domain.Commmand, error) {
	rows, err := s.db.Query(SELECT_COMMAND_SQL, key)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if rows.Next() {
		command := &domain.Commmand{}
		err := rows.Scan(&command.Key, &command.Text)
		if err != nil {
			return nil, err
		}
		return command, nil
	}

	return nil, nil
}

func (s *commandStore) List(ctx context.Context, query string) ([]string, error) {
	query = "%" + query + "%"
	rows, err := s.db.Query(LIST_COMMAND_SQL, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	keys := make([]string, 0)
	for rows.Next() {
		var key string
		err := rows.Scan(&key)
		if err != nil {
			return nil, err
		}
		keys = append(keys, key)
	}
	return keys, nil
}
