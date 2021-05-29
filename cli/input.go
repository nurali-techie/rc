package cli

type Input interface {
	GetContent() (string, error)
}
