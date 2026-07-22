/*
Package projects provides a way to manipulate projects for Ansible AAP
*/
package projects

import (
	"fmt"

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

// GetProject gets a project by name
//
//	:param name: The name of the project to get
func (project *Project) GetProject(name string) (schemaResponse ProjectResponseSchema, err error) {
	schemaResponse = ProjectResponseSchema{}

	params := map[string]string{
		"name": name,
	}

	response, err := project.connection.Get(project.URI, params)

	if err != nil {
		return schemaResponse, err
	}

	err = project.DataConversion.ResponseBodyToStruct(&schemaResponse, *response)

	if err != nil {
		return schemaResponse, err
	}

	return schemaResponse, nil
}

// GetProjectID gets a job template ID by name
//
//	:param name: The name of the project to get the ID for
func (project *Project) GetProjectID(name string) (id int32, err error) {
	schemaResponse, err := project.GetProject(name)

	if err != nil {
		return 0, err
	}

	if len(schemaResponse.Results) == 0 {
		return 0, fmt.Errorf("no project found with name %s", name)
	} else if len(schemaResponse.Results) > 1 {
		return 0, fmt.Errorf("more than one project found with name %s", name)
	}

	return schemaResponse.Results[0].ID, nil
}

// SyncProject syncs project by name
//
//	:param name: The name of the project to get
func (project *Project) SyncProject(name string) (schemaResponse ProjectResponseSchema, err error) {
	schemaResponse, err = project.GetProject(name)

	if err != nil {
		return schemaResponse, err
	}

	update := schemaResponse.Results[0].Related.Update

	var b []byte
	_, err = project.connection.PostFromRelated(update, b)

	if err != nil {
		return schemaResponse, err
	}

	return schemaResponse, nil
}
