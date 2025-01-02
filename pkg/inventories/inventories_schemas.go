package inventories

// InventoryRequestSchema is the schema for an inventory request
type InventoryRequestSchema struct {
	Name                         string `json:"name"`
	Description                  string `json:"description"`
	Organization                 int32  `json:"organization"`
	Kind                         string `json:"kind"`
	HostFilter                   string `json:"host_filter"`
	Variables                    string `json:"variables"`
	PreventInstanceGroupFallback bool   `json:"prevent_instance_group_fallback"`
}

// InventoryHostRequestSchema is the schema for an inventory host request
type InventoryHostRequestSchema struct {
	Name        string            `json:"name"`
	Description string            `json:"description"`
	Enabled     bool              `json:"enabled"`
	InstanceID  string            `json:"instance_id"`
	Variables   map[string]string `json:"variables"`
}

// InventoryRelatedResponseSchema is the schema for the related section of a response
type InventoryRelatedResponseSchema struct {
	Hosts                  string `json:"hosts"`
	VariableData           string `json:"variable_data"`
	Script                 string `json:"script"`
	ActivityStream         string `json:"activity_stream"`
	JobTemplates           string `json:"job_templates"`
	AdHocCommands          string `json:"ad_hoc_commands"`
	AccessList             string `json:"access_list"`
	ObjectRoles            string `json:"object_roles"`
	InstanceGroups         string `json:"instance_groups"`
	Copy                   string `json:"copy"`
	Labels                 string `json:"labels"`
	Groups                 string `json:"groups"`
	RootGroups             string `json:"root_groups"`
	UpdateInventorySources string `json:"update_inventory_sources"`
	InventorySources       string `json:"inventory_sources"`
	Tree                   string `json:"tree"`
	Organization           string `json:"organization"`
}

// InventoryResponseSingleSchema is the schema for a single inventory response item
type InventoryResponseSingleSchema struct {
	ID      int32                          `json:"id"`
	Type    string                         `json:"type"`
	URL     string                         `json:"url"`
	Related InventoryRelatedResponseSchema `json:"related"`
	InventoryRequestSchema
	Created                      string `json:"created"`
	Modified                     string `json:"modified"`
	HasActiveFailures            bool   `json:"has_active_failures"`
	TotalHosts                   int32  `json:"total_hosts"`
	HostsWithActiveFailures      int32  `json:"hosts_with_active_failures"`
	TotalGroups                  int32  `json:"total_groups"`
	HasInventorySources          bool   `json:"has_inventory_sources"`
	TotalInventorySources        int32  `json:"total_inventory_sources"`
	InventorySourcesWithFailures int32  `json:"inventory_sources_with_failures"`
	PendingDeletion              bool   `json:"pending_deletion"`
}

// InventoryResponseSchema is the schema for an inventory response
type InventoryResponseSchema struct {
	Count    int32                           `json:"count"`
	Next     string                          `json:"next"`
	Previous string                          `json:"previous"`
	Results  []InventoryResponseSingleSchema `json:"results"`
}
