package domain

import "context"

// Command entity
type Commmand struct {
	Key  string
	Text string
}

// Command service
type CommandService interface {
	Add(ctx context.Context, command *Commmand) error
	Get(ctx context.Context, key string) (*Commmand, error)
	List(ctx context.Context, query string) ([]string, error)
}

// Command repository
type CommandRepository interface {
	Add(ctx context.Context, command *Commmand) error
	Get(ctx context.Context, key string) (*Commmand, error)
	List(ctx context.Context, query string) ([]string, error)
}
