package test

import (
	"bytes"
	"os/exec"
	"testing"

	"github.com/stretchr/testify/assert"
)

var rcBinary = "../out/rc"

func TestRcCli(t *testing.T) {
	cmd, display := cmdRc()
	defer display(t)
	err := cmd.Run()
	assert.NoError(t, err)
}

func cmdRc(args ...string) (*exec.Cmd, func(t *testing.T)) {
	cmd := exec.Command(rcBinary, args...)

	var out bytes.Buffer
	cmd.Stderr = &out
	cmd.Stdout = &out

	display := func(t *testing.T) {
		if out.Len() == 0 {
			return
		}
		t.Log(out.String())
	}

	return cmd, display
}
