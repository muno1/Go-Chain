/*
SPDX-License-Identifier: Apache-2.0
*/

package main

import (
	"log"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	"github.com/muno1/Go-Chain/tree/prototype/smart-contract/chaincode-go"
)

func main() {
	auctionChaincode, err := contractapi.NewChaincode(&chaincode.SmartContract{})
	if err != nil {
		log.Panicf("Error creating auction chaincode: %v", err)
	}

	if err := auctionChaincode.Start(); err != nil {
		log.Panicf("Error auction chaincode: %v", err)
	}


}
