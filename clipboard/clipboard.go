package clipboard

type clipboard struct {
}

func NewClipboard() *clipboard {
	return new(clipboard)
}

func (input *clipboard) GetContent() []byte {
	return []byte{}
}

func (output *clipboard) SetContent(content []byte) error {
	return nil
}
