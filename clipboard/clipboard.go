package clipboard

import "fmt"

type clipboard struct {
}

func NewClipboard() *clipboard {
	return new(clipboard)
}

func (input *clipboard) GetContent() []byte {
	return []byte{}
}

func (output *clipboard) SetContent(content []byte) error {
	fmt.Println("clipboard:", string(content))
	return nil
}
