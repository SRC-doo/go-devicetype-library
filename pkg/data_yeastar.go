// Code generated by go generate; DO NOT EDIT.
package devices

var DeviceTypesMapyeastar = map[string]*DeviceData{
    "TE200": {
        Manufacturer: "Yeastar",
        Model: "TE200",
        Slug: "yeastar-te200",
        UHeight: 1,
        PartNumber: "Yeastar-TE200",
        IsFullDepth: false,
        Airflow: "passive",
        FrontImage: false,
        RearImage: false,
        SubdeviceRole: "",
        Weight: 1.8,
        WeightUnit: "",
        IsPowered: false,
        ConsolePorts: []ConsolePort{
            { Name: "console", Type: "rj-45", Label: "", Poe: false },
        },
        ConsoleServerPorts: []ConsoleServerPort{
        },
        PowerPorts: []PowerPort{
            { Name: "PSU", Label: "", Type: "iec-60320-c14", MaximumDraw: 6, AllocatedDraw: 0 },
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
            { Name: "wan", Label: "", Type: "1000base-t", MgmtOnly: false },
            { Name: "lan", Label: "", Type: "1000base-t", MgmtOnly: false },
            { Name: "e1/1", Label: "", Type: "e1", MgmtOnly: false },
            { Name: "e1/2", Label: "", Type: "e1", MgmtOnly: false },
        },
    },
}
