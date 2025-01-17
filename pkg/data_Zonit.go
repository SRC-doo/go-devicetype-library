// Code generated by go generate; DO NOT EDIT.
package devices

var DeviceTypesMapzonit = map[string]*DeviceData{
    "uATS1-HV-C14-2": {
        Manufacturer: "Zonit",
        Model: "uATS1-HV-C14-2",
        Slug: "zonit-uats1-hv-c14-2",
        UHeight: 0,
        PartNumber: "UATS1-HV-C14-2",
        IsFullDepth: false,
        Airflow: "passive",
        FrontImage: false,
        RearImage: false,
        SubdeviceRole: "",
        Weight: 0.9,
        WeightUnit: "",
        IsPowered: false,
        ConsolePorts: []ConsolePort{
        },
        ConsoleServerPorts: []ConsoleServerPort{
        },
        PowerPorts: []PowerPort{
            { Name: "Input A", Label: "", Type: "iec-60320-c14", MaximumDraw: 0, AllocatedDraw: 0 },
            { Name: "Input B", Label: "", Type: "iec-60320-c14", MaximumDraw: 0, AllocatedDraw: 0 },
        },
        PowerOutlets: []PowerOutlet{
            { Name: "Output", Type: "iec-60320-c13", Label: "", PowerPort: "", FeedLeg: "", MaximumDraw: 0, AllocatedDraw: 0 },
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
        },
    },
}
