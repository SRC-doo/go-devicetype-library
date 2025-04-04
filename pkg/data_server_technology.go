// Code generated by go generate; DO NOT EDIT.
package devices

var DeviceTypesMapserver_technology = map[string]*DeviceData{
    "C1W08HC-0ABA2BAC": {
        Manufacturer: "Server Technology",
        Model: "C1W08HC-0ABA2BAC",
        Slug: "server-technology-c1w08hc-0aba2bac",
        UHeight: 1,
        PartNumber: "C1W08HC-0ABA2BAC",
        IsFullDepth: false,
        Airflow: "passive",
        FrontImage: false,
        RearImage: false,
        SubdeviceRole: "",
        Weight: 0,
        WeightUnit: "",
        IsPowered: false,
        ConsolePorts: []ConsolePort{
            { Name: "ser", Type: "rj-45", Label: "SER", Poe: false },
            { Name: "th1", Type: "rj-11", Label: "T/H1", Poe: false },
            { Name: "th2", Type: "rj-11", Label: "T/H2", Poe: false },
        },
        ConsoleServerPorts: []ConsoleServerPort{
            { Name: "link", Type: "rj-12", Label: "LINK" },
        },
        PowerPorts: []PowerPort{
            { Name: "Power Port 1", Label: "", Type: "iec-60320-c20", MaximumDraw: 3300, AllocatedDraw: 0 },
        },
        PowerOutlets: []PowerOutlet{
            { Name: "Outlet 1", Type: "iec-60320-c13", Label: "", PowerPort: "Power Port 1", FeedLeg: "A", MaximumDraw: 0, AllocatedDraw: 0 },
            { Name: "Outlet 2", Type: "iec-60320-c13", Label: "", PowerPort: "Power Port 1", FeedLeg: "A", MaximumDraw: 0, AllocatedDraw: 0 },
            { Name: "Outlet 3", Type: "iec-60320-c13", Label: "", PowerPort: "Power Port 1", FeedLeg: "A", MaximumDraw: 0, AllocatedDraw: 0 },
            { Name: "Outlet 4", Type: "iec-60320-c13", Label: "", PowerPort: "Power Port 1", FeedLeg: "A", MaximumDraw: 0, AllocatedDraw: 0 },
            { Name: "Outlet 5", Type: "iec-60320-c13", Label: "", PowerPort: "Power Port 1", FeedLeg: "A", MaximumDraw: 0, AllocatedDraw: 0 },
            { Name: "Outlet 6", Type: "iec-60320-c13", Label: "", PowerPort: "Power Port 1", FeedLeg: "A", MaximumDraw: 0, AllocatedDraw: 0 },
            { Name: "Outlet 7", Type: "iec-60320-c13", Label: "", PowerPort: "Power Port 1", FeedLeg: "A", MaximumDraw: 0, AllocatedDraw: 0 },
            { Name: "Outlet 8", Type: "iec-60320-c13", Label: "", PowerPort: "Power Port 1", FeedLeg: "A", MaximumDraw: 0, AllocatedDraw: 0 },
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
            { Name: "net", Label: "NET", Type: "1000base-t", MgmtOnly: true },
        },
    },
    "C2WG24SN-4PJN5D (Primary)": {
        Manufacturer: "Server Technology",
        Model: "C2WG24SN-4PJN5D (Primary)",
        Slug: "server-technology-c2wg24sn-4pjn5d",
        UHeight: 2,
        PartNumber: "C2WG24SN-4PJN5D",
        IsFullDepth: true,
        Airflow: "passive",
        FrontImage: false,
        RearImage: false,
        SubdeviceRole: "",
        Weight: 0,
        WeightUnit: "",
        IsPowered: false,
        ConsolePorts: []ConsolePort{
            { Name: "ser", Type: "rj-45", Label: "SER", Poe: false },
            { Name: "th1", Type: "rj-11", Label: "T/H1", Poe: false },
            { Name: "th2", Type: "rj-11", Label: "T/H2", Poe: false },
        },
        ConsoleServerPorts: []ConsoleServerPort{
            { Name: "link", Type: "rj-12", Label: "LINK" },
        },
        PowerPorts: []PowerPort{
            { Name: "Power Port 1", Label: "", Type: "iec-60309-3p-n-e-6h", MaximumDraw: 0, AllocatedDraw: 0 },
        },
        PowerOutlets: []PowerOutlet{
            { Name: "Outlet 1", Type: "iec-60320-c19", Label: "", PowerPort: "Power Port 1", FeedLeg: "A", MaximumDraw: 0, AllocatedDraw: 0 },
            { Name: "Outlet 2", Type: "iec-60320-c13", Label: "", PowerPort: "Power Port 1", FeedLeg: "A", MaximumDraw: 0, AllocatedDraw: 0 },
            { Name: "Outlet 3", Type: "iec-60320-c13", Label: "", PowerPort: "Power Port 1", FeedLeg: "A", MaximumDraw: 0, AllocatedDraw: 0 },
            { Name: "Outlet 4", Type: "iec-60320-c13", Label: "", PowerPort: "Power Port 1", FeedLeg: "A", MaximumDraw: 0, AllocatedDraw: 0 },
            { Name: "Outlet 5", Type: "iec-60320-c19", Label: "", PowerPort: "Power Port 1", FeedLeg: "A", MaximumDraw: 0, AllocatedDraw: 0 },
            { Name: "Outlet 6", Type: "iec-60320-c13", Label: "", PowerPort: "Power Port 1", FeedLeg: "A", MaximumDraw: 0, AllocatedDraw: 0 },
            { Name: "Outlet 7", Type: "iec-60320-c13", Label: "", PowerPort: "Power Port 1", FeedLeg: "A", MaximumDraw: 0, AllocatedDraw: 0 },
            { Name: "Outlet 8", Type: "iec-60320-c13", Label: "", PowerPort: "Power Port 1", FeedLeg: "A", MaximumDraw: 0, AllocatedDraw: 0 },
            { Name: "Outlet 9", Type: "iec-60320-c19", Label: "", PowerPort: "Power Port 1", FeedLeg: "B", MaximumDraw: 0, AllocatedDraw: 0 },
            { Name: "Outlet 10", Type: "iec-60320-c13", Label: "", PowerPort: "Power Port 1", FeedLeg: "B", MaximumDraw: 0, AllocatedDraw: 0 },
            { Name: "Outlet 11", Type: "iec-60320-c13", Label: "", PowerPort: "Power Port 1", FeedLeg: "B", MaximumDraw: 0, AllocatedDraw: 0 },
            { Name: "Outlet 12", Type: "iec-60320-c13", Label: "", PowerPort: "Power Port 1", FeedLeg: "B", MaximumDraw: 0, AllocatedDraw: 0 },
            { Name: "Outlet 13", Type: "iec-60320-c19", Label: "", PowerPort: "Power Port 1", FeedLeg: "B", MaximumDraw: 0, AllocatedDraw: 0 },
            { Name: "Outlet 14", Type: "iec-60320-c13", Label: "", PowerPort: "Power Port 1", FeedLeg: "B", MaximumDraw: 0, AllocatedDraw: 0 },
            { Name: "Outlet 15", Type: "iec-60320-c13", Label: "", PowerPort: "Power Port 1", FeedLeg: "B", MaximumDraw: 0, AllocatedDraw: 0 },
            { Name: "Outlet 16", Type: "iec-60320-c13", Label: "", PowerPort: "Power Port 1", FeedLeg: "B", MaximumDraw: 0, AllocatedDraw: 0 },
            { Name: "Outlet 17", Type: "iec-60320-c19", Label: "", PowerPort: "Power Port 1", FeedLeg: "C", MaximumDraw: 0, AllocatedDraw: 0 },
            { Name: "Outlet 18", Type: "iec-60320-c13", Label: "", PowerPort: "Power Port 1", FeedLeg: "C", MaximumDraw: 0, AllocatedDraw: 0 },
            { Name: "Outlet 19", Type: "iec-60320-c13", Label: "", PowerPort: "Power Port 1", FeedLeg: "C", MaximumDraw: 0, AllocatedDraw: 0 },
            { Name: "Outlet 20", Type: "iec-60320-c13", Label: "", PowerPort: "Power Port 1", FeedLeg: "C", MaximumDraw: 0, AllocatedDraw: 0 },
            { Name: "Outlet 21", Type: "iec-60320-c19", Label: "", PowerPort: "Power Port 1", FeedLeg: "C", MaximumDraw: 0, AllocatedDraw: 0 },
            { Name: "Outlet 22", Type: "iec-60320-c13", Label: "", PowerPort: "Power Port 1", FeedLeg: "C", MaximumDraw: 0, AllocatedDraw: 0 },
            { Name: "Outlet 23", Type: "iec-60320-c13", Label: "", PowerPort: "Power Port 1", FeedLeg: "C", MaximumDraw: 0, AllocatedDraw: 0 },
            { Name: "Outlet 24", Type: "iec-60320-c13", Label: "", PowerPort: "Power Port 1", FeedLeg: "C", MaximumDraw: 0, AllocatedDraw: 0 },
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
            { Name: "net", Label: "NET", Type: "1000base-t", MgmtOnly: true },
        },
    },
    "C2WG24SN-YCLN5D6 (Primary)": {
        Manufacturer: "Server Technology",
        Model: "C2WG24SN-YCLN5D6 (Primary)",
        Slug: "server-technology-c2wg24sn-ycln5d6",
        UHeight: 2,
        PartNumber: "C2WG24SN-YCLN5D6",
        IsFullDepth: true,
        Airflow: "passive",
        FrontImage: false,
        RearImage: false,
        SubdeviceRole: "",
        Weight: 0,
        WeightUnit: "",
        IsPowered: false,
        ConsolePorts: []ConsolePort{
            { Name: "ser", Type: "rj-45", Label: "SER", Poe: false },
            { Name: "th1", Type: "rj-11", Label: "T/H1", Poe: false },
            { Name: "th2", Type: "rj-11", Label: "T/H2", Poe: false },
        },
        ConsoleServerPorts: []ConsoleServerPort{
            { Name: "link", Type: "rj-12", Label: "LINK" },
        },
        PowerPorts: []PowerPort{
            { Name: "Power Port 1", Label: "", Type: "nema-l21-30p", MaximumDraw: 0, AllocatedDraw: 0 },
        },
        PowerOutlets: []PowerOutlet{
            { Name: "Outlet 1", Type: "iec-60320-c19", Label: "", PowerPort: "Power Port 1", FeedLeg: "A", MaximumDraw: 0, AllocatedDraw: 0 },
            { Name: "Outlet 2", Type: "iec-60320-c13", Label: "", PowerPort: "Power Port 1", FeedLeg: "A", MaximumDraw: 0, AllocatedDraw: 0 },
            { Name: "Outlet 3", Type: "iec-60320-c13", Label: "", PowerPort: "Power Port 1", FeedLeg: "A", MaximumDraw: 0, AllocatedDraw: 0 },
            { Name: "Outlet 4", Type: "iec-60320-c13", Label: "", PowerPort: "Power Port 1", FeedLeg: "A", MaximumDraw: 0, AllocatedDraw: 0 },
            { Name: "Outlet 5", Type: "iec-60320-c19", Label: "", PowerPort: "Power Port 1", FeedLeg: "A", MaximumDraw: 0, AllocatedDraw: 0 },
            { Name: "Outlet 6", Type: "iec-60320-c13", Label: "", PowerPort: "Power Port 1", FeedLeg: "A", MaximumDraw: 0, AllocatedDraw: 0 },
            { Name: "Outlet 7", Type: "iec-60320-c13", Label: "", PowerPort: "Power Port 1", FeedLeg: "A", MaximumDraw: 0, AllocatedDraw: 0 },
            { Name: "Outlet 8", Type: "iec-60320-c13", Label: "", PowerPort: "Power Port 1", FeedLeg: "A", MaximumDraw: 0, AllocatedDraw: 0 },
            { Name: "Outlet 9", Type: "iec-60320-c19", Label: "", PowerPort: "Power Port 1", FeedLeg: "B", MaximumDraw: 0, AllocatedDraw: 0 },
            { Name: "Outlet 10", Type: "iec-60320-c13", Label: "", PowerPort: "Power Port 1", FeedLeg: "B", MaximumDraw: 0, AllocatedDraw: 0 },
            { Name: "Outlet 11", Type: "iec-60320-c13", Label: "", PowerPort: "Power Port 1", FeedLeg: "B", MaximumDraw: 0, AllocatedDraw: 0 },
            { Name: "Outlet 12", Type: "iec-60320-c13", Label: "", PowerPort: "Power Port 1", FeedLeg: "B", MaximumDraw: 0, AllocatedDraw: 0 },
            { Name: "Outlet 13", Type: "iec-60320-c19", Label: "", PowerPort: "Power Port 1", FeedLeg: "B", MaximumDraw: 0, AllocatedDraw: 0 },
            { Name: "Outlet 14", Type: "iec-60320-c13", Label: "", PowerPort: "Power Port 1", FeedLeg: "B", MaximumDraw: 0, AllocatedDraw: 0 },
            { Name: "Outlet 15", Type: "iec-60320-c13", Label: "", PowerPort: "Power Port 1", FeedLeg: "B", MaximumDraw: 0, AllocatedDraw: 0 },
            { Name: "Outlet 16", Type: "iec-60320-c13", Label: "", PowerPort: "Power Port 1", FeedLeg: "B", MaximumDraw: 0, AllocatedDraw: 0 },
            { Name: "Outlet 17", Type: "iec-60320-c19", Label: "", PowerPort: "Power Port 1", FeedLeg: "C", MaximumDraw: 0, AllocatedDraw: 0 },
            { Name: "Outlet 18", Type: "iec-60320-c13", Label: "", PowerPort: "Power Port 1", FeedLeg: "C", MaximumDraw: 0, AllocatedDraw: 0 },
            { Name: "Outlet 19", Type: "iec-60320-c13", Label: "", PowerPort: "Power Port 1", FeedLeg: "C", MaximumDraw: 0, AllocatedDraw: 0 },
            { Name: "Outlet 20", Type: "iec-60320-c13", Label: "", PowerPort: "Power Port 1", FeedLeg: "C", MaximumDraw: 0, AllocatedDraw: 0 },
            { Name: "Outlet 21", Type: "iec-60320-c19", Label: "", PowerPort: "Power Port 1", FeedLeg: "C", MaximumDraw: 0, AllocatedDraw: 0 },
            { Name: "Outlet 22", Type: "iec-60320-c13", Label: "", PowerPort: "Power Port 1", FeedLeg: "C", MaximumDraw: 0, AllocatedDraw: 0 },
            { Name: "Outlet 23", Type: "iec-60320-c13", Label: "", PowerPort: "Power Port 1", FeedLeg: "C", MaximumDraw: 0, AllocatedDraw: 0 },
            { Name: "Outlet 24", Type: "iec-60320-c13", Label: "", PowerPort: "Power Port 1", FeedLeg: "C", MaximumDraw: 0, AllocatedDraw: 0 },
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
            { Name: "net", Label: "NET", Type: "1000base-t", MgmtOnly: true },
        },
    },
    "C2XG24SN-4PJN5D (Link)": {
        Manufacturer: "Server Technology",
        Model: "C2XG24SN-4PJN5D (Link)",
        Slug: "server-technology-c2xg24sn-4pjn5d",
        UHeight: 2,
        PartNumber: "C2XG24SN-4PJN5D",
        IsFullDepth: true,
        Airflow: "passive",
        FrontImage: false,
        RearImage: false,
        SubdeviceRole: "",
        Weight: 0,
        WeightUnit: "",
        IsPowered: false,
        ConsolePorts: []ConsolePort{
            { Name: "link", Type: "rj-12", Label: "LINK", Poe: true },
            { Name: "th1", Type: "rj-11", Label: "T/H1", Poe: false },
            { Name: "th2", Type: "rj-11", Label: "T/H2", Poe: false },
        },
        ConsoleServerPorts: []ConsoleServerPort{
        },
        PowerPorts: []PowerPort{
            { Name: "Power Port 1", Label: "", Type: "iec-60309-3p-n-e-6h", MaximumDraw: 0, AllocatedDraw: 0 },
        },
        PowerOutlets: []PowerOutlet{
            { Name: "Outlet 1", Type: "iec-60320-c19", Label: "", PowerPort: "Power Port 1", FeedLeg: "A", MaximumDraw: 0, AllocatedDraw: 0 },
            { Name: "Outlet 2", Type: "iec-60320-c13", Label: "", PowerPort: "Power Port 1", FeedLeg: "A", MaximumDraw: 0, AllocatedDraw: 0 },
            { Name: "Outlet 3", Type: "iec-60320-c13", Label: "", PowerPort: "Power Port 1", FeedLeg: "A", MaximumDraw: 0, AllocatedDraw: 0 },
            { Name: "Outlet 4", Type: "iec-60320-c13", Label: "", PowerPort: "Power Port 1", FeedLeg: "A", MaximumDraw: 0, AllocatedDraw: 0 },
            { Name: "Outlet 5", Type: "iec-60320-c19", Label: "", PowerPort: "Power Port 1", FeedLeg: "A", MaximumDraw: 0, AllocatedDraw: 0 },
            { Name: "Outlet 6", Type: "iec-60320-c13", Label: "", PowerPort: "Power Port 1", FeedLeg: "A", MaximumDraw: 0, AllocatedDraw: 0 },
            { Name: "Outlet 7", Type: "iec-60320-c13", Label: "", PowerPort: "Power Port 1", FeedLeg: "A", MaximumDraw: 0, AllocatedDraw: 0 },
            { Name: "Outlet 8", Type: "iec-60320-c13", Label: "", PowerPort: "Power Port 1", FeedLeg: "A", MaximumDraw: 0, AllocatedDraw: 0 },
            { Name: "Outlet 9", Type: "iec-60320-c19", Label: "", PowerPort: "Power Port 1", FeedLeg: "B", MaximumDraw: 0, AllocatedDraw: 0 },
            { Name: "Outlet 10", Type: "iec-60320-c13", Label: "", PowerPort: "Power Port 1", FeedLeg: "B", MaximumDraw: 0, AllocatedDraw: 0 },
            { Name: "Outlet 11", Type: "iec-60320-c13", Label: "", PowerPort: "Power Port 1", FeedLeg: "B", MaximumDraw: 0, AllocatedDraw: 0 },
            { Name: "Outlet 12", Type: "iec-60320-c13", Label: "", PowerPort: "Power Port 1", FeedLeg: "B", MaximumDraw: 0, AllocatedDraw: 0 },
            { Name: "Outlet 13", Type: "iec-60320-c19", Label: "", PowerPort: "Power Port 1", FeedLeg: "B", MaximumDraw: 0, AllocatedDraw: 0 },
            { Name: "Outlet 14", Type: "iec-60320-c13", Label: "", PowerPort: "Power Port 1", FeedLeg: "B", MaximumDraw: 0, AllocatedDraw: 0 },
            { Name: "Outlet 15", Type: "iec-60320-c13", Label: "", PowerPort: "Power Port 1", FeedLeg: "B", MaximumDraw: 0, AllocatedDraw: 0 },
            { Name: "Outlet 16", Type: "iec-60320-c13", Label: "", PowerPort: "Power Port 1", FeedLeg: "B", MaximumDraw: 0, AllocatedDraw: 0 },
            { Name: "Outlet 17", Type: "iec-60320-c19", Label: "", PowerPort: "Power Port 1", FeedLeg: "C", MaximumDraw: 0, AllocatedDraw: 0 },
            { Name: "Outlet 18", Type: "iec-60320-c13", Label: "", PowerPort: "Power Port 1", FeedLeg: "C", MaximumDraw: 0, AllocatedDraw: 0 },
            { Name: "Outlet 19", Type: "iec-60320-c13", Label: "", PowerPort: "Power Port 1", FeedLeg: "C", MaximumDraw: 0, AllocatedDraw: 0 },
            { Name: "Outlet 20", Type: "iec-60320-c13", Label: "", PowerPort: "Power Port 1", FeedLeg: "C", MaximumDraw: 0, AllocatedDraw: 0 },
            { Name: "Outlet 21", Type: "iec-60320-c19", Label: "", PowerPort: "Power Port 1", FeedLeg: "C", MaximumDraw: 0, AllocatedDraw: 0 },
            { Name: "Outlet 22", Type: "iec-60320-c13", Label: "", PowerPort: "Power Port 1", FeedLeg: "C", MaximumDraw: 0, AllocatedDraw: 0 },
            { Name: "Outlet 23", Type: "iec-60320-c13", Label: "", PowerPort: "Power Port 1", FeedLeg: "C", MaximumDraw: 0, AllocatedDraw: 0 },
            { Name: "Outlet 24", Type: "iec-60320-c13", Label: "", PowerPort: "Power Port 1", FeedLeg: "C", MaximumDraw: 0, AllocatedDraw: 0 },
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
    "C2XG24SN-YCLN5D6 (Link)": {
        Manufacturer: "Server Technology",
        Model: "C2XG24SN-YCLN5D6 (Link)",
        Slug: "server-technology-c2xg24sn-ycln5d6",
        UHeight: 2,
        PartNumber: "C2XG24SN-YCLN5D6",
        IsFullDepth: true,
        Airflow: "passive",
        FrontImage: false,
        RearImage: false,
        SubdeviceRole: "",
        Weight: 0,
        WeightUnit: "",
        IsPowered: false,
        ConsolePorts: []ConsolePort{
            { Name: "link", Type: "rj-12", Label: "LINK", Poe: true },
            { Name: "th1", Type: "rj-11", Label: "T/H1", Poe: false },
            { Name: "th2", Type: "rj-11", Label: "T/H2", Poe: false },
        },
        ConsoleServerPorts: []ConsoleServerPort{
        },
        PowerPorts: []PowerPort{
            { Name: "Power Port 1", Label: "", Type: "nema-l21-30p", MaximumDraw: 0, AllocatedDraw: 0 },
        },
        PowerOutlets: []PowerOutlet{
            { Name: "Outlet 1", Type: "iec-60320-c19", Label: "", PowerPort: "Power Port 1", FeedLeg: "A", MaximumDraw: 0, AllocatedDraw: 0 },
            { Name: "Outlet 2", Type: "iec-60320-c13", Label: "", PowerPort: "Power Port 1", FeedLeg: "A", MaximumDraw: 0, AllocatedDraw: 0 },
            { Name: "Outlet 3", Type: "iec-60320-c13", Label: "", PowerPort: "Power Port 1", FeedLeg: "A", MaximumDraw: 0, AllocatedDraw: 0 },
            { Name: "Outlet 4", Type: "iec-60320-c13", Label: "", PowerPort: "Power Port 1", FeedLeg: "A", MaximumDraw: 0, AllocatedDraw: 0 },
            { Name: "Outlet 5", Type: "iec-60320-c19", Label: "", PowerPort: "Power Port 1", FeedLeg: "A", MaximumDraw: 0, AllocatedDraw: 0 },
            { Name: "Outlet 6", Type: "iec-60320-c13", Label: "", PowerPort: "Power Port 1", FeedLeg: "A", MaximumDraw: 0, AllocatedDraw: 0 },
            { Name: "Outlet 7", Type: "iec-60320-c13", Label: "", PowerPort: "Power Port 1", FeedLeg: "A", MaximumDraw: 0, AllocatedDraw: 0 },
            { Name: "Outlet 8", Type: "iec-60320-c13", Label: "", PowerPort: "Power Port 1", FeedLeg: "A", MaximumDraw: 0, AllocatedDraw: 0 },
            { Name: "Outlet 9", Type: "iec-60320-c19", Label: "", PowerPort: "Power Port 1", FeedLeg: "B", MaximumDraw: 0, AllocatedDraw: 0 },
            { Name: "Outlet 10", Type: "iec-60320-c13", Label: "", PowerPort: "Power Port 1", FeedLeg: "B", MaximumDraw: 0, AllocatedDraw: 0 },
            { Name: "Outlet 11", Type: "iec-60320-c13", Label: "", PowerPort: "Power Port 1", FeedLeg: "B", MaximumDraw: 0, AllocatedDraw: 0 },
            { Name: "Outlet 12", Type: "iec-60320-c13", Label: "", PowerPort: "Power Port 1", FeedLeg: "B", MaximumDraw: 0, AllocatedDraw: 0 },
            { Name: "Outlet 13", Type: "iec-60320-c19", Label: "", PowerPort: "Power Port 1", FeedLeg: "B", MaximumDraw: 0, AllocatedDraw: 0 },
            { Name: "Outlet 14", Type: "iec-60320-c13", Label: "", PowerPort: "Power Port 1", FeedLeg: "B", MaximumDraw: 0, AllocatedDraw: 0 },
            { Name: "Outlet 15", Type: "iec-60320-c13", Label: "", PowerPort: "Power Port 1", FeedLeg: "B", MaximumDraw: 0, AllocatedDraw: 0 },
            { Name: "Outlet 16", Type: "iec-60320-c13", Label: "", PowerPort: "Power Port 1", FeedLeg: "B", MaximumDraw: 0, AllocatedDraw: 0 },
            { Name: "Outlet 17", Type: "iec-60320-c19", Label: "", PowerPort: "Power Port 1", FeedLeg: "C", MaximumDraw: 0, AllocatedDraw: 0 },
            { Name: "Outlet 18", Type: "iec-60320-c13", Label: "", PowerPort: "Power Port 1", FeedLeg: "C", MaximumDraw: 0, AllocatedDraw: 0 },
            { Name: "Outlet 19", Type: "iec-60320-c13", Label: "", PowerPort: "Power Port 1", FeedLeg: "C", MaximumDraw: 0, AllocatedDraw: 0 },
            { Name: "Outlet 20", Type: "iec-60320-c13", Label: "", PowerPort: "Power Port 1", FeedLeg: "C", MaximumDraw: 0, AllocatedDraw: 0 },
            { Name: "Outlet 21", Type: "iec-60320-c19", Label: "", PowerPort: "Power Port 1", FeedLeg: "C", MaximumDraw: 0, AllocatedDraw: 0 },
            { Name: "Outlet 22", Type: "iec-60320-c13", Label: "", PowerPort: "Power Port 1", FeedLeg: "C", MaximumDraw: 0, AllocatedDraw: 0 },
            { Name: "Outlet 23", Type: "iec-60320-c13", Label: "", PowerPort: "Power Port 1", FeedLeg: "C", MaximumDraw: 0, AllocatedDraw: 0 },
            { Name: "Outlet 24", Type: "iec-60320-c13", Label: "", PowerPort: "Power Port 1", FeedLeg: "C", MaximumDraw: 0, AllocatedDraw: 0 },
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
    "PRO3X-C3WG36RL-DQJE2NT2": {
        Manufacturer: "Server Technology",
        Model: "PRO3X-C3WG36RL-DQJE2NT2",
        Slug: "server-technology-pro3x-c3wg36rl-dqje2nt2",
        UHeight: 0,
        PartNumber: "PRO3X-C3WG36RL-DQJE2NT2",
        IsFullDepth: false,
        Airflow: "passive",
        FrontImage: false,
        RearImage: false,
        SubdeviceRole: "",
        Weight: 0,
        WeightUnit: "",
        IsPowered: false,
        ConsolePorts: []ConsolePort{
            { Name: "console/modem", Type: "rj-45", Label: "SER", Poe: false },
            { Name: "usb-a", Type: "usb-a", Label: "USB-A", Poe: false },
        },
        ConsoleServerPorts: []ConsoleServerPort{
            { Name: "aux", Type: "rj-12", Label: "AUX" },
        },
        PowerPorts: []PowerPort{
            { Name: "Power Port 1", Label: "", Type: "iec-60309-3p-e-9h", MaximumDraw: 0, AllocatedDraw: 0 },
        },
        PowerOutlets: []PowerOutlet{
            { Name: "Outlet 1", Type: "hdot-cx", Label: "", PowerPort: "Power Port 1", FeedLeg: "A", MaximumDraw: 0, AllocatedDraw: 0 },
            { Name: "Outlet 2", Type: "hdot-cx", Label: "", PowerPort: "Power Port 1", FeedLeg: "A", MaximumDraw: 0, AllocatedDraw: 0 },
            { Name: "Outlet 3", Type: "hdot-cx", Label: "", PowerPort: "Power Port 1", FeedLeg: "A", MaximumDraw: 0, AllocatedDraw: 0 },
            { Name: "Outlet 4", Type: "iec-60320-c13", Label: "", PowerPort: "Power Port 1", FeedLeg: "A", MaximumDraw: 0, AllocatedDraw: 0 },
            { Name: "Outlet 5", Type: "iec-60320-c19", Label: "", PowerPort: "Power Port 1", FeedLeg: "A", MaximumDraw: 0, AllocatedDraw: 0 },
            { Name: "Outlet 6", Type: "iec-60320-c13", Label: "", PowerPort: "Power Port 1", FeedLeg: "A", MaximumDraw: 0, AllocatedDraw: 0 },
            { Name: "Outlet 7", Type: "hdot-cx", Label: "", PowerPort: "Power Port 1", FeedLeg: "B", MaximumDraw: 0, AllocatedDraw: 0 },
            { Name: "Outlet 8", Type: "hdot-cx", Label: "", PowerPort: "Power Port 1", FeedLeg: "B", MaximumDraw: 0, AllocatedDraw: 0 },
            { Name: "Outlet 9", Type: "hdot-cx", Label: "", PowerPort: "Power Port 1", FeedLeg: "B", MaximumDraw: 0, AllocatedDraw: 0 },
            { Name: "Outlet 10", Type: "iec-60320-c13", Label: "", PowerPort: "Power Port 1", FeedLeg: "B", MaximumDraw: 0, AllocatedDraw: 0 },
            { Name: "Outlet 11", Type: "iec-60320-c13", Label: "", PowerPort: "Power Port 1", FeedLeg: "B", MaximumDraw: 0, AllocatedDraw: 0 },
            { Name: "Outlet 12", Type: "iec-60320-c13", Label: "", PowerPort: "Power Port 1", FeedLeg: "B", MaximumDraw: 0, AllocatedDraw: 0 },
            { Name: "Outlet 13", Type: "hdot-cx", Label: "", PowerPort: "Power Port 1", FeedLeg: "C", MaximumDraw: 0, AllocatedDraw: 0 },
            { Name: "Outlet 14", Type: "hdot-cx", Label: "", PowerPort: "Power Port 1", FeedLeg: "C", MaximumDraw: 0, AllocatedDraw: 0 },
            { Name: "Outlet 15", Type: "hdot-cx", Label: "", PowerPort: "Power Port 1", FeedLeg: "C", MaximumDraw: 0, AllocatedDraw: 0 },
            { Name: "Outlet 16", Type: "iec-60320-c13", Label: "", PowerPort: "Power Port 1", FeedLeg: "C", MaximumDraw: 0, AllocatedDraw: 0 },
            { Name: "Outlet 17", Type: "iec-60320-c19", Label: "", PowerPort: "Power Port 1", FeedLeg: "C", MaximumDraw: 0, AllocatedDraw: 0 },
            { Name: "Outlet 18", Type: "iec-60320-c13", Label: "", PowerPort: "Power Port 1", FeedLeg: "C", MaximumDraw: 0, AllocatedDraw: 0 },
            { Name: "Outlet 19", Type: "hdot-cx", Label: "", PowerPort: "Power Port 1", FeedLeg: "A", MaximumDraw: 0, AllocatedDraw: 0 },
            { Name: "Outlet 20", Type: "hdot-cx", Label: "", PowerPort: "Power Port 1", FeedLeg: "A", MaximumDraw: 0, AllocatedDraw: 0 },
            { Name: "Outlet 21", Type: "hdot-cx", Label: "", PowerPort: "Power Port 1", FeedLeg: "A", MaximumDraw: 0, AllocatedDraw: 0 },
            { Name: "Outlet 22", Type: "iec-60320-c13", Label: "", PowerPort: "Power Port 1", FeedLeg: "A", MaximumDraw: 0, AllocatedDraw: 0 },
            { Name: "Outlet 23", Type: "iec-60320-c13", Label: "", PowerPort: "Power Port 1", FeedLeg: "A", MaximumDraw: 0, AllocatedDraw: 0 },
            { Name: "Outlet 24", Type: "iec-60320-c13", Label: "", PowerPort: "Power Port 1", FeedLeg: "A", MaximumDraw: 0, AllocatedDraw: 0 },
            { Name: "Outlet 25", Type: "hdot-cx", Label: "", PowerPort: "Power Port 1", FeedLeg: "B", MaximumDraw: 0, AllocatedDraw: 0 },
            { Name: "Outlet 26", Type: "hdot-cx", Label: "", PowerPort: "Power Port 1", FeedLeg: "B", MaximumDraw: 0, AllocatedDraw: 0 },
            { Name: "Outlet 27", Type: "hdot-cx", Label: "", PowerPort: "Power Port 1", FeedLeg: "B", MaximumDraw: 0, AllocatedDraw: 0 },
            { Name: "Outlet 28", Type: "iec-60320-c13", Label: "", PowerPort: "Power Port 1", FeedLeg: "B", MaximumDraw: 0, AllocatedDraw: 0 },
            { Name: "Outlet 29", Type: "iec-60320-c13", Label: "", PowerPort: "Power Port 1", FeedLeg: "B", MaximumDraw: 0, AllocatedDraw: 0 },
            { Name: "Outlet 30", Type: "iec-60320-c13", Label: "", PowerPort: "Power Port 1", FeedLeg: "B", MaximumDraw: 0, AllocatedDraw: 0 },
            { Name: "Outlet 31", Type: "hdot-cx", Label: "", PowerPort: "Power Port 1", FeedLeg: "C", MaximumDraw: 0, AllocatedDraw: 0 },
            { Name: "Outlet 32", Type: "hdot-cx", Label: "", PowerPort: "Power Port 1", FeedLeg: "C", MaximumDraw: 0, AllocatedDraw: 0 },
            { Name: "Outlet 33", Type: "hdot-cx", Label: "", PowerPort: "Power Port 1", FeedLeg: "C", MaximumDraw: 0, AllocatedDraw: 0 },
            { Name: "Outlet 34", Type: "iec-60320-c13", Label: "", PowerPort: "Power Port 1", FeedLeg: "C", MaximumDraw: 0, AllocatedDraw: 0 },
            { Name: "Outlet 35", Type: "iec-60320-c13", Label: "", PowerPort: "Power Port 1", FeedLeg: "C", MaximumDraw: 0, AllocatedDraw: 0 },
            { Name: "Outlet 36", Type: "iec-60320-c13", Label: "", PowerPort: "Power Port 1", FeedLeg: "C", MaximumDraw: 0, AllocatedDraw: 0 },
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
            { Name: "eth1", Label: "ETH1", Type: "1000base-t", MgmtOnly: true },
            { Name: "eth2", Label: "ETH2", Type: "1000base-t", MgmtOnly: true },
        },
    },
}
