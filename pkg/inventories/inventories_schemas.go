package inventories

import "github.com/btr1975/go-ansible-aap-api-client/pkg/hosts"

// InventoryRequestSchema is the schema for an inventory request
type InventoryRequestSchema struct {
	Name                         string `json:"name" yaml:"name"`
	Description                  string `json:"description" yaml:"description"`
	Organization                 int32  `json:"organization" yaml:"organization"`
	Kind                         string `json:"kind" yaml:"kind"`
	HostFilter                   string `json:"host_filter" yaml:"host_filter"`
	Variables                    string `json:"variables" yaml:"variables"`
	PreventInstanceGroupFallback bool   `json:"prevent_instance_group_fallback" yaml:"prevent_instance_group_fallback"`
}

// InventoryHostRequestSchema is the schema for an inventory host request
type InventoryHostRequestSchema struct {
	Name        string            `json:"name" yaml:"name"`
	Description string            `json:"description" yaml:"description"`
	Enabled     bool              `json:"enabled" yaml:"enabled"`
	InstanceID  string            `json:"instance_id" yaml:"instance_id"`
	Variables   map[string]string `json:"variables" yaml:"variables"`
}

// InventoryRelatedResponseSchema is the schema for the related section of a response
type InventoryRelatedResponseSchema struct {
	Hosts                  string `json:"hosts" yaml:"hosts"`
	VariableData           string `json:"variable_data" yaml:"variable_data"`
	Script                 string `json:"script" yaml:"script"`
	ActivityStream         string `json:"activity_stream" yaml:"activity_stream"`
	JobTemplates           string `json:"job_templates" yaml:"job_templates"`
	AdHocCommands          string `json:"ad_hoc_commands" yaml:"ad_hoc_commands"`
	AccessList             string `json:"access_list" yaml:"access_list"`
	ObjectRoles            string `json:"object_roles" yaml:"object_roles"`
	InstanceGroups         string `json:"instance_groups" yaml:"instance_groups"`
	Copy                   string `json:"copy" yaml:"copy"`
	Labels                 string `json:"labels" yaml:"labels"`
	Groups                 string `json:"groups" yaml:"groups"`
	RootGroups             string `json:"root_groups" yaml:"root_groups"`
	UpdateInventorySources string `json:"update_inventory_sources" yaml:"update_inventory_sources"`
	InventorySources       string `json:"inventory_sources" yaml:"inventory_sources"`
	Tree                   string `json:"tree" yaml:"tree"`
	Organization           string `json:"organization" yaml:"organization"`
}

// InventoryResponseSingleSchema is the schema for a single inventory response item
type InventoryResponseSingleSchema struct {
	ID      int32                          `json:"id" yaml:"id"`
	Type    string                         `json:"type" yaml:"type"`
	URL     string                         `json:"url" yaml:"url"`
	Related InventoryRelatedResponseSchema `json:"related" yaml:"related"`
	InventoryRequestSchema
	Created                      string `json:"created" yaml:"created"`
	Modified                     string `json:"modified" yaml:"modified"`
	HasActiveFailures            bool   `json:"has_active_failures" yaml:"has_active_failures"`
	TotalHosts                   int32  `json:"total_hosts" yaml:"total_hosts"`
	HostsWithActiveFailures      int32  `json:"hosts_with_active_failures" yaml:"hosts_with_active_failures"`
	TotalGroups                  int32  `json:"total_groups" yaml:"total_groups"`
	HasInventorySources          bool   `json:"has_inventory_sources" yaml:"has_inventory_sources"`
	TotalInventorySources        int32  `json:"total_inventory_sources" yaml:"total_inventory_sources"`
	InventorySourcesWithFailures int32  `json:"inventory_sources_with_failures" yaml:"inventory_sources_with_failures"`
	PendingDeletion              bool   `json:"pending_deletion" yaml:"pending_deletion"`
}

// InventoryResponseSchema is the schema for an inventory response
type InventoryResponseSchema struct {
	Count    int32                           `json:"count" yaml:"count"`
	Next     string                          `json:"next" yaml:"next"`
	Previous string                          `json:"previous" yaml:"previous"`
	Results  []InventoryResponseSingleSchema `json:"results" yaml:"results"`
}

// CustomGroupHostSchema is the schema for a custom group host
type CustomGroupHostSchema struct {
	GroupName string
	Host      hosts.HostRequestSchema
}

// CustomGroupsIDSchema is the schema for a custom groups IDs
type CustomGroupsIDSchema struct {
	GroupName string
	GroupID   int32
}
