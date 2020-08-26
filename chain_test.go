package filecoin

import (
	"context"
	"github.com/ipfs/go-cid"
	"testing"
)

// 根据消息Cid获取消息
func TestClient_ChainGetMessage(t *testing.T) {
	c := testClient()

	id, err := cid.Parse("bafy2bzacebrx3sb5do2b7cqgsnys2lkxtdq3pvjtgmdt2wclwmrtjeraa7x3q")
	if err != nil {
		t.Error(err)
	}

	msg, err := c.ChainGetMessage(context.Background(), id)
	if err != nil {
		t.Error(err)
	}

	t.Log(msg)
}

// 获取当前头部高度
func TestClient_ChainHead(t *testing.T) {
	c := testClient()

	ts, err := c.ChainHead(context.Background())
	if err != nil {
		t.Error(err)
	}

	t.Log(ts.Height)

	for _, n := range ts.Cids {
		bm, err := c.ChainGetBlockMessages(context.Background(), n)
		if err != nil {
			t.Error(err)
		}
		for index, msg := range bm.BlsMessages {
			t.Log(bm.Cids[index], msg)
		}
	}
}

// 根据高度遍历区块所有交易
func TestClient_ChainGetTipSetByHeight(t *testing.T) {
	c := testClient()

	ts, err := c.ChainGetTipSetByHeight(context.Background(), 1201, nil)
	if err != nil {
		t.Error(err)
	}
	for _, n := range ts.Cids {
		bm, err := c.ChainGetBlockMessages(context.Background(), n)
		if err != nil {
			t.Error(err)
		}
		for index, msg := range bm.BlsMessages {
			t.Log(bm.Cids[index], msg)
		}
	}
}
