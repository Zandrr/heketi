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
	// "github.com/heketi/heketi/apps/glusterfs"
	// "github.com/heketi/heketi/client/go/utils"
	// "net/http"
)

type GetClusterInfoCommand struct {
	// Generic stuff.  This is called
	// embedding.  In other words, the members in
	// the struct below are here also
	Cmd
}

func NewGetClusterInfoCommand() *GetClusterInfoCommand {
	cmd := &GetClusterInfoCommand{}
	cmd.name = "info"

	cmd.flags = flag.NewFlagSet(cmd.name, flag.ExitOnError)
	cmd.flags.Usage = func() {
		fmt.Println("Hello from my info")
	}

	return cmd
}

func (a *GetClusterInfoCommand) Name() string {
	return a.name

}

func (a *GetClusterInfoCommand) Parse(args []string) error {
	// a.flags.Parse(args)
	if len(args) > 0 {
		fmt.Println("Too many arguments!")
		return errors.New("Too many arguments!")
	}
	fmt.Println(len(args))
	return nil

}

func (a *GetClusterInfoCommand) Do() error {
	//create var that is http server of heketi server. var httpServer.
	//maybe pass server as arugment?
	//do a post to the server's URL/clusters and pass it the request object, {}
	//r, err := http.Post(httpServer.URL+"/clusters", "application/json", REQUEST)
	return nil
}
