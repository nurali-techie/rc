package command

import "context"

type Command interface {
	Execute(ctx context.Context, in Input, out Output, args []string) error
}
