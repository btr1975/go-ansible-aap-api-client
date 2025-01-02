package jobs

// JobRelatedResponseSchema is the schema for the related section of a response
type JobRelatedResponseSchema struct {
	CreatedBy            string `json:"created_by" yaml:"created_by"`
	Labels               string `json:"labels" yaml:"labels"`
	Inventory            string `json:"inventory" yaml:"inventory"`
	Project              string `json:"project" yaml:"project"`
	Organization         string `json:"organization" yaml:"organization"`
	Credentials          string `json:"credentials" yaml:"credentials"`
	UnifiedJobTemplate   string `json:"unified_job_template" yaml:"unified_job_template"`
	StdOut               string `json:"stdout" yaml:"stdout"`
	ExecutionEnvironment string `json:"execution_environment" yaml:"execution_environment"`
	JobEvents            string `json:"job_events" yaml:"job_events"`
	JobHostSummaries     string `json:"job_host_summaries" yaml:"job_host_summaries"`
	ActivityStream       string `json:"activity_stream" yaml:"activity_stream"`
	Notifications        string `json:"notifications" yaml:"notifications"`
	CreateSchedule       string `json:"create_schedule" yaml:"create_schedule"`
	JobTemplate          string `json:"job_template" yaml:"job_template"`
	Cancel               string `json:"cancel" yaml:"cancel"`
	Relaunch             string `json:"relaunch" yaml:"relaunch"`
}

// JobResponseSingleSchema is the schema for a single job response item
type JobResponseSingleSchema struct {
	ID                   int32                    `json:"id" yaml:"id"`
	Type                 string                   `json:"type" yaml:"type"`
	URL                  string                   `json:"url" yaml:"url"`
	Related              JobRelatedResponseSchema `json:"related" yaml:"related"`
	Created              string                   `json:"created" yaml:"created"`
	Modified             string                   `json:"modified" yaml:"modified"`
	Name                 string                   `json:"name" yaml:"name"`
	Description          string                   `json:"description" yaml:"description"`
	UnifiedJobTemplate   int32                    `json:"unified_job_template" yaml:"unified_job_template"`
	LaunchType           string                   `json:"launch_type" yaml:"launch_type"`
	Status               string                   `json:"status" yaml:"status"`
	ExecutionEnvironment int32                    `json:"execution_environment" yaml:"execution_environment"`
	Failed               bool                     `json:"failed" yaml:"failed"`
	Started              string                   `json:"started" yaml:"started"`
	Finished             string                   `json:"finished" yaml:"finished"`
	CanceledOn           string                   `json:"canceled_on" yaml:"canceled_on"`
	Elapsed              float32                  `json:"elapsed" yaml:"elapsed"`
	JobExplanation       string                   `json:"job_explanation" yaml:"job_explanation"`
	ExecutionNode        string                   `json:"execution_node" yaml:"execution_node"`
	ControllerNode       string                   `json:"controller_node" yaml:"controller_node"`
}

// JobResponseSchema is the schema for an job response
type JobResponseSchema struct {
	Count    int32                     `json:"count" yaml:"count"`
	Next     string                    `json:"next" yaml:"next"`
	Previous string                    `json:"previous" yaml:"previous"`
	Results  []JobResponseSingleSchema `json:"results" yaml:"results"`
}
