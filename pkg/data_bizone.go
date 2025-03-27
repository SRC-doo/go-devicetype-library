// Code generated by go generate; DO NOT EDIT.
package devices

var DeviceTypesMapbizone = map[string]*DeviceData{
    "sdwan-ce300n-4t2s2xs": {
        Manufacturer: "BiZone",
        Model: "sdwan-ce300n-4t2s2xs",
        Slug: "bizone-sdwan-ce300n-4t2s2xs",
        UHeight: 1,
        PartNumber: "sdwan-ce300n-4t2s2xs",
        IsFullDepth: false,
        Airflow: "front-to-rear",
        FrontImage: false,
        RearImage: false,
        SubdeviceRole: "",
        Weight: 7,
        WeightUnit: "",
        IsPowered: false,
        ConsolePorts: []ConsolePort{
            { Name: "CONSOLE", Type: "rj-45", Label: "", Poe: false },
        },
        ConsoleServerPorts: []ConsoleServerPort{
        },
        PowerPorts: []PowerPort{
            { Name: "PSU", Label: "", Type: "iec-60320-c14", MaximumDraw: 300, AllocatedDraw: 200 },
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
            { Name: "port1", Label: "", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "port2", Label: "", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "port3", Label: "", Type: "1000base-x-sfp", MgmtOnly: false },
            { Name: "port4", Label: "", Type: "1000base-x-sfp", MgmtOnly: false },
            { Name: "port5", Label: "", Type: "1000base-t", MgmtOnly: false },
            { Name: "port6", Label: "", Type: "1000base-t", MgmtOnly: false },
            { Name: "port7", Label: "", Type: "1000base-t", MgmtOnly: false },
            { Name: "port8", Label: "", Type: "1000base-t", MgmtOnly: false },
        },
    },
    "sdwan-ce50n-4t1sm": {
        Manufacturer: "BiZone",
        Model: "sdwan-ce50n-4t1sm",
        Slug: "bizone-sdwan-ce50n-4t1sm",
        UHeight: 1,
        PartNumber: "sdwan-ce50n-4t1sm",
        IsFullDepth: false,
        Airflow: "passive",
        FrontImage: false,
        RearImage: false,
        SubdeviceRole: "",
        Weight: 1.4,
        WeightUnit: "",
        IsPowered: false,
        ConsolePorts: []ConsolePort{
            { Name: "CONSOLE", Type: "rj-45", Label: "", Poe: false },
        },
        ConsoleServerPorts: []ConsoleServerPort{
        },
        PowerPorts: []PowerPort{
            { Name: "PSU 12v 3a", Label: "", Type: "dc-terminal", MaximumDraw: 40, AllocatedDraw: 30 },
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
            { Name: "port2", Label: "", Type: "1000base-t", MgmtOnly: false },
            { Name: "port3", Label: "", Type: "1000base-t", MgmtOnly: false },
            { Name: "port4", Label: "", Type: "1000base-t", MgmtOnly: false },
            { Name: "port5", Label: "", Type: "1000base-t", MgmtOnly: false },
            { Name: "port1", Label: "", Type: "1000base-x-sfp", MgmtOnly: false },
        },
    },
}
