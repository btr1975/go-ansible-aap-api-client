package jobtemplates

import (
	"fmt"
	"github.com/btr1975/go-ansible-aap-api-client/pkg/connection"
	"github.com/btr1975/go-ansible-aap-api-client/pkg/inventories"
	"github.com/btr1975/go-ansible-aap-api-client/pkg/jobs"
	"time"
)

// JobManagement represents an AAP job management object
type JobManagement struct {
	jobTemplate     *JobTemplate
	job             *jobs.Job
	jobID           int32
	jobTemplateName string
	jobTemplateID   int32
	inventoryName   string
	inventoryID     int32
}

// NewJobManagement creates a new job management instance
//
//	:param basicConnection: The basic connection to use
func NewJobManagement(basicConnection connection.BasicConnection, jobTemplateName string, inventoryName string) (*JobManagement, error) {
	inventory := inventories.NewInventory(basicConnection)
	inventoryID, err := inventory.GetInventoryID(inventoryName)

	if err != nil {
		return nil, err
	}

	jobTemplate := NewJobTemplate(basicConnection)

	jobTemplateID, err := jobTemplate.GetJobTemplateID(jobTemplateName)

	if err != nil {
		return nil, err
	}

	return &JobManagement{
		inventoryID:     inventoryID,
		jobTemplate:     jobTemplate,
		job:             jobs.NewJob(basicConnection),
		jobTemplateName: jobTemplateName,
		jobTemplateID:   jobTemplateID,
		inventoryName:   inventoryName,
	}, nil
}

// Run runs a job
//
//	:param launchData: The launch data
func (jobManagement *JobManagement) Run(launchData JobTemplateSimpleRequestSchema) (err error) {
	launchData.Inventory = jobManagement.inventoryID

	jobData, err := jobManagement.jobTemplate.LaunchJobTemplate(jobManagement.jobTemplateID, launchData)

	if err != nil {
		return err
	}

	jobManagement.jobID = jobData.ID

	return nil

}

// PollCompletion runs a job and polls for completion
//
//	:param printStatus: Whether to print the status
//	:param launchData: The launch data
func (jobManagement *JobManagement) PollCompletion(printStatus bool, launchData JobTemplateSimpleRequestSchema) (jobStatus string, err error) {
	jobStatus = "new"

	if jobManagement.jobID == 0 {

		err = jobManagement.Run(launchData)

		if err != nil {
			return jobStatus, err
		}
	}

	if printStatus {
		fmt.Printf("Polling Job ID %d current status %s\n", jobManagement.jobID, jobStatus)
	}

	for jobStatus != "successful" && jobStatus != "failed" && jobStatus != "error" && jobStatus != "cancelled" {
		currentStatus, err := jobManagement.job.GetJobStatus(jobManagement.jobID)

		if err != nil {
			return jobStatus, err
		}

		jobStatus = currentStatus

		if printStatus {
			fmt.Printf("Polling Job ID %d current status %s\n", jobManagement.jobID, jobStatus)
		}

		time.Sleep(5 * time.Second)
	}

	if printStatus {
		fmt.Printf("Polling Job ID %d completed status %s\n", jobManagement.jobID, jobStatus)
	}

	return jobStatus, nil
}
