package rpc

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"

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
	return block.Transactions, nil
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
	return transaction, nil
}

// func (c *Client) GetMemoryPoolTransactions() ([]types.Transactions, error) {
// 	rpcEndpoint := "/memoryPool/transactions"
// 	requestUrl := c.url + rpcEndpoint

// 	response, err := http.Get(requestUrl)
// 	if err != nil {
// 		return nil, err
// 	}
// 	t, err := io.ReadAll(response.Body)

// 	if err != nil {
// 		return nil, err
// 	}
// 	transaction := &types.Transaction{}

// 	err = json.Unmarshal(t, transaction)
// 	return transaction, err
// }

// returns the abi of the aleo program and saves the abi to the desired path
func (c *Client) GetProgram(programId, path string) error {
	rpcEndpoint := "/program/" + programId
	requestUrl := c.url + rpcEndpoint

	response, err := http.Get(requestUrl)
	if err != nil {
		return err
	}
	pr, err := io.ReadAll(response.Body)

	if err != nil {
		return err
	}

	program := string(pr)
	lengthOfRootState := len(program)

	if string(program[0]) == "\"" && string(program[lengthOfRootState-1]) == "\"" {
		program = program[1 : lengthOfRootState-1]
	}

	programIntermidiate := strings.Split(program, "\\n")

	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	for _, line := range programIntermidiate {
		_, err := file.WriteString(line + "\n")
		if err != nil {
			return err
		}
	}

	return nil

}

func (c *Client) GetMappingNames(programId string) ([]string, error) {
	var mapping []string
	rpcEndpoint := "/program/" + programId + "/mappings"
	requestUrl := c.url + rpcEndpoint

	response, err := http.Get(requestUrl)
	if err != nil {
		return nil, err
	}
	t, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(t, &mapping)

	fmt.Println(mapping)
	return mapping, nil
}

// returns the value in a key-value mapping corresponding to the supplied mappingKey
func (c *Client) GetMappingValue(programId, mappingName, mappingKey string) (map[string]interface{}, error) {
	rpcEndpoint := "/program/" + programId + "/mapping/" + mappingName + "/" + mappingKey
	requestUrl := c.url + rpcEndpoint

	response, err := http.Get(requestUrl)
	if err != nil {
		return nil, err
	}
	t, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}

	value := make(map[string]interface{})

	value[mappingKey] = string(t)

	return value, nil
}

// Returns the state path for the given commitment. The state path proves existence of the transition leaf to either a global or local state root.
func (c *Client) GetStatePathForCommitment(commitment string) {}

// returns the list of current beacon node addresses
func (c *Client) GetBeacons() ([]string, error) {
	rpcEndpoint := "/beacons"
	requestUrl := c.url + rpcEndpoint

	response, err := http.Get(requestUrl)
	if err != nil {
		return nil, err
	}
	t, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}

	var beacons []string
	err = json.Unmarshal(t, &beacons)
	if err != nil {
		return nil, err
	}
	return beacons, nil
}

func (c *Client) GetPeersCount() (int, error) {
	rpcEndpoint := "/peers/count"
	requestUrl := c.url + rpcEndpoint

	response, err := http.Get(requestUrl)
	if err != nil {
		return 0, err
	}
	t, err := io.ReadAll(response.Body)

	if err != nil {
		return 0, err
	}

	count, err := strconv.Atoi(string(t))
	if err != nil {
		return 0, err
	}

	return count, nil
}

// Returns the peers connected to the node.
func (c *Client) GetAllPeers() ([]string, error) {
	rpcEndpoint := "/peers/all"
	requestUrl := c.url + rpcEndpoint

	response, err := http.Get(requestUrl)
	if err != nil {
		return nil, err
	}
	t, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}

	var peers []string
	err = json.Unmarshal(t, &peers)
	if err != nil {
		return nil, err
	}
	return peers, nil
}

func (c *Client) GetNodeAddress() (string, error) {
	rpcEndpoint := "/node/address"
	requestUrl := c.url + rpcEndpoint

	response, err := http.Get(requestUrl)
	if err != nil {
		return "", err
	}
	t, err := io.ReadAll(response.Body)

	if err != nil {
		return "", err
	}

	return string(t), nil
}

// returns the block hash related to the transaction id
func (c *Client) GetBlockHashByTransactionId(transactionId string) (string, error) {
	rpcEndpoint := "/find/blockHash/" + transactionId
	requestUrl := c.url + rpcEndpoint

	response, err := http.Get(requestUrl)
	if err != nil {
		return "", err
	}
	t, err := io.ReadAll(response.Body)

	if err != nil {
		return "", err
	}

	hash := string(t)
	lengthOfRootState := len(hash)

	if string(hash[0]) == "\"" && string(hash[lengthOfRootState-1]) == "\"" {
		hash = hash[1 : lengthOfRootState-1]
	}
	return hash, err
}

// returns transaction id related to the program id
func (c *Client) FindTransactionIDByProgramID(programId string) (string, error) {
	rpcEndpoint := "/find/transactionID/deployment/" + programId
	requestUrl := c.url + rpcEndpoint

	response, err := http.Get(requestUrl)
	if err != nil {
		return "", err
	}
	t, err := io.ReadAll(response.Body)

	if err != nil {
		return "", err
	}

	transactionId := string(t)
	lengthOfRootState := len(transactionId)

	if string(transactionId[0]) == "\"" && string(transactionId[lengthOfRootState-1]) == "\"" {
		transactionId = transactionId[1 : lengthOfRootState-1]
	}
	return transactionId, err
}

// return transaction id related to the transition id
func (c *Client) FindTransactionIDByTransitionID(transitionId string) (string, error) {
	rpcEndpoint := "/find/transactionID/" + transitionId
	requestUrl := c.url + rpcEndpoint

	response, err := http.Get(requestUrl)
	if err != nil {
		return "", err
	}
	t, err := io.ReadAll(response.Body)

	if err != nil {
		return "", err
	}

	transactionId := string(t)
	lengthOfRootState := len(transactionId)

	if string(transactionId[0]) == "\"" && string(transactionId[lengthOfRootState-1]) == "\"" {
		transactionId = transactionId[1 : lengthOfRootState-1]
	}
	return transactionId, err
}

func (c *Client) FindTransitionIDByInputOrOutputID(ioId string) (string, error) {
	rpcEndpoint := "/find/transitionID/" + ioId
	requestUrl := c.url + rpcEndpoint

	response, err := http.Get(requestUrl)
	if err != nil {
		return "", err
	}
	t, err := io.ReadAll(response.Body)

	if err != nil {
		return "", err
	}

	transitionId := string(t)
	lengthOfRootState := len(transitionId)

	if string(transitionId[0]) == "\"" && string(transitionId[lengthOfRootState-1]) == "\"" {
		transitionId = transitionId[1 : lengthOfRootState-1]
	}
	return transitionId, err
}
