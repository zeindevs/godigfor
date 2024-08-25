package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	"github.com/google/gopacket/pcapgo"
)

// tcpdump -i eth0 -w my_capture.pcap

var (
	device            = "eth0"
	snapshotLen int32 = 1024
	promiseuous       = false
	timeout           = -1 * time.Second
	packetCount       = 0
	err         error
	handle      *pcap.Handle
)

func main() {
	f, _ := os.Create("test.pcap")
	w := pcapgo.NewWriter(f)
	w.WriteFileHeader(uint32(snapshotLen), layers.LinkTypeEthernet)
	defer f.Close()

	handle, err = pcap.OpenLive(device, snapshotLen, promiseuous, timeout)
	if err != nil {
		fmt.Printf("Error opening device %s: %v", device, err)
		log.Fatal(err)
	}
	defer handle.Close()

	packetCapture := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetCapture.Packets() {
		fmt.Println(packet)
		w.WritePacket(packet.Metadata().CaptureInfo, packet.Data())
		packetCount++

		if packetCount > 100 {
			break
		}
	}
}
