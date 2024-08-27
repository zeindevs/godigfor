package main

import (
	"fmt"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
)

func main() {
	payload := []byte{2, 4, 6}
	options := gopacket.SerializeOptions{}
	buffer := gopacket.NewSerializeBuffer()
	gopacket.SerializeLayers(buffer, options,
		&layers.Ethernet{},
		&layers.IPv4{},
		&layers.TCP{},
		gopacket.Payload(payload),
	)
	rawBytes := buffer.Bytes()

	ethPacket := gopacket.NewPacket(rawBytes, layers.LayerTypeEthernet, gopacket.Default)

	ipPacket := gopacket.NewPacket(rawBytes, layers.LayerTypeIPv4, gopacket.Lazy)

	tcpPacket := gopacket.NewPacket(rawBytes, layers.LayerTypeTCP, gopacket.NoCopy)

	fmt.Println(ethPacket)
	fmt.Println(ipPacket)
	fmt.Println(tcpPacket)
}
