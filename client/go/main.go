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
	"io"
	"os"
)

var (
	stdout  io.Writer = os.Stdout
	options commands.Options
)

func init() {

	flag.StringVar(&options.Url, "server", "", "server url goes here.")

	flag.Usage = func() {
		fmt.Println("USAGE: \n")
		fmt.Println("heketi cluster <n>\n")
		fmt.Println("where n can be one of the following: \n")
		fmt.Println("create <id> \n info <id> \n list \n destroy <id>")
	}
}

// ------ Main
func main() {
	flag.Parse()
	if options.Url == "" {
		fmt.Fprintf(stdout, "You need a server!\n")
		os.Exit(1)
	}
	cmds := commands.Commands{
		commands.NewClusterCommand(&options),
	}

	for _, cmd := range cmds {
		if flag.Arg(0) == cmd.Name() {

			//check for parse err
			err := cmd.Parse(flag.Args()[1:])
			if err != nil {
				fmt.Fprintf(stdout, "Parse Error: %v\n", err)
				os.Exit(1)
			}

			//check for do err
			err = cmd.Do()
			if err != nil {
				fmt.Fprintf(stdout, "Do Error: %v\n", err)
				fmt.Println(err)
				os.Exit(1)
			}
			return
		}
	}

	fmt.Println("Command not found")
}
