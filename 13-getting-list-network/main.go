package main

import (
	"fmt"
	"log"

	"github.com/google/gopacket/pcap"
)

// tcpdump -D
// tcpdump --list-interface

func main() {
  devices, err := pcap.FindAllDevs()
  if err != nil {
    log.Fatal(err)
  }

  fmt.Println("Devices found:")
  for _, device := range devices {
    fmt.Println("\nName:", device.Name)
    fmt.Println("Description:", device.Description)
    fmt.Println("Devices addresses:")
    for _, address := range device.Addresses {
      fmt.Println("- IP address:",address.IP )
      fmt.Println("- Subnet mask:", address.Netmask)
    }
  }
}
