package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
	"github.com/bondar-aleksandr/cisco_parser"
)

var (
	InfoLogger  *log.Logger = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger *log.Logger = log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
)

func main() {

	start := time.Now()
	InfoLogger.Println("Starting...")

	var firstFileName = flag.String("d1", "", "input configuration file to parse data from")
	var firstPlatform = flag.String("p1", "ios", "cisco OS family, possible values are 'ios', 'nxos'. Default is 'ios'")
	var secondFileName = flag.String("d2", "", "input configuration file to parse data from")
	var secondPlatform = flag.String("p2", "ios", "cisco OS family, possible values are 'ios', 'nxos'. Default is 'ios'")
	var action = flag.String("a", "parse", "action to perform. Possible values are 'parse', 'subnets'. Default is 'parse'" )
	var format = flag.String("f", "csv", "output format, possible values are 'csv', 'json'. Default is csv")

	if len(os.Args) < 2 {
		ErrorLogger.Fatalf("No input data provided, use -h flag for help. Exiting...")
	}
	flag.Parse()

	// read 1st device config
	iFile, err := os.Open(*firstFileName)
	if err != nil {
		ErrorLogger.Fatalf("Can not open file %q because of: %q", *firstFileName, err)
	}
	defer iFile.Close()

	// create 1st device
	device1, err := cisco_parser.NewDevice(iFile, *firstPlatform)
	if err != nil {
		ErrorLogger.Fatal("Cannot create 1st device:", err)
	}

	if *action == "parse" {
		// prepare output file

		oFileName := FileExtReplace(*firstFileName, *format)
		
		oFile, err := os.Create(oFileName)
		if err != nil {
			ErrorLogger.Fatalf("can't create output file: %q", err)
		}
		defer oFile.Close()
	
		// prepare serializer
		serializer, err := cisco_parser.NewSerializer(oFile, device1, *format)
		if err != nil {
			ErrorLogger.Fatal(err)
		}
		// serialize 
		if err = serializer.Serialize(); err != nil {
			ErrorLogger.Fatal(err)
		}
		InfoLogger.Printf("saved as %s\n", oFileName)

	} else if *action == "subnets" {
		// check whether all required flags provided
		if *secondFileName == "" {
			ErrorLogger.Fatal("not enough arguments provided for 2nd device!")
		}
		// read 2nd device config
		iFile, err := os.Open(*secondFileName)
		if err != nil {
			ErrorLogger.Fatalf("Can not open file %q because of: %q", *secondFileName, err)
		}
		defer iFile.Close()
	
		// create 2тв device
		device2, err := cisco_parser.NewDevice(iFile, *secondPlatform)
		if err != nil {
			ErrorLogger.Fatal("Cannot create 2nd device:", err)
		}
		GetCommonSubnets(device1, device2)
	} else {
		fmt.Println("Wrong action specified. Exiting...")
	}
	InfoLogger.Printf("Finished! Time taken: %s\n", time.Since(start))
}

// GetCommonSubnets finds subnets in common for 2 devices.
func GetCommonSubnets(d1, d2 *cisco_parser.Device) {
	dev1Subnets, err := d1.GetSubnets()
	if err != nil {
		ErrorLogger.Fatal(err)
	}
	var counter uint16
	for _, p := range dev1Subnets {
		d2Data, err := d2.GetSubnetData(p)
		if err != nil {
			ErrorLogger.Fatal("Cannot get 2nd device subnets data:", err)
		}
		// if no matching subnet found on d2
		if d2Data == nil {
			continue
		} else {
			counter++
			d1Data,_ := d1.GetSubnetData(p)		// get d1 subnet data as well
			fmt.Printf("\t%s\n", p)
			fmt.Println("--------")
			fmt.Printf("%q device interfaces:\n", d1.Hostname)
			fmt.Println(d1Data)
			fmt.Println("--------")
			fmt.Printf("%q device interfaces:\n", d2.Hostname)
			fmt.Println(d2Data)
			fmt.Println("--------")
		}
	}
	InfoLogger.Printf("Found %d common subnets\n", counter)
}

// utility function for file extension replacement.
func FileExtReplace(f string, ex string) string {
	bareName := strings.TrimSuffix(f, filepath.Ext(f))
	return fmt.Sprintf("%s.%s", bareName, ex)
}