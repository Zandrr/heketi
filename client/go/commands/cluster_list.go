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
	"github.com/heketi/heketi/apps/glusterfs"
	"github.com/heketi/heketi/utils"
	"net/http"
	// "github.com/heketi/heketi/client/go/utils"
)

type GetClusterListCommand struct {
	// Generic stuff.  This is called
	// embedding.  In other words, the members in
	// the struct below are here also
	Cmd
	options *Options
}

func NewGetClusterListCommand(options *Options) *GetClusterListCommand {
	cmd := &GetClusterListCommand{}
	cmd.name = "list"
	cmd.options = options
	cmd.flags = flag.NewFlagSet(cmd.name, flag.ExitOnError)

	return cmd
}

func (a *GetClusterListCommand) Name() string {
	return a.name

}

func (a *GetClusterListCommand) Parse(args []string) error {
	// a.flags.Parse(args)
	if len(args) > 0 {
		return errors.New("Too many arguments!")
	}
	return nil

}

func (a *GetClusterListCommand) Do() error {
	//set url
	url := a.options.Url

	//do http GET and check if sent to server
	r, err := http.Get(url + "/clusters")
	if err != nil {
		fmt.Fprintf(stdout, "Unable to send command to server: %v", err)
		return err
	}

	//check status code
	if r.StatusCode != http.StatusOK {
		fmt.Println("status not ok")
		return errors.New("returned with bad response")
	}

	//check json response
	var body glusterfs.ClusterListResponse
	err = utils.GetJsonFromResponse(r, &body)
	if err != nil {
		fmt.Println("bad json response from server")
		return err
	}

	//if all is well, print stuff
	fmt.Fprintf(stdout, "Clusters: \n")
	for _, cluster := range body.Clusters {
		fmt.Println(cluster)
	}
	return nil
}
