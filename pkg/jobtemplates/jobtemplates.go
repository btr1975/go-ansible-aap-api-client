/*
Package jobtemplates provides a way to manipulate jobs templates for Ansible AAP
*/
package jobtemplates

import (
	"encoding/json"
	"fmt"
	"github.com/btr1975/go-ansible-aap-api-client/pkg/connection"
	"github.com/btr1975/go-ansible-aap-api-client/pkg/dataconversion"
)

// JobTemplate represents an AAP job template
type JobTemplate struct {
	URI            string
	connection     connection.BasicConnection
	DataConversion dataconversion.DataConverterInterface
}

// NewJobTemplate creates a new job template instance
//
//	:param basicConnection: The basic connection to use
func NewJobTemplate(basicConnection connection.BasicConnection) *JobTemplate {
	return &JobTemplate{
		URI:            "job_templates/",
		connection:     basicConnection,
		DataConversion: dataconversion.NewDataConverter(),
	}
}

// GetAllJobTemplates gets all job templates
func (jobTemplate *JobTemplate) GetAllJobTemplates() (schemaResponse JobTemplateResponseSchema, err error) {
	schemaResponse = JobTemplateResponseSchema{}

	response, err := jobTemplate.connection.Get(jobTemplate.URI, nil)

	if err != nil {
		return schemaResponse, err
	}

	err = jobTemplate.DataConversion.ResponseBodyToStruct(&schemaResponse, *response)

	if err != nil {
		return schemaResponse, err
	}

	return schemaResponse, nil
}

// GetJobTemplate gets a job template by name
//
//	:param name: The name of the job template to get
func (jobTemplate *JobTemplate) GetJobTemplate(name string) (schemaResponse JobTemplateResponseSchema, err error) {
	schemaResponse = JobTemplateResponseSchema{}

	params := map[string]string{
		"name": name,
	}

	response, err := jobTemplate.connection.Get(jobTemplate.URI, params)

	if err != nil {
		return schemaResponse, err
	}

	err = jobTemplate.DataConversion.ResponseBodyToStruct(&schemaResponse, *response)

	if err != nil {
		return schemaResponse, err
	}

	return schemaResponse, nil
}

// GetJobTemplateID gets a job template ID by name
//
//	:param name: The name of the job template to get the ID for
func (jobTemplate *JobTemplate) GetJobTemplateID(name string) (id int32, err error) {
	schemaResponse, err := jobTemplate.GetJobTemplate(name)

	if err != nil {
		return 0, err
	}

	if len(schemaResponse.Results) == 0 {
		return 0, fmt.Errorf("no job template found with name %s", name)
	} else if len(schemaResponse.Results) > 1 {
		return 0, fmt.Errorf("more than one job template found with name %s", name)
	}

	return schemaResponse.Results[0].ID, nil
}

// LaunchJobTemplate launches a job template by ID
//
//	:param id: The ID of the job template to launch
//	:param launchData: The struct to use for the launch data
func (jobTemplate *JobTemplate) LaunchJobTemplate(id int32, launchData any) (schemaResponse JobTemplateResponseSingleSchema, err error) {
	schemaResponse = JobTemplateResponseSingleSchema{}

	uri := fmt.Sprintf("%s%d/launch/", jobTemplate.URI, id)

	jsonData, err := json.Marshal(launchData)

	if err != nil {
		return schemaResponse, err
	}

	response, err := jobTemplate.connection.Post(uri, jsonData)

	if err != nil {
		return schemaResponse, err
	}

	err = jobTemplate.DataConversion.ResponseBodyToStruct(&schemaResponse, *response)

	if err != nil {
		return schemaResponse, err
	}

	return schemaResponse, nil
}
