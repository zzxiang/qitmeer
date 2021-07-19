// Copyright 2021 The qitmeer developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.
package main

import (
	"flag"
	"fmt"
	"github.com/Qitmeer/qitmeer/common/encode/base58"
	"github.com/Qitmeer/qitmeer/params"
	"math/rand"
	"os"
	"strings"
	"time"
)

const addrSize int = 35
const alphabet string = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"
const defaultNetwork string = "testnet"

// The generator of the Qitmeer burn addresses. The tool which generate a valid
// qitmeer-base58check encoded address for the specified network (the default
// is testnet).
// Note: The template need to be long enough to remain the strong security.
// (recommend at least 16 words)
// See https://en.bitcoin.it/wiki/Vanitygen for the details
//
func main() {
	var template string
	var network string
	var generate bool
	flag.StringVar(&template, "t", "","template")
	flag.StringVar(&network, "n",defaultNetwork ,"network [mainnet|testnet|mixnet|privnet]")
	flag.BoolVar(&generate, "new", false, "generate new address")
	flag.Parse()
	p, err := getParams(network);
	exitIfErr(err)
	if template == "" {
		template = genTemplateByParams(p)
	}
	addr, err := getAddr(template,p,generate)
	exitIfErr(err)
	fmt.Printf("template = %s \n", template)
	fmt.Printf("    addr = %v \n", string(addr));
}

func exitIfErr(err error){
	if err != nil {
		fmt.Printf("error: %v\n", err);
		os.Exit(-1)
	}
}
func getParams(network string) (*params.Params, error) {
	switch network {
	case "testnet":
		return &params.TestNetParams, nil
	case "privnet":
		return &params.PrivNetParams, nil
	case "mainnet":
		return &params.MainNetParams, nil
	case "mixnet":
		return &params.MixNetParams, nil
	default:
		return nil, fmt.Errorf("unknown network %s",network)
	}
}
func genTemplateByParams(p *params.Params) string {
	var sb strings.Builder
	sb.WriteString(p.NetworkAddressPrefix)
	sb.WriteString("mQitmeer")
	sb.WriteString(strings.Title(p.Name))
	sb.WriteString("BurnAddress")
	return sb.String()
}

func getAddr(template string, p *params.Params, randomSuffix bool) ([]byte, error) {
	pickSize := addrSize-len(template)
	var sb strings.Builder
	sb.WriteString(template);
	if randomSuffix {
		rand.Seed(time.Now().UnixNano())
		for i := 0; i < pickSize; i++ {
			randomIndex := rand.Intn(58)
			pick := alphabet[randomIndex]
			sb.WriteString(string(pick))
		}
	} else {
		for i := 0; i < pickSize; i++ {
			sb.WriteString("X");
		}
	}
	decoded := base58.Decode([]byte(sb.String()))
	if len(decoded) == 0 {
		return nil, fmt.Errorf("incorrect base58 encoded template %s", sb.String());
	}
	addr, err := base58.QitmeerCheckEncode(decoded[2:22],p.PubKeyHashAddrID[:])
 	if err!=nil {
 		return nil, err
	}
	_,_,err = base58.QitmeerCheckDecode(string(addr));
	if err!=nil {
		return nil, err
	}
	return addr,nil
}
