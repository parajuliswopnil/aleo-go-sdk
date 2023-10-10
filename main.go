package main

import (
	"fmt"

	"github.com/parajuliswopnil/aleo-go-sdk/rpc"
)

func main() {
	// rpc := "/testnet3/block/48512"
	// requestUrl := fmt.Sprintf("https://vm.aleo.org/api" + rpc)
	// response, err := http.Get(requestUrl)

	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// body, err := io.ReadAll(response.Body)

	// if err != nil {
	// 	return
	// }

	// block := &types.Block{}
	// err = json.Unmarshal(body, block)

	// fmt.Println(block.Transactions[0].Transaction.Id)

	client, err := rpc.NewClient("https://vm.aleo.org/api","testnet3")
	if err != nil {
		return
	}

	latestHeight, err := client.GetLatestRootState()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(latestHeight)

}
