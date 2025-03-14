// Code generated by go generate; DO NOT EDIT.
package devices

var DeviceTypesMapinspur = map[string]*DeviceData{
    "NF5180M6": {
        Manufacturer: "Inspur",
        Model: "NF5180M6",
        Slug: "inspur-nf5180m6",
        UHeight: 1,
        PartNumber: "NF5180M6",
        IsFullDepth: true,
        Airflow: "front-to-rear",
        FrontImage: false,
        RearImage: false,
        SubdeviceRole: "",
        Weight: 31,
        WeightUnit: "",
        IsPowered: false,
        ConsolePorts: []ConsolePort{
            { Name: "BMC/System Serial", Type: "other", Label: "", Poe: false },
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
            { Name: "PSU0", Label: "", Position: "PSU0" },
            { Name: "PSU1", Label: "", Position: "PSU1" },
            { Name: "OCP3", Label: "", Position: "OCP3" },
            { Name: "PCIe0", Label: "Left PCIe Riser-Card", Position: "PCIe0" },
            { Name: "PCIe1", Label: "Left PCIe Riser-Card", Position: "PCIe1" },
            { Name: "PCIe2", Label: "Right PCIe Riser-Card", Position: "PCIe2" },
        },
			  DeviceBays: []DeviceBay{
        },
        InventoryItems: []InventoryItem{
        },
        Interfaces: []Interface{
            { Name: "LAN0", Label: "", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "LAN1", Label: "", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "BMC", Label: "", Type: "1000base-t", MgmtOnly: true },
        },
    },
    "NF8260M6": {
        Manufacturer: "Inspur",
        Model: "NF8260M6",
        Slug: "inspur-nf8260m6",
        UHeight: 2,
        PartNumber: "NF8260M6",
        IsFullDepth: true,
        Airflow: "front-to-rear",
        FrontImage: false,
        RearImage: false,
        SubdeviceRole: "",
        Weight: 0,
        WeightUnit: "",
        IsPowered: false,
        ConsolePorts: []ConsolePort{
            { Name: "BMC Debug Serial Port", Type: "other", Label: "", Poe: false },
            { Name: "System Serial Port", Type: "other", Label: "", Poe: false },
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
            { Name: "PSU0", Label: "", Position: "Down" },
            { Name: "PSU1", Label: "", Position: "Top" },
            { Name: "OCP3", Label: "", Position: "OCP3" },
            { Name: "PCIe0", Label: "Motherboard", Position: "PCIe0" },
            { Name: "PCIe1", Label: "Motherboard", Position: "PCIe1" },
            { Name: "PCIe2", Label: "Motherboard", Position: "PCIe2" },
            { Name: "PCIe3", Label: "Motherboard", Position: "PCIe3" },
            { Name: "PCIe4", Label: "Motherboard", Position: "PCIe4" },
            { Name: "PCIe5", Label: "Motherboard", Position: "PCIe5" },
            { Name: "PCIe6", Label: "PCIe riser module 0", Position: "PCIe6" },
            { Name: "PCIe7", Label: "PCIe riser module 0", Position: "PCIe7" },
            { Name: "PCIe8", Label: "PCIe riser module 0", Position: "PCIe8" },
            { Name: "PCIe9", Label: "PCIe riser module 1", Position: "PCIe9" },
            { Name: "PCIe10", Label: "PCIe riser module 1", Position: "PCIe10" },
            { Name: "PCIe11", Label: "PCIe riser module 1", Position: "PCIe11" },
        },
			  DeviceBays: []DeviceBay{
        },
        InventoryItems: []InventoryItem{
        },
        Interfaces: []Interface{
            { Name: "BMC", Label: "", Type: "1000base-t", MgmtOnly: true },
        },
    },
}
