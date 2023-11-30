package main

import (
	"log"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	"github.com/hyperledger/fabric-samples/smart-contract/chaincode-go/chaincode"
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
