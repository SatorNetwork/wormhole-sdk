package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"

	publicrpcv1 "github.com/evgeniy-scherbina/sandbox/wormhole-sdk/proto/publicrpc/v1"
)

const (
	address = "localhost:7070"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := publicrpcv1.NewPublicRPCServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	resp, err := client.GetSignedVAA(ctx, &publicrpcv1.GetSignedVAARequest{
		MessageId: &publicrpcv1.MessageID{
			EmitterChain:   publicrpcv1.ChainID_CHAIN_ID_ETHEREUM,
			EmitterAddress: "0000000000000000000000000290fb167208af455bb137780163b7b7a9a10c16",
			Sequence:       1,
		},
	})
	if err != nil {
		log.Fatalf("can't get signed vaa: %v", err)
	}

	displayVaa(resp.VaaBytes)
}

func displayVaa(vaa []byte) {
	fmt.Print("[")
	for idx, nextByte := range vaa {
		fmt.Print(nextByte)
		if idx != len(vaa)-1 {
			fmt.Print(", ")
		}
	}
	fmt.Printf("]\n")
}
