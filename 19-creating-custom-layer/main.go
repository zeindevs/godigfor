package main

import (
	"fmt"

	"github.com/google/gopacket"
)

type CustomLayer struct {
	SomeByte    byte
	AnotherByte byte
	restOfData  []byte
}

var CustomLayerType = gopacket.RegisterLayerType(
	2001,
	gopacket.LayerTypeMetadata{
		"CustomLayerType",
		gopacket.DecodeFunc(decodeCustomLayer),
	},
)

func (l CustomLayer) LayerType() gopacket.LayerType {
	return CustomLayerType
}

func (l CustomLayer) LayerContents() []byte {
	return []byte{l.SomeByte, l.AnotherByte}
}

func (l CustomLayer) LayerPayload() []byte {
	return l.restOfData
}

func decodeCustomLayer(data []byte, p gopacket.PacketBuilder) error {
	p.AddLayer(&CustomLayer{data[0], data[1], data[2:]})

	return p.NextDecoder(gopacket.DecodePayload)
}

func main() {
	rawBytes := []byte{0xF0, 0x0F, 65, 65, 66, 67, 68}
	packet := gopacket.NewPacket(rawBytes, CustomLayerType, gopacket.Default)

	fmt.Println("Created packet out of raw bytes.")
	fmt.Println(packet)

	customLayer := packet.Layer(CustomLayerType)
	if customLayer != nil {
		fmt.Println("Packet was successfully decoded.")
		customLayerContent, _ := customLayer.(*CustomLayer)
		fmt.Println("Payload:", customLayerContent.LayerPayload())
		fmt.Println("SomeByte element:", customLayerContent.SomeByte)
		fmt.Println("AnotherByte element:", customLayerContent.AnotherByte)
	}
}
