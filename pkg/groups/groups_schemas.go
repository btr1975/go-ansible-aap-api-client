package groups

// GroupRequestSchema is the schema for an group request
type GroupRequestSchema struct {
	Name        string `json:"name" yaml:"name"`
	Description string `json:"description" yaml:"description"`
	Inventory   int32  `json:"inventory" yaml:"inventory"`
	Variables   string `json:"variables" yaml:"variables"`
}

// GroupRelatedResponseSchema is the schema for the related section of a response
type GroupRelatedResponseSchema struct {
	CreatedBy         string `json:"created_by" yaml:"created_by"`
	ModifiedBy        string `json:"modified_by" yaml:"modified_by"`
	VariableData      string `json:"variable_data" yaml:"variable_data"`
	Hosts             string `json:"hosts" yaml:"hosts"`
	PotentialChildren string `json:"potential_children" yaml:"potential_children"`
	Children          string `json:"children" yaml:"children"`
	AllHosts          string `json:"all_hosts" yaml:"all_hosts"`
	JobEvents         string `json:"job_events" yaml:"job_events"`
	JobHostSummaries  string `json:"job_host_summaries" yaml:"job_host_summaries"`
	ActivityStream    string `json:"activity_stream" yaml:"activity_stream"`
	InventorySources  string `json:"inventory_sources" yaml:"inventory_sources"`
	AdHocCommands     string `json:"ad_hoc_commands" yaml:"ad_hoc_commands"`
	Inventory         string `json:"inventory" yaml:"inventory"`
}

// GroupResponseSingleSchema is the schema for a single group response item
type GroupResponseSingleSchema struct {
	ID      int32                      `json:"id" yaml:"id"`
	Type    string                     `json:"type" yaml:"type"`
	URL     string                     `json:"url" yaml:"url"`
	Related GroupRelatedResponseSchema `json:"related" yaml:"related"`
	GroupRequestSchema
	Created  string `json:"created" yaml:"created"`
	Modified string `json:"modified" yaml:"modified"`
}

// GroupResponseSchema is the schema for an groups response
type GroupResponseSchema struct {
	Count    int32                       `json:"count" yaml:"count"`
	Next     string                      `json:"next" yaml:"next"`
	Previous string                      `json:"previous" yaml:"previous"`
	Results  []GroupResponseSingleSchema `json:"results" yaml:"results"`
}

// GroupGeneralNetwork is the schema for general network device group vars
type GroupGeneralNetwork struct {
	AnsibleConnection   string `json:"ansible_connection" yaml:"ansible_connection"`
	AnsibleBecome       bool   `json:"ansible_become" yaml:"ansible_become"`
	AnsibleBecomeMethod string `json:"ansible_become_method" yaml:"ansible_become_method"`
	AnsibleNetworkOS    string `json:"ansible_network_os " yaml:"ansible_network_os"`
}
