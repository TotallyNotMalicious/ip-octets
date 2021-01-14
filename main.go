package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var address string

	fmt.Print("IP to convert: ")
	fmt.Scan(&address)

	octets := strings.Split(address, ".")

	octet0, _ := strconv.Atoi(octets[0])
	octet1, _ := strconv.Atoi(octets[1])
	octet2, _ := strconv.Atoi(octets[2])
	octet3, _ := strconv.Atoi(octets[3])

	complete := [4]byte{byte(octet0), byte(octet1), byte(octet2), byte(octet3)}

	fmt.Printf("%b\n", complete)
}
