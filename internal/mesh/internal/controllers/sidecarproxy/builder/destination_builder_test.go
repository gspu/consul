// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

package builder

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/hashicorp/consul/internal/catalog"
	"github.com/hashicorp/consul/internal/mesh/internal/types/intermediate"
	"github.com/hashicorp/consul/internal/resource"
	"github.com/hashicorp/consul/internal/resource/resourcetest"
	pbcatalog "github.com/hashicorp/consul/proto-public/pbcatalog/v1alpha1"
	pbmesh "github.com/hashicorp/consul/proto-public/pbmesh/v1alpha1"
	"github.com/hashicorp/consul/proto-public/pbresource"
)

var (
	endpointsData = &pbcatalog.ServiceEndpoints{
		Endpoints: []*pbcatalog.Endpoint{
			{
				Addresses: []*pbcatalog.WorkloadAddress{
					{Host: "10.0.0.1"},
				},
				Ports: map[string]*pbcatalog.WorkloadPort{
					"tcp":  {Port: 8080, Protocol: pbcatalog.Protocol_PROTOCOL_TCP},
					"http": {Port: 8080, Protocol: pbcatalog.Protocol_PROTOCOL_HTTP},
					"mesh": {Port: 20000, Protocol: pbcatalog.Protocol_PROTOCOL_MESH},
				},
			},
		},
	}
)

func TestBuildExplicitDestinations(t *testing.T) {
	api1Endpoints := resourcetest.Resource(catalog.ServiceEndpointsType, "api-1").
		WithData(t, endpointsData).Build()

	api2Endpoints := resourcetest.Resource(catalog.ServiceEndpointsType, "api-2").
		WithData(t, endpointsData).Build()

	api1Identity := &pbresource.Reference{
		Name:    "api1-identity",
		Tenancy: api1Endpoints.Id.Tenancy,
	}

	api2Identity := &pbresource.Reference{
		Name:    "api2-identity",
		Tenancy: api2Endpoints.Id.Tenancy,
	}

	destinationIpPort := &intermediate.Destination{
		Explicit: &pbmesh.Upstream{
			DestinationRef:  resource.Reference(api1Endpoints.Id, ""),
			DestinationPort: "tcp",
			Datacenter:      "dc1",
			ListenAddr: &pbmesh.Upstream_IpPort{
				IpPort: &pbmesh.IPPortAddress{Ip: "1.1.1.1", Port: 1234},
			},
		},
		ServiceEndpoints: &intermediate.ServiceEndpoints{
			Resource:  api1Endpoints,
			Endpoints: endpointsData,
		},
		Identities: []*pbresource.Reference{api1Identity},
	}

	destinationUnix := &intermediate.Destination{
		Explicit: &pbmesh.Upstream{
			DestinationRef:  resource.Reference(api2Endpoints.Id, ""),
			DestinationPort: "tcp",
			Datacenter:      "dc1",
			ListenAddr: &pbmesh.Upstream_Unix{
				Unix: &pbmesh.UnixSocketAddress{Path: "/path/to/socket", Mode: "0666"},
			},
		},
		ServiceEndpoints: &intermediate.ServiceEndpoints{
			Resource:  api2Endpoints,
			Endpoints: endpointsData,
		},
		Identities: []*pbresource.Reference{api2Identity},
	}

	cases := map[string]struct {
		destinations []*intermediate.Destination
	}{
		"destination/l4-single-destination-ip-port-bind-address": {
			destinations: []*intermediate.Destination{destinationIpPort},
		},
		"destination/l4-single-destination-unix-socket-bind-address": {
			destinations: []*intermediate.Destination{destinationUnix},
		},
		"destination/l4-multi-destination": {
			destinations: []*intermediate.Destination{destinationIpPort, destinationUnix},
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			proxyTmpl := New(testProxyStateTemplateID(), testIdentityRef(), "foo.consul", "dc1", nil).
				BuildDestinations(c.destinations).
				Build()

			//sort routers because JSON does not guarantee ordering and it causes flakes
			actualRouters := proxyTmpl.ProxyState.Listeners[0].Routers
			sort.Slice(actualRouters, func(i, j int) bool {
				return actualRouters[i].String() < actualRouters[j].String()
			})

			actual := protoToJSON(t, proxyTmpl)
			expected := JSONToProxyTemplate(t, goldenValueBytes(t, name, actual, *update))

			//sort routers on listener from golden file
			expectedRouters := expected.ProxyState.Listeners[0].Routers
			sort.Slice(expectedRouters, func(i, j int) bool {
				return expectedRouters[i].String() < expectedRouters[j].String()
			})
			require.Equal(t, protoToJSON(t, expected), protoToJSON(t, proxyTmpl))
		})
	}
}

func TestBuildImplicitDestinations(t *testing.T) {
	api1Endpoints := resourcetest.Resource(catalog.ServiceEndpointsType, "api-1").
		WithOwner(resourcetest.Resource(catalog.ServiceType, "api-1").ID()).
		WithData(t, endpointsData).Build()

	api2Endpoints := resourcetest.Resource(catalog.ServiceEndpointsType, "api-2").
		WithOwner(resourcetest.Resource(catalog.ServiceType, "api-2").ID()).
		WithData(t, endpointsData).Build()

	api1Identity := &pbresource.Reference{
		Name:    "api1-identity",
		Tenancy: api1Endpoints.Id.Tenancy,
	}

	api2Identity := &pbresource.Reference{
		Name:    "api2-identity",
		Tenancy: api2Endpoints.Id.Tenancy,
	}

	proxyCfg := &pbmesh.ProxyConfiguration{
		DynamicConfig: &pbmesh.DynamicConfig{
			Mode: pbmesh.ProxyMode_PROXY_MODE_TRANSPARENT,
			TransparentProxy: &pbmesh.TransparentProxy{
				OutboundListenerPort: 15001,
			},
		},
	}

	destination1 := &intermediate.Destination{
		ServiceEndpoints: &intermediate.ServiceEndpoints{
			Resource:  api1Endpoints,
			Endpoints: endpointsData,
		},
		Identities: []*pbresource.Reference{api1Identity},
		VirtualIPs: []string{"1.1.1.1"},
	}

	destination2 := &intermediate.Destination{
		ServiceEndpoints: &intermediate.ServiceEndpoints{
			Resource:  api2Endpoints,
			Endpoints: endpointsData,
		},
		Identities: []*pbresource.Reference{api2Identity},
		VirtualIPs: []string{"2.2.2.2", "3.3.3.3"},
	}

	destination3 := &intermediate.Destination{
		Explicit: &pbmesh.Upstream{
			DestinationRef:  resource.Reference(api1Endpoints.Id, ""),
			DestinationPort: "tcp",
			Datacenter:      "dc1",
			ListenAddr: &pbmesh.Upstream_IpPort{
				IpPort: &pbmesh.IPPortAddress{Ip: "1.1.1.1", Port: 1234},
			},
		},
		ServiceEndpoints: &intermediate.ServiceEndpoints{
			Resource:  api1Endpoints,
			Endpoints: endpointsData,
		},
		Identities: []*pbresource.Reference{api1Identity},
	}

	cases := map[string]struct {
		destinations []*intermediate.Destination
	}{
		"destination/l4-single-implicit-destination-tproxy": {
			destinations: []*intermediate.Destination{destination1},
		},
		"destination/l4-multiple-implicit-destinations-tproxy": {
			destinations: []*intermediate.Destination{destination1, destination2},
		},
		"destination/l4-implicit-and-explicit-destinations-tproxy": {
			destinations: []*intermediate.Destination{destination2, destination3},
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			proxyTmpl := New(testProxyStateTemplateID(), testIdentityRef(), "foo.consul", "dc1", proxyCfg).
				BuildDestinations(c.destinations).
				Build()

			//sort routers because JSON does not guarantee ordering and it causes flakes
			actualRouters := proxyTmpl.ProxyState.Listeners[0].Routers
			sort.Slice(actualRouters, func(i, j int) bool {
				return actualRouters[i].String() < actualRouters[j].String()
			})

			actual := protoToJSON(t, proxyTmpl)
			expected := JSONToProxyTemplate(t, goldenValueBytes(t, name, actual, *update))

			//sort routers on listener from golden file
			expectedRouters := expected.ProxyState.Listeners[0].Routers
			sort.Slice(expectedRouters, func(i, j int) bool {
				return expectedRouters[i].String() < expectedRouters[j].String()
			})

			require.Equal(t, protoToJSON(t, expected), protoToJSON(t, proxyTmpl))
		})
	}
}

func Test_isMeshPort(t *testing.T) {
	cases := map[string]struct {
		protocol       pbcatalog.Protocol
		expectedResult bool
	}{
		"mesh protocol returns true": {
			protocol:       pbcatalog.Protocol_PROTOCOL_MESH,
			expectedResult: true,
		},
		"grpc protocol returns false": {
			protocol:       pbcatalog.Protocol_PROTOCOL_GRPC,
			expectedResult: false,
		},
		"tcp protocol returns false": {
			protocol:       pbcatalog.Protocol_PROTOCOL_TCP,
			expectedResult: false,
		},
		"http protocol returns false": {
			protocol:       pbcatalog.Protocol_PROTOCOL_HTTP,
			expectedResult: false,
		},
		"http2 protocol returns false": {
			protocol:       pbcatalog.Protocol_PROTOCOL_HTTP2,
			expectedResult: false,
		},
	}
	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			require.Equal(t, tc.expectedResult, isMeshPort(&pbcatalog.WorkloadPort{Protocol: tc.protocol}))
		})
	}
}