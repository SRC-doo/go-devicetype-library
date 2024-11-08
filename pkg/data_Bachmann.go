// Code generated by go generate; DO NOT EDIT.
package devices

var DeviceTypesMapBachmann = map[string]*DeviceData{
    "Bachmann 6x Power Socket": {
        Manufacturer: "Bachmann",
        Model: "Bachmann 6x Power Socket",
        Slug: "bachmann-6x-power-socket",
        UHeight: 1,
        PartNumber: "4016514017232",
        IsFullDepth: false,
        Airflow: "passive",
        FrontImage: false,
        RearImage: false,
        SubdeviceRole: "",
        Weight: 2,
        WeightUnit: "",
        IsPowered: false,
        ConsolePorts: []ConsolePort{
        },
        ConsoleServerPorts: []ConsoleServerPort{
        },
        PowerPorts: []PowerPort{
            { Name: "Power 1", Label: "", Type: "iec-60320-c14", MaximumDraw: 0, AllocatedDraw: 0 },
        },
        PowerOutlets: []PowerOutlet{
            { Name: "Outlet 1", Type: "ita-f", Label: "", PowerPort: "Power 1", FeedLeg: "", MaximumDraw: 0, AllocatedDraw: 0 },
            { Name: "Outlet 2", Type: "ita-f", Label: "", PowerPort: "Power 1", FeedLeg: "", MaximumDraw: 0, AllocatedDraw: 0 },
            { Name: "Outlet 3", Type: "ita-f", Label: "", PowerPort: "Power 1", FeedLeg: "", MaximumDraw: 0, AllocatedDraw: 0 },
            { Name: "Outlet 4", Type: "ita-f", Label: "", PowerPort: "Power 1", FeedLeg: "", MaximumDraw: 0, AllocatedDraw: 0 },
            { Name: "Outlet 5", Type: "ita-f", Label: "", PowerPort: "Power 1", FeedLeg: "", MaximumDraw: 0, AllocatedDraw: 0 },
            { Name: "Outlet 6", Type: "ita-f", Label: "", PowerPort: "Power 1", FeedLeg: "", MaximumDraw: 0, AllocatedDraw: 0 },
            { Name: "Outlet 7", Type: "ita-f", Label: "", PowerPort: "Power 1", FeedLeg: "", MaximumDraw: 0, AllocatedDraw: 0 },
            { Name: "Outlet 8", Type: "ita-f", Label: "", PowerPort: "Power 1", FeedLeg: "", MaximumDraw: 0, AllocatedDraw: 0 },
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
