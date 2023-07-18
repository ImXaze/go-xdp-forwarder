```go
package main

import (
	"net"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
)

type Forwarder struct {
	SourcePort      int
	DestinationPort int
	BindAddress     net.IP
	InternalAddress net.IP
}

func CreateForwarder(config Forwarder) (*Forwarder, error) {
	return &config, nil
}

func (f *Forwarder) ForwardPacket(packet gopacket.Packet) {
	udpLayer := packet.Layer(layers.LayerTypeUDP)
	if udpLayer != nil {
		udp, _ := udpLayer.(*layers.UDP)

		if int(udp.SrcPort) == f.SourcePort {
			udp.SrcPort = layers.UDPPort(f.DestinationPort)
			udp.DstPort = layers.UDPPort(f.DestinationPort)
		}
	}

	ipLayer := packet.Layer(layers.LayerTypeIPv4)
	if ipLayer != nil {
		ip, _ := ipLayer.(*layers.IPv4)

		if ip.SrcIP.Equal(f.BindAddress) {
			ip.SrcIP = f.InternalAddress
			ip.DstIP = f.InternalAddress
		}
	}
}
```