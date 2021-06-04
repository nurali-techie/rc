package io

type Input interface {
	GetContent() (string, error)
}
