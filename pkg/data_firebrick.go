// Code generated by go generate; DO NOT EDIT.
package devices

var DeviceTypesMapfirebrick = map[string]*DeviceData{
    "FB2900": {
        Manufacturer: "FireBrick",
        Model: "FB2900",
        Slug: "firebrick-fb2900",
        UHeight: 1,
        PartNumber: "FB2900",
        IsFullDepth: false,
        Airflow: "passive",
        FrontImage: false,
        RearImage: false,
        SubdeviceRole: "",
        Weight: 0,
        WeightUnit: "",
        IsPowered: false,
        ConsolePorts: []ConsolePort{
        },
        ConsoleServerPorts: []ConsoleServerPort{
        },
        PowerPorts: []PowerPort{
            { Name: "PSU1", Label: "", Type: "iec-60320-c8", MaximumDraw: 15, AllocatedDraw: 0 },
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
            { Name: "1", Label: "", Type: "1000base-t", MgmtOnly: false },
            { Name: "2", Label: "", Type: "1000base-t", MgmtOnly: false },
            { Name: "3", Label: "", Type: "1000base-t", MgmtOnly: false },
            { Name: "4", Label: "", Type: "1000base-t", MgmtOnly: false },
            { Name: "5", Label: "", Type: "1000base-x-sfp", MgmtOnly: false },
        },
    },
    "FB6202": {
        Manufacturer: "FireBrick",
        Model: "FB6202",
        Slug: "firebrick-fb6202",
        UHeight: 1,
        PartNumber: "FB6202",
        IsFullDepth: false,
        Airflow: "passive",
        FrontImage: false,
        RearImage: false,
        SubdeviceRole: "",
        Weight: 0,
        WeightUnit: "",
        IsPowered: false,
        ConsolePorts: []ConsolePort{
            { Name: "Console", Type: "de-9", Label: "", Poe: false },
        },
        ConsoleServerPorts: []ConsoleServerPort{
        },
        PowerPorts: []PowerPort{
            { Name: "PSU1", Label: "", Type: "iec-60320-c14", MaximumDraw: 30, AllocatedDraw: 0 },
            { Name: "PSU2", Label: "", Type: "iec-60320-c14", MaximumDraw: 30, AllocatedDraw: 0 },
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
            { Name: "1", Label: "", Type: "1000base-t", MgmtOnly: false },
            { Name: "2", Label: "", Type: "1000base-t", MgmtOnly: false },
        },
    },
}
