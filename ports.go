```go
package main

import (
	"net"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
)

type Port struct {
	Source      uint16 `json:"source"`
	Destination uint16 `json:"destination"`
}

func (p *Port) ForwardPacket(packet gopacket.Packet) {
	udpLayer := packet.Layer(layers.LayerTypeUDP)
	if udpLayer != nil {
		udp, _ := udpLayer.(*layers.UDP)

		if udp.SrcPort == layers.UDPPort(p.Source) {
			udp.SrcPort = layers.UDPPort(p.Destination)
		}

		if udp.DstPort == layers.UDPPort(p.Source) {
			udp.DstPort = layers.UDPPort(p.Destination)
		}
	}
}

func GetPortsFromConfig(config Config) ([]Port, error) {
	ports := make([]Port, len(config.Ports))
	for i, portConfig := range config.Ports {
		ports[i] = Port{
			Source:      portConfig.Source,
			Destination: portConfig.Destination,
		}
	}
	return ports, nil
}

func ForwardTrafficToPorts(packet gopacket.Packet, ports []Port) {
	for _, port := range ports {
		port.ForwardPacket(packet)
	}
}
```