package jobtemplates

// JobTemplateRequestSchema is the schema for an job templates request
type JobTemplateRequestSchema struct {
	Name                            string `json:"name"`
	Description                     string `json:"description"`
	JobType                         string `json:"job_type"`
	Inventory                       int32  `json:"inventory"`
	Project                         int32  `json:"project"`
	Playbook                        string `json:"playbook"`
	ScmBranch                       string `json:"scm_branch"`
	Forks                           int32  `json:"forks"`
	Limit                           string `json:"limit"`
	Verbosity                       int32  `json:"verbosity"`
	ExtraVars                       string `json:"extra_vars"`
	JobTags                         string `json:"job_tags"`
	ForceHandlers                   bool   `json:"force_handlers"`
	SkipTags                        string `json:"skip_tags"`
	StartAtTask                     string `json:"start_at_task"`
	Timeout                         int32  `json:"timeout"`
	UseFactCache                    bool   `json:"use_fact_cache"`
	ExecutionEnvironment            string `json:"execution_environment"`
	HostConfigKey                   string `json:"host_config_key"`
	AskScmBranchOnLaunch            bool   `json:"ask_scm_branch_on_launch"`
	AskDiffModeOnLaunch             bool   `json:"ask_diff_mode_on_launch"`
	AskVariablesOnLaunch            bool   `json:"ask_variables_on_launch"`
	AskLimitOnLaunch                bool   `json:"ask_limit_on_launch"`
	AskTagsOnLaunch                 bool   `json:"ask_tags_on_launch"`
	AskSkipTagsOnLaunch             bool   `json:"ask_skip_tags_on_launch"`
	AskJobTypeOnLaunch              bool   `json:"ask_job_type_on_launch"`
	AskVerbosityOnLaunch            bool   `json:"ask_verbosity_on_launch"`
	AskInventoryOnLaunch            bool   `json:"ask_inventory_on_launch"`
	AskCredentialOnLaunch           bool   `json:"ask_credential_on_launch"`
	AskExecutionEnvironmentOnLaunch bool   `json:"ask_execution_environment_on_launch"`
	AskLabelsOnLaunch               bool   `json:"ask_labels_on_launch"`
	AskForksOnLaunch                bool   `json:"ask_forks_on_launch"`
	AskJobSliceCountOnLaunch        bool   `json:"ask_job_slice_count_on_launch"`
	AskTimeoutOnLaunch              bool   `json:"ask_timeout_on_launch"`
	AskInstanceGroupsOnLaunch       bool   `json:"ask_instance_groups_on_launch"`
	SurveyEnabled                   bool   `json:"survey_enabled"`
	BecomeEnabled                   bool   `json:"become_enabled"`
	DiffMode                        bool   `json:"diff_mode"`
	AllowSimultaneous               bool   `json:"allow_simultaneous"`
	JobSliceCount                   int32  `json:"job_slice_count"`
	WebhookService                  string `json:"webhook_service"`
	WebhookCredential               string `json:"webhook_credential"`
	PreventInstanceGroupFallback    bool   `json:"prevent_instance_group_fallback"`
}

// JobTemplateRelatedResponseSchema is the schema for the related section of a response
type JobTemplateRelatedResponseSchema struct {
	Labels                       string `json:"labels"`
	Inventory                    string `json:"inventory"`
	Project                      string `json:"project"`
	Organization                 string `json:"organization"`
	Credentials                  string `json:"credentials"`
	LastJob                      string `json:"last_job"`
	Jobs                         string `json:"jobs"`
	Schedules                    string `json:"schedules"`
	ActivityStream               string `json:"activity_stream"`
	Launch                       string `json:"launch"`
	WebhookKey                   string `json:"webhook_key"`
	WebhookReceiver              string `json:"webhook_receiver"`
	NotificationTemplatesStarted string `json:"notification_templates_started"`
	NotificationTemplatesSuccess string `json:"notification_templates_success"`
	NotificationTemplatesError   string `json:"notification_templates_error"`
	AccessList                   string `json:"access_list"`
	SurveySpec                   string `json:"survey_spec"`
	ObjectRoles                  string `json:"object_roles"`
	InstanceGroups               string `json:"instance_groups"`
	SliceWorkflowJobs            string `json:"slice_workflow_jobs"`
	Copy                         string `json:"copy"`
}

// JobTemplateResponseSingleSchema is the schema for a single job template response item
type JobTemplateResponseSingleSchema struct {
	ID      int32                            `json:"id"`
	Type    string                           `json:"type"`
	URL     string                           `json:"url"`
	Related JobTemplateRelatedResponseSchema `json:"related"`
	JobTemplateRequestSchema
	Created  string `json:"created"`
	Modified string `json:"modified"`
}

// JobTemplateResponseSchema is the schema for an job response
type JobTemplateResponseSchema struct {
	Count    int32                             `json:"count"`
	Next     string                            `json:"next"`
	Previous string                            `json:"previous"`
	Results  []JobTemplateResponseSingleSchema `json:"results"`
}
