// Code generated by go generate; DO NOT EDIT.
package devices

var DeviceTypesMapzpe = map[string]*DeviceData{
    "GSR-T8-BASE": {
        Manufacturer: "ZPE",
        Model: "GSR-T8-BASE",
        Slug: "zpe-gsr-t8-base",
        UHeight: 1,
        PartNumber: "GSR-T8-BASE",
        IsFullDepth: false,
        Airflow: "",
        FrontImage: false,
        RearImage: false,
        SubdeviceRole: "",
        Weight: 0,
        WeightUnit: "",
        IsPowered: false,
        ConsolePorts: []ConsolePort{
            { Name: "Console", Type: "rj-45", Label: "", Poe: false },
        },
        ConsoleServerPorts: []ConsoleServerPort{
            { Name: "ttyS1", Type: "rj-45", Label: "" },
            { Name: "ttyS2", Type: "rj-45", Label: "" },
            { Name: "ttyS3", Type: "rj-45", Label: "" },
            { Name: "ttyS4", Type: "rj-45", Label: "" },
            { Name: "ttyS5", Type: "rj-45", Label: "" },
            { Name: "ttyS6", Type: "rj-45", Label: "" },
            { Name: "ttyS7", Type: "rj-45", Label: "" },
            { Name: "ttyS8", Type: "rj-45", Label: "" },
            { Name: "usbS0-1", Type: "usb-a", Label: "" },
            { Name: "usbS0-2", Type: "usb-a", Label: "" },
            { Name: "usbS0-3", Type: "usb-a", Label: "" },
            { Name: "usbS0-4", Type: "usb-a", Label: "" },
        },
        PowerPorts: []PowerPort{
            { Name: "PSU1", Label: "", Type: "dc-terminal", MaximumDraw: 0, AllocatedDraw: 45 },
            { Name: "PSU2", Label: "", Type: "dc-terminal", MaximumDraw: 0, AllocatedDraw: 45 },
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
            { Name: "ETH0", Label: "", Type: "1000base-t", MgmtOnly: false },
            { Name: "SFP0", Label: "", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "SFP1", Label: "", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "netS1", Label: "", Type: "1000base-t", MgmtOnly: false },
            { Name: "netS2", Label: "", Type: "1000base-t", MgmtOnly: false },
            { Name: "netS3", Label: "", Type: "1000base-t", MgmtOnly: false },
            { Name: "netS4", Label: "", Type: "1000base-t", MgmtOnly: false },
            { Name: "netS5", Label: "", Type: "1000base-t", MgmtOnly: false },
            { Name: "netS6", Label: "", Type: "1000base-t", MgmtOnly: false },
            { Name: "netS7", Label: "", Type: "1000base-t", MgmtOnly: false },
            { Name: "netS8", Label: "", Type: "1000base-t", MgmtOnly: false },
        },
    },
    "NSC-T48R-STND-DAC": {
        Manufacturer: "ZPE",
        Model: "NSC-T48R-STND-DAC",
        Slug: "zpe-nsc-t48r-stnd-dac",
        UHeight: 1,
        PartNumber: "NSC-T48R-STND-DAC",
        IsFullDepth: false,
        Airflow: "",
        FrontImage: false,
        RearImage: false,
        SubdeviceRole: "",
        Weight: 0,
        WeightUnit: "",
        IsPowered: false,
        ConsolePorts: []ConsolePort{
            { Name: "Console", Type: "rj-45", Label: "", Poe: false },
        },
        ConsoleServerPorts: []ConsoleServerPort{
            { Name: "Port 1", Type: "rj-45", Label: "" },
            { Name: "Port 2", Type: "rj-45", Label: "" },
            { Name: "Port 3", Type: "rj-45", Label: "" },
            { Name: "Port 4", Type: "rj-45", Label: "" },
            { Name: "Port 5", Type: "rj-45", Label: "" },
            { Name: "Port 6", Type: "rj-45", Label: "" },
            { Name: "Port 7", Type: "rj-45", Label: "" },
            { Name: "Port 8", Type: "rj-45", Label: "" },
            { Name: "Port 9", Type: "rj-45", Label: "" },
            { Name: "Port 10", Type: "rj-45", Label: "" },
            { Name: "Port 11", Type: "rj-45", Label: "" },
            { Name: "Port 12", Type: "rj-45", Label: "" },
            { Name: "Port 13", Type: "rj-45", Label: "" },
            { Name: "Port 14", Type: "rj-45", Label: "" },
            { Name: "Port 15", Type: "rj-45", Label: "" },
            { Name: "Port 16", Type: "rj-45", Label: "" },
            { Name: "Port 17", Type: "rj-45", Label: "" },
            { Name: "Port 18", Type: "rj-45", Label: "" },
            { Name: "Port 19", Type: "rj-45", Label: "" },
            { Name: "Port 20", Type: "rj-45", Label: "" },
            { Name: "Port 21", Type: "rj-45", Label: "" },
            { Name: "Port 22", Type: "rj-45", Label: "" },
            { Name: "Port 23", Type: "rj-45", Label: "" },
            { Name: "Port 24", Type: "rj-45", Label: "" },
            { Name: "Port 25", Type: "rj-45", Label: "" },
            { Name: "Port 26", Type: "rj-45", Label: "" },
            { Name: "Port 27", Type: "rj-45", Label: "" },
            { Name: "Port 28", Type: "rj-45", Label: "" },
            { Name: "Port 29", Type: "rj-45", Label: "" },
            { Name: "Port 30", Type: "rj-45", Label: "" },
            { Name: "Port 31", Type: "rj-45", Label: "" },
            { Name: "Port 32", Type: "rj-45", Label: "" },
            { Name: "Port 33", Type: "rj-45", Label: "" },
            { Name: "Port 34", Type: "rj-45", Label: "" },
            { Name: "Port 35", Type: "rj-45", Label: "" },
            { Name: "Port 36", Type: "rj-45", Label: "" },
            { Name: "Port 37", Type: "rj-45", Label: "" },
            { Name: "Port 38", Type: "rj-45", Label: "" },
            { Name: "Port 39", Type: "rj-45", Label: "" },
            { Name: "Port 40", Type: "rj-45", Label: "" },
            { Name: "Port 41", Type: "rj-45", Label: "" },
            { Name: "Port 42", Type: "rj-45", Label: "" },
            { Name: "Port 43", Type: "rj-45", Label: "" },
            { Name: "Port 44", Type: "rj-45", Label: "" },
            { Name: "Port 45", Type: "rj-45", Label: "" },
            { Name: "Port 46", Type: "rj-45", Label: "" },
            { Name: "Port 47", Type: "rj-45", Label: "" },
            { Name: "Port 48", Type: "rj-45", Label: "" },
        },
        PowerPorts: []PowerPort{
            { Name: "PSU1", Label: "", Type: "iec-60320-c14", MaximumDraw: 0, AllocatedDraw: 30 },
            { Name: "PSU2", Label: "", Type: "iec-60320-c14", MaximumDraw: 0, AllocatedDraw: 30 },
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
            { Name: "eth0", Label: "", Type: "1000base-t", MgmtOnly: true },
            { Name: "eth1", Label: "", Type: "1000base-t", MgmtOnly: true },
        },
    },
    "NSR-BASE-DAC": {
        Manufacturer: "ZPE",
        Model: "NSR-BASE-DAC",
        Slug: "zpe-nsr-base-dac",
        UHeight: 1,
        PartNumber: "NSR-BASE-DAC",
        IsFullDepth: true,
        Airflow: "",
        FrontImage: false,
        RearImage: false,
        SubdeviceRole: "",
        Weight: 0,
        WeightUnit: "",
        IsPowered: false,
        ConsolePorts: []ConsolePort{
            { Name: "Console", Type: "rj-45", Label: "", Poe: false },
            { Name: "HDMI", Type: "other", Label: "", Poe: false },
        },
        ConsoleServerPorts: []ConsoleServerPort{
            { Name: "usbS0-1", Type: "usb-a", Label: "" },
            { Name: "usbS0-2", Type: "usb-a", Label: "" },
            { Name: "usbS0-3", Type: "usb-a", Label: "" },
        },
        PowerPorts: []PowerPort{
            { Name: "PS1", Label: "", Type: "dc-terminal", MaximumDraw: 0, AllocatedDraw: 45 },
            { Name: "PS2", Label: "", Type: "dc-terminal", MaximumDraw: 0, AllocatedDraw: 45 },
        },
        PowerOutlets: []PowerOutlet{
        },
        FrontPorts: []FrontPort{
        },
        RearPorts: []RearPort{
        },
        ModuleBays: []ModuleBay{
            { Name: "slot-1", Label: "", Position: "1" },
            { Name: "slot-2", Label: "", Position: "2" },
            { Name: "slot-3", Label: "", Position: "3" },
            { Name: "slot-4", Label: "", Position: "4" },
            { Name: "slot-5", Label: "", Position: "5" },
        },
			  DeviceBays: []DeviceBay{
        },
        InventoryItems: []InventoryItem{
        },
        Interfaces: []Interface{
            { Name: "ETH0", Label: "", Type: "1000base-t", MgmtOnly: false },
            { Name: "ETH1", Label: "", Type: "1000base-t", MgmtOnly: false },
            { Name: "SFP0", Label: "", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "SFP1", Label: "", Type: "10gbase-x-sfpp", MgmtOnly: false },
        },
    },
}
