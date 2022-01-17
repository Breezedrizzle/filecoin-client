package filecoin

import (
	"context"
	"fmt"
	"github.com/filecoin-project/go-address"
	"github.com/ipfs/go-cid"
	"testing"
)

// 查询消息/交易执行状态
func TestClient_StateGetReceipt(t *testing.T) {
	c := testClient()

	id, err := cid.Parse("bafy2bzacebrx3sb5do2b7cqgsnys2lkxtdq3pvjtgmdt2wclwmrtjeraa7x3q")
	if err != nil {
		t.Error(err)
	}

	mr, err := c.StateGetReceipt(context.Background(), id, nil)
	if err != nil {
		t.Error(err)
	}

	t.Log(mr)
}

// 查询消息状态
// Receipt 为空表示未执行
func TestClient_StateSearchMsg(t *testing.T) {
	c := testClient()

	id, err := cid.Parse("bafy2bzacebrx3sb5do2b7cqgsnys2lkxtdq3pvjtgmdt2wclwmrtjeraa7x3q")
	if err != nil {
		t.Error(err)
	}

	msg, err := c.StateSearchMsg(context.Background(), id)
	if err != nil {
		t.Error(err)
	}

	if msg == nil {
		t.Log("nil")
	} else {
		t.Log(msg)
	}
}

func TestClient_StateGetActor(t *testing.T) {
	c := testClient()

	address.CurrentNetwork = address.Mainnet

	addr, _ := address.NewFromString("f3qx3jo74v6d6z35qhfeax3xozsegzliowrrchuyumshnwb2kz66xajhl55pxjr5xvvpeggioytv7uko5hpzga")

	actor, err := c.StateGetActor(context.Background(), addr, nil)
	if err != nil {
		t.Error(err)
	}

	t.Log(actor.Nonce)

	nonce, err := c.MpoolGetNonce(context.Background(), addr)
	if err != nil {
		t.Error(err)
	}

	t.Log(nonce)
}


func TestClient_StateLookup(t *testing.T) {
	c := testClient()
	address.CurrentNetwork = address.Mainnet

	id, err := c.StateLookUp(context.Background(), "f1tybjx2ri2khuugqpsc4f34vbvcb64gavhfgu6pq")
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("%+v", id)
}

func TestClient_StateSectors(t *testing.T) {
	c := testClient()
	address.CurrentNetwork = address.Mainnet

	secotrs, err := c.StateSectors(context.Background(), "f020436")
	if err != nil {
		t.Error(err)
		return
	}

	for _,v:=range secotrs{
		fmt.Println(v.SectorNumber)
	}

}

func TestClient_StateListMiners(t *testing.T) {
	c := testClient()
	address.CurrentNetwork = address.Mainnet

	res, err := c.StateListMiners(context.Background())
	if err != nil {
		t.Error(err)
		return
	}

	for _,v:=range res {
		fmt.Println(v)
	}

}


func TestClient_StateReplay(t *testing.T) {
	c := testClient()
	address.CurrentNetwork = address.Mainnet

	cid,err:=cid.Parse("bafy2bzacebq2b7ltqct5prf6lqgaksmwlnl6phunujoko33wwqwrziyywldcm")
	if err != nil {
		t.Error(err)
		return
	}

	res, err := c.StateReplay(context.Background(),nil,cid)
	if err != nil {
		t.Error(err)
		return
	}

	fmt.Printf("%+v\n",res.Msg)
	fmt.Println(string(res.Msg.Params))
	fmt.Printf("%x",res.Msg.Params)
}

