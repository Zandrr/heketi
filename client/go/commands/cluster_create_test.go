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
	"os"
	// "net/http"
	"net/http/httptest"
	"strings"
	"testing"
	// "github.com/heketi/heketi/utils"
)

func TestNewClusterCreate(t *testing.T) {
	options := &Options{
		Url: "home",
	}
	cluster := NewCreateNewClusterCommand(options)

	tests.Assert(t, cluster != nil)
	tests.Assert(t, cluster.name == "create")
	tests.Assert(t, cluster.options == options, *options)
}

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
	tests.Assert(t, strings.Contains(b.String(), "Cluster id: "))
}

func TestClusterPostFailureServer(t *testing.T) {
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

// func TestClusterPostFailureStatusResponse(t *testing.T) {
// 	defer os.Remove("heketi.db")
// 	// Create the app
// 	app := glusterfs.NewApp()
// 	defer app.Close()
// 	router := mux.NewRouter()
// 	app.SetRoutes(router)

// 	// Setup the server
// 	ts := httptest.NewServer(router)
// 	defer ts.Close()

// 	options := &Options{
// 		Url: ts.URL,
// 	}

// 	var b bytes.Buffer
// 	defer tests.Patch(&stdout, &b).Restore()

// 	cluster := NewCreateNewClusterCommand(options)
// 	tests.Assert(t, cluster != nil)
// 	err := cluster.Do()
// 	tests.Assert(t, err != nil, err)
// 	tests.Assert(t, strings.Contains(b.String(), "returned with bad response"))
// }
