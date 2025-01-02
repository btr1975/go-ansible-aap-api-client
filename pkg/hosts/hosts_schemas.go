package hosts

// HostRequestSchema is the schema for an host request
type HostRequestSchema struct {
	Name        string `json:"name" yaml:"name"`
	Description string `json:"description" yaml:"description"`
	Inventory   int32  `json:"inventory" yaml:"inventory"`
	Enabled     bool   `json:"enabled" yaml:"enabled"`
	InstanceID  string `json:"instance_id" yaml:"instance_id"`
	Variables   string `json:"variables" yaml:"variables"`
}

// HostRelatedResponseSchema is the schema for the related section of a response
type HostRelatedResponseSchema struct {
	CreatedBy          string `json:"created_by" yaml:"created_by"`
	ModifiedBy         string `json:"modified_by" yaml:"modified_by"`
	VariableData       string `json:"variable_data" yaml:"variable_data"`
	Groups             string `json:"groups" yaml:"groups"`
	AllGroups          string `json:"all_groups" yaml:"all_groups"`
	JobEvents          string `json:"job_events" yaml:"job_events"`
	JobHostSummaries   string `json:"job_host_summaries" yaml:"job_host_summaries"`
	ActivityStream     string `json:"activity_stream" yaml:"activity_stream"`
	InventorySources   string `json:"inventory_sources" yaml:"inventory_sources"`
	SmartInventories   string `json:"smart_inventories" yaml:"smart_inventories"`
	AdHocCommands      string `json:"ad_hoc_commands" yaml:"ad_hoc_commands"`
	AdHocCommandEvents string `json:"ad_hoc_command_events" yaml:"ad_hoc_command_events"`
	AnsibleFacts       string `json:"ansible_facts" yaml:"ansible_facts"`
	Inventory          string `json:"inventory" yaml:"inventory"`
	LastJob            string `json:"last_job" yaml:"last_job"`
	LastJobHostSummary string `json:"last_job_host_summary" yaml:"last_job_host_summary"`
}

// HostResponseSingleSchema is the schema for a single host response item
type HostResponseSingleSchema struct {
	ID      int32                     `json:"id" yaml:"id"`
	Type    string                    `json:"type" yaml:"type"`
	URL     string                    `json:"url" yaml:"url"`
	Related HostRelatedResponseSchema `json:"related" yaml:"related"`
	HostRequestSchema
	Created              string `json:"created" yaml:"created"`
	Modified             string `json:"modified" yaml:"modified"`
	HasActiveFailures    bool   `json:"has_active_failures" yaml:"has_active_failures"`
	HasInventorySources  bool   `json:"has_inventory_sources" yaml:"has_inventory_sources"`
	LastJob              int32  `json:"last_job" yaml:"last_job"`
	LastJobHostSummary   int32  `json:"last_job_host_summary" yaml:"last_job_host_summary"`
	AnsibleFactsModified string `json:"ansible_facts_modified" yaml:"ansible_facts_modified"`
}

// HostResponseSchema is the schema for an hosts response
type HostResponseSchema struct {
	Count    int32                      `json:"count" yaml:"count"`
	Next     string                     `json:"next" yaml:"next"`
	Previous string                     `json:"previous" yaml:"previous"`
	Results  []HostResponseSingleSchema `json:"results" yaml:"results"`
}
