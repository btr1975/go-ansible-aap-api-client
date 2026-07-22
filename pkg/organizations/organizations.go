/*
Package organizations provides a way to manipulate organizations for Ansible AAP
*/
package organizations

import (
	"encoding/json"
	"fmt"

	"github.com/btr1975/go-ansible-aap-api-client/pkg/connection"
	"github.com/btr1975/go-ansible-aap-api-client/pkg/dataconversion"
	"github.com/btr1975/go-ansible-aap-api-client/pkg/inventories"
	"github.com/btr1975/go-ansible-aap-api-client/pkg/jobtemplates"
	"github.com/btr1975/go-ansible-aap-api-client/pkg/projects"
)

// Organization represents an AAP organization
type Organization struct {
	URI            string
	connection     connection.BasicConnection
	DataConversion dataconversion.DataConverterInterface
}

// NewOrganization creates a new organization instance
//
//	:param basicConnection: The basic connection to use
func NewOrganization(basicConnection connection.BasicConnection) *Organization {
	return &Organization{
		URI:            "organizations/",
		connection:     basicConnection,
		DataConversion: dataconversion.NewDataConverter(),
	}
}

// GetAllOrganizations gets all organizations
func (organization *Organization) GetAllOrganizations() (schemaResponse OrganizationResponseSchema, err error) {
	schemaResponse = OrganizationResponseSchema{}

	response, err := organization.connection.Get(organization.URI, nil)

	if err != nil {
		return schemaResponse, err
	}

	err = organization.DataConversion.ResponseBodyToStruct(&schemaResponse, *response)

	if err != nil {
		return schemaResponse, err
	}

	return schemaResponse, nil
}

// GetOrganization gets an organization by name
//
//	:param name: The name of the organization to get
func (organization *Organization) GetOrganization(name string) (schemaResponse OrganizationResponseSchema, err error) {
	schemaResponse = OrganizationResponseSchema{}

	params := map[string]string{
		"name": name,
	}

	response, err := organization.connection.Get(organization.URI, params)

	if err != nil {
		return schemaResponse, err
	}

	err = organization.DataConversion.ResponseBodyToStruct(&schemaResponse, *response)

	if err != nil {
		return schemaResponse, err
	}

	return schemaResponse, nil
}

// GetOrganizationID gets an organization ID by name
//
//	:param name: The name of the organization to get
func (organization *Organization) GetOrganizationID(name string) (id int32, err error) {
	schemaResponse, err := organization.GetOrganization(name)

	if err != nil {
		return 0, err
	}

	if len(schemaResponse.Results) > 1 {
		return 0, fmt.Errorf("more than one organization found with name %s", name)
	}

	if len(schemaResponse.Results) == 0 {
		return 0, fmt.Errorf("no organization found with name %s", name)
	}

	return schemaResponse.Results[0].ID, nil
}

// DeleteOrganization deletes an organization by ID
//
//	:param id: The ID of the organization to delete
func (organization *Organization) DeleteOrganization(id int32) (statusCode int, err error) {
	uri := fmt.Sprintf("%s%d/", organization.URI, id)

	response, err := organization.connection.Delete(uri, nil)

	if err != nil {
		return 0, err
	}

	return response.StatusCode, nil
}

// UpdateOrganization updates an organization by ID
//
//	:param id: The ID of the organization to update
//	:param orgRequest: The organization request schema to use
func (organization *Organization) UpdateOrganization(id int32, orgRequest OrganizationRequestSchema) (schemaResponse OrganizationResponseSingleSchema, err error) {
	schemaResponse = OrganizationResponseSingleSchema{}

	uri := fmt.Sprintf("%s%d/", organization.URI, id)

	data, err := json.Marshal(orgRequest)

	if err != nil {
		return schemaResponse, err
	}

	response, err := organization.connection.Patch(uri, data)

	if err != nil {
		return schemaResponse, err
	}

	err = organization.DataConversion.ResponseBodyToStruct(&schemaResponse, *response)

	if err != nil {
		return schemaResponse, err
	}

	return schemaResponse, nil
}

// CreateOrganization creates an organization
//
//	:param orgRequest: The organization request schema to use
func (organization *Organization) CreateOrganization(orgRequest OrganizationRequestSchema) (schemaResponse OrganizationResponseSingleSchema, err error) {
	schemaResponse = OrganizationResponseSingleSchema{}

	data, err := json.Marshal(orgRequest)

	if err != nil {
		return schemaResponse, err
	}

	response, err := organization.connection.Post(organization.URI, data)

	if err != nil {
		return schemaResponse, err
	}

	err = organization.DataConversion.ResponseBodyToStruct(&schemaResponse, *response)

	if err != nil {
		return schemaResponse, err
	}

	return schemaResponse, nil
}

// GetOrganizationProjects gets an organizations projects
//
//	:param name: The name of the organization to get
func (organization *Organization) GetOrganizationProjects(name string) (projectResponse projects.ProjectResponseSchema, err error) {
	var projectResponseSchema projects.ProjectResponseSchema
	schemaResponse, err := organization.GetOrganization(name)

	if err != nil {
		return projectResponseSchema, err
	}

	if len(schemaResponse.Results) > 1 {
		return projectResponseSchema, fmt.Errorf("more than one organization found with name %s", name)
	}

	if len(schemaResponse.Results) == 0 {
		return projectResponseSchema, fmt.Errorf("no organization found with name %s", name)
	}

	response, err := organization.connection.GetFromRelated(schemaResponse.Results[0].Related.Projects)

	if err != nil {
		return projectResponseSchema, err
	}

	err = organization.DataConversion.ResponseBodyToStruct(&projectResponseSchema, *response)

	if err != nil {
		return projectResponseSchema, err
	}

	return projectResponseSchema, nil
}

// GetOrganizationInventories gets an organizations inventories
//
//	:param name: The name of the organization to get
func (organization *Organization) GetOrganizationInventories(name string) (inventoriesResponse inventories.InventoryResponseSchema, err error) {
	var inventoriesResponseSchema inventories.InventoryResponseSchema
	schemaResponse, err := organization.GetOrganization(name)

	if err != nil {
		return inventoriesResponseSchema, err
	}

	if len(schemaResponse.Results) > 1 {
		return inventoriesResponseSchema, fmt.Errorf("more than one organization found with name %s", name)
	}

	if len(schemaResponse.Results) == 0 {
		return inventoriesResponseSchema, fmt.Errorf("no organization found with name %s", name)
	}

	response, err := organization.connection.GetFromRelated(schemaResponse.Results[0].Related.Inventories)

	if err != nil {
		return inventoriesResponseSchema, err
	}

	err = organization.DataConversion.ResponseBodyToStruct(&inventoriesResponseSchema, *response)

	if err != nil {
		return inventoriesResponseSchema, err
	}

	return inventoriesResponseSchema, nil
}

// GetOrganizationJobTemplates gets an organizations job templates
//
//	:param name: The name of the organization to get
func (organization *Organization) GetOrganizationJobTemplates(name string) (jobTemplatesResponse jobtemplates.JobTemplateResponseSchema, err error) {
	var jobTemplatesResponseSchema jobtemplates.JobTemplateResponseSchema
	schemaResponse, err := organization.GetOrganization(name)

	if err != nil {
		return jobTemplatesResponseSchema, err
	}

	if len(schemaResponse.Results) > 1 {
		return jobTemplatesResponseSchema, fmt.Errorf("more than one organization found with name %s", name)
	}

	if len(schemaResponse.Results) == 0 {
		return jobTemplatesResponseSchema, fmt.Errorf("no organization found with name %s", name)
	}

	response, err := organization.connection.GetFromRelated(schemaResponse.Results[0].Related.JobTemplates)

	if err != nil {
		return jobTemplatesResponseSchema, err
	}

	err = organization.DataConversion.ResponseBodyToStruct(&jobTemplatesResponseSchema, *response)

	if err != nil {
		return jobTemplatesResponseSchema, err
	}

	return jobTemplatesResponseSchema, nil
}
