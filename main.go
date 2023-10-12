package main

import (
	"fmt"

	"github.com/parajuliswopnil/aleo-go-sdk/rpc"
)

func main() {

	client, err := rpc.NewClient("https://vm.aleo.org/api", "testnet3")
	if err != nil {
		return
	}

	block, err := client.GetBlock("68100")

	transaction, err := client.FindTransitionIDByInputOrOutputID(block.Transactions[0].Transaction.Execution_.Transitions[0].Outputs[0].Id)

	fmt.Println(transaction)
}
