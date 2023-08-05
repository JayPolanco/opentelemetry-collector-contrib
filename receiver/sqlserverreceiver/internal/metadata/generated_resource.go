// Code generated by mdatagen. DO NOT EDIT.

package metadata

import (
	"go.opentelemetry.io/collector/pdata/pcommon"
)

// ResourceBuilder is a helper struct to build resources predefined in metadata.yaml.
// The ResourceBuilder is not thread-safe and must not to be used in multiple goroutines.
type ResourceBuilder struct {
	config ResourceAttributesConfig
	res    pcommon.Resource
}

// NewResourceBuilder creates a new ResourceBuilder. This method should be called on the start of the application.
func NewResourceBuilder(rac ResourceAttributesConfig) *ResourceBuilder {
	return &ResourceBuilder{
		config: rac,
		res:    pcommon.NewResource(),
	}
}

// SetSqlserverComputerName sets provided value as "sqlserver.computer.name" attribute.
func (rb *ResourceBuilder) SetSqlserverComputerName(val string) {
	if rb.config.SqlserverComputerName.Enabled {
		rb.res.Attributes().PutStr("sqlserver.computer.name", val)
	}
}

// SetSqlserverDatabaseName sets provided value as "sqlserver.database.name" attribute.
func (rb *ResourceBuilder) SetSqlserverDatabaseName(val string) {
	if rb.config.SqlserverDatabaseName.Enabled {
		rb.res.Attributes().PutStr("sqlserver.database.name", val)
	}
}

// SetSqlserverInstanceName sets provided value as "sqlserver.instance.name" attribute.
func (rb *ResourceBuilder) SetSqlserverInstanceName(val string) {
	if rb.config.SqlserverInstanceName.Enabled {
		rb.res.Attributes().PutStr("sqlserver.instance.name", val)
	}
}

// Emit returns the built resource and resets the internal builder state.
func (rb *ResourceBuilder) Emit() pcommon.Resource {
	r := rb.res
	rb.res = pcommon.NewResource()
	return r
}