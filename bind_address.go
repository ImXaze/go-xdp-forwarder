```go
package main

import (
	"net"
	"fmt"
	"os"
)

type BindAddress struct {
	Address string `json:"address"`
}

func (ba *BindAddress) Bind() error {
	ln, err := net.Listen("tcp", ba.Address)
	if err != nil {
		fmt.Println("Error binding address: ", err)
		os.Exit(1)
	}
	defer ln.Close()

	fmt.Println("Successfully binded to address: ", ba.Address)
	return nil
}
```