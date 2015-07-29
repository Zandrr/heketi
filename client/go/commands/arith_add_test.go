package commands

import (
	// "errors"
	// "fmt"
	"github.com/heketi/heketi/client/go/commands"
	"github.com/heketi/heketi/tests"
	"testing"
)

func TestArithAdd(t *testing.T) {
	adder := commands.NewArithAddCommand()
	tests.Assert(t, adder.Do() == 4)
}
