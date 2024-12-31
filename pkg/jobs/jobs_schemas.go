package jobs

// JobRelatedResponseSchema is the schema for the related section of a response
type JobRelatedResponseSchema struct {
	CreatedBy            string `json:"created_by"`
	Labels               string `json:"labels"`
	Inventory            string `json:"inventory"`
	Project              string `json:"project"`
	Organization         string `json:"organization"`
	Credentials          string `json:"credentials"`
	UnifiedJobTemplate   string `json:"unified_job_template"`
	StdOut               string `json:"stdout"`
	ExecutionEnvironment string `json:"execution_environment"`
	JobEvents            string `json:"job_events"`
	JobHostSummaries     string `json:"job_host_summaries"`
	ActivityStream       string `json:"activity_stream"`
	Notifications        string `json:"notifications"`
	CreateSchedule       string `json:"create_schedule"`
	JobTemplate          string `json:"job_template"`
	Cancel               string `json:"cancel"`
	Relaunch             string `json:"relaunch"`
}

// JobResponseSingleSchema is the schema for a single job response item
type JobResponseSingleSchema struct {
	ID                   int32                    `json:"id"`
	Type                 string                   `json:"type"`
	URL                  string                   `json:"url"`
	Related              JobRelatedResponseSchema `json:"related"`
	Created              string                   `json:"created"`
	Modified             string                   `json:"modified"`
	Name                 string                   `json:"name"`
	Description          string                   `json:"description"`
	UnifiedJobTemplate   int32                    `json:"unified_job_template"`
	LaunchType           string                   `json:"launch_type"`
	Status               string                   `json:"status"`
	ExecutionEnvironment int32                    `json:"execution_environment"`
	Failed               bool                     `json:"failed"`
	Started              string                   `json:"started"`
	Finished             string                   `json:"finished"`
	CanceledOn           string                   `json:"canceled_on"`
	Elapsed              float32                  `json:"elapsed"`
	JobExplanation       string                   `json:"job_explanation"`
	ExecutionNode        string                   `json:"execution_node"`
	ControllerNode       string                   `json:"controller_node"`
}

// JobResponseSchema is the schema for an job response
type JobResponseSchema struct {
	Count    int32                     `json:"count"`
	Next     string                    `json:"next"`
	Previous string                    `json:"previous"`
	Results  []JobResponseSingleSchema `json:"results"`
}
