// Code generated by go generate; DO NOT EDIT.
package devices

var DeviceTypesMapnetgate = map[string]*DeviceData{
    "1100 Security Gateway": {
        Manufacturer: "Netgate",
        Model: "1100 Security Gateway",
        Slug: "netgate-1100-security-gateway",
        UHeight: 1,
        PartNumber: "",
        IsFullDepth: false,
        Airflow: "front-to-rear",
        FrontImage: false,
        RearImage: false,
        SubdeviceRole: "",
        Weight: 0.45,
        WeightUnit: "",
        IsPowered: false,
        ConsolePorts: []ConsolePort{
            { Name: "console-usb", Type: "usb-micro-b", Label: "", Poe: false },
        },
        ConsoleServerPorts: []ConsoleServerPort{
        },
        PowerPorts: []PowerPort{
            { Name: "psu1", Label: "", Type: "dc-terminal", MaximumDraw: 0, AllocatedDraw: 0 },
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
            { Name: "mvneta0.4090", Label: "WAN", Type: "1000base-t", MgmtOnly: false },
            { Name: "mvneta0.4091", Label: "LAN", Type: "1000base-t", MgmtOnly: false },
            { Name: "mvneta0.4092", Label: "OPT", Type: "1000base-t", MgmtOnly: false },
        },
    },
    "6100 Security Gateway": {
        Manufacturer: "Netgate",
        Model: "6100 Security Gateway",
        Slug: "netgate-6100-security-gateway",
        UHeight: 1,
        PartNumber: "",
        IsFullDepth: false,
        Airflow: "front-to-rear",
        FrontImage: false,
        RearImage: false,
        SubdeviceRole: "",
        Weight: 1.9,
        WeightUnit: "",
        IsPowered: false,
        ConsolePorts: []ConsolePort{
            { Name: "console-usb", Type: "usb-micro-b", Label: "", Poe: false },
            { Name: "console-rj45", Type: "rj-45", Label: "", Poe: false },
        },
        ConsoleServerPorts: []ConsoleServerPort{
        },
        PowerPorts: []PowerPort{
            { Name: "psu1", Label: "", Type: "dc-terminal", MaximumDraw: 0, AllocatedDraw: 0 },
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
            { Name: "ix3", Label: "WAN1", Type: "1000base-x-sfp", MgmtOnly: false },
            { Name: "ix2", Label: "WAN2", Type: "1000base-x-sfp", MgmtOnly: false },
            { Name: "ix0", Label: "WAN3", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "ix1", Label: "WAN4", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "igc0", Label: "LAN1", Type: "2.5gbase-t", MgmtOnly: false },
            { Name: "igc1", Label: "LAN2", Type: "2.5gbase-t", MgmtOnly: false },
            { Name: "igc2", Label: "LAN3", Type: "2.5gbase-t", MgmtOnly: false },
            { Name: "igc3", Label: "LAN4", Type: "2.5gbase-t", MgmtOnly: false },
        },
    },
    "7100 Security Gateway": {
        Manufacturer: "Netgate",
        Model: "7100 Security Gateway",
        Slug: "netgate-7100-security-gateway",
        UHeight: 1,
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
            { Name: "console", Type: "usb-mini-a", Label: "", Poe: false },
        },
        ConsoleServerPorts: []ConsoleServerPort{
        },
        PowerPorts: []PowerPort{
            { Name: "PSU1", Label: "", Type: "iec-60320-c14", MaximumDraw: 360, AllocatedDraw: 0 },
        },
        PowerOutlets: []PowerOutlet{
        },
        FrontPorts: []FrontPort{
        },
        RearPorts: []RearPort{
        },
        ModuleBays: []ModuleBay{
            { Name: "pci1", Label: "", Position: "PCI1" },
        },
			  DeviceBays: []DeviceBay{
        },
        InventoryItems: []InventoryItem{
        },
        Interfaces: []Interface{
            { Name: "eth1", Label: "", Type: "1000base-t", MgmtOnly: false },
            { Name: "eth2", Label: "", Type: "1000base-t", MgmtOnly: false },
            { Name: "eth3", Label: "", Type: "1000base-t", MgmtOnly: false },
            { Name: "eth4", Label: "", Type: "1000base-t", MgmtOnly: false },
            { Name: "eth5", Label: "", Type: "1000base-t", MgmtOnly: false },
            { Name: "eth6", Label: "", Type: "1000base-t", MgmtOnly: false },
            { Name: "eth7", Label: "", Type: "1000base-t", MgmtOnly: false },
            { Name: "eth8", Label: "", Type: "1000base-t", MgmtOnly: false },
            { Name: "ix0", Label: "", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "ix1", Label: "", Type: "10gbase-x-sfpp", MgmtOnly: false },
        },
    },
    "8200 Max PFSense&#43; Security Gateway": {
        Manufacturer: "Netgate",
        Model: "8200 Max PFSense&#43; Security Gateway",
        Slug: "netgate-8200-max-pfsense-plus-security-gateway",
        UHeight: 1,
        PartNumber: "",
        IsFullDepth: false,
        Airflow: "front-to-rear",
        FrontImage: false,
        RearImage: false,
        SubdeviceRole: "",
        Weight: 0,
        WeightUnit: "",
        IsPowered: false,
        ConsolePorts: []ConsolePort{
            { Name: "console-usb", Type: "usb-micro-b", Label: "", Poe: false },
            { Name: "console-rj45", Type: "rj-45", Label: "", Poe: false },
        },
        ConsoleServerPorts: []ConsoleServerPort{
        },
        PowerPorts: []PowerPort{
            { Name: "psu1", Label: "", Type: "dc-terminal", MaximumDraw: 0, AllocatedDraw: 0 },
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
            { Name: "ix0", Label: "WAN3", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "ix1", Label: "WAN4", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "ix2", Label: "WAN2", Type: "1000base-x-sfp", MgmtOnly: false },
            { Name: "ix3", Label: "WAN1", Type: "1000base-x-sfp", MgmtOnly: false },
            { Name: "igc0", Label: "LAN1", Type: "2.5gbase-t", MgmtOnly: false },
            { Name: "igc1", Label: "LAN2", Type: "2.5gbase-t", MgmtOnly: false },
            { Name: "igc2", Label: "LAN3", Type: "2.5gbase-t", MgmtOnly: false },
            { Name: "igc3", Label: "LAN4", Type: "2.5gbase-t", MgmtOnly: false },
        },
    },
    "8300 Security Gateway": {
        Manufacturer: "Netgate",
        Model: "8300 Security Gateway",
        Slug: "netgate-8300-security-gateway",
        UHeight: 1,
        PartNumber: "",
        IsFullDepth: false,
        Airflow: "front-to-rear",
        FrontImage: false,
        RearImage: false,
        SubdeviceRole: "",
        Weight: 6.05,
        WeightUnit: "",
        IsPowered: false,
        ConsolePorts: []ConsolePort{
            { Name: "console-rj45", Type: "rj-45", Label: "", Poe: false },
        },
        ConsoleServerPorts: []ConsoleServerPort{
        },
        PowerPorts: []PowerPort{
            { Name: "psu1", Label: "", Type: "iec-60320-c16", MaximumDraw: 0, AllocatedDraw: 0 },
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
            { Name: "ice0", Label: "P0", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "ice1", Label: "P1", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "ice2", Label: "P2", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "ice3", Label: "P3", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "ice4", Label: "P4", Type: "1000base-x-sfp", MgmtOnly: false },
            { Name: "ice5", Label: "P5", Type: "1000base-x-sfp", MgmtOnly: false },
            { Name: "ice6", Label: "P5", Type: "1000base-x-sfp", MgmtOnly: false },
            { Name: "ice7", Label: "P6", Type: "1000base-x-sfp", MgmtOnly: false },
            { Name: "igc2", Label: "P7", Type: "2.5gbase-t", MgmtOnly: false },
            { Name: "igc1", Label: "P8", Type: "2.5gbase-t", MgmtOnly: false },
            { Name: "igc0", Label: "P7", Type: "2.5gbase-t", MgmtOnly: false },
        },
    },
}
