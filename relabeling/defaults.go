package relabeling

import "github.com/itshikanov/prometheus-nginxlog-exporter/config"

// DefaultRelabelings are hardcoded relabeling configs that are always there
// and do not need to be explicitly configured
var DefaultRelabelings = []*Relabeling{
	{
		config.RelabelConfig{
			TargetLabel: "method",
			SourceValue: "request",
			Split:       1,
		},
	},
	{
		config.RelabelConfig{
			TargetLabel: "status",
			SourceValue: "status",
		},
	},
}
