package rpc

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/parajuliswopnil/aleo-go-sdk/types"
)

type Client struct {
	url string
}

func NewClient(RpcEndPoint, Network string) (*Client, error) {
	client := &Client{
		url: RpcEndPoint + "/" + Network,
	}
	return client, nil
}

func (c *Client) GetLatestHeight() (int64, error) {
	latestHeight := "/latest/height"
	requestUrl := c.url + latestHeight

	response, err := http.Get(requestUrl)
	if err != nil {
		return 0, err
	}
	ht, err := io.ReadAll(response.Body)

	height, err := strconv.Atoi(string(ht))
	if err != nil {
		return 0, err
	}
	return int64(height), err
}

func (c *Client) GetLatestHash() (string, error) {
	latestHash := "/latest/hash"
	requestUrl := c.url + latestHash

	response, err := http.Get(requestUrl)
	if err != nil {
		return "", err
	}
	ht, err := io.ReadAll(response.Body)

	if err != nil {
		return "", err
	}
	return string(ht), err
}

func (c *Client) GetLatestBlock() (*types.Block, error) {
	latestBlock := "/latest/block"
	requestUrl := c.url + latestBlock

	response, err := http.Get(requestUrl)
	if err != nil {
		return nil, err
	}
	bl, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}
	block := &types.Block{}

	err = json.Unmarshal(bl, block)
	if err != nil {
		return nil, err
	}
	return block, nil
}

func (c *Client) GetLatestRootState() (string, error) {
	latestRootState := "/latest/stateRoot"
	requestUrl := c.url + latestRootState

	response, err := http.Get(requestUrl)
	if err != nil {
		return "", err
	}
	rs, err := io.ReadAll(response.Body)

	if err != nil {
		return "", err
	}

	rootState := string(rs)
	lengthOfRootState := len(rootState)

	if string(rootState[0]) == "\"" && string(rootState[lengthOfRootState-1]) == "\"" {
		rootState = rootState[1 : lengthOfRootState-1]
	}
	return rootState, err
}

// get block by hash or height
func (c *Client) GetBlock(id string) (*types.Block, error) {
	blockEndpoint := "/block/" + id
	requestUrl := c.url + blockEndpoint

	response, err := http.Get(requestUrl)
	if err != nil {
		return nil, err
	}
	bl, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}

	block := &types.Block{}
	err = json.Unmarshal(bl, block)
	if err != nil {
		return nil, err
	}
	return block, nil
}

// gets block height of a given blocks hash
func (c *Client) GetHeightByHash(hash string) (int64, error) {
	rpcEndpoint := "/height/" + hash
	requestUrl := c.url + rpcEndpoint

	response, err := http.Get(requestUrl)
	if err != nil {
		return 0, err
	}
	ht, err := io.ReadAll(response.Body)

	if err != nil {
		return 0, err
	}

	height, err := strconv.Atoi(string(ht))
	if err != nil {
		return 0, err
	}
	return int64(height), nil
}

// gets the blocks transactions
func (c *Client) GetBlocksTransactions(height int64) ([]types.Transactions, error) {
	block, err := c.GetBlock(strconv.Itoa(int(height)))
	if err != nil {
		return nil, err
	}
	return block.Transactions, err
}

func (c *Client) GetTransactionById(transactionId string) (*types.Transaction, error) {
	rpcEndpoint := "/transaction/" + transactionId
	requestUrl := c.url + rpcEndpoint

	response, err := http.Get(requestUrl)
	if err != nil {
		return nil, err
	}
	t, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}
	transaction := &types.Transaction{}

	err = json.Unmarshal(t, transaction)
	return transaction, err 
}

