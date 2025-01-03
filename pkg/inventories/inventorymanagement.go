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
	IOSHosts            []hosts.HostRequestSchema
	IOSGroupID          int32
	IOSXRGroupVars      groups.GroupGeneralNetwork
	IOSXKHosts          []hosts.HostRequestSchema
	IOSXRGroupID        int32
	NXOSGroupVars       groups.GroupGeneralNetwork
	NXOSHosts           []hosts.HostRequestSchema
	NXOSGroupID         int32
	EOSGroupVars        groups.GroupGeneralNetwork
	EOSHosts            []hosts.HostRequestSchema
	EOSGroupID          int32
	CustomGroups        []groups.GroupRequestSchema
	CustomGroupHosts    []CustomGroupHostSchema
	CustomGroupsIDs     []CustomGroupsIDSchema
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

// Run runs the inventory builder
func (ib *InventoryBuilder) Run() (err error) {

	thisInventory, err := ib.inventoryManagement.Inventory.CreateInventory(ib.Inventory)

	if err != nil {
		return err
	}

	ib.InventoryID = thisInventory.ID

	err = ib.createBasicGroups()

	if err != nil {
		return err
	}

	for _, host := range ib.IOSHosts {
		_, err = ib.inventoryManagement.Group.AddHostToGroup(ib.IOSGroupID, host)

		if err != nil {
			return err
		}
	}

	for _, host := range ib.IOSXKHosts {
		_, err = ib.inventoryManagement.Group.AddHostToGroup(ib.IOSXRGroupID, host)

		if err != nil {
			return err
		}
	}

	for _, host := range ib.NXOSHosts {
		_, err = ib.inventoryManagement.Group.AddHostToGroup(ib.NXOSGroupID, host)

		if err != nil {
			return err
		}
	}

	for _, host := range ib.EOSHosts {
		_, err = ib.inventoryManagement.Group.AddHostToGroup(ib.EOSGroupID, host)

		if err != nil {
			return err
		}
	}

	for _, group := range ib.CustomGroups {
		groupData, err := ib.inventoryManagement.Inventory.AddGroupToInventory(ib.InventoryID, group)

		if err != nil {
			return err
		}

		ib.CustomGroupsIDs = append(ib.CustomGroupsIDs, CustomGroupsIDSchema{GroupName: groupData.Name, GroupID: groupData.ID})

	}

	for _, customGroupHost := range ib.CustomGroupHosts {
		for _, group := range ib.CustomGroupsIDs {
			if customGroupHost.GroupName == group.GroupName {
				_, err = ib.inventoryManagement.Group.AddHostToGroup(group.GroupID, customGroupHost.Host)

				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}

// AddIOSHost adds an IOS host to the inventory builder
//
// :param host: The host to add
func (ib *InventoryBuilder) AddIOSHost(host hosts.HostRequestSchema) {
	ib.IOSHosts = append(ib.IOSHosts, host)
}

// AddIOSHosts adds multiple IOS hosts to the inventory builder
//
// :param hosts: The hosts to add
func (ib *InventoryBuilder) AddIOSHosts(hosts []hosts.HostRequestSchema) {
	ib.IOSHosts = append(ib.IOSHosts, hosts...)
}

// AddIOSXRHost adds an IOS XR host to the inventory builder
//
// :param host: The host to add
func (ib *InventoryBuilder) AddIOSXRHost(host hosts.HostRequestSchema) {
	ib.IOSXKHosts = append(ib.IOSXKHosts, host)
}

// AddIOSXRHosts adds multiple IOS XR hosts to the inventory builder
//
// :param hosts: The hosts to add
func (ib *InventoryBuilder) AddIOSXRHosts(hosts []hosts.HostRequestSchema) {
	ib.IOSXKHosts = append(ib.IOSXKHosts, hosts...)
}

// AddNXOSHost adds an NX-OS host to the inventory builder
//
// :param host: The host to add
func (ib *InventoryBuilder) AddNXOSHost(host hosts.HostRequestSchema) {
	ib.NXOSHosts = append(ib.NXOSHosts, host)
}

// AddNXOSHosts adds multiple NX-OS hosts to the inventory builder
//
// :param hosts: The hosts to add
func (ib *InventoryBuilder) AddNXOSHosts(hosts []hosts.HostRequestSchema) {
	ib.NXOSHosts = append(ib.NXOSHosts, hosts...)
}

// AddEOSHost adds an EOS host to the inventory builder
//
// :param host: The host to add
func (ib *InventoryBuilder) AddEOSHost(host hosts.HostRequestSchema) {
	ib.EOSHosts = append(ib.EOSHosts, host)
}

// AddEOSHosts adds multiple EOS hosts to the inventory builder
//
// :param hosts: The hosts to add
func (ib *InventoryBuilder) AddEOSHosts(hosts []hosts.HostRequestSchema) {
	ib.EOSHosts = append(ib.EOSHosts, hosts...)
}

// customGroupExists checks if a custom group exists
//
// :param groupName: The group name to check
func (ib *InventoryBuilder) customGroupExists(groupName string) bool {
	for _, group := range ib.CustomGroups {
		if group.Name == groupName {
			return true
		}
	}

	return false
}

// AddCustomGroup adds a custom group to the inventory builder
//
// :param group: The group to add
func (ib *InventoryBuilder) AddCustomGroup(group groups.GroupRequestSchema) (err error) {
	if ib.customGroupExists(group.Name) {
		return fmt.Errorf("custom group %s already exists", group.Name)
	}

	ib.CustomGroups = append(ib.CustomGroups, group)

	return nil
}

// AddCustomGroups adds multiple custom groups to the inventory builder
//
// :param groups: The groups to add
func (ib *InventoryBuilder) AddCustomGroups(groups []groups.GroupRequestSchema) (err error) {
	for _, group := range groups {
		if ib.customGroupExists(group.Name) {
			return fmt.Errorf("custom group %s already exists", group.Name)
		}

	}

	ib.CustomGroups = append(ib.CustomGroups, groups...)

	return nil
}

// AddHostToCustomGroup adds a host to a custom group to the inventory builder
//
// :param groupName: The group name to add the host to
// :param host: The host to add
func (ib *InventoryBuilder) AddHostToCustomGroup(groupName string, host hosts.HostRequestSchema) (err error) {
	if !ib.customGroupExists(groupName) {
		return fmt.Errorf("custom group %s does not exists", groupName)
	}

	ib.CustomGroupHosts = append(ib.CustomGroupHosts, CustomGroupHostSchema{
		GroupName: groupName,
		Host:      host,
	})

	return nil
}

// AddHostsToCustomGroup adds a multiple hosts to a custom group to the inventory builder
//
// :param groupName: The group name to add the host to
// :param hosts: The hosts to add
func (ib *InventoryBuilder) AddHostsToCustomGroup(groupName string, hosts []hosts.HostRequestSchema) (err error) {
	if !ib.customGroupExists(groupName) {
		return fmt.Errorf("custom group %s does not exists", groupName)
	}

	for _, host := range hosts {
		ib.CustomGroupHosts = append(ib.CustomGroupHosts, CustomGroupHostSchema{
			GroupName: groupName,
			Host:      host,
		})

	}

	return nil
}

// createBasicGroups creates the basic groups for the inventory
func (ib *InventoryBuilder) createBasicGroups() (err error) {
	var basicNOSGroups = []string{"ios", "iosxr", "nxos", "eos"}

	for _, nos := range basicNOSGroups {
		groupRequest, err := ib.getBasicGroupsRequestSchema(nos)

		if err != nil {
			return err
		}

		groupResponse, err := ib.inventoryManagement.Inventory.AddGroupToInventory(ib.InventoryID, groupRequest)

		if err != nil {
			return err
		}

		switch nos {
		case "ios":
			ib.IOSGroupID = groupResponse.ID
		case "iosxr":
			ib.IOSXRGroupID = groupResponse.ID
		case "nxos":
			ib.NXOSGroupID = groupResponse.ID
		case "eos":
			ib.EOSGroupID = groupResponse.ID
		default:
			return fmt.Errorf("unsupported NOS: %s", nos)
		}

	}

	return nil
}

// getBasicGroupsRequestSchema gets the basic group request schema for a given NOS
//
// :param nos: The NOS to get the group request schema for
func (ib *InventoryBuilder) getBasicGroupsRequestSchema(nos string) (groupSchema groups.GroupRequestSchema, err error) {
	dc := dataconversion.NewDataConverter()

	switch nos {
	case "ios":
		groupVars, _ := dc.StructToYAMLString(ib.IOSGroupVars)

		return groups.GroupRequestSchema{
			Name:        fmt.Sprintf("%s-%s", ib.InventoryName, nos),
			Description: fmt.Sprintf("Inventory %s Group for %s", ib.InventoryName, nos),
			Variables:   groupVars,
		}, nil

	case "iosxr":
		groupVars, _ := dc.StructToYAMLString(ib.IOSXRGroupVars)

		return groups.GroupRequestSchema{
			Name:        fmt.Sprintf("%s-%s", ib.InventoryName, nos),
			Description: fmt.Sprintf("Inventory %s Group for %s", ib.InventoryName, nos),
			Variables:   groupVars,
		}, nil

	case "nxos":
		groupVars, _ := dc.StructToYAMLString(ib.NXOSGroupVars)

		return groups.GroupRequestSchema{
			Name:        fmt.Sprintf("%s-%s", ib.InventoryName, nos),
			Description: fmt.Sprintf("Inventory %s Group for %s", ib.InventoryName, nos),
			Variables:   groupVars,
		}, nil

	case "eos":
		groupVars, _ := dc.StructToYAMLString(ib.EOSGroupVars)

		return groups.GroupRequestSchema{
			Name:        fmt.Sprintf("%s-%s", ib.InventoryName, nos),
			Description: fmt.Sprintf("Inventory %s Group for %s", ib.InventoryName, nos),
			Variables:   groupVars,
		}, nil

	default:
		return groups.GroupRequestSchema{}, fmt.Errorf("unsupported NOS: %s", nos)
	}

}
