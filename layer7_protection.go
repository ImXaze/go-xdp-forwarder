```go
package main

import (
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"net"
)

func ProtectLayer7(packet gopacket.Packet, config Config) bool {
	applicationLayer := packet.ApplicationLayer()
	if applicationLayer != nil {
		for _, rule := range config.Layer7Rules {
			if rule.Protocol == "http" {
				http := layers.HTTP{}
				err := http.DecodeFromBytes(applicationLayer.Payload(), gopacket.NilDecodeFeedback)
				if err != nil {
					continue
				}
				if http.RequestMethod == rule.Method && http.RequestURI == rule.URI {
					return false
				}
			}
		}
	}
	return true
}

type Config struct {
	Layer7Rules []Layer7Rule `json:"layer7_rules"`
}

type Layer7Rule struct {
	Protocol string `json:"protocol"`
	Method   string `json:"method"`
	URI      string `json:"uri"`
}
```