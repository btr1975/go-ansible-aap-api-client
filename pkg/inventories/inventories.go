/*
Package inventories provides a way to manipulate inventories for Ansible AAP
*/
package inventories

import (
	"encoding/json"
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

func (inventory *Inventory) GetAllInventories() (response *http.Response, err error) {
	return inventory.connection.Get(inventory.URI, nil)
}

func (inventory *Inventory) GetInventory(name string) (response *http.Response, err error) {
	params := map[string]string{
		"name": name,
	}

	return inventory.connection.Get(inventory.URI, params)

}

func (inventory *Inventory) CreateInventory(inventoryRequest InventoryRequestSchema) (response *http.Response, err error) {
	data, err := json.Marshal(inventoryRequest)

	if err != nil {
		return nil, err
	}

	return inventory.connection.Post(inventory.URI, data)
}
