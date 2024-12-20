/*
Package inventories provides a way to manipulate inventories for Ansible AAP
*/
package inventories

import (
	"encoding/json"
	"fmt"
	"github.com/btr1975/go-ansible-aap-api-client/pkg/connection"
	"net/http"
)

type Inventory struct {
	URI        string
	connection connection.BasicConnection
}

func NewInventory(basicConnection connection.BasicConnection) *Inventory {
	return &Inventory{
		URI:        "inventories/",
		connection: basicConnection,
	}
}

func (inventory *Inventory) GetAllInventories() (schemaResponse InventoryResponseSchema, err error) {
	response, err := inventory.connection.Get(inventory.URI, nil)

	if err != nil {
		return InventoryResponseSchema{}, err
	}

	schemaResponse = InventoryResponseSchema{}

	err = json.NewDecoder(response.Body).Decode(&schemaResponse)

	if err != nil {
		return InventoryResponseSchema{}, err
	}

	return schemaResponse, nil
}

func (inventory *Inventory) GetInventory(name string) (schemaResponse InventoryResponseSchema, err error) {
	params := map[string]string{
		"name": name,
	}

	response, err := inventory.connection.Get(inventory.URI, params)

	if err != nil {
		return InventoryResponseSchema{}, err
	}

	schemaResponse = InventoryResponseSchema{}

	err = json.NewDecoder(response.Body).Decode(&schemaResponse)

	if err != nil {
		return InventoryResponseSchema{}, err
	}

	return schemaResponse, nil

}

func (inventory *Inventory) GetInventoryID(name string) (id int, err error) {
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

func (inventory *Inventory) CreateInventory(inventoryRequest InventoryRequestSchema) (response *http.Response, err error) {
	data, err := json.Marshal(inventoryRequest)

	if err != nil {
		return nil, err
	}

	return inventory.connection.Post(inventory.URI, data)
}

func (inventory *Inventory) DeleteInventory(id int) (statusCode int, err error) {
	uri := fmt.Sprintf("%s%d/", inventory.URI, id)

	response, err := inventory.connection.Delete(uri, nil)

	if err != nil {
		return 0, err
	}

	return response.StatusCode, nil
}
