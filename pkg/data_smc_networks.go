// Code generated by go generate; DO NOT EDIT.
package devices

var DeviceTypesMapsmc_networks = map[string]*DeviceData{
    "SMCGS8P-SMART": {
        Manufacturer: "SMC Networks",
        Model: "SMCGS8P-SMART",
        Slug: "smc-networks-smcgs8p-smart",
        UHeight: 1,
        PartNumber: "SMCGS8P-SMART",
        IsFullDepth: false,
        Airflow: "left-to-right",
        FrontImage: true,
        RearImage: true,
        SubdeviceRole: "",
        Weight: 2.06,
        WeightUnit: "",
        IsPowered: false,
        ConsolePorts: []ConsolePort{
        },
        ConsoleServerPorts: []ConsoleServerPort{
        },
        PowerPorts: []PowerPort{
            { Name: "PSU0", Label: "", Type: "iec-60320-c14", MaximumDraw: 165, AllocatedDraw: 0 },
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
            { Name: "GigabitEthernet1", Label: "", Type: "1000base-t", MgmtOnly: false },
            { Name: "GigabitEthernet2", Label: "", Type: "1000base-t", MgmtOnly: false },
            { Name: "GigabitEthernet3", Label: "", Type: "1000base-t", MgmtOnly: false },
            { Name: "GigabitEthernet4", Label: "", Type: "1000base-t", MgmtOnly: false },
            { Name: "GigabitEthernet5", Label: "", Type: "1000base-t", MgmtOnly: false },
            { Name: "GigabitEthernet6", Label: "", Type: "1000base-t", MgmtOnly: false },
            { Name: "GigabitEthernet7", Label: "", Type: "1000base-t", MgmtOnly: false },
            { Name: "GigabitEthernet8", Label: "", Type: "1000base-t", MgmtOnly: false },
            { Name: "GigabitEthernetSFP1", Label: "", Type: "1000base-x-sfp", MgmtOnly: false },
        },
    },
}
