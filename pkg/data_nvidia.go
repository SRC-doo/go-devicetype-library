// Code generated by go generate; DO NOT EDIT.
package devices

var DeviceTypesMapnvidia = map[string]*DeviceData{
    "Jetson Xavier NX Developer Kit": {
        Manufacturer: "Nvidia",
        Model: "Jetson Xavier NX Developer Kit",
        Slug: "nvidia-jetson-xavier-nx-developer-kit",
        UHeight: 0,
        PartNumber: "",
        IsFullDepth: false,
        Airflow: "",
        FrontImage: false,
        RearImage: false,
        SubdeviceRole: "",
        Weight: 0,
        WeightUnit: "",
        IsPowered: false,
        ConsolePorts: []ConsolePort{
            { Name: "HDMI", Type: "other", Label: "", Poe: false },
            { Name: "Display Port", Type: "other", Label: "", Poe: false },
        },
        ConsoleServerPorts: []ConsoleServerPort{
        },
        PowerPorts: []PowerPort{
            { Name: "PSU", Label: "", Type: "dc-terminal", MaximumDraw: 0, AllocatedDraw: 0 },
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
            { Name: "eth0", Label: "", Type: "1000base-t", MgmtOnly: false },
        },
    },
}
