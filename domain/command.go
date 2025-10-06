package domain

import "context"

// Command entity
type Commmand struct {
	Key  string
	Text string
}

type Memory struct {
	Seq int
	Key string
}

// Command service
type CommandService interface {
	Add(ctx context.Context, command *Commmand) error
	Get(ctx context.Context, key string) (*Commmand, error)
	List(ctx context.Context, query string) ([]Memory, error)
	Move(ctx context.Context, oldKey, newKey string) error
	Remove(ctx context.Context, key string) error
}

// Command repository
type CommandRepository interface {
	Add(ctx context.Context, command *Commmand) error
	Get(ctx context.Context, key string) (*Commmand, error)
	List(ctx context.Context, query string) ([]string, error)
	Move(ctx context.Context, oldKey, newKey string) error
	Remove(ctx context.Context, key string) error
}

// Memory service
type MemroyService interface {
	AddAll(ctx context.Context, keys []string) ([]Memory, error)
	Get(ctx context.Context, seq int) (*Memory, error)
	Clear(ctx context.Context) error
}

// Memory repository
type MemoryRepository interface {
	AddAll(ctx context.Context, keys []string) ([]Memory, error)
	Get(ctx context.Context, seq int) (*Memory, error)
	Clear(ctx context.Context) error
}
