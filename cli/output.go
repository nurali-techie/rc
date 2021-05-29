package cli

type Output interface {
	SetContent(content string) error
}
