/*
Package jobs provides a way to manipulate jobs for Ansible AAP
*/
package jobs

import (
	"fmt"
	"github.com/btr1975/go-ansible-aap-api-client/pkg/connection"
	"github.com/btr1975/go-ansible-aap-api-client/pkg/dataconversion"
	"io"
)

// Job represents an AAP job
type Job struct {
	URI            string
	connection     connection.BasicConnection
	DataConversion dataconversion.DataConverterInterface
}

// NewJob creates a new job instance
//
//	:param basicConnection: The basic connection to use
func NewJob(basicConnection connection.BasicConnection) *Job {
	return &Job{
		URI:            "jobs/",
		connection:     basicConnection,
		DataConversion: dataconversion.NewDataConverter(),
	}
}

// GetAllJobs gets all jobs
func (job *Job) GetAllJobs() (schemaResponse JobResponseSchema, err error) {
	schemaResponse = JobResponseSchema{}

	response, err := job.connection.Get(job.URI, nil)

	if err != nil {
		return schemaResponse, err
	}

	err = job.DataConversion.ResponseBodyToStruct(&schemaResponse, *response)

	if err != nil {
		return schemaResponse, err
	}

	return schemaResponse, nil
}

// GetJob gets a job by ID
//
//	:param id: The ID of the job to get
func (job *Job) GetJob(id int32) (schemaResponse JobResponseSingleSchema, err error) {
	schemaResponse = JobResponseSingleSchema{}

	uri := fmt.Sprintf("%s%d/", job.URI, id)
	response, err := job.connection.Get(uri, nil)

	if err != nil {
		return schemaResponse, err
	}

	err = job.DataConversion.ResponseBodyToStruct(&schemaResponse, *response)

	if err != nil {
		return schemaResponse, err
	}

	return schemaResponse, nil
}

// GetJobStdOut gets the standard output of a job by ID
//
//	:param id: The ID of the job to get the standard output for
//	:param outputFormat: The format to get the output in ("txt", "json", "html")
func (job *Job) GetJobStdOut(id int32, outputFormat string) (response string, err error) {
	params := map[string]string{
		"format": outputFormat,
	}

	uri := fmt.Sprintf("%s%d/stdout/", job.URI, id)

	resp, err := job.connection.Get(uri, params)

	if err != nil {
		return "", err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil

}

// GetJobStatus gets the status of a job by ID
//
//	:param id: The ID of the job to get the status for
func (job *Job) GetJobStatus(id int32) (status string, err error) {
	response, err := job.GetJob(id)

	if err != nil {
		return "", err
	}

	if response.Status == "" {
		return "", fmt.Errorf("status not found for job %d", id)
	}

	return response.Status, nil
}
