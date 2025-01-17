// Code generated by go generate; DO NOT EDIT.
package devices

var DeviceTypesMapaten = map[string]*DeviceData{
    "8-Port-PS/2-USB-VGA-KVM-Switch": {
        Manufacturer: "Aten",
        Model: "8-Port-PS/2-USB-VGA-KVM-Switch",
        Slug: "aten-cs1308",
        UHeight: 1,
        PartNumber: "CS1308",
        IsFullDepth: false,
        Airflow: "",
        FrontImage: false,
        RearImage: false,
        SubdeviceRole: "",
        Weight: 1.9,
        WeightUnit: "",
        IsPowered: false,
        ConsolePorts: []ConsolePort{
            { Name: "console", Type: "other", Label: "console", Poe: false },
        },
        ConsoleServerPorts: []ConsoleServerPort{
            { Name: "KVM/1", Type: "other", Label: "1" },
            { Name: "KVM/2", Type: "other", Label: "2" },
            { Name: "KVM/3", Type: "other", Label: "3" },
            { Name: "KVM/4", Type: "other", Label: "4" },
            { Name: "KVM/5", Type: "other", Label: "5" },
            { Name: "KVM/6", Type: "other", Label: "6" },
            { Name: "KVM/7", Type: "other", Label: "7" },
            { Name: "KVM/8", Type: "other", Label: "8" },
        },
        PowerPorts: []PowerPort{
            { Name: "DC In", Label: "", Type: "dc-terminal", MaximumDraw: 3, AllocatedDraw: 0 },
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
        },
    },
}
