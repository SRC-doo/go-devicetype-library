// Code generated by go generate; DO NOT EDIT.
package devices

var DeviceTypesMaparctic_wolf = map[string]*DeviceData{
    "AWN1000 10G": {
        Manufacturer: "Arctic Wolf",
        Model: "AWN1000 10G",
        Slug: "arctic-wolf-awn1000-10g",
        UHeight: 1,
        PartNumber: "AWN1000 10G",
        IsFullDepth: true,
        Airflow: "front-to-rear",
        FrontImage: false,
        RearImage: false,
        SubdeviceRole: "",
        Weight: 0,
        WeightUnit: "",
        IsPowered: false,
        ConsolePorts: []ConsolePort{
            { Name: "Console", Type: "rj-45", Label: "", Poe: false },
        },
        ConsoleServerPorts: []ConsoleServerPort{
        },
        PowerPorts: []PowerPort{
        },
        PowerOutlets: []PowerOutlet{
        },
        FrontPorts: []FrontPort{
        },
        RearPorts: []RearPort{
        },
        ModuleBays: []ModuleBay{
            { Name: "PS-A", Label: "", Position: "A" },
            { Name: "PS-B", Label: "", Position: "B" },
        },
			  DeviceBays: []DeviceBay{
        },
        InventoryItems: []InventoryItem{
        },
        Interfaces: []Interface{
            { Name: "LAN0", Label: "", Type: "10gbase-t", MgmtOnly: false },
            { Name: "LAN1", Label: "", Type: "10gbase-t", MgmtOnly: false },
            { Name: "LAN2", Label: "", Type: "1000base-t", MgmtOnly: false },
            { Name: "LAN3", Label: "", Type: "1000base-t", MgmtOnly: false },
            { Name: "LAN4", Label: "", Type: "1000base-t", MgmtOnly: false },
            { Name: "LAN5", Label: "", Type: "1000base-t", MgmtOnly: false },
            { Name: "Management", Label: "", Type: "1000base-t", MgmtOnly: true },
            { Name: "LAN6", Label: "1", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "LAN7", Label: "2", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "LAN8", Label: "3", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "LAN9", Label: "4", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "WAN0", Label: "", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "WAN1", Label: "", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "WAN2", Label: "", Type: "1000base-t", MgmtOnly: false },
            { Name: "WAN3", Label: "", Type: "1000base-t", MgmtOnly: false },
            { Name: "WAN4", Label: "", Type: "1000base-t", MgmtOnly: false },
            { Name: "WAN5", Label: "", Type: "1000base-t", MgmtOnly: false },
        },
    },
}
