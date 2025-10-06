package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/nurali-techie/rc/domain"
)

const (
	INSERT_MEMORY_SQL = "INSERT INTO memory (seq, key) VALUES (?, ?)"
	SELECT_MEMORY_SQL = "SELECT seq, key from memory WHERE seq = ?"
	DELETE_MEMORY_SQL = "DELETE FROM memory"
)

type memoryRepository struct {
	db *sql.DB
}

func NewMemoryRepository(db *sql.DB) domain.MemoryRepository {
	memoryRepo := new(memoryRepository)
	memoryRepo.db = db
	return memoryRepo
}

func (s *memoryRepository) AddAll(ctx context.Context, keys []string) ([]domain.Memory, error) {
	stmt, err := s.db.Prepare(INSERT_MEMORY_SQL)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	memories := make([]domain.Memory, 0)
	for ind, key := range keys {
		seq := ind + 1
		_, err = stmt.Exec(seq, key)
		if err != nil {
			return nil, err
		}
		memories = append(memories, domain.Memory{Seq: seq, Key: key})
	}
	return memories, nil
}

func (s *memoryRepository) Get(ctx context.Context, seq int) (*domain.Memory, error) {
	rows, err := s.db.Query(SELECT_MEMORY_SQL, seq)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if rows.Next() {
		memory := &domain.Memory{}
		err := rows.Scan(&memory.Seq, &memory.Key)
		if err != nil {
			return nil, err
		}
		return memory, nil
	}

	return nil, fmt.Errorf("'%d' command number not found", seq)
}

func (s *memoryRepository) Clear(ctx context.Context) error {
	stmt, err := s.db.Prepare(DELETE_MEMORY_SQL)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec()
	return err
}
