// Copyright 2022 The Amesh Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package types

import "errors"

var (
	ErrorRequireFurtherEDS = errors.New("require further eds")
)

var (
	// RouteConfigurationUrl is the RDS type url.
	RouteConfigurationUrl = "type.googleapis.com/envoy.config.route.v3.RouteConfiguration"
	// ClusterUrl is the Cluster type url.
	ClusterUrl = "type.googleapis.com/envoy.config.cluster.v3.Cluster"
	// ClusterLoadAssignmentUrl is the Cluster type url.
	ClusterLoadAssignmentUrl = "type.googleapis.com/envoy.config.endpoint.v3.ClusterLoadAssignment"
	// ListenerUrl is the Listener type url.
	ListenerUrl = "type.googleapis.com/envoy.config.listener.v3.Listener"
)

// Provisioner provisions config event.
// The source type can be xDS or UDPA or whatever anything else.
type Provisioner interface {
	// Channel returns a readonly channel where caller can get events.
	EventsChannel() <-chan []Event
	// Run launches the provisioner.
	Run(<-chan struct{}) error
	// GetData returns the specific data for status server
	GetData(dataType string) (string, error)
}

// EventType is the kind of event.
type EventType string

var (
	// EventAdd represents the add event.
	EventAdd = EventType("add")
	// EventUpdate represents the update event.
	EventUpdate = EventType("update")
	// EventDelete represents the delete event.
	EventDelete = EventType("delete")
)

// Event describes a specific event generated from the provisioner.
type Event struct {
	Type   EventType
	Key    string
	Object interface{}
	// Tombstone is only valid for delete event,
	// in such a case it stands for the final state
	// of the object.
	Tombstone interface{}

	// Revision is the revision that the event happened
	Revision int64
}

type AmeshPluginProvider interface {
	GetPlugins() map[string]*ApisixPlugin
}

type ApisixExtPluginConfig struct {
	Conf []*ApisixExtPlugin `json:"conf,omitempty" yaml:"conf,omitempty"`
}

type ApisixExtPlugin struct {
	Name  string                 `json:"name,omitempty" yaml:"name,omitempty"`
	Value map[string]interface{} `json:"value,omitempty" yaml:"value,omitempty"`
}

type ApisixPlugin struct {
	// The plugin type
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
	// The plugin name
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
	// The plugin config
	Config map[string]interface{} `json:"config,omitempty" yaml:"config,omitempty"`
}
