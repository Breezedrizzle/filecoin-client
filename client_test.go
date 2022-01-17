package filecoin

import (
	"context"
	"testing"
)

// The Lotus Node
// The default token is in ~/.lotus/token
func testClient() *Client {
	return New("https://1lB5G4SmGdSTikOo7l6vYlsktdd:b58884915362a99b4fc18c2bf8af8358@filecoin.infura.io")
}

// The Lotus Node
// The default token is in ~/.lotus/token
func testLocalClient() *Client {
	return New("http://127.0.0.1:1234/rpc/v0").SetToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJBbGxvdyI6WyJyZWFkIiwid3JpdGUiLCJzaWduIiwiYWRtaW4iXX0.uOZVVMfnYo2Dos9ovPQDnN3sfy0CMilWeol0u0XE4_0")
}

// 测试RpcClient
func TestClient_Request(t *testing.T) {
	c := NewClient("https://eth-mainnet.token.im", "")
	var blockNumber string
	if err := c.Request(context.Background(), "eth_blockNumber", &blockNumber); err != nil {
		t.Fatal(err)
	}

	t.Log(blockNumber)

	var tr struct {
		BlockHash   string `json:"blockHash"`
		BlockNumber string `json:"blockNumber"`
	}
	if err := c.Request(context.Background(), "eth_getTransactionReceipt", &tr, "0xbb3a336e3f823ec18197f1e13ee875700f08f03e2cab75f0d0b118dabb44cba0"); err != nil {
		t.Error(err)
	}

	t.Log(tr.BlockHash)
	t.Log(tr.BlockNumber)
}
