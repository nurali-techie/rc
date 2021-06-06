package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/nurali-techie/rc/domain"
)

const (
	INSERT_COMMAND_SQL = "INSERT INTO commands (key, text) VALUES (?, ?)"
	SELECT_COMMAND_SQL = "SELECT key, text from commands WHERE key = ?"
	LIST_COMMAND_SQL   = "SELECT key from commands WHERE key like ?"
)

type commandRepository struct {
	db *sql.DB
}

func NewCommandRepository(db *sql.DB) domain.CommandRepository {
	commandRepo := new(commandRepository)
	commandRepo.db = db
	return commandRepo
}

func (s *commandRepository) Add(ctx context.Context, command *domain.Commmand) error {
	stmt, err := s.db.Prepare(INSERT_COMMAND_SQL)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(command.Key, command.Text)
	return err
}

func (s *commandRepository) Get(ctx context.Context, key string) (*domain.Commmand, error) {
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

	return nil, fmt.Errorf("'%s' command key not found", key)
}

func (s *commandRepository) List(ctx context.Context, query string) ([]string, error) {
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
