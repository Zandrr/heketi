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
	// "github.com/heketi/heketi/utils"
	// "github.com/heketi/heketi/apps/glusterfs"
	"net/http"
	// "github.com/heketi/heketi/client/go/utils"
	// "net/http"
)

type DestroyClusterCommand struct {
	// Generic stuff.  This is called
	// embedding.  In other words, the members in
	// the struct below are here also
	Cmd
	options   *Options
	clusterId string
}

func NewDestroyClusterCommand(options *Options) *DestroyClusterCommand {
	cmd := &DestroyClusterCommand{}
	cmd.name = "destroy"
	cmd.options = options
	cmd.flags = flag.NewFlagSet(cmd.name, flag.ExitOnError)
	cmd.flags.Usage = func() {
		fmt.Println("Hello from my destroy")
	}

	return cmd
}

func (a *DestroyClusterCommand) Name() string {
	return a.name

}

func (a *DestroyClusterCommand) Parse(args []string) error {
	a.flags.Parse(args)
	a.clusterId = a.flags.Arg(0)
	return nil

}

func (a *DestroyClusterCommand) Do() error {
	//set url
	url := a.options.Url

	//create destroy request object
	req, err := http.NewRequest("DELETE", url+"/clusters/"+a.clusterId, nil)
	if err != nil {
		fmt.Fprintf(stdout, "Unable to initiate destroy: %v", err)
		return err
	}

	//destroy cluster
	r, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Fprintf(stdout, "Unable to send command to server: %v", err)
		return err
	}

	//check status code
	if r.StatusCode != http.StatusOK {
		fmt.Println("status not ok")
		return errors.New("returned with bad response")
	}

	//if all is well, print stuff
	fmt.Fprintf(stdout, "Successfully destroyed cluster with id: %v ", a.clusterId)

	return nil
}
