package jobtemplates

// JobTemplateRequestSchema is the schema for a job templates request
type JobTemplateRequestSchema struct {
	Name                            string `json:"name" yaml:"name"`
	Description                     string `json:"description" yaml:"description"`
	JobType                         string `json:"job_type" yaml:"job_type"`
	Inventory                       int32  `json:"inventory" yaml:"inventory"`
	Project                         int32  `json:"project" yaml:"project"`
	Playbook                        string `json:"playbook" yaml:"playbook"`
	ScmBranch                       string `json:"scm_branch" yaml:"scm_branch"`
	Forks                           int32  `json:"forks" yaml:"forks"`
	Limit                           string `json:"limit" yaml:"limit"`
	Verbosity                       int32  `json:"verbosity" yaml:"verbosity"`
	ExtraVars                       string `json:"extra_vars" yaml:"extra_vars"`
	JobTags                         string `json:"job_tags" yaml:"job_tags"`
	ForceHandlers                   bool   `json:"force_handlers" yaml:"force_handlers"`
	SkipTags                        string `json:"skip_tags" yaml:"skip_tags"`
	StartAtTask                     string `json:"start_at_task" yaml:"start_at_task"`
	Timeout                         int32  `json:"timeout" yaml:"timeout"`
	UseFactCache                    bool   `json:"use_fact_cache" yaml:"use_fact_cache"`
	ExecutionEnvironment            string `json:"execution_environment" yaml:"execution_environment"`
	HostConfigKey                   string `json:"host_config_key" yaml:"host_config_key"`
	AskScmBranchOnLaunch            bool   `json:"ask_scm_branch_on_launch" yaml:"ask_scm_branch_on_launch"`
	AskDiffModeOnLaunch             bool   `json:"ask_diff_mode_on_launch" yaml:"ask_diff_mode_on_launch"`
	AskVariablesOnLaunch            bool   `json:"ask_variables_on_launch" yaml:"ask_variables_on_launch"`
	AskLimitOnLaunch                bool   `json:"ask_limit_on_launch" yaml:"ask_limit_on_launch"`
	AskTagsOnLaunch                 bool   `json:"ask_tags_on_launch" yaml:"ask_tags_on_launch"`
	AskSkipTagsOnLaunch             bool   `json:"ask_skip_tags_on_launch" yaml:"ask_skip_tags_on_launch"`
	AskJobTypeOnLaunch              bool   `json:"ask_job_type_on_launch" yaml:"ask_job_type_on_launch"`
	AskVerbosityOnLaunch            bool   `json:"ask_verbosity_on_launch" yaml:"ask_verbosity_on_launch"`
	AskInventoryOnLaunch            bool   `json:"ask_inventory_on_launch" yaml:"ask_inventory_on_launch"`
	AskCredentialOnLaunch           bool   `json:"ask_credential_on_launch" yaml:"ask_credential_on_launch"`
	AskExecutionEnvironmentOnLaunch bool   `json:"ask_execution_environment_on_launch" yaml:"ask_execution_environment_on_launch"`
	AskLabelsOnLaunch               bool   `json:"ask_labels_on_launch" yaml:"ask_labels_on_launch"`
	AskForksOnLaunch                bool   `json:"ask_forks_on_launch" yaml:"ask_forks_on_launch"`
	AskJobSliceCountOnLaunch        bool   `json:"ask_job_slice_count_on_launch" yaml:"ask_job_slice_count_on_launch"`
	AskTimeoutOnLaunch              bool   `json:"ask_timeout_on_launch" yaml:"ask_timeout_on_launch"`
	AskInstanceGroupsOnLaunch       bool   `json:"ask_instance_groups_on_launch" yaml:"ask_instance_groups_on_launch"`
	SurveyEnabled                   bool   `json:"survey_enabled" yaml:"survey_enabled"`
	BecomeEnabled                   bool   `json:"become_enabled" yaml:"become_enabled"`
	DiffMode                        bool   `json:"diff_mode" yaml:"diff_mode"`
	AllowSimultaneous               bool   `json:"allow_simultaneous" yaml:"allow_simultaneous"`
	JobSliceCount                   int32  `json:"job_slice_count" yaml:"job_slice_count"`
	WebhookService                  string `json:"webhook_service" yaml:"webhook_service"`
	WebhookCredential               string `json:"webhook_credential" yaml:"webhook_credential"`
	PreventInstanceGroupFallback    bool   `json:"prevent_instance_group_fallback" yaml:"prevent_instance_group_fallback"`
}

// JobTemplateSimpleRequestSchema is the schema for a simple job templates request
type JobTemplateSimpleRequestSchema struct {
	Inventory int32  `json:"inventory" yaml:"inventory"`
	ExtraVars string `json:"extra_vars" yaml:"extra_vars"`
}

// JobTemplateRelatedResponseSchema is the schema for the related section of a response
type JobTemplateRelatedResponseSchema struct {
	Labels                       string `json:"labels" yaml:"labels"`
	Inventory                    string `json:"inventory" yaml:"inventory"`
	Project                      string `json:"project" yaml:"project"`
	Organization                 string `json:"organization" yaml:"organization"`
	Credentials                  string `json:"credentials" yaml:"credentials"`
	LastJob                      string `json:"last_job" yaml:"last_job"`
	Jobs                         string `json:"jobs" yaml:"jobs"`
	Schedules                    string `json:"schedules" yaml:"schedules"`
	ActivityStream               string `json:"activity_stream" yaml:"activity_stream"`
	Launch                       string `json:"launch" yaml:"launch"`
	WebhookKey                   string `json:"webhook_key" yaml:"webhook_key"`
	WebhookReceiver              string `json:"webhook_receiver" yaml:"webhook_receiver"`
	NotificationTemplatesStarted string `json:"notification_templates_started" yaml:"notification_templates_started"`
	NotificationTemplatesSuccess string `json:"notification_templates_success" yaml:"notification_templates_success"`
	NotificationTemplatesError   string `json:"notification_templates_error" yaml:"notification_templates_error"`
	AccessList                   string `json:"access_list" yaml:"access_list"`
	SurveySpec                   string `json:"survey_spec" yaml:"survey_spec"`
	ObjectRoles                  string `json:"object_roles" yaml:"object_roles"`
	InstanceGroups               string `json:"instance_groups" yaml:"instance_groups"`
	SliceWorkflowJobs            string `json:"slice_workflow_jobs" yaml:"slice_workflow_jobs"`
	Copy                         string `json:"copy" yaml:"copy"`
}

// JobTemplateResponseSingleSchema is the schema for a single job template response item
type JobTemplateResponseSingleSchema struct {
	ID      int32                            `json:"id" yaml:"id"`
	Type    string                           `json:"type" yaml:"type"`
	URL     string                           `json:"url" yaml:"url"`
	Related JobTemplateRelatedResponseSchema `json:"related" yaml:"related"`
	JobTemplateRequestSchema
	Created  string `json:"created" yaml:"created"`
	Modified string `json:"modified" yaml:"modified"`
}

// JobTemplateResponseSchema is the schema for an job response
type JobTemplateResponseSchema struct {
	Count    int32                             `json:"count" yaml:"count"`
	Next     string                            `json:"next" yaml:"next"`
	Previous string                            `json:"previous" yaml:"previous"`
	Results  []JobTemplateResponseSingleSchema `json:"results" yaml:"results"`
}
