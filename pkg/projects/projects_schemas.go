package projects

// ProjectRequestSchema is the schema for a project request
type ProjectRequestSchema struct {
	Name                          string `json:"name" yaml:"name"`
	Description                   string `json:"description" yaml:"description"`
	LocalPath                     string `json:"local_path" yaml:"local_path"`
	ScmType                       string `json:"scm_type" yaml:"scm_type"`
	ScmURL                        string `json:"scm_url" yaml:"scm_url"`
	ScmBranch                     string `json:"scm_branch" yaml:"scm_branch"`
	ScmRefSpec                    string `json:"scm_refspec" yaml:"scm_refspec"`
	ScmClean                      bool   `json:"scm_clean" yaml:"scm_clean"`
	ScmTrackSubmodules            bool   `json:"scm_track_submodules" yaml:"scm_track_submodules"`
	ScmDeleteOnUpdate             bool   `json:"scm_delete_on_update" yaml:"scm_delete_on_update"`
	Credential                    string `json:"credential" yaml:"credential"`
	Timeout                       int32  `json:"timeout" yaml:"timeout"`
	Organization                  string `json:"organization" yaml:"organization"`
	ScmUpdateOnLaunch             bool   `json:"scm_update_on_launch" yaml:"scm_update_on_launch"`
	ScmUpdateCacheTimeout         int32  `json:"scm_update_cache_timeout" yaml:"scm_update_cache_timeout"`
	AllowOverride                 bool   `json:"allow_override" yaml:"allow_override"`
	DefaultEnvironment            string `json:"default_environment" yaml:"default_environment"`
	SignatureValidationCredential string `json:"signature_validation_credential" yaml:"signature_validation_credential"`
}

// ProjectRelatedResponseSchema is the schema for the related section of a response
type ProjectRelatedResponseSchema struct {
	CreatedBy                    string `json:"created_by" yaml:"created_by"`
	ModifiedBy                   string `json:"modified_by" yaml:"modified_by"`
	LastJob                      string `json:"last_job" yaml:"last_job"`
	Teams                        string `json:"teams" yaml:"teams"`
	Playbooks                    string `json:"playbooks" yaml:"playbooks"`
	InventoryFiles               string `json:"inventory_files" yaml:"inventory_files"`
	Update                       string `json:"update" yaml:"update"`
	ProjectUpdates               string `json:"project_updates" yaml:"project_updates"`
	ScmInventorySources          string `json:"scm_inventory_sources" yaml:"scm_inventory_sources"`
	Schedules                    string `json:"schedules" yaml:"schedules"`
	ActivityStream               string `json:"activity_stream" yaml:"activity_stream"`
	NotificationTemplatesStarted string `json:"notification_templates_started" yaml:"notification_templates_started"`
	NotificationTemplatesSuccess string `json:"notification_templates_success" yaml:"notification_templates_success"`
	NotificationTemplatesError   string `json:"notification_templates_error" yaml:"notification_templates_error"`
	AccessList                   string `json:"access_list" yaml:"access_list"`
	ObjectRoles                  string `json:"object_roles" yaml:"object_roles"`
	Copy                         string `json:"copy" yaml:"copy"`
	Organization                 string `json:"organization" yaml:"organization"`
	LastUpdate                   string `json:"last_update" yaml:"last_update"`
}

// ProjectResponseSingleSchema is the schema for a single project response item
type ProjectResponseSingleSchema struct {
	ID      int32                        `json:"id" yaml:"id"`
	Type    string                       `json:"type" yaml:"type"`
	URL     string                       `json:"url" yaml:"url"`
	Related ProjectRelatedResponseSchema `json:"related" yaml:"related"`
	ProjectRequestSchema
	Created  string `json:"created" yaml:"created"`
	Modified string `json:"modified" yaml:"modified"`
}

// ProjectResponseSchema is the schema for a project response
type ProjectResponseSchema struct {
	Count    int32                         `json:"count" yaml:"count"`
	Next     string                        `json:"next" yaml:"next"`
	Previous string                        `json:"previous" yaml:"previous"`
	Results  []ProjectResponseSingleSchema `json:"results" yaml:"results"`
}
