package organizations

import (
	"encoding/json"
	"fmt"
	"github.com/btr1975/go-ansible-aap-api-client/pkg/connection"
	"net/http"
)

// Organization represents an AAP organization
type Organization struct {
	URI        string
	connection connection.BasicConnection
}

// NewOrganization creates a new organization instance
//
//	:param basicConnection: The basic connection to use
func NewOrganization(basicConnection connection.BasicConnection) *Organization {
	return &Organization{
		URI:        "organizations/",
		connection: basicConnection,
	}
}

// GetAllOrganizations gets all organizations
func (organization *Organization) GetAllOrganizations() (schemaResponse OrganizationResponseSchema, err error) {
	response, err := organization.connection.Get(organization.URI, nil)

	if err != nil {
		return OrganizationResponseSchema{}, err
	}

	schemaResponse = OrganizationResponseSchema{}

	err = json.NewDecoder(response.Body).Decode(&schemaResponse)

	if err != nil {
		return OrganizationResponseSchema{}, err
	}

	return schemaResponse, nil
}

// GetOrganization gets an organization by name
//
//	:param name: The name of the organization to get
func (organization *Organization) GetOrganization(name string) (schemaResponse OrganizationResponseSchema, err error) {
	params := map[string]string{
		"name": name,
	}

	response, err := organization.connection.Get(organization.URI, params)

	if err != nil {
		return OrganizationResponseSchema{}, err
	}

	schemaResponse = OrganizationResponseSchema{}

	err = json.NewDecoder(response.Body).Decode(&schemaResponse)

	if err != nil {
		return OrganizationResponseSchema{}, err
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
func (organization *Organization) UpdateOrganization(id int32, orgRequest OrganizationRequestSchema) (response *http.Response, err error) {
	uri := fmt.Sprintf("%s%d/", organization.URI, id)

	data, err := json.Marshal(orgRequest)

	if err != nil {
		return nil, err
	}

	return organization.connection.Patch(uri, data)
}

// CreateOrganization creates an organization
//
//	:param orgRequest: The organization request schema to use
func (organization *Organization) CreateOrganization(orgRequest OrganizationRequestSchema) (response *http.Response, err error) {
	data, err := json.Marshal(orgRequest)

	if err != nil {
		return nil, err
	}

	return organization.connection.Post(organization.URI, data)
}
