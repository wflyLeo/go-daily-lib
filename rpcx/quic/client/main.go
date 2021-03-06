package main

import (
	"context"
	"crypto/tls"
	"flag"
	"fmt"
	"log"

	"github.com/smallnest/rpcx/client"
)

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

var (
	addr = flag.String("addr", "localhost:8972", "service address")
)

func main() {
	flag.Parse()

	conf := &tls.Config{
		InsecureSkipVerify: true,
	}

	option := client.DefaultOption
	option.TLSConfig = conf
	d := client.NewPeer2PeerDiscovery("quic@"+*addr, "")
	xclient := client.NewXClient("Arith", client.Failtry, client.RandomSelect, d, option)
	defer xclient.Close()

	args := &Args{A: 10, B: 20}
	var reply int

	err := xclient.Call(context.Background(), "Mul", args, &reply)
	if err != nil {
		log.Fatalf("failed to call: %v", err)
	}

	fmt.Printf("%d * %d = %d\n", args.A, args.B, reply)

	args = &Args{50, 20}
	var quo Quotient
	err = xclient.Call(context.Background(), "Div", args, &quo)
	if err != nil {
		log.Fatalf("failed to call: %v", err)
	}

	fmt.Printf("%d * %d = %d...%d\n", args.A, args.B, quo.Quo, quo.Rem)
}
