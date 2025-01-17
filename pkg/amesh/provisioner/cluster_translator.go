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
package provisioner

import (
	"github.com/api7/gopkg/pkg/id"
	clusterv3 "github.com/envoyproxy/go-control-plane/envoy/config/cluster/v3"
	"go.uber.org/zap"

	"github.com/api7/amesh/pkg/apisix"
)

// TranslateCluster generates ApisixUpstream from Cluster
func (p *xdsProvisioner) TranslateCluster(c *clusterv3.Cluster) (*apisix.Upstream, error) {
	ups := &apisix.Upstream{
		Name:  c.Name,
		Id:    id.GenID(c.Name),
		Nodes: []*apisix.Node{},
	}
	if err := p.translateClusterLbPolicy(c, ups); err != nil {
		return nil, err
	}
	if err := p.translateClusterTimeoutSettings(c, ups); err != nil {
		return nil, err
	}
	if err := p.translateClusterLoadAssignments(c, ups); err != nil {
		return nil, err
	}

	p.logger.Debugw("got upstream after parsing cluster",
		zap.String("cluster_name", c.Name),
		zap.Any("upstream", ups),
	)

	return ups, nil
}

// translateClusterLbPolicy generates load balance Type settings for ApisixUpstream from Cluster
func (p *xdsProvisioner) translateClusterLbPolicy(c *clusterv3.Cluster, ups *apisix.Upstream) error {
	switch c.GetLbPolicy() {
	case clusterv3.Cluster_ROUND_ROBIN:
		ups.Type = "roundrobin"
	case clusterv3.Cluster_LEAST_REQUEST:
		// Apache APISIX's lease_conn policy is same to lease request.
		// But is doesn't expose configuration items. So LbConfig field
		// is ignored.
		ups.Type = "least_conn"
	default:
		// Apache APISIX doesn't support Random, Manglev. In addition,
		// also RinghHash (Consistent Hash) is available but the configurations
		// like key is in RouteConfiguration, so we cannot use it either.
		p.logger.Warnw("ignore cluster with unsupported load balancer",
			zap.String("cluster_name", c.Name),
			zap.String("lb_policy", c.GetLbPolicy().String()),
		)
		return nil
	}
	return nil
}

// translateClusterTimeoutSettings generates Timeout settings for ApisixUpstream from Cluster
func (p *xdsProvisioner) translateClusterTimeoutSettings(c *clusterv3.Cluster, ups *apisix.Upstream) error {
	if c.GetConnectTimeout() != nil {
		ups.Timeout = &apisix.Timeout{
			Connect: float64((*c.GetConnectTimeout()).Seconds),
			Read:    60,
			Send:    60,
		}
	}
	return nil
}

// translateClusterLoadAssignments generates Nodes for ApisixUpstream from Cluster.LoadAssignment
func (p *xdsProvisioner) translateClusterLoadAssignments(c *clusterv3.Cluster, ups *apisix.Upstream) error {
	if c.GetClusterType() != nil {
		p.logger.Warnw("ignore cluster with unsupported cluster type",
			zap.String("cluster_type", c.GetClusterType().Name),
			zap.Any("cluster", c),
		)
		return nil
	}
	switch c.GetType() {
	case clusterv3.Cluster_EDS:
		p.logger.Debugw("cluster depends on another EDS config, an upstream without nodes setting was generated",
			zap.Any("upstream", ups),
		)
		return nil
	default:
		nodes, err := p.TranslateClusterLoadAssignment(c.GetLoadAssignment())
		if err != nil {
			return err
		}
		ups.Nodes = nodes
		return nil
	}
}
