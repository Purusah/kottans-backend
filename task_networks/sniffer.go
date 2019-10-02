package main

import (
	"flag"
	"fmt"
	"net"
	"regexp"

	"message"
)

func isFlagSet(fl *flag.Flag) bool {
	isSet := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == fl.Name {
			isSet = true
		}
	})
	return isSet
}

func validateHost(host *string) bool {
	// TODO Validate IPv6
	var regexIPAddr string = `^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}(25[0-5]|2[0-4][0-9]|1[0-9]{2}|[1-9][0-9]|[0-9])$`

	if *host == "localhost" {
		return true
	} else if matched, _ := regexp.Match(regexIPAddr, []byte(*host)); matched == true {
		return true
	} else {
		return false
	}
}

func validatePortsRange(ports *string) bool {
	// TODO Validate Ports Range
	return true
}

func checkConnection(host, port string) bool {
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%s", host, port))
	if err != nil {
		fmt.Println(err)
	}
	return true
}

func main() {
	help := flag.Bool("help", false, "Print this message")
	host := flag.String("host", "", "IP address of machine to check ports")
	portsRange := flag.String("ports", "0-65535", "Range of ports to scan")
	_ = help

	helpFlag := flag.Lookup("help")
	hostFlag := flag.Lookup("host")

	flag.Parse()

	// Optional help flag
	if isFlagSet(helpFlag) {
		fmt.Printf(message.HelpMessage)
		return
	}
	// Required host flag
	if !isFlagSet(hostFlag) {
		fmt.Printf(message.ErrorHostRequiredMessage)
		return
	}

	var isHostValid bool = validateHost(host)
	if !isHostValid {
		fmt.Printf(message.ErrorHostRequiredMessage)
		return
	}

	var startPort, lastPort int
	err := parsePorts(*portsRange, &startPort, &lastPort)
	if err != nil {
		fmt.Printf(message.ErrorPortRangeInvalid)
		return
	}

	// for ports
	for _, port := range ports {
		checkConnection(*host, port)
	}

	_ = isHostValid
	_ = host
	_ = portsRange
	// fmt.Println("helpFlag:", *help)
	// fmt.Println("hostFlag:", *host)
	// fmt.Println("helpFlag:", *helpFlag)
	// fmt.Println("portsRangeFlag:", *portsRange)
}
