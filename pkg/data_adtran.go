// Code generated by go generate; DO NOT EDIT.
package devices

var DeviceTypesMapadtran = map[string]*DeviceData{
    "MX2800": {
        Manufacturer: "Adtran",
        Model: "MX2800",
        Slug: "adtran-mx2800",
        UHeight: 1,
        PartNumber: "1200290L1",
        IsFullDepth: false,
        Airflow: "",
        FrontImage: false,
        RearImage: false,
        SubdeviceRole: "",
        Weight: 0,
        WeightUnit: "",
        IsPowered: false,
        ConsolePorts: []ConsolePort{
            { Name: "LAN", Type: "rj-45", Label: "", Poe: false },
            { Name: "MODEM", Type: "rj-45", Label: "", Poe: false },
        },
        ConsoleServerPorts: []ConsoleServerPort{
        },
        PowerPorts: []PowerPort{
            { Name: "PS1", Label: "", Type: "iec-60320-c14", MaximumDraw: 27, AllocatedDraw: 0 },
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
            { Name: "DS3 A IN", Label: "", Type: "t3", MgmtOnly: false },
            { Name: "DS3 A OUT", Label: "", Type: "t3", MgmtOnly: false },
            { Name: "DS3 B IN", Label: "", Type: "t3", MgmtOnly: false },
            { Name: "DS3 B OUT", Label: "", Type: "t3", MgmtOnly: false },
        },
    },
    "NetVanta 4660": {
        Manufacturer: "Adtran",
        Model: "NetVanta 4660",
        Slug: "adtran-netvanta-4660",
        UHeight: 1,
        PartNumber: "17004660F1",
        IsFullDepth: false,
        Airflow: "",
        FrontImage: true,
        RearImage: false,
        SubdeviceRole: "",
        Weight: 3.2,
        WeightUnit: "",
        IsPowered: false,
        ConsolePorts: []ConsolePort{
            { Name: "Console", Type: "de-9", Label: "", Poe: false },
        },
        ConsoleServerPorts: []ConsoleServerPort{
        },
        PowerPorts: []PowerPort{
            { Name: "Power", Label: "", Type: "other", MaximumDraw: 40, AllocatedDraw: 0 },
        },
        PowerOutlets: []PowerOutlet{
        },
        FrontPorts: []FrontPort{
        },
        RearPorts: []RearPort{
        },
        ModuleBays: []ModuleBay{
            { Name: "Slot 1", Label: "", Position: "1" },
        },
			  DeviceBays: []DeviceBay{
        },
        InventoryItems: []InventoryItem{
        },
        Interfaces: []Interface{
            { Name: "GIG 0/1", Label: "", Type: "1000base-x-sfp", MgmtOnly: false },
            { Name: "GIG 0/2", Label: "", Type: "1000base-x-sfp", MgmtOnly: false },
            { Name: "GIG 0/3", Label: "", Type: "1000base-x-sfp", MgmtOnly: false },
            { Name: "GIG 0/4", Label: "", Type: "1000base-x-sfp", MgmtOnly: false },
            { Name: "GIG 0/5", Label: "", Type: "1000base-x-sfp", MgmtOnly: false },
            { Name: "T4", Label: "", Type: "other", MgmtOnly: false },
        },
    },
    "Total Access 5000 23-Inch Chassis": {
        Manufacturer: "Adtran",
        Model: "Total Access 5000 23-Inch Chassis",
        Slug: "adtran-total-access-5000-23-inch-chassis",
        UHeight: 8,
        PartNumber: "1187001G1",
        IsFullDepth: false,
        Airflow: "",
        FrontImage: false,
        RearImage: false,
        SubdeviceRole: "parent",
        Weight: 0,
        WeightUnit: "",
        IsPowered: false,
        ConsolePorts: []ConsolePort{
        },
        ConsoleServerPorts: []ConsoleServerPort{
        },
        PowerPorts: []PowerPort{
            { Name: "DC-Input-A", Label: "DC-A", Type: "dc-terminal", MaximumDraw: 1440, AllocatedDraw: 1000 },
            { Name: "DC-Input-B", Label: "DC-B", Type: "dc-terminal", MaximumDraw: 1440, AllocatedDraw: 1000 },
        },
        PowerOutlets: []PowerOutlet{
        },
        FrontPorts: []FrontPort{
        },
        RearPorts: []RearPort{
        },
        ModuleBays: []ModuleBay{
            { Name: "1", Label: "Slot 1", Position: "1" },
            { Name: "2", Label: "Slot 2", Position: "2" },
            { Name: "3", Label: "Slot 3", Position: "3" },
            { Name: "4", Label: "Slot 4", Position: "4" },
            { Name: "5", Label: "Slot 5", Position: "5" },
            { Name: "6", Label: "Slot 6", Position: "6" },
            { Name: "7", Label: "Slot 7", Position: "7" },
            { Name: "8", Label: "Slot 8", Position: "8" },
            { Name: "9", Label: "Slot 9", Position: "9" },
            { Name: "10", Label: "Slot 10", Position: "10" },
            { Name: "11", Label: "Slot 11", Position: "11" },
            { Name: "13", Label: "Slot 13", Position: "13" },
            { Name: "14", Label: "Slot 14", Position: "14" },
            { Name: "15", Label: "Slot 15", Position: "15" },
            { Name: "16", Label: "Slot 16", Position: "16" },
            { Name: "17", Label: "Slot 17", Position: "17" },
            { Name: "18", Label: "Slot 18", Position: "18" },
            { Name: "19", Label: "Slot 19", Position: "19" },
            { Name: "20", Label: "Slot 20", Position: "20" },
            { Name: "21", Label: "Slot 21", Position: "21" },
            { Name: "22", Label: "Slot 22", Position: "22" },
            { Name: "A", Label: "SM-A", Position: "A" },
            { Name: "B", Label: "SM-B", Position: "B" },
            { Name: "SCM", Label: "SCM", Position: "SCM" },
        },
			  DeviceBays: []DeviceBay{
        },
        InventoryItems: []InventoryItem{
        },
        Interfaces: []Interface{
            { Name: "Ethernet 1/S/1", Label: "", Type: "1000base-t", MgmtOnly: false },
        },
    },
}
