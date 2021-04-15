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
	cmd, display := cmdRc("add", "k1", "v1")
	defer display(t)
	err := cmd.Run()
	assert.NoError(t, err)

	cmd, display = cmdRc("k1")
	defer display(t, "GetCmd", "v1")
	err = cmd.Run()
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
		output := out.String()
		for _, want := range wants {
			assert.Contains(t, output, want)
		}
	}

	return cmd, display
}
