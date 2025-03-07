// Code generated by go generate; DO NOT EDIT.
package devices

var DeviceTypesMapdepo = map[string]*DeviceData{
    "DEPO-Storage-1304": {
        Manufacturer: "DEPO",
        Model: "DEPO-Storage-1304",
        Slug: "depo-storage-1304",
        UHeight: 1,
        PartNumber: "",
        IsFullDepth: true,
        Airflow: "front-to-rear",
        FrontImage: false,
        RearImage: false,
        SubdeviceRole: "",
        Weight: 10.9,
        WeightUnit: "",
        IsPowered: false,
        ConsolePorts: []ConsolePort{
        },
        ConsoleServerPorts: []ConsoleServerPort{
        },
        PowerPorts: []PowerPort{
            { Name: "PSU1", Label: "", Type: "iec-60320-c14", MaximumDraw: 250, AllocatedDraw: 100 },
            { Name: "PSU2", Label: "", Type: "iec-60320-c14", MaximumDraw: 250, AllocatedDraw: 100 },
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
            { Name: "Eth1", Label: "", Type: "1000base-t", MgmtOnly: false },
            { Name: "Eth2", Label: "", Type: "1000base-t", MgmtOnly: false },
        },
    },
    "Race-SR518": {
        Manufacturer: "DEPO",
        Model: "Race-SR518",
        Slug: "depo-race-sr518",
        UHeight: 2,
        PartNumber: "Race-SR518",
        IsFullDepth: true,
        Airflow: "front-to-rear",
        FrontImage: false,
        RearImage: false,
        SubdeviceRole: "",
        Weight: 8.7,
        WeightUnit: "",
        IsPowered: false,
        ConsolePorts: []ConsolePort{
        },
        ConsoleServerPorts: []ConsoleServerPort{
        },
        PowerPorts: []PowerPort{
            { Name: "Power", Label: "", Type: "iec-60320-c14", MaximumDraw: 500, AllocatedDraw: 200 },
        },
        PowerOutlets: []PowerOutlet{
        },
        FrontPorts: []FrontPort{
        },
        RearPorts: []RearPort{
        },
        ModuleBays: []ModuleBay{
            { Name: "PCIe-1 LP x16", Label: "", Position: "PCIe-1" },
            { Name: "PCIe-2 LP x16", Label: "", Position: "PCIe-2" },
            { Name: "PCIe-3 LP x16", Label: "", Position: "PCIe-3" },
            { Name: "PCIe-4 LP x8", Label: "", Position: "PCIe-4" },
            { Name: "PCIe-5 LP x8", Label: "", Position: "PCIe-7" },
        },
			  DeviceBays: []DeviceBay{
        },
        InventoryItems: []InventoryItem{
        },
        Interfaces: []Interface{
            { Name: "Eth-1", Label: "", Type: "1000base-t", MgmtOnly: false },
            { Name: "Eth-2", Label: "", Type: "1000base-t", MgmtOnly: false },
            { Name: "IPMI", Label: "", Type: "1000base-t", MgmtOnly: true },
        },
    },
}
