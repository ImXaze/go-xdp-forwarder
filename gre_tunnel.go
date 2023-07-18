```go
package main

import (
	"net"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
)

type GRE struct {
	Tunnel *net.IPConn
	Source net.IP
	Destination net.IP
}

func CreateTunnel(config Config) (*GRE, error) {
	srcIP := net.ParseIP(config.BindAddress)
	dstIP := net.ParseIP(config.InternalAddress)

	rAddr := &net.IPAddr{IP: dstIP}
	lAddr := &net.IPAddr{IP: srcIP}

	conn, err := net.DialIP("ip4:gre", lAddr, rAddr)
	if err != nil {
		return nil, err
	}

	gre := &GRE{
		Tunnel: conn,
		Source: srcIP,
		Destination: dstIP,
	}

	return gre, nil
}

func (gre *GRE) ForwardPacket(packet gopacket.Packet) error {
	greLayer := &layers.GRE{
		Protocol: layers.EthernetTypeIPv4,
	}

	buffer := gopacket.NewSerializeBuffer()
	opts := gopacket.SerializeOptions{
		ComputeChecksums: true,
		FixLengths:       true,
	}

	err := gopacket.SerializeLayers(buffer, opts, greLayer, gopacket.Payload(packet.Data()))
	if err != nil {
		return err
	}

	_, err = gre.Tunnel.Write(buffer.Bytes())
	return err
}
```