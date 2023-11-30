package chaincode

import (
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// SmartContract provides functions for managing an auction
type SmartContract struct {
	contractapi.Contract
}

type Auction struct {
	DocType        string `json:"docType"` //docType is used to distinguish the various types of objects in state database
	ID             string `json:"id"`
	Owner          string `json:"owner"`
	StartingPrice  int    `json:"startingprice"`
	IsOpen		   bool   `json:"isopen"`
	ItemId		   string `json:"itemid"`
}

func (s *SmartContract) InitLedger(ctx contractapi.TransactionContextInterface) error {
	auctions := []Auction{
		{DocType : "auction",ID: "auction1",   Owner: "Tomoko", StartingPrice: 300, IsOpen:false,ItemId:"item1"},
		{DocType : "auction",ID: "auction2",   Owner: "Brad", StartingPrice: 400, IsOpen:false,ItemId:"item2"},
		{DocType : "auction",ID: "auction3",   Owner: "Jin Soo", StartingPrice: 500, IsOpen:false,ItemId:"item3"},
	}

	for _, auction := range auctions {
		auctionJSON, err := json.Marshal(auction)
		if err!= nil{
			return err
		}
		err = ctx.GetStub().PutState(auction.ID,auctionJSON)
		if err != nil{
			return fmt.Errorf("%s marshaling auction",err)
		}
	}
	
	items :=[]Item{
		{ID: "item1",   Owner: "Ugo", Category:"Cars" , Size:10},
		{ID: "item2",   Owner: "Mario", Category: "Boats", Size:20},
		{ID: "item3",   Owner: "Luigi", Category: "Real estate", Size:30},

	}

	for _,item := range items{
		itemJSON, err := json.Marshal(item)
		if err!=nil{
			return err
		}
		err = ctx.GetStub().PutState(item.ID,itemJSON)
		if err!= nil{
			return fmt.Errorf("%s marshaling item",err)
		}
	}


	return nil
}

func (s *SmartContract) CreateAuction(ctx contractapi.TransactionContextInterface,id string,owner string, startingprice int) error{
	exists, err:=s.AuctionExists (ctx, id)
	if err != nil{
		return err
	}
	if exists{
		return fmt.Errorf("the auction %s already exists", id)
	}

	auction := Auction{
		ID:				id,
		Owner:			owner,
		StartingPrice:	startingprice,
	}
	auctionJSON, err:=json.Marshal(auction)
	if err != nil {
		return err
	}
	return ctx.GetStub().PutState(id,auctionJSON)
}

func (s *SmartContract) AuctionExists(ctx contractapi.TransactionContextInterface, id string) (bool,error){
	auctionJSON,err := ctx.GetStub().GetState(id)
	if err!=nil{
		return false, fmt.Errorf("failed to read from world state:%v",err)
	}
	return auctionJSON != nil,nil
}


func (s *SmartContract) GetAuctionById(ctx contractapi.TransactionContextInterface, id string)(*Auction,error){
	auctionJSON,err := ctx.GetStub().GetState(id)
	if err!=nil{
		return nil,fmt.Errorf("Failed to read from world state:%v",err)
	}
	if auctionJSON == nil{
		return nil,fmt.Errorf("The auction %s don't exists",id)
	}

	var auction Auction
	err = json.Unmarshal(auctionJSON, &auction)
	if err!= nil{
		return nil, err
	}
	return &auction,nil
}

func (s *SmartContract) GetAuction(ctx contractapi.TransactionContextInterface)([]*Auction,error){
	resultsIterator,err := ctx.GetStub().GetStateByRange("","")
	if err!= nil{
		return nil,err
	}
	defer resultsIterator.Close()

	var auctions []*Auction
	for resultsIterator.HasNext() {
		queryResponse,err := resultsIterator.Next()
		if err != nil{
			return nil, err
		}
		var auction Auction
		err = json.Unmarshal(queryResponse.Value, &auction)
		if err != nil {
			return nil,err
		}
		auctions = append(auctions, &auction)

	}

	return auctions,nil
} 



