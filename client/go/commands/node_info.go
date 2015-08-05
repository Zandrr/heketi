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
)

type GetNodeInfoCommand struct {
	Cmd
	options *Options
	nodeId  string
}

func NewGetNodeInfoCommand(options *Options) *GetNodeInfoCommand {
	cmd := &GetNodeInfoCommand{}
	cmd.name = "info"
	cmd.options = options
	cmd.flags = flag.NewFlagSet(cmd.name, flag.ExitOnError)

	return cmd
}

func (a *GetNodeInfoCommand) Name() string {
	return a.name

}

func (a *GetNodeInfoCommand) Exec(args []string) error {
	if len(args) < 1 {
		return errors.New("Not enough arguments!")
	}
	if len(args) >= 2 {
		return errors.New("Too many arguments!")
	}
	a.flags.Parse(args)
	a.nodeId = a.flags.Arg(0)
	url := a.options.Url

	//do http GET and check if sent to server
	r, err := http.Get(url + "/nodes/" + a.nodeId)
	if err != nil {
		fmt.Fprintf(stdout, "Unable to send command to server: %v", err)
		return err
	}

	//check status code
	if r.StatusCode != http.StatusOK {
		s, err := utils.GetStringFromResponse(r)
		if err != nil {
			return err
		}
		return errors.New(s)
	}

	//check json response
	var body glusterfs.NodeInfoResponse
	err = utils.GetJsonFromResponse(r, &body)
	if err != nil {
		fmt.Println("Error: Bad json response from server")
		return err
	}

	//print revelent results
	s := "For node: " + a.nodeId + " \n in cluster " + body.ClusterId + "\n in zone " + string(body.Zone) + "\n Manage hostnames are: \n"
	for _, hostname := range body.Hostnames.Manage {
		s += hostname + "\n"
	}
	s += "Storage hostnames are: \n"
	for _, hostname := range body.Hostnames.Storage {
		s += hostname + "\n"
	}
	s += "Devices are: \n"
	for _, device := range body.DevicesInfo {
		s += "Name is: " + device.Name

	}

	fmt.Fprintf(stdout, s)
	return nil

}
