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
	"bytes"
	"errors"
	"flag"
	"fmt"
	"github.com/heketi/heketi/apps/glusterfs"
	"github.com/heketi/heketi/utils"
	"io"
	"net/http"
	"os"
)

type CreateNewClusterCommand struct {
	// Generic stuff.  This is called
	// embedding.  In other words, the members in
	// the struct below are here also
	Cmd

	options *Options
}

var (
	stdout io.Writer = os.Stdout
)

func NewCreateNewClusterCommand(options *Options) *CreateNewClusterCommand {
	cmd := &CreateNewClusterCommand{}
	cmd.name = "create"
	cmd.options = options
	cmd.flags = flag.NewFlagSet(cmd.name, flag.ExitOnError)
	cmd.flags.Usage = func() {
		fmt.Println("Hello from my create")
	}

	return cmd
}

func (a *CreateNewClusterCommand) Name() string {
	return a.name

}

func (a *CreateNewClusterCommand) Parse(args []string) error {
	// a.flags.Parse(args)
	if len(args) > 0 {
		fmt.Println("Too many arguments!")
		return errors.New("Too many arguments!")
	}
	fmt.Println(len(args))
	return nil

}

func (a *CreateNewClusterCommand) Do() error {
	//set url
	url := a.options.Url

	//do http POST and check if sent to server
	r, err := http.Post(url+"/clusters", "application/json", bytes.NewBuffer([]byte("{}")))
	if err != nil {
		fmt.Fprintf(stdout, "Unable to send command to server: %v", err)
		return err
	}

	if r.StatusCode != http.StatusCreated {
		fmt.Println("status not ok")
		return errors.New("returned with bad response")
	}

	var body glusterfs.ClusterInfoResponse
	err = utils.GetJsonFromResponse(r, &body)
	if err != nil {
		fmt.Println("bad json response from server")
		return err
	}
	fmt.Fprintf(stdout, "Cluster id: %v\n", body.Id)
	//create var that is http server of heketi server. var httpServer.
	//maybe pass server as argument?
	//do a post to the server's URL/clusters and pass it the request object, {}
	//r, err := http.Post(httpServer.URL+"/clusters", "application/json", REQUEST)
	return nil
}
