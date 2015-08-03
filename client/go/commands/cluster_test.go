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
	"github.com/gorilla/mux"
	"github.com/heketi/heketi/apps/glusterfs"
	"github.com/heketi/heketi/tests"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

//tests cluster info and destroy
func TestNewGetClusterInfoAndDestroy(t *testing.T) {
	defer os.Remove("heketi.db")

	// Create the app
	app := glusterfs.NewApp()
	defer app.Close()
	router := mux.NewRouter()
	app.SetRoutes(router)

	// Setup the server
	ts := httptest.NewServer(router)
	defer ts.Close()

	//set options
	options := &Options{
		Url: ts.URL,
	}

	//create b to get values of stdout
	var b bytes.Buffer
	defer tests.Patch(&stdout, &b).Restore()

	//create mock cluster and mock destroy
	mockCluster := NewCreateNewClusterCommand(options)

	//create new cluster
	err := mockCluster.Do()
	tests.Assert(t, err == nil)

	//get cluster id
	MockClusterId := strings.SplitAfter(b.String(), "id:")[1]
	b.Reset()

	//set destroy id to our id
	clusterInfo := NewGetClusterInfoCommand(options)
	clusterInfo.clusterId = MockClusterId

	//assert that cluster info Do succeeds and prints correctly
	err = clusterInfo.Do()
	tests.Assert(t, err == nil, err)
	tests.Assert(t, strings.Contains(b.String(), "For cluster:"), b.String())

	//create destroy struct and destroy it
	mockClusterDestroy := NewDestroyClusterCommand(options)
	mockClusterDestroy.clusterId = MockClusterId
	err = mockClusterDestroy.Do()
	tests.Assert(t, err == nil)

	//assert that we cannot get info on destroyed cluster
	err = clusterInfo.Do()
	tests.Assert(t, err != nil)

}

//test cluster list
func TestNewGetClusterList(t *testing.T) {
	defer os.Remove("heketi.db")

	// Create the app
	app := glusterfs.NewApp()
	defer app.Close()
	router := mux.NewRouter()
	app.SetRoutes(router)

	// Setup the server
	ts := httptest.NewServer(router)
	defer ts.Close()

	//set options
	options := &Options{
		Url: ts.URL,
	}

	//create b to get values of stdout
	var b bytes.Buffer
	defer tests.Patch(&stdout, &b).Restore()

	//create mock cluster and mock destroy
	mockCluster := NewCreateNewClusterCommand(options)

	//create new cluster
	err := mockCluster.Do()
	tests.Assert(t, err == nil)

	//assert cluster was created
	tests.Assert(t, strings.Contains(b.String(), "Cluster id:"), b.String())
	b.Reset()

	//create new list command
	listCommand := NewGetClusterListCommand(options)
	err = listCommand.Do()
	tests.Assert(t, err == nil)

	//asert stdout is correct
	tests.Assert(t, strings.Contains(b.String(), "Clusters: "), b.String())
}

//test cluster create
func TestClusterPostSuccess(t *testing.T) {
	defer os.Remove("heketi.db")
	// Create the app
	app := glusterfs.NewApp()
	defer app.Close()
	router := mux.NewRouter()
	app.SetRoutes(router)

	// Setup the server
	ts := httptest.NewServer(router)
	defer ts.Close()

	options := &Options{
		Url: ts.URL,
	}

	var b bytes.Buffer
	defer tests.Patch(&stdout, &b).Restore()

	cluster := NewCreateNewClusterCommand(options)
	tests.Assert(t, cluster != nil)
	err := cluster.Do()
	tests.Assert(t, err == nil)
	tests.Assert(t, strings.Contains(b.String(), "Cluster id:"), b.String())
}

func TestClusterPostFailure(t *testing.T) {
	defer os.Remove("heketi.db")

	// Create the app
	app := glusterfs.NewApp()
	defer app.Close()
	router := mux.NewRouter()
	app.SetRoutes(router)

	// Setup the server
	ts := httptest.NewServer(router)
	defer ts.Close()

	options := &Options{
		Url: "http://nottherightthing:8080",
	}

	var b bytes.Buffer
	defer tests.Patch(&stdout, &b).Restore()

	cluster := NewCreateNewClusterCommand(options)
	tests.Assert(t, cluster != nil)
	err := cluster.Do()
	tests.Assert(t, err != nil)
	tests.Assert(t, strings.Contains(b.String(), "Unable to send "))
}
