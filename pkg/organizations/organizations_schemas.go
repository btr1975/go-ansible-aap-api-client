package organizations

// OrganizationRequestSchema is the schema for an organization request
type OrganizationRequestSchema struct {
	Name               string `json:"name" yaml:"name"`
	Description        string `json:"description" yaml:"description"`
	MaxHosts           int32  `json:"max_hosts" yaml:"max_hosts"`
	DefaultEnvironment string `json:"default_environment" yaml:"default_environment"`
}

// OrganizationRelatedResponseSchema is the schema for the related section of a response
type OrganizationRelatedResponseSchema struct {
	ExecutionEnvironments          string `json:"execution_environments" yaml:"execution_environments"`
	Projects                       string `json:"projects" yaml:"projects"`
	Inventories                    string `json:"inventories" yaml:"inventories"`
	JobTemplates                   string `json:"job_templates" yaml:"job_templates"`
	WorkflowJobTemplates           string `json:"workflow_job_templates" yaml:"workflow_job_templates"`
	Users                          string `json:"users" yaml:"users"`
	Admins                         string `json:"admins" yaml:"admins"`
	Teams                          string `json:"teams" yaml:"teams"`
	Credentials                    string `json:"credentials" yaml:"credentials"`
	Applications                   string `json:"applications" yaml:"applications"`
	ActivityStream                 string `json:"activity_stream" yaml:"activity_stream"`
	NotificationTemplates          string `json:"notification_templates" yaml:"notification_templates"`
	NotificationTemplatesStarted   string `json:"notification_templates_started" yaml:"notification_templates_started"`
	NotificationTemplatesSuccess   string `json:"notification_templates_success" yaml:"notification_templates_success"`
	NotificationTemplatesError     string `json:"notification_templates_error" yaml:"notification_templates_error"`
	NotificationTemplatesApprovals string `json:"notification_templates_approvals" yaml:"notification_templates_approvals"`
	ObjectRoles                    string `json:"object_roles" yaml:"object_roles"`
	AccessList                     string `json:"access_list" yaml:"access_list"`
	InstanceGroups                 string `json:"instance_groups" yaml:"instance_groups"`
	GalaxyCredentials              string `json:"galaxy_credentials" yaml:"galaxy_credentials"`
}

// OrganizationResponseSingleSchema is the schema for a single organization response item
type OrganizationResponseSingleSchema struct {
	ID      int32                             `json:"id" yaml:"id"`
	Type    string                            `json:"type" yaml:"type"`
	URL     string                            `json:"url" yaml:"url"`
	Related OrganizationRelatedResponseSchema `json:"related" yaml:"related"`
	OrganizationRequestSchema
	Created          string `json:"created" yaml:"created"`
	Modified         string `json:"modified" yaml:"modified"`
	CustomVirtualenv bool   `json:"custom_virtualenv" yaml:"custom_virtualenv"`
}

// OrganizationResponseSchema is the schema for an organization response
type OrganizationResponseSchema struct {
	Count    int32                              `json:"count" yaml:"count"`
	Next     string                             `json:"next" yaml:"next"`
	Previous string                             `json:"previous" yaml:"previous"`
	Results  []OrganizationResponseSingleSchema `json:"results" yaml:"results"`
}
