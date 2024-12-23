package hosts

// HostRequestSchema is the schema for an host request
type HostRequestSchema struct {
	Name        string            `json:"name"`
	Description string            `json:"description"`
	Inventory   int32             `json:"inventory"`
	Enabled     bool              `json:"enabled"`
	InstanceID  string            `json:"instance_id"`
	Variables   map[string]string `json:"variables"`
}

// HostRelatedResponseSchema is the schema for the related section of a response
type HostRelatedResponseSchema struct {
	CreatedBy          string `json:"created_by"`
	ModifiedBy         string `json:"modified_by"`
	VariableData       string `json:"variable_data"`
	Groups             string `json:"groups"`
	AllGroups          string `json:"all_groups"`
	JobEvents          string `json:"job_events"`
	JobHostSummaries   string `json:"job_host_summaries"`
	ActivityStream     string `json:"activity_stream"`
	InventorySources   string `json:"inventory_sources"`
	SmartInventories   string `json:"smart_inventories"`
	AdHocCommands      string `json:"ad_hoc_commands"`
	AdHocCommandEvents string `json:"ad_hoc_command_events"`
	AnsibleFacts       string `json:"ansible_facts"`
	Inventory          string `json:"inventory"`
	LastJob            string `json:"last_job"`
	LastJobHostSummary string `json:"last_job_host_summary"`
}

// HostResponseSingleSchema is the schema for a single host response item
type HostResponseSingleSchema struct {
	ID      int32                     `json:"id"`
	Type    string                    `json:"type"`
	URL     string                    `json:"url"`
	Related HostRelatedResponseSchema `json:"related"`
	HostRequestSchema
	Variables            string `json:"variables"` // Redefine this as a string
	Created              string `json:"created"`
	Modified             string `json:"modified"`
	HasActiveFailures    bool   `json:"has_active_failures"`
	HasInventorySources  bool   `json:"has_inventory_sources"`
	LastJob              int32  `json:"last_job"`
	LastJobHostSummary   int32  `json:"last_job_host_summary"`
	AnsibleFactsModified string `json:"ansible_facts_modified"`
}

// HostResponseSchema is the schema for an hosts response
type HostResponseSchema struct {
	Count    int32                      `json:"count"`
	Next     string                     `json:"next"`
	Previous string                     `json:"previous"`
	Results  []HostResponseSingleSchema `json:"results"`
}
