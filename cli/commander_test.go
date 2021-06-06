package cli_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/nurali-techie/rc/cli"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type CommanderSuite struct {
	suite.Suite
	commander *cli.Commander
}

func TestCommander(t *testing.T) {
	commander := cli.NewCommander(newHelpCmd(), newDefaultCmd(t, []string{"1"}))
	s := &CommanderSuite{
		commander: commander,
	}
	suite.Run(t, s)
}

func (s *CommanderSuite) TestDefaultCommand() {
	err := s.commander.ServeCommand([]string{"1"})
	assert.NoError(s.T(), err)
}

func (s *CommanderSuite) TestHelpCommand() {
	err := s.commander.ServeCommand([]string{"help"})
	assert.EqualError(s.T(), err, "help not provided")
}

func (s *CommanderSuite) TestZeroArgs() {
	err := s.commander.ServeCommand([]string{})
	assert.EqualError(s.T(), err, "help not provided")
}

func (s *CommanderSuite) TestOtherCommand() {
	s.commander.Register("other", newOtherCmd(s.T(), []string{"10", "20"}))
	err := s.commander.ServeCommand([]string{"other", "10", "20"})
	assert.NoError(s.T(), err)
}

func (s *CommanderSuite) TestCommandReturnError() {
	s.commander.Register("error", newErrorCmd(s.T(), "sample error"))
	err := s.commander.ServeCommand([]string{"error"})
	assert.EqualError(s.T(), err, "sample error")
}

// test commands
type helpCmd struct {
}

func (c *helpCmd) Execute(ctx context.Context, args []string) error {
	return fmt.Errorf("help not provided") // check help command, also check command return error
}

type argCheckCmd struct {
	t        *testing.T
	wantArgs []string
	errMsg   string
}

func (c *argCheckCmd) Execute(ctx context.Context, args []string) error {
	if c.errMsg != "" {
		return fmt.Errorf(c.errMsg)
	}
	assert.Equal(c.t, c.wantArgs, args) // check input arg, in case of default command and in case of registred command
	return nil
}

func newHelpCmd() cli.Command {
	return &helpCmd{}
}
func newDefaultCmd(t *testing.T, wantArgs []string) cli.Command {
	return &argCheckCmd{t, wantArgs, ""}
}
func newOtherCmd(t *testing.T, wantArgs []string) cli.Command {
	return &argCheckCmd{t, wantArgs, ""}
}
func newErrorCmd(t *testing.T, errMsg string) cli.Command {
	return &argCheckCmd{t, nil, errMsg}
}

// ---
