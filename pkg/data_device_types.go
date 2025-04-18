// Code generated by go generate; DO NOT EDIT.
package devices

var DeviceTypesMapdevice_types = map[string]*DeviceData{
    "Meraki MX75": {
        Manufacturer: "Cisco",
        Model: "Meraki MX75",
        Slug: "cisco-meraki-mx75",
        UHeight: 1,
        PartNumber: "MX75",
        IsFullDepth: false,
        Airflow: "passive",
        FrontImage: false,
        RearImage: false,
        SubdeviceRole: "",
        Weight: 1.87,
        WeightUnit: "",
        IsPowered: false,
        ConsolePorts: []ConsolePort{
        },
        ConsoleServerPorts: []ConsoleServerPort{
        },
        PowerPorts: []PowerPort{
            { Name: "Slot 1", Label: "", Type: "dc-terminal", MaximumDraw: 79, AllocatedDraw: 11 },
        },
        PowerOutlets: []PowerOutlet{
        },
        FrontPorts: []FrontPort{
        },
        RearPorts: []RearPort{
        },
        ModuleBays: []ModuleBay{
        },
			  DeviceBays: []DeviceBay{
        },
        InventoryItems: []InventoryItem{
        },
        Interfaces: []Interface{
            { Name: "USB", Label: "", Type: "lte", MgmtOnly: false },
            { Name: "Internet 1", Label: "", Type: "1000base-x-sfp", MgmtOnly: false },
            { Name: "Internet 2", Label: "", Type: "1000base-t", MgmtOnly: false },
            { Name: "Internet 3", Label: "", Type: "1000base-t", MgmtOnly: false },
            { Name: "Port 4", Label: "", Type: "1000base-t", MgmtOnly: false },
            { Name: "Port 5", Label: "", Type: "1000base-t", MgmtOnly: false },
            { Name: "Port 6", Label: "", Type: "1000base-t", MgmtOnly: false },
            { Name: "Port 7", Label: "", Type: "1000base-t", MgmtOnly: false },
            { Name: "Port 8", Label: "", Type: "1000base-t", MgmtOnly: false },
            { Name: "Port 9", Label: "", Type: "1000base-t", MgmtOnly: false },
            { Name: "Port 10", Label: "", Type: "1000base-t", MgmtOnly: false },
            { Name: "Port 11", Label: "", Type: "1000base-t", MgmtOnly: false },
            { Name: "Port 12", Label: "", Type: "1000base-t", MgmtOnly: false },
            { Name: "Port 13", Label: "", Type: "1000base-t", MgmtOnly: false },
        },
    },
}
