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
	connection          connection.BasicConnection
	inventoryManagement *inventories.InventoryManagement
	jobTemplate         *JobTemplate
	job                 *jobs.Job
	jobID               int32
	jobTemplateName     string
	jobTemplateID       int32
	inventoryName       string
	inventoryID         int32
}

// NewJobManagement creates a new job management instance
//
//	:param basicConnection: The basic connection to use
func NewJobManagement(basicConnection connection.BasicConnection, jobTemplateName string, inventoryName string) *JobManagement {
	return &JobManagement{
		connection:          basicConnection,
		inventoryManagement: inventories.NewInventoryManagement(basicConnection),
		jobTemplate:         NewJobTemplate(basicConnection),
		job:                 jobs.NewJob(basicConnection),
		jobTemplateName:     jobTemplateName,
		inventoryName:       inventoryName,
	}
}

// Run runs a job
func (jobManagement *JobManagement) Run(launchData JobTemplateSimpleRequestSchema) (err error) {
	jobManagement.jobTemplateID, err = jobManagement.jobTemplate.GetJobTemplateID(jobManagement.jobTemplateName)

	fmt.Printf("Template ID %d\n", jobManagement.jobTemplateID)

	if err != nil {
		return err
	}

	jobManagement.inventoryID, err = jobManagement.inventoryManagement.Inventory.GetInventoryID(jobManagement.inventoryName)

	fmt.Printf("Inventory ID %d\n", jobManagement.inventoryID)

	if err != nil {
		return err
	}

	launchData.Inventory = jobManagement.inventoryID

	fmt.Printf("launchData %v\n", launchData)

	jobData, err := jobManagement.jobTemplate.LaunchJobTemplate(jobManagement.jobTemplateID, launchData)

	if err != nil {
		return err
	}

	fmt.Printf("Job ID %v\n", jobData.ID)

	jobManagement.jobID = jobData.ID

	return nil

}

func (jobManagement *JobManagement) PollCompletion(printStatus bool, launchData JobTemplateSimpleRequestSchema) (jobStatus string, err error) {
	jobStatus = "new"

	fmt.Printf("POOP %d", jobManagement.jobID)

	if jobManagement.jobID == 0 {

		err = jobManagement.Run(launchData)

		if err != nil {
			return jobStatus, err
		}
	}

	if printStatus {
		fmt.Printf("Polling Job ID %d current status %s", jobManagement.jobID, jobStatus)
	}

	for jobStatus != "successful" && jobStatus != "failed" && jobStatus != "error" && jobStatus != "cancelled" {
		currentStatus, err := jobManagement.job.GetJobStatus(jobManagement.jobID)

		if err != nil {
			return jobStatus, err
		}

		jobStatus = currentStatus

		if printStatus {
			fmt.Printf("Polling Job ID %d current status %s", jobManagement.jobID, jobStatus)
		}

		time.Sleep(5 * time.Second)
	}

	if printStatus {
		fmt.Printf("Polling Job ID %d completed status %s", jobManagement.jobID, jobStatus)
	}

	return jobStatus, nil
}
