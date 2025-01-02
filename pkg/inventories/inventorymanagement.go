package inventories

import (
	"fmt"
	"github.com/btr1975/go-ansible-aap-api-client/pkg/connection"
	"github.com/btr1975/go-ansible-aap-api-client/pkg/dataconversion"
	"github.com/btr1975/go-ansible-aap-api-client/pkg/groups"
	"github.com/btr1975/go-ansible-aap-api-client/pkg/hosts"
)

// InventoryManagement represents an AAP inventory management object
type InventoryManagement struct {
	Inventory *Inventory
	Group     *groups.Group
	Host      *hosts.Host
}

// NewInventoryManagement creates a new inventory management instance
//
// :param basicConnection: The basic connection to use
func NewInventoryManagement(basicConnection connection.BasicConnection) *InventoryManagement {
	return &InventoryManagement{
		Inventory: NewInventory(basicConnection),
		Group:     groups.NewGroup(basicConnection),
		Host:      hosts.NewHost(basicConnection),
	}
}

// InventoryBuilder represents an AAP inventory builder object
type InventoryBuilder struct {
	inventoryManagement *InventoryManagement
	Inventory           InventoryRequestSchema
	InventoryName       string
	InventoryID         int32
	IOSGroupVars        groups.GroupGeneralNetwork
	IOSHosts            []InventoryHostRequestSchema
	IOSGroupID          int32
	IOSXRGroupVars      groups.GroupGeneralNetwork
	IOSXKHosts          []InventoryHostRequestSchema
	IOSXRGroupID        int32
	NXOSGroupVars       groups.GroupGeneralNetwork
	NXOSHosts           []InventoryHostRequestSchema
	NXOSGroupID         int32
	EOSGroupVars        groups.GroupGeneralNetwork
	EOSHosts            []InventoryHostRequestSchema
	EOSGroupID          int32
}

// NewInventoryBuilder creates a new inventory builder instance
//
// :param inventoryManagement: The inventory management object to use
func NewInventoryBuilder(inventoryManagement *InventoryManagement, inventory InventoryRequestSchema) *InventoryBuilder {
	iosGroupVars := groups.GroupGeneralNetwork{
		AnsibleConnection:   "ansible.netcommon.network_cli",
		AnsibleBecome:       true,
		AnsibleBecomeMethod: "enable",
		AnsibleNetworkOS:    "cisco.ios.ios",
	}

	iosxrGroupVars := groups.GroupGeneralNetwork{
		AnsibleConnection:   "ansible.netcommon.network_cli",
		AnsibleBecome:       true,
		AnsibleBecomeMethod: "enable",
		AnsibleNetworkOS:    "cisco.iosxr.iosxr",
	}

	nxosGroupVars := groups.GroupGeneralNetwork{
		AnsibleConnection:   "ansible.netcommon.network_cli",
		AnsibleBecome:       true,
		AnsibleBecomeMethod: "enable",
		AnsibleNetworkOS:    "cisco.nxos.nxos",
	}

	eosGroupVars := groups.GroupGeneralNetwork{
		AnsibleConnection:   "ansible.netcommon.network_cli",
		AnsibleBecome:       true,
		AnsibleBecomeMethod: "enable",
		AnsibleNetworkOS:    "arista.eos.eos",
	}

	return &InventoryBuilder{
		inventoryManagement: inventoryManagement,
		Inventory:           inventory,
		InventoryName:       inventory.Name,
		IOSGroupVars:        iosGroupVars,
		IOSXRGroupVars:      iosxrGroupVars,
		NXOSGroupVars:       nxosGroupVars,
		EOSGroupVars:        eosGroupVars,
	}
}

func (ib *InventoryBuilder) Run() (err error) {

	thisInventory, err := ib.inventoryManagement.Inventory.CreateInventory(ib.Inventory)

	if err != nil {
		return err
	}

	ib.InventoryID = thisInventory.ID

	return nil
}

// AddIOSHost adds an IOS host to the inventory builder
//
// :param host: The host to add
func (ib *InventoryBuilder) AddIOSHost(host InventoryHostRequestSchema) {
	ib.IOSHosts = append(ib.IOSHosts, host)
}

// AddIOSHosts adds multiple IOS hosts to the inventory builder
//
// :param hosts: The hosts to add
func (ib *InventoryBuilder) AddIOSHosts(hosts []InventoryHostRequestSchema) {
	ib.IOSHosts = append(ib.IOSHosts, hosts...)
}

// AddIOSXRHost adds an IOS XR host to the inventory builder
//
// :param host: The host to add
func (ib *InventoryBuilder) AddIOSXRHost(host InventoryHostRequestSchema) {
	ib.IOSXKHosts = append(ib.IOSXKHosts, host)
}

// AddIOSXRHosts adds multiple IOS XR hosts to the inventory builder
//
// :param hosts: The hosts to add
func (ib *InventoryBuilder) AddIOSXRHosts(hosts []InventoryHostRequestSchema) {
	ib.IOSXKHosts = append(ib.IOSXKHosts, hosts...)
}

// AddNXOSHost adds an NX-OS host to the inventory builder
//
// :param host: The host to add
func (ib *InventoryBuilder) AddNXOSHost(host InventoryHostRequestSchema) {
	ib.NXOSHosts = append(ib.NXOSHosts, host)
}

// AddNXOSHosts adds multiple NX-OS hosts to the inventory builder
//
// :param hosts: The hosts to add
func (ib *InventoryBuilder) AddNXOSHosts(hosts []InventoryHostRequestSchema) {
	ib.NXOSHosts = append(ib.NXOSHosts, hosts...)
}

// AddEOSHost adds an EOS host to the inventory builder
//
// :param host: The host to add
func (ib *InventoryBuilder) AddEOSHost(host InventoryHostRequestSchema) {
	ib.EOSHosts = append(ib.EOSHosts, host)
}

// AddEOSHosts adds multiple EOS hosts to the inventory builder
//
// :param hosts: The hosts to add
func (ib *InventoryBuilder) AddEOSHosts(hosts []InventoryHostRequestSchema) {
	ib.EOSHosts = append(ib.EOSHosts, hosts...)
}

func (ib *InventoryBuilder) getGroupRequestSchema(nos string) (groupSchema groups.GroupRequestSchema, err error) {
	dc := dataconversion.NewDataConverter()

	switch nos {
	case "ios":
		groupVars, _ := dc.StructToYAMLString(ib.IOSGroupVars)

		return groups.GroupRequestSchema{
			Name:        fmt.Sprintf("%s-%s", ib.InventoryName, nos),
			Description: fmt.Sprintf("Inventory %s Groups fo %s", ib.InventoryName, nos),
			Variables:   groupVars,
		}, nil

	case "iosxr":
		groupVars, _ := dc.StructToYAMLString(ib.IOSXRGroupVars)

		return groups.GroupRequestSchema{
			Name:        fmt.Sprintf("%s-%s", ib.InventoryName, nos),
			Description: fmt.Sprintf("Inventory %s Groups fo %s", ib.InventoryName, nos),
			Variables:   groupVars,
		}, nil

	case "nxos":
		groupVars, _ := dc.StructToYAMLString(ib.NXOSGroupVars)

		return groups.GroupRequestSchema{
			Name:        fmt.Sprintf("%s-%s", ib.InventoryName, nos),
			Description: fmt.Sprintf("Inventory %s Groups fo %s", ib.InventoryName, nos),
			Variables:   groupVars,
		}, nil

	case "eos":
		groupVars, _ := dc.StructToYAMLString(ib.EOSGroupVars)

		return groups.GroupRequestSchema{
			Name:        fmt.Sprintf("%s-%s", ib.InventoryName, nos),
			Description: fmt.Sprintf("Inventory %s Groups fo %s", ib.InventoryName, nos),
			Variables:   groupVars,
		}, nil

	default:
		return groups.GroupRequestSchema{}, fmt.Errorf("unsupported NOS: %s", nos)
	}

}
