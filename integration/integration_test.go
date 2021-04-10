package test

import (
	"bytes"
	"os/exec"
	"testing"

	"github.com/stretchr/testify/assert"
)

var rcBinary = "../out/rc"

func TestHelpCmd(t *testing.T) {
	cmd, display := cmdRc()
	defer display(t, "rc stands for recall")
	err := cmd.Run()
	assert.NoError(t, err)
}

func TestGetCmd(t *testing.T) {
	cmd, display := cmdRc("1")
	defer display(t, "GetCmd", "1")
	err := cmd.Run()
	assert.NoError(t, err)
}

func cmdRc(args ...string) (*exec.Cmd, func(t *testing.T, wants ...string)) {
	cmd := exec.Command(rcBinary, args...)

	var out bytes.Buffer
	cmd.Stderr = &out
	cmd.Stdout = &out

	display := func(t *testing.T, wants ...string) {
		if out.Len() == 0 {
			return
		}
		for _, want := range wants {
			assert.Contains(t, out.String(), want)
		}
	}

	return cmd, display
}
