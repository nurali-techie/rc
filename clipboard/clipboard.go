package clipboard

import (
	cb "github.com/atotto/clipboard"
)

type clipboard struct {
}

func NewClipboard() *clipboard {
	return new(clipboard)
}

func (input *clipboard) GetContent() (string, error) {
	return cb.ReadAll()
}

func (output *clipboard) SetContent(content string) error {
	return cb.WriteAll(content)
}
