package io

type Output interface {
	SetContent(content string) error
}
