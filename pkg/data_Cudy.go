// Code generated by go generate; DO NOT EDIT.
package devices

var DeviceTypesMapcudy = map[string]*DeviceData{
    "WR3000 v1": {
        Manufacturer: "Cudy",
        Model: "WR3000 v1",
        Slug: "cudy-wr3000-v1",
        UHeight: 0,
        PartNumber: "",
        IsFullDepth: false,
        Airflow: "passive",
        FrontImage: false,
        RearImage: false,
        SubdeviceRole: "",
        Weight: 295,
        WeightUnit: "",
        IsPowered: false,
        ConsolePorts: []ConsolePort{
        },
        ConsoleServerPorts: []ConsoleServerPort{
        },
        PowerPorts: []PowerPort{
            { Name: "PS1", Label: "", Type: "dc-terminal", MaximumDraw: 7, AllocatedDraw: 0 },
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
            { Name: "lan1", Label: "", Type: "1000base-t", MgmtOnly: false },
            { Name: "lan2", Label: "", Type: "1000base-t", MgmtOnly: false },
            { Name: "lan3", Label: "", Type: "1000base-t", MgmtOnly: false },
            { Name: "wlan0", Label: "", Type: "ieee802.11ax", MgmtOnly: false },
            { Name: "wlan1", Label: "", Type: "ieee802.11ac", MgmtOnly: false },
        },
    },
}
