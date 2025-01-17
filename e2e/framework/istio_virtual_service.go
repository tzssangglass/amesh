package framework

import (
	"github.com/api7/amesh/e2e/framework/utils"
)

type RouteFaultAbortRule struct {
	StatusCode int
	Percentage int
}

type RouteFaultDelayRule struct {
	Duration   float32
	Percentage int
}

type RouteFaultRule struct {
	Abort *RouteFaultAbortRule
	Delay *RouteFaultDelayRule
}

type RouteMatchRule struct {
	Headers map[string]string
}

type RouteDestinationConfig struct {
	Weight int
}

type RouteConfig struct {
	Match        *RouteMatchRule
	Fault        *RouteFaultRule
	Destinations map[string]*RouteDestinationConfig
	Timeout      float32
}

type VirtualServiceConfig struct {
	Host         string
	Destinations []string
	Routes       []*RouteConfig
}

const (
	virtualServiceTemplate = `
apiVersion: networking.istio.io/v1alpha3
kind: DestinationRule
metadata:
  name: {{ .Host }}
spec:
  host: {{ .Host }}
  subsets:
{{- range $idx, $destKey := .Destinations }}
  - name: {{ $destKey }}
    labels:
      version: {{ $destKey }}
{{- end }}
---
apiVersion: networking.istio.io/v1alpha3
kind: VirtualService
metadata:
  name: {{ .Host }}
spec:
  hosts:
  - {{ .Host }}
  http:
{{- $host := .Host }}
{{- range $index, $routeConfig := .Routes }}
  - name: route-{{ $index }}

{{- if and $routeConfig.Match (gt (len $routeConfig.Match.Headers) 0) }}
    match:
{{- if $routeConfig.Match.Headers }}
    - headers:
{{- range $headerKey, $headerValue := $routeConfig.Match.Headers }}
        {{ $headerKey }}:
          exact: {{ $headerValue }}
{{- end }}
{{- end }}
{{- end }}
    route:

{{- range $dest, $destConfig := $routeConfig.Destinations }}
    - destination:
        host: {{ $host }}
        subset: {{ $dest }}
      weight: {{ .Weight }}
{{- end }}

{{- if $routeConfig.Timeout }}
    timeout: {{ $routeConfig.Timeout }}s
{{- end }}

{{- if and $routeConfig.Fault (or $routeConfig.Fault.Abort $routeConfig.Fault.Delay) }}
    fault:
{{- if $routeConfig.Fault.Abort }}
      abort:
        httpStatus: {{ $routeConfig.Fault.Abort.StatusCode }}
        percentage:
          value: {{ $routeConfig.Fault.Abort.Percentage }}
{{- end }}
{{- if $routeConfig.Fault.Delay }}
      delay:
        fixedDelay: {{ $routeConfig.Fault.Delay.Duration }}s
        percentage:
          value: {{ $routeConfig.Fault.Delay.Percentage }}
{{- end }}
{{- end }}

{{- end }}
`
)

func (vs *VirtualServiceConfig) GenerateYAML() string {
	artifact, err := utils.RenderManifest(virtualServiceTemplate, vs)
	utils.AssertNil(err, "render virtual service template")
	return artifact
}

func (f *Framework) ApplyVirtualService(vs *VirtualServiceConfig) error {
	yaml := vs.GenerateYAML()
	return f.ApplyResourceFromString(yaml)
}
