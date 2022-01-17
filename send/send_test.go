package send

import (
	"fmt"
	"github.com/myxtype/filecoin-client"
	"log"
	"testing"
)

var trustPriv = "7b2254797065223a22736563703235366b31222c22507269766174654b6579223a22376f634a5873376735466d594233425a3443623576464679444b53353376567550493231624657526636383d227d"
var trustAddr = "f0821344"

var test01Priv= "7b2254797065223a22626c73222c22507269766174654b6579223a227276322f50384a5255426a57756e7562325946743447493644366a554e586773654541714879514562556f3d227d"
var test01Addr = "f01638052"

var miner="f01646210"
var client = filecoin.New("https://1lB5G4SmGdSTikOo7l6vYlsktdd:b58884915362a99b4fc18c2bf8af8358@filecoin.infura.io")


func Test_SendFile(t *testing.T) {
	hash,err:= SendFile(test01Priv,trustAddr,0.5,client)
	if err!=nil{
		log.Fatal(err)
	}

	fmt.Println("Send",hash)
}


func Test_CreateMiners(t *testing.T) {
	hash,err:=CreateMiner(test01Priv, test01Addr,client)
	if err!=nil{
		log.Fatal(err)
	}

	fmt.Println(hash)
	return
}

func Test_WithdrawFromMiner(t *testing.T) {
	hash,err:=WithdrawFromMiner(trustPriv,miner,0.0998,client)
	if err!=nil{
		log.Fatal(err)
	}

	fmt.Println(hash)
	return
}

// create miner with lotus ,序列化 data
//lotus send  --gas-premium 1000000 --gas-feecap 2000000000 --gas-limit 133404453   --from f3qcvyi3o74h2ahdgf342vg5bpddwvhw5lfmjrqhcsgnrcfbynawq4d6y76vteihpxwilcwx25jibglx3sp75a --method 2   --params-hex 854400a4fd634400a4fd63084080 f04 0
//
func Test_ChangeOwner(t *testing.T) {
	hash,err:=ChangeOwner(trustPriv,trustAddr,miner,client)
	if err!=nil{
		log.Fatal(err)
	}

	fmt.Println(hash)
	return
}
