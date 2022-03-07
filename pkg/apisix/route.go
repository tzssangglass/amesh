package apisix

// RouteStatus Enumerations.
type RouteStatus int32

const (
	RouteDisable RouteStatus = 0
	RouteEnable  RouteStatus = 1
)

// A Route contains multiple parts but basically can be grouped
// into three:
// 1). Route match, fields like uris, hosts, remote_addrs are the
// predicates to indicate whether a request can hit the route.
// 2). Route action, upstream_id specifies the backend upstream
// object, which guides Apache APISIX how to route request.
// 3). Plugins, plugins will run before/after the route action,
// some plugins are "terminated" so may be requests will be returned
// on the APISIX side (like authentication failures).
type Route struct {
	// URI array used to do the route match.
	// At least one item should be configured and each of them cannot be
	// duplicated.
	Uris []string `json:"uris,omitempty"`
	// The route name, it's useful for the logging but it's not required.
	Name string `json:"name,omitempty"`
	// The route id.
	Id string `json:"id,omitempty"`
	// Textual descriptions used to describe the route use.
	Desc string `json:"desc,omitempty"`
	// Priority of this route, used to decide which route should be used when
	// multiple routes contains same URI.
	// Larger value means higher priority. The default value is 0.
	Priority int32 `json:"priority,omitempty"`
	// HTTP Methods used to do the route match.
	Methods []string `json:"methods,omitempty"`
	// Host array used to do the route match.
	Hosts []string `json:"hosts,omitempty"`
	// Remote address array used to do the route match.
	RemoteAddrs []string `json:"remote_addrs,omitempty"`
	// Nginx vars used to do the route match.
	Vars []*Var `json:"vars,omitempty"`
	// Embedded plugins.
	Plugins map[string]interface{} `json:"plugins,omitempty"`
	// The referred service id.
	ServiceId string `json:"service_id,omitempty"`
	// The referred upstream id.
	UpstreamId string `json:"upstream_id,omitempty"`
	// The route status.
	Status RouteStatus `json:"status,omitempty"`
	// Timeout sets the I/O operations timeouts on the route level.
	Timeout *Timeout `json:"timeout,omitempty"`
	// enable_websocket indicates whether the websocket proxy is enabled.
	EnableWebsocket bool `json:"enable_websocket,omitempty"`
	// Labels contains some labels for the sake of management.
	Labels map[string]string `json:"labels,omitempty"  `
	// create_time indicate the create timestamp of this route.
	CreateTime int64 `json:"create_time,omitempty"`
	// update_time indicate the last update timestamp of this route.
	UpdateTime int64 `json:"update_time,omitempty"`
}
