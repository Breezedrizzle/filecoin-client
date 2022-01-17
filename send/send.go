package send

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/lotus/chain/actors"
	"github.com/filecoin-project/lotus/chain/actors/builtin/power"
	miner0 "github.com/filecoin-project/specs-actors/actors/builtin/miner"
	power0 "github.com/filecoin-project/specs-actors/actors/builtin/power"
	builtin6 "github.com/filecoin-project/specs-actors/v6/actors/builtin"
	"github.com/myxtype/filecoin-client"
	"github.com/myxtype/filecoin-client/local"
	"github.com/myxtype/filecoin-client/types"
	"github.com/shopspring/decimal"
	"math/rand"
)
//FIL	1
//milliFIL	1000
//microFIL	1000000
//nanoFIL	1000000000
//picoFIL	1000000000000
//femtoFIL	1000000000000000
//attoFIL	1000000000000000000
//func SendToMiner

//fee 默认设置成 0.003
func SendFile(priv string,toAddr string ,amount float64,client  *filecoin.Client)(string,error){
	address.CurrentNetwork = address.Mainnet

	privData,err:=hex.DecodeString(priv)
	if err!=nil{
		return "",err
	}

	var key types.KeyInfo
	if err:=json.Unmarshal(privData,&key);err!=nil{
		return "",err
	}

	from,err:=local.WalletPrivateToAddress(local.ActSigType(key.Type),key.PrivateKey)
	if err!=nil{
		return "",err
	}

	actor,err:=client.StateGetActor(context.Background(),*from,nil)
	if err!=nil{
		return "",err
	}

	b,err:=client.WalletBalance(context.Background(),*from)
	if err!=nil{
		return "",err
	}

	balance,_:=filecoin.ToFil(b).Float64()
	if balance<amount{
		return "",fmt.Errorf("amount %f balance %f need ",amount,balance)
	}

	to, err := address.NewFromString(toAddr)
	if err!=nil{
		return "",err
	}

	msg := &types.Message{
		Version:    0,
		To:         to,
		From:       *from,
		Nonce:      actor.Nonce,
		Value:      filecoin.FromFil(decimal.NewFromFloat(amount)),
		GasLimit:   1333745+int64(rand.Int31n(10)),
		GasFeeCap:  abi.NewTokenAmount(2000000000+int64(rand.Int31n(10))),
		GasPremium: abi.NewTokenAmount(1000000+int64(rand.Int31n(10))),
		Method:     0,
		Params:     nil,
	}


	s, err := local.WalletSignMessage(key.Type, key.PrivateKey, msg)
	if err != nil {
		return "",err
	}

	mid, err := client.MpoolPush(context.Background(), s)
	if err != nil {
		return "",err
	}

	return mid.String(),nil
}



func SendMethod(priv string,to address.Address ,method uint64,parma []byte, client  *filecoin.Client)(string,error){
	fmt.Printf("0x%x",parma) //todo
	return "",nil

	address.CurrentNetwork = address.Mainnet

	privData,err:=hex.DecodeString(priv)
	if err!=nil{
		return "",err
	}

	var key types.KeyInfo
	if err:=json.Unmarshal(privData,&key);err!=nil{
		return "",err
	}

	from,err:=local.WalletPrivateToAddress(local.ActSigType(key.Type),key.PrivateKey)
	if err!=nil{
		return "",err
	}

	actor,err:=client.StateGetActor(context.Background(),*from,nil)
	if err!=nil{
		return "",err
	}


	msg := &types.Message{
		Version:    0,
		To:         to,
		From:       *from,
		Nonce:      actor.Nonce,
		Value:      abi.NewTokenAmount(0),
		GasLimit:		133404453+int64(rand.Int31n(104453)),
		GasFeeCap:  abi.NewTokenAmount(2000000000+int64(rand.Int31n(1000))),
		GasPremium: abi.NewTokenAmount(1000000+int64(rand.Int31n(10))),
		Method:     method,
		Params:     parma,
	}


	s, err := local.WalletSignMessage(key.Type, key.PrivateKey, msg)
	if err != nil {
		return "",err
	}

	mid, err := client.MpoolPush(context.Background(), s)
	if err != nil {
		return "",err
	}

	fmt.Printf("%+v\n",msg)

	return mid.String(),nil
}



func CreateMiner(priv , ownerStr string,client *filecoin.Client) (string,error) {
	creator:="f04"
	to, err := address.NewFromString(creator)
	if err!=nil{
		return "",err
	}

	owner, err := address.NewFromString(ownerStr)
	if err!=nil{
		return "",err
	}

	constructorParams := &power0.CreateMinerParams{
		Owner:         owner,
		Worker:        owner,
		Peer:          nil,
		SealProofType: 8,
	}

	enc, err := actors.SerializeParams(constructorParams)
	if err!=nil{
		return "",err
	}

	return  SendMethod(priv,to,uint64(power.Methods.CreateMiner),enc,client)
}

func WithdrawFromMiner(ownerPriv, minerStr string,amount float64,client *filecoin.Client) (string,error) {
	miner, err := address.NewFromString(minerStr)
	if err!=nil{
		return "",err
	}

	fil:=filecoin.FromFil(decimal.NewFromFloat(amount))
	constructorParams := &miner0.WithdrawBalanceParams{
		AmountRequested:fil,
	}

	enc, err := actors.SerializeParams(constructorParams)
	if err!=nil{
		return "",err
	}

	return  SendMethod(ownerPriv, miner,uint64(builtin6.MethodsMiner.WithdrawBalance),enc,client)
}

func ChangeOwner(optPriv,newOwnerStr, minerStr string,client *filecoin.Client) (string,error) {
	miner, err := address.NewFromString(minerStr)
	if err!=nil{
		return "",err
	}

	newOwner, err := address.NewFromString(newOwnerStr)
	if err!=nil{
		return "",err
	}

	enc, err := actors.SerializeParams(&newOwner)
	if err != nil {
		return "",err
	}

	return  SendMethod(optPriv, miner,uint64(builtin6.MethodsMiner.ChangeOwnerAddress),enc,client)
}


func Terminate(optPriv,newOwnerStr, minerStr string,client *filecoin.Client) (string,error) {
	miner, err := address.NewFromString(minerStr)
	if err!=nil{
		return "",err
	}

	newOwner, err := address.NewFromString(newOwnerStr)
	if err!=nil{
		return "",err
	}

	enc, err := actors.SerializeParams(&newOwner)
	if err != nil {
		return "",err
	}

	return  SendMethod(optPriv, miner,uint64(builtin6.MethodsMiner.ChangeOwnerAddress),enc,client)
}