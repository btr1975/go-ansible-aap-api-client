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
