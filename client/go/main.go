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

package main

import (
	"flag"
	"fmt"
	"github.com/heketi/heketi/client/go/commands"
	"os"
)

var options commands.Options

func init() {

	flag.StringVar(&options.Url, "server", "", "server url goes here.")

}

// ------ Main
func main() {
	flag.Parse()
	if options.Url == "" {
		fmt.Println("You need a server!")
		os.Exit(1)
	}
	cmds := commands.Commands{
		commands.NewArithCommand(),
		commands.NewEchoCommand(),
		commands.NewClusterCommand(&options),
	}

	for _, cmd := range cmds {
		if flag.Arg(0) == cmd.Name() {
			cmd.Parse(flag.Args()[1:])
			cmd.Do()
			return
		}
	}

	fmt.Println("Command not found")
}
