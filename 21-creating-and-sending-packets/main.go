package main

import (
	"log"
	"net"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)

var (
	device             = "eth0"
	snapshotLen  int32 = 1024
	promisecuous       = false
	timeout            = 30 * time.Second
	err          error
	handle       *pcap.Handle
	buffer       gopacket.SerializeBuffer
	options      gopacket.SerializeOptions
)

func main() {
	handle, err = pcap.OpenLive(device, snapshotLen, promisecuous, timeout)
	if err != nil {
		log.Fatal("Error opening device.", err.Error())
	}
	defer handle.Close()

	rawBytes := []byte{10, 20, 30}
	err = handle.WritePacketData(rawBytes)
	if err != nil {
		log.Fatal("Error writing bytes to network device.", err.Error())
	}

	buffer := gopacket.NewSerializeBuffer()
	gopacket.SerializeLayers(buffer, options,
		&layers.Ethernet{},
		&layers.IPv4{},
		&layers.TCP{},
		gopacket.Payload(rawBytes),
	)
	outgoingPacket := buffer.Bytes()
	err = handle.WritePacketData(outgoingPacket)
	if err != nil {
		log.Fatal("Error seding packet to network device.", err.Error())
	}

	ipLayer := &layers.IPv4{
		SrcIP: net.IP{127, 0, 0, 1},
		DstIP: net.IP{8, 8, 8, 8},
	}
	ethernetLayer := &layers.Ethernet{
		SrcMAC: net.HardwareAddr{0xFF, 0xAA, 0xFA, 0xAA, 0xFF, 0xAA},
		DstMAC: net.HardwareAddr{0xBD, 0xBD, 0xBD, 0xBD, 0xBD, 0xBD},
	}
	tcpLayer := &layers.TCP{
		SrcPort: layers.TCPPort(4321),
		DstPort: layers.TCPPort(80),
	}
	buffer = gopacket.NewSerializeBuffer()
	gopacket.SerializeLayers(buffer, options,
		ethernetLayer,
		ipLayer,
		tcpLayer,
		gopacket.Payload(rawBytes),
	)

	outgoingPacket = buffer.Bytes()
}
