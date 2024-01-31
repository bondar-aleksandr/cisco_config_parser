module github.com/bondar-aleksandr/cisco_config_parser

go 1.20

replace github.com/bondar-aleksandr/cisco_parser => ../cisco_parser_pkg

require github.com/bondar-aleksandr/cisco_parser v0.0.0-20240125173609-2c7d77cdae80

require (
	github.com/sirupsen/logrus v1.9.3 // indirect
	golang.org/x/sys v0.16.0 // indirect
)
