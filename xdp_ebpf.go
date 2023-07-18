```go
package main

import (
	"fmt"
	"github.com/iovisor/gobpf/bcc"
	"os"
)

const source string = `
#include <uapi/linux/bpf.h>
#include <linux/if_ether.h>
#include <linux/ip.h>
#include <linux/udp.h>

BPF_TABLE("prog", int, int, jmp, 1);

int forwarder(struct __sk_buff *skb) {
	bpf_trace_printk("Packet received\\n");
	return -1;
}

int forwardTraffic(struct __sk_buff *skb) {
	int key = 0;
	jmp.call(skb, &key);
	return 0;
}
`

func ForwardTraffic() {
	m := bcc.NewModule(source, []string{})
	defer m.Close()

	fn, err := m.Load("forwardTraffic", bcc.BPF_PROG_TYPE_SCHED_CLS)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to load forwardTraffic: %s\n", err)
		os.Exit(1)
	}

	err = m.AttachClassifier(fn)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to attach forwardTraffic: %s\n", err)
		os.Exit(1)
	}

	fmt.Println("Successfully loaded and attached BPF program. Forwarding traffic...")
}
```