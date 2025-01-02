/*
Package inventories provides a way to manipulate inventories for Ansible AAP
*/
package inventories

import (
	"encoding/json"
	"fmt"
	"github.com/btr1975/go-ansible-aap-api-client/pkg/connection"
	"github.com/btr1975/go-ansible-aap-api-client/pkg/dataconversion"
	"github.com/btr1975/go-ansible-aap-api-client/pkg/groups"
	"github.com/btr1975/go-ansible-aap-api-client/pkg/hosts"
)

// Inventory represents an AAP inventory
type Inventory struct {
	URI            string
	connection     connection.BasicConnection
	DataConversion dataconversion.DataConverterInterface
}

// NewInventory creates a new inventory instance
//
// :param basicConnection: The basic connection to use
func NewInventory(basicConnection connection.BasicConnection) *Inventory {
	return &Inventory{
		URI:            "inventories/",
		connection:     basicConnection,
		DataConversion: dataconversion.NewDataConverter(),
	}
}

// GetAllInventories gets all inventories
func (inventory *Inventory) GetAllInventories() (schemaResponse InventoryResponseSchema, err error) {
	schemaResponse = InventoryResponseSchema{}

	response, err := inventory.connection.Get(inventory.URI, nil)

	if err != nil {
		return schemaResponse, err
	}

	err = inventory.DataConversion.ResponseBodyToStruct(&schemaResponse, *response)

	if err != nil {
		return schemaResponse, err
	}

	return schemaResponse, nil
}

// GetInventory gets an inventory by name
//
//	:param name: The name of the inventory to get
func (inventory *Inventory) GetInventory(name string) (schemaResponse InventoryResponseSchema, err error) {
	schemaResponse = InventoryResponseSchema{}

	params := map[string]string{
		"name": name,
	}

	response, err := inventory.connection.Get(inventory.URI, params)

	if err != nil {
		return schemaResponse, err
	}

	err = inventory.DataConversion.ResponseBodyToStruct(&schemaResponse, *response)

	if err != nil {
		return schemaResponse, err
	}

	return schemaResponse, nil

}

// GetInventoryID gets an inventory ID by name
//
//	:param name: The name of the inventory to get
func (inventory *Inventory) GetInventoryID(name string) (id int32, err error) {
	schemaResponse, err := inventory.GetInventory(name)

	if err != nil {
		return 0, err
	}

	if len(schemaResponse.Results) > 1 {
		return 0, fmt.Errorf("more than one inventory found with name %s", name)
	}

	if len(schemaResponse.Results) == 0 {
		return 0, fmt.Errorf("no inventory found with name %s", name)
	}

	return schemaResponse.Results[0].ID, nil

}

// DeleteInventory deletes an inventory by ID
//
//	:param id: The ID of the inventory to delete
func (inventory *Inventory) DeleteInventory(id int32) (statusCode int, err error) {
	uri := fmt.Sprintf("%s%d/", inventory.URI, id)

	response, err := inventory.connection.Delete(uri, nil)

	if err != nil {
		return 0, err
	}

	return response.StatusCode, nil
}

// UpdateInventory updates an inventory by ID
//
//	:param id: The ID of the inventory to update
//	:param inventoryRequest: The inventory request schema to use
func (inventory *Inventory) UpdateInventory(id int32, inventoryRequest InventoryRequestSchema) (schemaResponse InventoryResponseSingleSchema, err error) {
	schemaResponse = InventoryResponseSingleSchema{}

	uri := fmt.Sprintf("%s%d/", inventory.URI, id)

	data, err := json.Marshal(inventoryRequest)

	if err != nil {
		return schemaResponse, err
	}

	response, err := inventory.connection.Patch(uri, data)

	if err != nil {
		return schemaResponse, err
	}

	err = inventory.DataConversion.ResponseBodyToStruct(&schemaResponse, *response)

	if err != nil {
		return schemaResponse, err
	}

	return schemaResponse, nil
}

// CreateInventory creates a new inventory
//
//	:param inventoryRequest: The inventory request schema to use
func (inventory *Inventory) CreateInventory(inventoryRequest InventoryRequestSchema) (schemaResponse InventoryResponseSingleSchema, err error) {
	schemaResponse = InventoryResponseSingleSchema{}

	data, err := json.Marshal(inventoryRequest)

	if err != nil {
		return schemaResponse, err
	}

	response, err := inventory.connection.Post(inventory.URI, data)

	if err != nil {
		return schemaResponse, err
	}

	err = inventory.DataConversion.ResponseBodyToStruct(&schemaResponse, *response)

	if err != nil {
		return schemaResponse, err
	}

	return schemaResponse, nil
}

// AddHostToInventory adds a host to an inventory
//
//	:param id: The ID of the inventory to add the host to
//	:param hostRequest: The host request schema to use
func (inventory *Inventory) AddHostToInventory(id int32, hostRequest hosts.HostRequestSchema) (schemaResponse hosts.HostResponseSingleSchema, err error) {
	schemaResponse = hosts.HostResponseSingleSchema{}

	uri := fmt.Sprintf("%s%d/hosts/", inventory.URI, id)

	data, err := json.Marshal(hostRequest)

	if err != nil {
		return schemaResponse, err
	}

	response, err := inventory.connection.Post(uri, data)

	if err != nil {
		return schemaResponse, err
	}

	err = inventory.DataConversion.ResponseBodyToStruct(&schemaResponse, *response)

	if err != nil {
		return schemaResponse, err
	}

	return schemaResponse, nil

}

// AddGroupToInventory adds a group to an inventory
//
//	:param id: The ID of the inventory to add the group to
//	:param groupRequest: The group request schema to use
func (inventory *Inventory) AddGroupToInventory(id int32, groupRequest groups.GroupRequestSchema) (schemaResponse groups.GroupResponseSingleSchema, err error) {
	schemaResponse = groups.GroupResponseSingleSchema{}

	uri := fmt.Sprintf("%s%d/groups/", inventory.URI, id)

	data, err := json.Marshal(groupRequest)

	if err != nil {
		return schemaResponse, err
	}

	response, err := inventory.connection.Post(uri, data)

	if err != nil {
		return schemaResponse, err
	}

	err = inventory.DataConversion.ResponseBodyToStruct(&schemaResponse, *response)

	if err != nil {
		return schemaResponse, err
	}

	return schemaResponse, nil

}
