// Code generated by go generate; DO NOT EDIT.
package devices

var DeviceTypesMapribbon_communications = map[string]*DeviceData{
    "SBC 1000": {
        Manufacturer: "Ribbon Communications",
        Model: "SBC 1000",
        Slug: "ribbon-communications-sbc-1000",
        UHeight: 1,
        PartNumber: "SBC 1000",
        IsFullDepth: false,
        Airflow: "front-to-rear",
        FrontImage: false,
        RearImage: false,
        SubdeviceRole: "",
        Weight: 5.67,
        WeightUnit: "",
        IsPowered: false,
        ConsolePorts: []ConsolePort{
        },
        ConsoleServerPorts: []ConsoleServerPort{
        },
        PowerPorts: []PowerPort{
            { Name: "PSU0", Label: "", Type: "iec-60320-c14", MaximumDraw: 144, AllocatedDraw: 0 },
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
            { Name: "Port 1", Label: "", Type: "1000base-t", MgmtOnly: false },
            { Name: "Port 2", Label: "", Type: "1000base-t", MgmtOnly: false },
            { Name: "Port 3", Label: "", Type: "1000base-t", MgmtOnly: false },
        },
    },
    "SBC 2000": {
        Manufacturer: "Ribbon Communications",
        Model: "SBC 2000",
        Slug: "ribbon-communications-sbc-2000",
        UHeight: 1,
        PartNumber: "SBC 2000",
        IsFullDepth: true,
        Airflow: "front-to-rear",
        FrontImage: false,
        RearImage: false,
        SubdeviceRole: "",
        Weight: 10.43,
        WeightUnit: "",
        IsPowered: false,
        ConsolePorts: []ConsolePort{
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
            { Name: "PSU0", Label: "", Position: "0" },
            { Name: "PSU1", Label: "", Position: "1" },
            { Name: "Module0", Label: "", Position: "0" },
            { Name: "Module1", Label: "", Position: "1" },
        },
			  DeviceBays: []DeviceBay{
        },
        InventoryItems: []InventoryItem{
        },
        Interfaces: []Interface{
            { Name: "Admin", Label: "", Type: "1000base-t", MgmtOnly: true },
            { Name: "Port 1", Label: "", Type: "1000base-t", MgmtOnly: false },
            { Name: "Port 2", Label: "", Type: "1000base-t", MgmtOnly: false },
            { Name: "Port 3", Label: "", Type: "1000base-t", MgmtOnly: false },
            { Name: "Port 4", Label: "", Type: "1000base-t", MgmtOnly: false },
        },
    },
}
