/*
Package hosts provides a way to manipulate hosts for Ansible AAP
*/

package hosts

import (
	"encoding/json"
	"fmt"
	"github.com/btr1975/go-ansible-aap-api-client/pkg/connection"
	"net/http"
)

// Host represents an AAP host
type Host struct {
	URI        string
	connection connection.BasicConnection
}

// NewHost creates a new host instance
//
//	:param basicConnection: The basic connection to use
func NewHost(basicConnection connection.BasicConnection) *Host {
	return &Host{
		URI:        "hosts/",
		connection: basicConnection,
	}
}

// GetAllHosts gets all hosts
func (host *Host) GetAllHosts() (schemaResponse HostResponseSchema, err error) {
	response, err := host.connection.Get(host.URI, nil)

	if err != nil {
		return HostResponseSchema{}, err
	}

	schemaResponse = HostResponseSchema{}

	err = json.NewDecoder(response.Body).Decode(&schemaResponse)

	if err != nil {
		return HostResponseSchema{}, err
	}

	return schemaResponse, nil
}

// GetHost gets a host by name
//
//	:param name: The name of the host to get
func (host *Host) GetHost(name string) (schemaResponse HostResponseSchema, err error) {
	params := map[string]string{
		"name": name,
	}

	response, err := host.connection.Get(host.URI, params)

	if err != nil {
		return HostResponseSchema{}, err
	}

	schemaResponse = HostResponseSchema{}

	err = json.NewDecoder(response.Body).Decode(&schemaResponse)

	if err != nil {
		return HostResponseSchema{}, err
	}

	return schemaResponse, nil
}

// GetHostID gets a host ID by name
//
//	:param name: The name of the host to get
func (host *Host) GetHostID(name string) (id int32, err error) {
	params := map[string]string{
		"name": name,
	}

	response, err := host.connection.Get(host.URI, params)

	if err != nil {
		return 0, err
	}

	schemaResponse := HostResponseSchema{}

	err = json.NewDecoder(response.Body).Decode(&schemaResponse)

	if err != nil {
		return 0, err
	}

	if len(schemaResponse.Results) == 0 {
		return 0, nil
	}

	return schemaResponse.Results[0].ID, nil

}

// DeleteHost deletes a host by ID
//
//	:param id: The ID of the host to delete
func (host *Host) DeleteHost(id int32) (statusCode int, err error) {
	uri := fmt.Sprintf("%s%d/", host.URI, id)

	response, err := host.connection.Delete(uri, nil)

	if err != nil {
		return 0, err
	}

	return response.StatusCode, nil

}

// UpdateHost updates a host by ID
//
//	:param id: The ID of the host to update
//	:param hostRequest: The host request to use
func (host *Host) UpdateHost(id int32, hostRequest HostRequestSchema) (response *http.Response, err error) {
	uri := fmt.Sprintf("%s%d/", host.URI, id)

	data, err := json.Marshal(hostRequest)

	if err != nil {
		return nil, err
	}

	return host.connection.Patch(uri, data)
}
