// Code generated by go generate; DO NOT EDIT.
package devices

var DeviceTypesMapicewhale = map[string]*DeviceData{
    "ZimaBoard": {
        Manufacturer: "IceWhale",
        Model: "ZimaBoard",
        Slug: "icewhale-zimaboard",
        UHeight: 0,
        PartNumber: "",
        IsFullDepth: false,
        Airflow: "passive",
        FrontImage: false,
        RearImage: false,
        SubdeviceRole: "",
        Weight: 278,
        WeightUnit: "",
        IsPowered: false,
        ConsolePorts: []ConsolePort{
        },
        ConsoleServerPorts: []ConsoleServerPort{
        },
        PowerPorts: []PowerPort{
            { Name: "DC-in", Label: "", Type: "dc-terminal", MaximumDraw: 0, AllocatedDraw: 0 },
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
            { Name: "eth1", Label: "", Type: "1000base-t", MgmtOnly: false },
        },
    },
}
