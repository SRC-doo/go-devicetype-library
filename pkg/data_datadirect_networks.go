// Code generated by go generate; DO NOT EDIT.
package devices

var DeviceTypesMapdatadirect_networks = map[string]*DeviceData{
    "SFA200NVX2 Controller": {
        Manufacturer: "DataDirect Networks",
        Model: "SFA200NVX2 Controller",
        Slug: "datadirect-networks-sfa200nvx2-controller",
        UHeight: 0,
        PartNumber: "CTLR-200NVX2E",
        IsFullDepth: false,
        Airflow: "front-to-rear",
        FrontImage: false,
        RearImage: false,
        SubdeviceRole: "child",
        Weight: 0,
        WeightUnit: "",
        IsPowered: false,
        ConsolePorts: []ConsolePort{
            { Name: "Console", Type: "other", Label: "", Poe: false },
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
        },
			  DeviceBays: []DeviceBay{
        },
        InventoryItems: []InventoryItem{
        },
        Interfaces: []Interface{
            { Name: "BMC", Label: "", Type: "1000base-t", MgmtOnly: true },
            { Name: "Mgmt 0", Label: "", Type: "10gbase-t", MgmtOnly: true },
            { Name: "Mgmt 1", Label: "", Type: "10gbase-t", MgmtOnly: true },
            { Name: "RP0 Port0", Label: "", Type: "infiniband-hdr", MgmtOnly: false },
            { Name: "RP0 Port1", Label: "", Type: "infiniband-hdr", MgmtOnly: false },
        },
    },
    "SFA200NVX2 Enclosure": {
        Manufacturer: "DataDirect Networks",
        Model: "SFA200NVX2 Enclosure",
        Slug: "datadirect-networks-sfa200nvx2-enclosure",
        UHeight: 2,
        PartNumber: "ENC-200NVX2",
        IsFullDepth: true,
        Airflow: "front-to-rear",
        FrontImage: false,
        RearImage: false,
        SubdeviceRole: "parent",
        Weight: 36.36,
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
            { Name: "PSU 1", Label: "PSU Slot 1", Position: "1" },
            { Name: "PSU 2", Label: "PSU Slot 2", Position: "2" },
        },
			  DeviceBays: []DeviceBay{
            { Name: "Controller 0", Label: "Controller Slot 0" },
            { Name: "Controller 1", Label: "Controller Slot 1" },
        },
        InventoryItems: []InventoryItem{
        },
        Interfaces: []Interface{
        },
    },
}
