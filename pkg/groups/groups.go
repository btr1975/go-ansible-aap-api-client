/*
Package groups provides a way to manipulate groups for Ansible AAP
*/
package groups

import (
	"encoding/json"
	"fmt"
	"github.com/btr1975/go-ansible-aap-api-client/pkg/connection"
	"github.com/btr1975/go-ansible-aap-api-client/pkg/dataconversion"
	"github.com/btr1975/go-ansible-aap-api-client/pkg/hosts"
	"net/http"
)

// Group represents an AAP group
type Group struct {
	URI            string
	connection     connection.BasicConnection
	DataConversion dataconversion.DataConverterInterface
}

// NewGroup creates a new group instance
//
//	:param basicConnection: The basic connection to use
func NewGroup(basicConnection connection.BasicConnection) *Group {
	return &Group{
		URI:            "groups/",
		connection:     basicConnection,
		DataConversion: dataconversion.NewDataConverter(),
	}
}

// GetAllGroups gets all groups
func (group *Group) GetAllGroups() (schemaResponse GroupResponseSchema, err error) {
	schemaResponse = GroupResponseSchema{}

	response, err := group.connection.Get(group.URI, nil)

	if err != nil {
		return schemaResponse, err
	}

	err = group.DataConversion.ResponseBodyToStruct(&schemaResponse, *response)

	if err != nil {
		return schemaResponse, err
	}

	return schemaResponse, nil
}

// GetGroup gets a group by name
//
//	:param name: The name of the group to get
func (group *Group) GetGroup(name string) (schemaResponse GroupResponseSchema, err error) {
	schemaResponse = GroupResponseSchema{}

	params := map[string]string{
		"name": name,
	}

	response, err := group.connection.Get(group.URI, params)

	if err != nil {
		return schemaResponse, err
	}

	err = group.DataConversion.ResponseBodyToStruct(&schemaResponse, *response)

	if err != nil {
		return schemaResponse, err
	}

	return schemaResponse, nil
}

// GetGroupID gets a group ID by name
//
//	:param name: The name of the group to get
func (group *Group) GetGroupID(name string) (id int32, err error) {
	schemaResponse, err := group.GetGroup(name)

	if err != nil {
		return 0, err
	}

	if len(schemaResponse.Results) > 1 {
		return 0, fmt.Errorf("more than one group found with name %s", name)
	}

	if len(schemaResponse.Results) == 0 {
		return 0, fmt.Errorf("no group found with name %s", name)
	}

	return schemaResponse.Results[0].ID, nil

}

// DeleteGroup deletes a group by ID
//
//	:param id: The ID of the group to delete
func (group *Group) DeleteGroup(id int32) (statusCode int, err error) {
	uri := fmt.Sprintf("%s%d/", group.URI, id)

	response, err := group.connection.Delete(uri, nil)

	if err != nil {
		return 0, err
	}

	return response.StatusCode, nil
}

// UpdateGroup updates a group by ID
//
//	:param id: The ID of the group to update
//	:param groupRequest: The group request to use
func (group *Group) UpdateGroup(id int32, groupRequest GroupRequestSchema) (schemaResponse GroupResponseSingleSchema, err error) {
	schemaResponse = GroupResponseSingleSchema{}

	uri := fmt.Sprintf("%s%d/", group.URI, id)

	data, err := json.Marshal(groupRequest)

	if err != nil {
		return schemaResponse, err
	}

	response, err := group.connection.Patch(uri, data)

	if err != nil {
		return schemaResponse, err
	}

	err = group.DataConversion.ResponseBodyToStruct(&schemaResponse, *response)

	if err != nil {
		return schemaResponse, err
	}

	return schemaResponse, nil
}

func (group *Group) AddHostToGroup(id int32, schema hosts.HostRequestSchema) (response *http.Response, err error) {
	uri := fmt.Sprintf("%s%d/hosts/", group.URI, id)

	data, err := json.Marshal(schema)

	if err != nil {
		return nil, err
	}

	return group.connection.Post(uri, data)
}
