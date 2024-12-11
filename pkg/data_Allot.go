// Code generated by go generate; DO NOT EDIT.
package devices

var DeviceTypesMapAllot = map[string]*DeviceData{
    "Secure Service Gateway SSG-500": {
        Manufacturer: "Allot",
        Model: "Secure Service Gateway SSG-500",
        Slug: "allot-ssg-500",
        UHeight: 2,
        PartNumber: "SSG-500",
        IsFullDepth: true,
        Airflow: "front-to-rear",
        FrontImage: false,
        RearImage: false,
        SubdeviceRole: "",
        Weight: 44,
        WeightUnit: "",
        IsPowered: false,
        ConsolePorts: []ConsolePort{
            { Name: "Console", Type: "rj-45", Label: "", Poe: false },
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
            { Name: "PSU1", Label: "", Position: "1" },
            { Name: "PSU2", Label: "", Position: "2" },
        },
			  DeviceBays: []DeviceBay{
        },
        InventoryItems: []InventoryItem{
        },
        Interfaces: []Interface{
            { Name: "Management 1", Label: "", Type: "1000base-t", MgmtOnly: true },
            { Name: "Management 2", Label: "", Type: "1000base-t", MgmtOnly: true },
            { Name: "Interface 1", Label: "", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "Interface 2", Label: "", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "Interface 3", Label: "", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "Interface 4", Label: "", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "Interface 5", Label: "", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "Interface 6", Label: "", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "Interface 7", Label: "", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "Interface 8", Label: "", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "Interface 9", Label: "", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "Interface 10", Label: "", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "Interface 11", Label: "", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "Interface 12", Label: "", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "Interface 13", Label: "", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "Interface 14", Label: "", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "Interface 15", Label: "", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "Interface 16", Label: "", Type: "10gbase-x-sfpp", MgmtOnly: false },
        },
    },
}
