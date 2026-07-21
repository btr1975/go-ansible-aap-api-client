/*
Package projects provides a way to manipulate projects for Ansible AAP
*/
package projects

import (
	"github.com/btr1975/go-ansible-aap-api-client/pkg/connection"
	"github.com/btr1975/go-ansible-aap-api-client/pkg/dataconversion"
)

// Project represents an AAP project
type Project struct {
	URI            string
	connection     connection.BasicConnection
	DataConversion dataconversion.DataConverterInterface
}

// NewProject creates a new project instance
//
//	:param basicConnection: The basic connection to use
func NewProject(basicConnection connection.BasicConnection) *Project {
	return &Project{
		URI:            "projects/",
		connection:     basicConnection,
		DataConversion: dataconversion.NewDataConverter(),
	}
}

// GetAllProjects gets all projects
func (project *Project) GetAllProjects() (schemaResponse ProjectResponseSchema, err error) {
	schemaResponse = ProjectResponseSchema{}

	response, err := project.connection.Get(project.URI, nil)

	if err != nil {
		return schemaResponse, err
	}

	err = project.DataConversion.ResponseBodyToStruct(&schemaResponse, *response)

	if err != nil {
		return schemaResponse, err
	}

	return schemaResponse, nil
}
