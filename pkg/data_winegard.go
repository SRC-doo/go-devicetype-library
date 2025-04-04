// Code generated by go generate; DO NOT EDIT.
package devices

var DeviceTypesMapwinegard = map[string]*DeviceData{
    "Gateway 4G": {
        Manufacturer: "Winegard",
        Model: "Gateway 4G",
        Slug: "winegard-gateway-4g",
        UHeight: 0,
        PartNumber: "GW-1000",
        IsFullDepth: false,
        Airflow: "passive",
        FrontImage: false,
        RearImage: false,
        SubdeviceRole: "",
        Weight: 0.8,
        WeightUnit: "",
        IsPowered: false,
        ConsolePorts: []ConsolePort{
        },
        ConsoleServerPorts: []ConsoleServerPort{
        },
        PowerPorts: []PowerPort{
            { Name: "PSU1", Label: "DC&#43;12V/1A", Type: "dc-terminal", MaximumDraw: 16, AllocatedDraw: 0 },
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
            { Name: "wlan0", Label: "Red - WiFi", Type: "ieee802.11n", MgmtOnly: false },
            { Name: "eth0", Label: "Ethernet - LAN/WAN", Type: "1000base-t", MgmtOnly: false },
            { Name: "ant4", Label: "Yellow - 4G/LTE", Type: "lte", MgmtOnly: false },
            { Name: "ant5", Label: "Green - 4G/LTE", Type: "lte", MgmtOnly: false },
            { Name: "sim0", Label: "", Type: "other", MgmtOnly: false },
            { Name: "sim1", Label: "", Type: "other", MgmtOnly: false },
        },
    },
    "Gateway 4G Pro": {
        Manufacturer: "Winegard",
        Model: "Gateway 4G Pro",
        Slug: "winegard-gateway-4g-pro",
        UHeight: 0,
        PartNumber: "GW-PRO1",
        IsFullDepth: false,
        Airflow: "passive",
        FrontImage: false,
        RearImage: false,
        SubdeviceRole: "",
        Weight: 0.8,
        WeightUnit: "",
        IsPowered: false,
        ConsolePorts: []ConsolePort{
        },
        ConsoleServerPorts: []ConsoleServerPort{
        },
        PowerPorts: []PowerPort{
            { Name: "PSU1", Label: "DC&#43;12V/1A", Type: "dc-terminal", MaximumDraw: 16, AllocatedDraw: 0 },
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
            { Name: "wlan0", Label: "Red - WiFi", Type: "ieee802.11ac", MgmtOnly: false },
            { Name: "eth0", Label: "Ethernet - LAN/WAN", Type: "1000base-t", MgmtOnly: false },
            { Name: "ant4", Label: "Yellow - 4G/LTE", Type: "lte", MgmtOnly: false },
            { Name: "ant5", Label: "Green - 4G/LTE", Type: "lte", MgmtOnly: false },
            { Name: "sim0", Label: "", Type: "other", MgmtOnly: false },
            { Name: "sim1", Label: "", Type: "other", MgmtOnly: false },
        },
    },
    "Gateway 5G": {
        Manufacturer: "Winegard",
        Model: "Gateway 5G",
        Slug: "winegard-gateway-5g",
        UHeight: 0,
        PartNumber: "GW-5G01",
        IsFullDepth: false,
        Airflow: "passive",
        FrontImage: false,
        RearImage: false,
        SubdeviceRole: "",
        Weight: 1,
        WeightUnit: "",
        IsPowered: false,
        ConsolePorts: []ConsolePort{
        },
        ConsoleServerPorts: []ConsoleServerPort{
        },
        PowerPorts: []PowerPort{
            { Name: "PSU1", Label: "DC&#43;12V/1.7A", Type: "dc-terminal", MaximumDraw: 27, AllocatedDraw: 0 },
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
            { Name: "wlan0", Label: "Red - WiFi", Type: "ieee802.11ac", MgmtOnly: false },
            { Name: "eth0", Label: "", Type: "1000base-t", MgmtOnly: false },
            { Name: "ant3", Label: "Green - 5G/LTE", Type: "5g", MgmtOnly: false },
            { Name: "ant4", Label: "Yellow - 5G/LTE", Type: "5g", MgmtOnly: false },
            { Name: "ant5", Label: "Green - 5G/LTE", Type: "5g", MgmtOnly: false },
            { Name: "ant6", Label: "Yellow - 5G/LTE", Type: "5g", MgmtOnly: false },
            { Name: "sim0", Label: "", Type: "other", MgmtOnly: false },
            { Name: "sim1", Label: "", Type: "other", MgmtOnly: false },
        },
    },
}
