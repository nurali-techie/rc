package command

type Output interface {
	SetContent(content []byte) error
}
