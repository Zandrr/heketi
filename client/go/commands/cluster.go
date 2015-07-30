//
// Copyright (c) 2015 The heketi Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package commands

import (
	"errors"
	"flag"
	"io"
	"os"
)

type ClusterCommand struct {
	Cmd

	// Subcommands available to this command
	cmds Commands

	options *Options

	// Subcommand
	cmd Command
}

var (
	stdout io.Writer = os.Stdout
)

func NewClusterCommand(options *Options) *ClusterCommand {
	cmd := &ClusterCommand{}
	cmd.name = "cluster"
	cmd.options = options
	cmd.cmds = Commands{
		NewCreateNewClusterCommand(options),
		NewGetClusterInfoCommand(options),
		NewGetClusterListCommand(options),
		NewDestroyClusterCommand(options),
	}

	cmd.flags = flag.NewFlagSet(cmd.name, flag.ExitOnError)

	return cmd
}

func (a *ClusterCommand) Name() string {
	return a.name

}

func (a *ClusterCommand) Parse(args []string) error {
	// Parse our flags here
	a.flags.Parse(args)

	//check number of args
	if len(a.flags.Args()) < 1 {
		return errors.New("Not enough arguments")
	}

	// Check which of the subcommands we need to call the .Parse function
	for _, cmd := range a.cmds {
		if a.flags.Arg(0) == cmd.Name() {
			err := cmd.Parse(a.flags.Args()[1:])
			if err != nil {
				return err
			}
			// Save this command for later use
			a.cmd = cmd

			return nil
		}
	}

	// Done
	return errors.New("Command not found")
}

func (a *ClusterCommand) Do() error {

	return a.cmd.Do()

}
