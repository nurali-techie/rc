package io_test

import (
	"testing"

	"github.com/nurali-techie/rc/io"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestClipboard(t *testing.T) {
	clipboard := io.NewClipboard()
	err := clipboard.SetContent("hello_test")
	require.NoError(t, err)

	// test
	content, err := clipboard.GetContent()

	assert.NoError(t, err)
	assert.Equal(t, "hello_test", content)
}

// Note: this test will not pass as clipboard is not thread safe
// func TestClipboardParallel(t *testing.T) {
// 	table := []struct {
// 		in   string
// 		want string
// 	}{
// 		{
// 			"hello_test1",
// 			"hello_test1",
// 		},
// 		{
// 			"hello_test2",
// 			"hello_test2",
// 		},
// 	}

// 	for i, tt := range table {
// 		tt := tt
// 		name := fmt.Sprintf("test-%d", i)
// 		t.Run(name, func(t *testing.T) {
// 			t.Parallel()
// 			clipboard := io.NewClipboard()
// 			err := clipboard.SetContent(tt.in)
// 			require.NoError(t, err)

// 			got, err := clipboard.GetContent()
// 			assert.NoError(t, err)
// 			assert.Equal(t, tt.want, got)
// 		})
// 	}
// }
