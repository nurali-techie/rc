package console

import (
	"os"
)

type console struct {
}

func NewConsole() *console {
	return new(console)
}

func (output *console) SetContent(content []byte) error {
	os.Stdout.Write(content)
	return nil
}
