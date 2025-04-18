// Code generated by go generate; DO NOT EDIT.
package devices

var DeviceTypesMapqsan = map[string]*DeviceData{
    "xn5008r": {
        Manufacturer: "QSAN",
        Model: "xn5008r",
        Slug: "qsan-xn5008r",
        UHeight: 2,
        PartNumber: "",
        IsFullDepth: true,
        Airflow: "front-to-rear",
        FrontImage: false,
        RearImage: false,
        SubdeviceRole: "",
        Weight: 13.5,
        WeightUnit: "",
        IsPowered: false,
        ConsolePorts: []ConsolePort{
            { Name: "Serial", Type: "rj-45", Label: "", Poe: false },
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
            { Name: "PSU1", Label: "", Position: "PSU1" },
            { Name: "PSU2", Label: "", Position: "PSU2" },
        },
			  DeviceBays: []DeviceBay{
        },
        InventoryItems: []InventoryItem{
        },
        Interfaces: []Interface{
            { Name: "Eth-1", Label: "", Type: "1000base-t", MgmtOnly: false },
            { Name: "Eth-2", Label: "", Type: "1000base-t", MgmtOnly: false },
            { Name: "Eth-3", Label: "", Type: "1000base-t", MgmtOnly: false },
            { Name: "Eth-4", Label: "", Type: "1000base-t", MgmtOnly: false },
        },
    },
}
