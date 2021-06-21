package main

import (
	"fmt"
	"regexp"
)

func main() {
	line := "    Router Capability:  Router ID 0.0.0.0, Flags: 0x00"
	ipv4Reg := regexp.MustCompile(`Router Capability:\s+Router ID\s+([a-zA-Z0-9_.-]+),`).FindAllStringSubmatch(line, -1)
	if len(ipv4Reg) > 0 {
		fmt.Printf("ipv4 regs: %v\n", ipv4Reg)
	}
}
