package organizations

// OrganizationRequestSchema is the schema for an organization request
type OrganizationRequestSchema struct {
	Name               string `json:"name"`
	Description        string `json:"description"`
	MaxHosts           int32  `json:"max_hosts"`
	DefaultEnvironment string `json:"default_environment"`
}

// OrganizationRelatedResponseSchema is the schema for the related section of a response
type OrganizationRelatedResponseSchema struct {
	ExecutionEnvironments          string `json:"execution_environments"`
	Projects                       string `json:"projects"`
	Inventories                    string `json:"inventories"`
	JobTemplates                   string `json:"job_templates"`
	WorkflowJobTemplates           string `json:"workflow_job_templates"`
	Users                          string `json:"users"`
	Admins                         string `json:"admins"`
	Teams                          string `json:"teams"`
	Credentials                    string `json:"credentials"`
	Applications                   string `json:"applications"`
	ActivityStream                 string `json:"activity_stream"`
	NotificationTemplates          string `json:"notification_templates"`
	NotificationTemplatesStarted   string `json:"notification_templates_started"`
	NotificationTemplatesSuccess   string `json:"notification_templates_success"`
	NotificationTemplatesError     string `json:"notification_templates_error"`
	NotificationTemplatesApprovals string `json:"notification_templates_approvals"`
	ObjectRoles                    string `json:"object_roles"`
	AccessList                     string `json:"access_list"`
	InstanceGroups                 string `json:"instance_groups"`
	GalaxyCredentials              string `json:"galaxy_credentials"`
}

// OrganizationResponseSingleSchema is the schema for a single organization response item
type OrganizationResponseSingleSchema struct {
	ID      int32                             `json:"id"`
	Type    string                            `json:"type"`
	URL     string                            `json:"url"`
	Related OrganizationRelatedResponseSchema `json:"related"`
	OrganizationRequestSchema
	Created          string `json:"created"`
	Modified         string `json:"modified"`
	CustomVirtualenv bool   `json:"custom_virtualenv"`
}

// OrganizationResponseSchema is the schema for an organization response
type OrganizationResponseSchema struct {
	Count    int32                              `json:"count"`
	Next     string                             `json:"next"`
	Previous string                             `json:"previous"`
	Results  []OrganizationResponseSingleSchema `json:"results"`
}
