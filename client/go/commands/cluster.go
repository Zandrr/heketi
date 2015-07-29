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
	"fmt"
)

type ClusterCommand struct {
	// Generic stuff.  This is called
	// embedding.  In other words, the members in
	// the struct below are here also
	Cmd

	// Subcommands available to this command
	cmds Commands

	// Subcommand
	cmd Command
}

func NewClusterCommand() *ClusterCommand {
	cmd := &ClusterCommand{}
	cmd.name = "cluster"

	cmd.cmds = Commands{
		NewCreateNewClusterCommand(),
		NewGetClusterInfoCommand(),
		NewGetClusterListCommand(),
		NewDestroyClusterCommand(),
	}

	cmd.flags = flag.NewFlagSet(cmd.name, flag.ExitOnError)
	cmd.flags.Usage = func() {
		fmt.Println("Hello from CLUSTER usage")
	}

	return cmd
}

func (a *ClusterCommand) Name() string {
	return a.name

}

func (a *ClusterCommand) Parse(args []string) error {

	// Parse our flags here

	// Check which of the subcommands we need to call the .Parse function
	for _, cmd := range a.cmds {
		if args[0] == cmd.Name() {
			cmd.Parse(args[1:])
			// Save this command for later use
			a.cmd = cmd

			return nil
		}
	}

	// Done
	return errors.New("Command not found")
}

func (a *ClusterCommand) Do() error {

	// Call cmd.Do()

	return a.cmd.Do()

}
