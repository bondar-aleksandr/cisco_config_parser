# Cisco config parser CLI

This app is aimed to parse Cisco devices config into structured format (csv or json).
Supported device types: IOS/IOS-XE/IOS-XR/NXOS.
We can further convert .csv result file to excel file for convenient analysis.
It's also possible to search for common subnets between two devices, might be useful for drawing logical network topology.
App uses [https://github.com/bondar-aleksandr/cisco_parser](https://github.com/bondar-aleksandr/cisco_parser) as a library.
___
## Usage

CLI flags description is below:


| Flag | Data Type | Mandatory | Description | 
| ------ | ----------- | --- | --|
| -d1 | string | yes | cisco config-file location (1st device) |
| -p1 | string | yes | 1st device platform (**ios/nxos**) |
| -d2 | string | no | cisco config-file location (2nd device) |
| -p2 | string | no | 2nd device platform (**ios/nxos**) |
| -a | string | no | action to perform. Possible values are **parse, subnets**. Default is **parse** |
| -f | string | no | output format. Default is **csv**. Possible values are **csv, json** |
| -h | -- | no | get help |

Upon startup app parses cli arguments in order to determine cisco config-file location(s), output data format, and action needed to be performed.

### Prasing

Currently, the following interface values are parsed:
- name
- description
- encapsulation
- ip address
- ip subnet
- vrf
- mtu
- input ACL
- output ACL

Device hostname parsed as well

Launch app for parsing:
```
app.exe -d1 ciscoConfig.cfg -p1 ios -f csv -a parse
```
In this case we will get the .csv file name "ciscoConfig.csv" in current working directory

Let's suppose we have the following interface configuration in config file:
```
!
interface GigabitEthernet0/0/2.2
 encapsulation dot1Q 2
 description TUNNEL-SOURCE_INET
 ip vrf forwarding INET
 ip address 1.2.3.4 255.255.255.224
 no ip redirects
 no ip proxy-arp
 ip access-group FROM_INET_IPSEC in
 negotiation auto
!
```
Output file in this case will look like:
```
Name,Description,Encapsulation,Ip_addr,Subnet,Vrf,Mtu,ACLin,ACLout
GigabitEthernet0/0/2,TUNNEL-SOURCE_INET,dot1q 2,1.2.3.4/27,1.2.3.0/27,INET,,FROM_INET_IPSEC,
```

### Common subnets search

Launch app for parsing:
```
app.exe -d1 ciscoConfig01.cfg -p1 ios -d2 ciscoConfig02.cfg -p2 nxos -a subnets
```
App output will look like
```
        10.10.5.0/30
--------
"ciscoConfig01" device interfaces:
interface: "Vlan20", vrf: ""

--------
"ciscoConfig02" device interfaces:
interface: "Vlan20", vrf: "WAN"

--------
INFO: 2024/01/31 17:58:43 main.go:123: Found 1 common subnets
```