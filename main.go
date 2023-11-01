package main

import (
	"context"
	"fmt"

	"github.com/parajuliswopnil/aleo-go-sdk/rpc"
)

func main() {

	client, err := rpc.NewClient("https://vm.aleo.org/api", "testnet3")
	if err != nil {
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()


	transaction, err := client.GetAllPeers(ctx)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(transaction)
}
