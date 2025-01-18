// Code generated by go generate; DO NOT EDIT.
package devices

var DeviceTypesMaprad = map[string]*DeviceData{
    "ETX-2i-10G": {
        Manufacturer: "RAD",
        Model: "ETX-2i-10G",
        Slug: "rad-etx-2i-10g",
        UHeight: 1,
        PartNumber: "",
        IsFullDepth: false,
        Airflow: "passive",
        FrontImage: false,
        RearImage: false,
        SubdeviceRole: "",
        Weight: 0,
        WeightUnit: "",
        IsPowered: false,
        ConsolePorts: []ConsolePort{
            { Name: "control", Type: "usb-mini-b", Label: "", Poe: false },
        },
        ConsoleServerPorts: []ConsoleServerPort{
        },
        PowerPorts: []PowerPort{
            { Name: "power", Label: "", Type: "iec-60320-c14", MaximumDraw: 90, AllocatedDraw: 0 },
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
            { Name: "ETH-1/1", Label: "", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "ETH-1/2", Label: "", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "ETH-1/3", Label: "", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "ETH-1/4", Label: "", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "ETH-1/5", Label: "", Type: "1000base-x-sfp", MgmtOnly: false },
            { Name: "ETH-1/6", Label: "", Type: "1000base-x-sfp", MgmtOnly: false },
            { Name: "ETH-1/7", Label: "", Type: "1000base-x-sfp", MgmtOnly: false },
            { Name: "ETH-1/8", Label: "", Type: "1000base-x-sfp", MgmtOnly: false },
            { Name: "ETH-1/9", Label: "", Type: "1000base-t", MgmtOnly: false },
            { Name: "ETH-1/10", Label: "", Type: "1000base-t", MgmtOnly: false },
            { Name: "ETH-1/11", Label: "", Type: "1000base-t", MgmtOnly: false },
            { Name: "ETH-1/12", Label: "", Type: "1000base-t", MgmtOnly: false },
            { Name: "MNG-ETH", Label: "", Type: "100base-tx", MgmtOnly: true },
        },
    },
}
