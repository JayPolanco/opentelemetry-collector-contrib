// Code generated by mdatagen. DO NOT EDIT.

package metadata

import "go.opentelemetry.io/collector/confmap"

// MetricConfig provides common config for a particular metric.
type MetricConfig struct {
	Enabled bool `mapstructure:"enabled"`

	enabledSetByUser bool
}

func (ms *MetricConfig) Unmarshal(parser *confmap.Conf) error {
	if parser == nil {
		return nil
	}
	err := parser.Unmarshal(ms, confmap.WithErrorUnused())
	if err != nil {
		return err
	}
	ms.enabledSetByUser = parser.IsSet("enabled")
	return nil
}

// MetricsConfig provides config for aerospikereceiver metrics.
type MetricsConfig struct {
	AerospikeNamespaceDiskAvailable                   MetricConfig `mapstructure:"aerospike.namespace.disk.available"`
	AerospikeNamespaceGeojsonRegionQueryCells         MetricConfig `mapstructure:"aerospike.namespace.geojson.region_query_cells"`
	AerospikeNamespaceGeojsonRegionQueryFalsePositive MetricConfig `mapstructure:"aerospike.namespace.geojson.region_query_false_positive"`
	AerospikeNamespaceGeojsonRegionQueryPoints        MetricConfig `mapstructure:"aerospike.namespace.geojson.region_query_points"`
	AerospikeNamespaceGeojsonRegionQueryRequests      MetricConfig `mapstructure:"aerospike.namespace.geojson.region_query_requests"`
	AerospikeNamespaceMemoryFree                      MetricConfig `mapstructure:"aerospike.namespace.memory.free"`
	AerospikeNamespaceMemoryUsage                     MetricConfig `mapstructure:"aerospike.namespace.memory.usage"`
	AerospikeNamespaceQueryCount                      MetricConfig `mapstructure:"aerospike.namespace.query.count"`
	AerospikeNamespaceScanCount                       MetricConfig `mapstructure:"aerospike.namespace.scan.count"`
	AerospikeNamespaceTransactionCount                MetricConfig `mapstructure:"aerospike.namespace.transaction.count"`
	AerospikeNodeConnectionCount                      MetricConfig `mapstructure:"aerospike.node.connection.count"`
	AerospikeNodeConnectionOpen                       MetricConfig `mapstructure:"aerospike.node.connection.open"`
	AerospikeNodeMemoryFree                           MetricConfig `mapstructure:"aerospike.node.memory.free"`
	AerospikeNodeQueryTracked                         MetricConfig `mapstructure:"aerospike.node.query.tracked"`
}

func DefaultMetricsConfig() MetricsConfig {
	return MetricsConfig{
		AerospikeNamespaceDiskAvailable: MetricConfig{
			Enabled: true,
		},
		AerospikeNamespaceGeojsonRegionQueryCells: MetricConfig{
			Enabled: true,
		},
		AerospikeNamespaceGeojsonRegionQueryFalsePositive: MetricConfig{
			Enabled: true,
		},
		AerospikeNamespaceGeojsonRegionQueryPoints: MetricConfig{
			Enabled: true,
		},
		AerospikeNamespaceGeojsonRegionQueryRequests: MetricConfig{
			Enabled: true,
		},
		AerospikeNamespaceMemoryFree: MetricConfig{
			Enabled: true,
		},
		AerospikeNamespaceMemoryUsage: MetricConfig{
			Enabled: true,
		},
		AerospikeNamespaceQueryCount: MetricConfig{
			Enabled: true,
		},
		AerospikeNamespaceScanCount: MetricConfig{
			Enabled: true,
		},
		AerospikeNamespaceTransactionCount: MetricConfig{
			Enabled: true,
		},
		AerospikeNodeConnectionCount: MetricConfig{
			Enabled: true,
		},
		AerospikeNodeConnectionOpen: MetricConfig{
			Enabled: true,
		},
		AerospikeNodeMemoryFree: MetricConfig{
			Enabled: true,
		},
		AerospikeNodeQueryTracked: MetricConfig{
			Enabled: true,
		},
	}
}

// ResourceAttributeConfig provides common config for a particular resource attribute.
type ResourceAttributeConfig struct {
	Enabled bool `mapstructure:"enabled"`
}

// ResourceAttributesConfig provides config for aerospikereceiver resource attributes.
type ResourceAttributesConfig struct {
	AerospikeNamespace ResourceAttributeConfig `mapstructure:"aerospike.namespace"`
	AerospikeNodeName  ResourceAttributeConfig `mapstructure:"aerospike.node.name"`
}

func DefaultResourceAttributesConfig() ResourceAttributesConfig {
	return ResourceAttributesConfig{
		AerospikeNamespace: ResourceAttributeConfig{
			Enabled: true,
		},
		AerospikeNodeName: ResourceAttributeConfig{
			Enabled: true,
		},
	}
}

// MetricsBuilderConfig is a configuration for aerospikereceiver metrics builder.
type MetricsBuilderConfig struct {
	Metrics            MetricsConfig            `mapstructure:"metrics"`
	ResourceAttributes ResourceAttributesConfig `mapstructure:"resource_attributes"`
}

func DefaultMetricsBuilderConfig() MetricsBuilderConfig {
	return MetricsBuilderConfig{
		Metrics:            DefaultMetricsConfig(),
		ResourceAttributes: DefaultResourceAttributesConfig(),
	}
}
