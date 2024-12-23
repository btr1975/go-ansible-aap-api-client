package groups

// GroupRequestSchema is the schema for an group request
type GroupRequestSchema struct {
	Name        string            `json:"name"`
	Description string            `json:"description"`
	Inventory   int32             `json:"inventory"`
	Variables   map[string]string `json:"variables"`
}

// GroupRelatedResponseSchema is the schema for the related section of a response
type GroupRelatedResponseSchema struct {
	CreatedBy         string `json:"created_by"`
	ModifiedBy        string `json:"modified_by"`
	VariableData      string `json:"variable_data"`
	Hosts             string `json:"hosts"`
	PotentialChildren string `json:"potential_children"`
	Children          string `json:"children"`
	AllHosts          string `json:"all_hosts"`
	JobEvents         string `json:"job_events"`
	JobHostSummaries  string `json:"job_host_summaries"`
	ActivityStream    string `json:"activity_stream"`
	InventorySources  string `json:"inventory_sources"`
	AdHocCommands     string `json:"ad_hoc_commands"`
	Inventory         string `json:"inventory"`
}

// GroupResponseSingleSchema is the schema for a single group response item
type GroupResponseSingleSchema struct {
	ID      int32                      `json:"id"`
	Type    string                     `json:"type"`
	URL     string                     `json:"url"`
	Related GroupRelatedResponseSchema `json:"related"`
	GroupRequestSchema
	Variables string `json:"variables"` // Redefine this as a string
	Created   string `json:"created"`
	Modified  string `json:"modified"`
}

// GroupResponseSchema is the schema for an groups response
type GroupResponseSchema struct {
	Count    int32                       `json:"count"`
	Next     string                      `json:"next"`
	Previous string                      `json:"previous"`
	Results  []GroupResponseSingleSchema `json:"results"`
}
