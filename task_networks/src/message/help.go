package message

const HelpMessage = `This tool for checking opened ports

Example: go run sniffer.go --host 127.0.0.1 --ports 10-1000

Usage:
	go run sniffer.go --host <address> [--ports <ports-range>]
Arguments:
	--help		See this message (optional)
	--host		IP address of server to check ports (required)
	--ports		Ports range to check (optional). Min value 0, max value 65535. Separated by dash.
`
