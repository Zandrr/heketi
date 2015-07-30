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

func TestNewGetClusterInfo(t *testing.T) {
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
	mockClustererr := mockCluster.Do()
	tests.Assert(t, mockClustererr == nil)

	//get cluster id
	MockClusterId := strings.SplitAfter(b.String(), "id:")[1]
	b.Reset()

	//set destroy id to our id
	clusterInfo := NewGetClusterInfoCommand(options)
	clusterInfo.clusterId = MockClusterId

	err := clusterInfo.Do()

	tests.Assert(t, err == nil, err)
	tests.Assert(t, strings.Contains(b.String(), "For cluster:"), b.String())

	//create destroy struct and destroy it
	mockClusterDestroy := NewDestroyClusterCommand(options)
	mockClusterDestroy.clusterId = MockClusterId
	err = mockClusterDestroy.Do()
	tests.Assert(t, err == nil)

	err = clusterInfo.Do()
	tests.Assert(t, err != nil)

}
