/*
Package hosts provides a way to manipulate hosts for Ansible AAP
*/

package hosts

import (
	"encoding/json"
	"github.com/btr1975/go-ansible-aap-api-client/pkg/connection"
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
