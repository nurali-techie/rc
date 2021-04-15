package cli

type Output interface {
	SetContent(content []byte) error
}
