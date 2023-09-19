# Cisco config parser CLI

This app is a CLI version of [https://github.com/bondar-aleksandr/cisco_parser](https://https://github.com/bondar-aleksandr/cisco_parser)
___
## Usage

Upon startup app parses cli arguments in order to determine cisco config-file location, output location, and output format. As a result, we got .csv file with device's interfaces, which may be used later for analysis (import to excell for example). Also it's possible to get .josn formatted output as well.

*Supported device types*: IOS/IOS-XE/IOS-XR/NXOS

CLI flags description is below:


| Flag | Data Type | Mandatory | Description | 
| ------ | ----------- | --- | --|
| -i | string | yes | Input cisco config-file location |
| -o | string | no | Output file location. Default is the same as input file, but with replaced extention |
| -t | string | no | OS type, possible values are "ios", "nxos". Default os "ios" |
| -j | -- | no | Whether we need additional json output. Default is "false" |
| -h | -- | no | Get help |

Launch app example:
```
main.exe -i ciscoConfig.cfg
```
In this case we will get the .csv file name "ciscoConfig.csv" in current working directory

___
## Output data format

Output file example is below:
```
Name,Description,Ip_addr,Subnet,Vrf,Mtu,ACLin,ACLout
Bundle-Ether10,BGP-PEERING,1.2.3.4/28,1.2.3.0/28,,,,
MgmtEth0/RSP0/CPU0/0,,10.0.0.1/24,10.0.0.0/24,management,,,
```

