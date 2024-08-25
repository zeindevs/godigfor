package main

import (
	"fmt"
	"log"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

var (
	// device             = "\\Device\\NPF_{D6000EB2-243D-41AC-9A30-95BE9B50CBA9}"
	// device             = "\\Device\\NPF_{98DFCE5D-D119-44C7-924F-52672CA77428}"
	device             = "\\Device\\NPF_{794B5CCE-32A7-4142-ADF6-DE9BF4410B88}"
	snapshotLen  int32 = 1024
	promisecuous       = false
	timeout            = 30 * time.Second

	err    error
	handle *pcap.Handle
)

func main() {
	handle, err = pcap.OpenLive(device, snapshotLen, promisecuous, timeout)
	if err != nil {
		log.Fatal(err)
	}
	defer handle.Close()

	fmt.Println("Start capturing")

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSource.Packets() {
		fmt.Println(packet)
	}
}
