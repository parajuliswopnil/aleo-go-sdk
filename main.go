package main

import (
	"fmt"


	"github.com/parajuliswopnil/aleo-go-sdk/rpc"
)

func main() {

	client, err := rpc.NewClient("https://vm.aleo.org/api","testnet3")
	if err != nil {
		return
	}

	transaction, err := client.GetTransactionById("at1flp4j2wsdrv0znqr9xk6ka2ujf5crgsru8djgv07hfe6hmap85fqttprc6")
	fmt.Println(transaction)

}
