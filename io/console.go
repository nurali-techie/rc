package io

import (
	"os"
)

type console struct {
}

func NewConsole() *console {
	return new(console)
}

func (output *console) SetContent(content string) error {
	os.Stdout.Write([]byte(content))
	return nil
}
