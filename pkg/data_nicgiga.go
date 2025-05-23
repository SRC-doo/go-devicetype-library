// Code generated by go generate; DO NOT EDIT.
package devices

var DeviceTypesMapnicgiga = map[string]*DeviceData{
    "S25-0402P": {
        Manufacturer: "NICGIGA",
        Model: "S25-0402P",
        Slug: "nicgiga-s25-0402p",
        UHeight: 1,
        PartNumber: "",
        IsFullDepth: false,
        Airflow: "passive",
        FrontImage: false,
        RearImage: false,
        SubdeviceRole: "",
        Weight: 2,
        WeightUnit: "",
        IsPowered: false,
        ConsolePorts: []ConsolePort{
        },
        ConsoleServerPorts: []ConsoleServerPort{
        },
        PowerPorts: []PowerPort{
            { Name: "AC Input", Label: "", Type: "iec-60320-c14", MaximumDraw: 100, AllocatedDraw: 0 },
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
            { Name: "2.5GE 1", Label: "", Type: "2.5gbase-t", MgmtOnly: false },
            { Name: "2.5GE 2", Label: "", Type: "2.5gbase-t", MgmtOnly: false },
            { Name: "2.5GE 3", Label: "", Type: "2.5gbase-t", MgmtOnly: false },
            { Name: "2.5GE 4", Label: "", Type: "2.5gbase-t", MgmtOnly: false },
            { Name: "SFP&#43; 1", Label: "", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "SFP&#43; 2", Label: "", Type: "10gbase-x-sfpp", MgmtOnly: false },
        },
    },
}
