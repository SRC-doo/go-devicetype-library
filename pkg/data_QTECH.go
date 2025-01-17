// Code generated by go generate; DO NOT EDIT.
package devices

var DeviceTypesMapqtech = map[string]*DeviceData{
    "QSW-6900-32H": {
        Manufacturer: "QTECH",
        Model: "QSW-6900-32H",
        Slug: "qtech-qsw-6900-32h",
        UHeight: 1,
        PartNumber: "QSW-6900-32H",
        IsFullDepth: true,
        Airflow: "",
        FrontImage: false,
        RearImage: false,
        SubdeviceRole: "",
        Weight: 0,
        WeightUnit: "",
        IsPowered: false,
        ConsolePorts: []ConsolePort{
            { Name: "console", Type: "rj-45", Label: "Console", Poe: false },
        },
        ConsoleServerPorts: []ConsoleServerPort{
        },
        PowerPorts: []PowerPort{
            { Name: "PSU1", Label: "", Type: "iec-60320-c14", MaximumDraw: 550, AllocatedDraw: 300 },
            { Name: "PSU2", Label: "", Type: "iec-60320-c14", MaximumDraw: 550, AllocatedDraw: 300 },
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
            { Name: "HundredGigabitEthernet 0/1", Label: "", Type: "100gbase-x-qsfp28", MgmtOnly: false },
            { Name: "HundredGigabitEthernet 0/2", Label: "", Type: "100gbase-x-qsfp28", MgmtOnly: false },
            { Name: "HundredGigabitEthernet 0/3", Label: "", Type: "100gbase-x-qsfp28", MgmtOnly: false },
            { Name: "HundredGigabitEthernet 0/4", Label: "", Type: "100gbase-x-qsfp28", MgmtOnly: false },
            { Name: "HundredGigabitEthernet 0/5", Label: "", Type: "100gbase-x-qsfp28", MgmtOnly: false },
            { Name: "HundredGigabitEthernet 0/6", Label: "", Type: "100gbase-x-qsfp28", MgmtOnly: false },
            { Name: "HundredGigabitEthernet 0/7", Label: "", Type: "100gbase-x-qsfp28", MgmtOnly: false },
            { Name: "HundredGigabitEthernet 0/8", Label: "", Type: "100gbase-x-qsfp28", MgmtOnly: false },
            { Name: "HundredGigabitEthernet 0/9", Label: "", Type: "100gbase-x-qsfp28", MgmtOnly: false },
            { Name: "HundredGigabitEthernet 0/10", Label: "", Type: "100gbase-x-qsfp28", MgmtOnly: false },
            { Name: "HundredGigabitEthernet 0/11", Label: "", Type: "100gbase-x-qsfp28", MgmtOnly: false },
            { Name: "HundredGigabitEthernet 0/12", Label: "", Type: "100gbase-x-qsfp28", MgmtOnly: false },
            { Name: "HundredGigabitEthernet 0/13", Label: "", Type: "100gbase-x-qsfp28", MgmtOnly: false },
            { Name: "HundredGigabitEthernet 0/14", Label: "", Type: "100gbase-x-qsfp28", MgmtOnly: false },
            { Name: "HundredGigabitEthernet 0/15", Label: "", Type: "100gbase-x-qsfp28", MgmtOnly: false },
            { Name: "HundredGigabitEthernet 0/16", Label: "", Type: "100gbase-x-qsfp28", MgmtOnly: false },
            { Name: "HundredGigabitEthernet 0/17", Label: "", Type: "100gbase-x-qsfp28", MgmtOnly: false },
            { Name: "HundredGigabitEthernet 0/18", Label: "", Type: "100gbase-x-qsfp28", MgmtOnly: false },
            { Name: "HundredGigabitEthernet 0/19", Label: "", Type: "100gbase-x-qsfp28", MgmtOnly: false },
            { Name: "HundredGigabitEthernet 0/20", Label: "", Type: "100gbase-x-qsfp28", MgmtOnly: false },
            { Name: "HundredGigabitEthernet 0/21", Label: "", Type: "100gbase-x-qsfp28", MgmtOnly: false },
            { Name: "HundredGigabitEthernet 0/22", Label: "", Type: "100gbase-x-qsfp28", MgmtOnly: false },
            { Name: "HundredGigabitEthernet 0/23", Label: "", Type: "100gbase-x-qsfp28", MgmtOnly: false },
            { Name: "HundredGigabitEthernet 0/24", Label: "", Type: "100gbase-x-qsfp28", MgmtOnly: false },
            { Name: "HundredGigabitEthernet 0/25", Label: "", Type: "100gbase-x-qsfp28", MgmtOnly: false },
            { Name: "HundredGigabitEthernet 0/26", Label: "", Type: "100gbase-x-qsfp28", MgmtOnly: false },
            { Name: "HundredGigabitEthernet 0/27", Label: "", Type: "100gbase-x-qsfp28", MgmtOnly: false },
            { Name: "HundredGigabitEthernet 0/28", Label: "", Type: "100gbase-x-qsfp28", MgmtOnly: false },
            { Name: "HundredGigabitEthernet 0/29", Label: "", Type: "100gbase-x-qsfp28", MgmtOnly: false },
            { Name: "HundredGigabitEthernet 0/30", Label: "", Type: "100gbase-x-qsfp28", MgmtOnly: false },
            { Name: "HundredGigabitEthernet 0/31", Label: "", Type: "100gbase-x-qsfp28", MgmtOnly: false },
            { Name: "HundredGigabitEthernet 0/32", Label: "", Type: "100gbase-x-qsfp28", MgmtOnly: false },
            { Name: "mgmt 0", Label: "MGMT", Type: "1000base-t", MgmtOnly: true },
        },
    },
    "QSW-6900-56F": {
        Manufacturer: "QTECH",
        Model: "QSW-6900-56F",
        Slug: "qtech-qsw-6900-56f",
        UHeight: 1,
        PartNumber: "QSW-6900-56F",
        IsFullDepth: false,
        Airflow: "front-to-rear",
        FrontImage: false,
        RearImage: false,
        SubdeviceRole: "",
        Weight: 8,
        WeightUnit: "",
        IsPowered: false,
        ConsolePorts: []ConsolePort{
            { Name: "console", Type: "rj-45", Label: "Console", Poe: false },
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
            { Name: "Power 1", Label: "PWR1", Position: "1" },
            { Name: "Power 2", Label: "PWR2", Position: "2" },
        },
			  DeviceBays: []DeviceBay{
        },
        InventoryItems: []InventoryItem{
        },
        Interfaces: []Interface{
            { Name: "HundredGigabitEthernet 0/49", Label: "49F", Type: "100gbase-x-qsfp28", MgmtOnly: false },
            { Name: "HundredGigabitEthernet 0/50", Label: "50F", Type: "100gbase-x-qsfp28", MgmtOnly: false },
            { Name: "HundredGigabitEthernet 0/51", Label: "51F", Type: "100gbase-x-qsfp28", MgmtOnly: false },
            { Name: "HundredGigabitEthernet 0/52", Label: "52F", Type: "100gbase-x-qsfp28", MgmtOnly: false },
            { Name: "HundredGigabitEthernet 0/53", Label: "53F", Type: "100gbase-x-qsfp28", MgmtOnly: false },
            { Name: "HundredGigabitEthernet 0/54", Label: "54F", Type: "100gbase-x-qsfp28", MgmtOnly: false },
            { Name: "HundredGigabitEthernet 0/55", Label: "55F", Type: "100gbase-x-qsfp28", MgmtOnly: false },
            { Name: "HundredGigabitEthernet 0/56", Label: "56F", Type: "100gbase-x-qsfp28", MgmtOnly: false },
            { Name: "TFGigabitEthernet 0/1", Label: "1F", Type: "25gbase-x-sfp28", MgmtOnly: false },
            { Name: "TFGigabitEthernet 0/2", Label: "2F", Type: "25gbase-x-sfp28", MgmtOnly: false },
            { Name: "TFGigabitEthernet 0/3", Label: "3F", Type: "25gbase-x-sfp28", MgmtOnly: false },
            { Name: "TFGigabitEthernet 0/4", Label: "4F", Type: "25gbase-x-sfp28", MgmtOnly: false },
            { Name: "TFGigabitEthernet 0/5", Label: "5F", Type: "25gbase-x-sfp28", MgmtOnly: false },
            { Name: "TFGigabitEthernet 0/6", Label: "6F", Type: "25gbase-x-sfp28", MgmtOnly: false },
            { Name: "TFGigabitEthernet 0/7", Label: "7F", Type: "25gbase-x-sfp28", MgmtOnly: false },
            { Name: "TFGigabitEthernet 0/8", Label: "8F", Type: "25gbase-x-sfp28", MgmtOnly: false },
            { Name: "TFGigabitEthernet 0/9", Label: "9F", Type: "25gbase-x-sfp28", MgmtOnly: false },
            { Name: "TFGigabitEthernet 0/10", Label: "10F", Type: "25gbase-x-sfp28", MgmtOnly: false },
            { Name: "TFGigabitEthernet 0/11", Label: "11F", Type: "25gbase-x-sfp28", MgmtOnly: false },
            { Name: "TFGigabitEthernet 0/12", Label: "12F", Type: "25gbase-x-sfp28", MgmtOnly: false },
            { Name: "TFGigabitEthernet 0/13", Label: "13F", Type: "25gbase-x-sfp28", MgmtOnly: false },
            { Name: "TFGigabitEthernet 0/14", Label: "14F", Type: "25gbase-x-sfp28", MgmtOnly: false },
            { Name: "TFGigabitEthernet 0/15", Label: "15F", Type: "25gbase-x-sfp28", MgmtOnly: false },
            { Name: "TFGigabitEthernet 0/16", Label: "16F", Type: "25gbase-x-sfp28", MgmtOnly: false },
            { Name: "TFGigabitEthernet 0/17", Label: "17F", Type: "25gbase-x-sfp28", MgmtOnly: false },
            { Name: "TFGigabitEthernet 0/18", Label: "18F", Type: "25gbase-x-sfp28", MgmtOnly: false },
            { Name: "TFGigabitEthernet 0/19", Label: "19F", Type: "25gbase-x-sfp28", MgmtOnly: false },
            { Name: "TFGigabitEthernet 0/20", Label: "20F", Type: "25gbase-x-sfp28", MgmtOnly: false },
            { Name: "TFGigabitEthernet 0/21", Label: "21F", Type: "25gbase-x-sfp28", MgmtOnly: false },
            { Name: "TFGigabitEthernet 0/22", Label: "22F", Type: "25gbase-x-sfp28", MgmtOnly: false },
            { Name: "TFGigabitEthernet 0/23", Label: "23F", Type: "25gbase-x-sfp28", MgmtOnly: false },
            { Name: "TFGigabitEthernet 0/24", Label: "24F", Type: "25gbase-x-sfp28", MgmtOnly: false },
            { Name: "TFGigabitEthernet 0/25", Label: "25F", Type: "25gbase-x-sfp28", MgmtOnly: false },
            { Name: "TFGigabitEthernet 0/26", Label: "26F", Type: "25gbase-x-sfp28", MgmtOnly: false },
            { Name: "TFGigabitEthernet 0/27", Label: "27F", Type: "25gbase-x-sfp28", MgmtOnly: false },
            { Name: "TFGigabitEthernet 0/28", Label: "28F", Type: "25gbase-x-sfp28", MgmtOnly: false },
            { Name: "TFGigabitEthernet 0/29", Label: "29F", Type: "25gbase-x-sfp28", MgmtOnly: false },
            { Name: "TFGigabitEthernet 0/30", Label: "30F", Type: "25gbase-x-sfp28", MgmtOnly: false },
            { Name: "TFGigabitEthernet 0/31", Label: "31F", Type: "25gbase-x-sfp28", MgmtOnly: false },
            { Name: "TFGigabitEthernet 0/32", Label: "32F", Type: "25gbase-x-sfp28", MgmtOnly: false },
            { Name: "TFGigabitEthernet 0/33", Label: "33F", Type: "25gbase-x-sfp28", MgmtOnly: false },
            { Name: "TFGigabitEthernet 0/34", Label: "34F", Type: "25gbase-x-sfp28", MgmtOnly: false },
            { Name: "TFGigabitEthernet 0/35", Label: "35F", Type: "25gbase-x-sfp28", MgmtOnly: false },
            { Name: "TFGigabitEthernet 0/36", Label: "36F", Type: "25gbase-x-sfp28", MgmtOnly: false },
            { Name: "TFGigabitEthernet 0/37", Label: "37F", Type: "25gbase-x-sfp28", MgmtOnly: false },
            { Name: "TFGigabitEthernet 0/38", Label: "38F", Type: "25gbase-x-sfp28", MgmtOnly: false },
            { Name: "TFGigabitEthernet 0/39", Label: "39F", Type: "25gbase-x-sfp28", MgmtOnly: false },
            { Name: "TFGigabitEthernet 0/40", Label: "40F", Type: "25gbase-x-sfp28", MgmtOnly: false },
            { Name: "TFGigabitEthernet 0/41", Label: "41F", Type: "25gbase-x-sfp28", MgmtOnly: false },
            { Name: "TFGigabitEthernet 0/42", Label: "42F", Type: "25gbase-x-sfp28", MgmtOnly: false },
            { Name: "TFGigabitEthernet 0/43", Label: "43F", Type: "25gbase-x-sfp28", MgmtOnly: false },
            { Name: "TFGigabitEthernet 0/44", Label: "44F", Type: "25gbase-x-sfp28", MgmtOnly: false },
            { Name: "TFGigabitEthernet 0/45", Label: "45F", Type: "25gbase-x-sfp28", MgmtOnly: false },
            { Name: "TFGigabitEthernet 0/46", Label: "46F", Type: "25gbase-x-sfp28", MgmtOnly: false },
            { Name: "TFGigabitEthernet 0/47", Label: "47F", Type: "25gbase-x-sfp28", MgmtOnly: false },
            { Name: "TFGigabitEthernet 0/48", Label: "48F", Type: "25gbase-x-sfp28", MgmtOnly: false },
            { Name: "mgmt 0", Label: "MGMT", Type: "1000base-t", MgmtOnly: true },
        },
    },
    "QSW-6900-56LF": {
        Manufacturer: "QTECH",
        Model: "QSW-6900-56LF",
        Slug: "qtech-qsw-6900-56lf",
        UHeight: 1,
        PartNumber: "QSW-6900-56LF",
        IsFullDepth: false,
        Airflow: "front-to-rear",
        FrontImage: false,
        RearImage: false,
        SubdeviceRole: "",
        Weight: 8,
        WeightUnit: "",
        IsPowered: false,
        ConsolePorts: []ConsolePort{
            { Name: "console", Type: "rj-45", Label: "Console", Poe: false },
        },
        ConsoleServerPorts: []ConsoleServerPort{
        },
        PowerPorts: []PowerPort{
            { Name: "PSU1", Label: "", Type: "iec-60320-c14", MaximumDraw: 550, AllocatedDraw: 300 },
            { Name: "PSU2", Label: "", Type: "iec-60320-c14", MaximumDraw: 550, AllocatedDraw: 300 },
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
            { Name: "TenGigabitEthernet 0/1", Label: "1F", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "TenGigabitEthernet 0/2", Label: "2F", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "TenGigabitEthernet 0/3", Label: "3F", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "TenGigabitEthernet 0/4", Label: "4F", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "TenGigabitEthernet 0/5", Label: "5F", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "TenGigabitEthernet 0/6", Label: "6F", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "TenGigabitEthernet 0/7", Label: "7F", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "TenGigabitEthernet 0/8", Label: "8F", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "TenGigabitEthernet 0/9", Label: "9F", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "TenGigabitEthernet 0/10", Label: "10F", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "TenGigabitEthernet 0/11", Label: "11F", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "TenGigabitEthernet 0/12", Label: "12F", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "TenGigabitEthernet 0/13", Label: "13F", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "TenGigabitEthernet 0/14", Label: "14F", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "TenGigabitEthernet 0/15", Label: "15F", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "TenGigabitEthernet 0/16", Label: "16F", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "TenGigabitEthernet 0/17", Label: "17F", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "TenGigabitEthernet 0/18", Label: "18F", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "TenGigabitEthernet 0/19", Label: "19F", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "TenGigabitEthernet 0/20", Label: "20F", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "TenGigabitEthernet 0/21", Label: "21F", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "TenGigabitEthernet 0/22", Label: "22F", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "TenGigabitEthernet 0/23", Label: "23F", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "TenGigabitEthernet 0/24", Label: "24F", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "TenGigabitEthernet 0/25", Label: "25F", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "TenGigabitEthernet 0/26", Label: "26F", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "TenGigabitEthernet 0/27", Label: "27F", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "TenGigabitEthernet 0/28", Label: "28F", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "TenGigabitEthernet 0/29", Label: "29F", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "TenGigabitEthernet 0/30", Label: "30F", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "TenGigabitEthernet 0/31", Label: "31F", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "TenGigabitEthernet 0/32", Label: "32F", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "TenGigabitEthernet 0/33", Label: "33F", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "TenGigabitEthernet 0/34", Label: "34F", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "TenGigabitEthernet 0/35", Label: "35F", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "TenGigabitEthernet 0/36", Label: "36F", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "TenGigabitEthernet 0/37", Label: "37F", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "TenGigabitEthernet 0/38", Label: "38F", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "TenGigabitEthernet 0/39", Label: "39F", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "TenGigabitEthernet 0/40", Label: "40F", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "TenGigabitEthernet 0/41", Label: "41F", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "TenGigabitEthernet 0/42", Label: "42F", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "TenGigabitEthernet 0/43", Label: "43F", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "TenGigabitEthernet 0/44", Label: "44F", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "TenGigabitEthernet 0/45", Label: "45F", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "TenGigabitEthernet 0/46", Label: "46F", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "TenGigabitEthernet 0/47", Label: "47F", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "TenGigabitEthernet 0/48", Label: "48F", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "HundredGigabitEthernet 0/49", Label: "49F", Type: "100gbase-x-qsfp28", MgmtOnly: false },
            { Name: "HundredGigabitEthernet 0/50", Label: "50F", Type: "100gbase-x-qsfp28", MgmtOnly: false },
            { Name: "HundredGigabitEthernet 0/51", Label: "51F", Type: "100gbase-x-qsfp28", MgmtOnly: false },
            { Name: "HundredGigabitEthernet 0/52", Label: "52F", Type: "100gbase-x-qsfp28", MgmtOnly: false },
            { Name: "HundredGigabitEthernet 0/53", Label: "53F", Type: "100gbase-x-qsfp28", MgmtOnly: false },
            { Name: "HundredGigabitEthernet 0/54", Label: "54F", Type: "100gbase-x-qsfp28", MgmtOnly: false },
            { Name: "HundredGigabitEthernet 0/55", Label: "55F", Type: "100gbase-x-qsfp28", MgmtOnly: false },
            { Name: "HundredGigabitEthernet 0/56", Label: "56F", Type: "100gbase-x-qsfp28", MgmtOnly: false },
            { Name: "mgmt 0", Label: "MGMT", Type: "1000base-t", MgmtOnly: true },
        },
    },
}
