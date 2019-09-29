package main

import (
	"flag"
	"fmt"
)

func isFlagSet(fl flag.Flag) bool {
	isSet := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == fl.Name {
			isSet = true
		}
	})
	return isSet
}

func validateHost(host *string) bool {
	return true
}

func validatePortsRange(ports *string) bool {
	return true
}

func main() {
	helpFlag := flag.Bool("help", false, "Print this message")
	hostFlag := flag.String("host", "", "IP address of machine to check ports")
	portsRangeFlag := flag.String("ports", "0-65535", "Range of ports to scan")
	flag.Parse()

	if isFlagSet(helpFlag) {
		fmt.Printf("HELP MESSAGE")
		return
	}

	fmt.Println("hostFlag:", *hostFlag)
	fmt.Println("hostFlag:", *hostFlag)
	fmt.Println("portsRangeFlag:", *portsRangeFlag)
}