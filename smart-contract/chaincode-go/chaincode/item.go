package chaincode

import (
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	"fmt"
	"encoding/json"
)


type Item struct {
	ID			string	`json:"id"`
	Owner		string	`json:"owner"`
	Category	string	`json:"category"`
	Size		int		`json:"size"`

}


func (s *SmartContract) CreateItem(ctx contractapi.TransactionContextInterface,id string,owner string,category string,size int)error{
	exists, err:= s.ItemExists(ctx, id)
	if err!=nil{
		return err
	}
	if exists{
		return fmt.Errorf("Item %s already exists", id)
	}

	item:= Item{
		ID:			id,
		Owner:		owner,
		Category: 	category,
		Size:		size,
	}
	itemJSON,err:=json.Marshal(item)
	if err!=nil {
		return err
	}
	return ctx.GetStub().PutState(id,itemJSON)
}

func(s *SmartContract) ItemExists(ctx contractapi.TransactionContextInterface,id string)(bool,error){

	itemJSON, err := ctx.GetStub().GetState(id)
	if err!=nil{
		return false,fmt.Errorf("Error reading item from world state:%v",err)
	}
	return itemJSON != nil,nil
}

func (s *SmartContract) GetItem(ctx contractapi.TransactionContextInterface)([]*Item,error){
	resultsIterator,err:=ctx.GetStub().GetStateByRange("","")
	if err!=nil{
		return nil,err
	}
	defer resultsIterator.Close()

	var items []*Item
	for resultsIterator.HasNext(){
		queryResponse,err :=resultsIterator.Next()
		if err!=nil{
			return nil,err
		}
		var item Item
		err = json.Unmarshal(queryResponse.Value,&item)
		if err!= nil{
			return nil,err
		}
		items= append(items, &item)
	}
	return items,nil
}


func (s *SmartContract) GetitemById(ctx contractapi.TransactionContextInterface, id string)(*Item,error){
	itemJSON,err := ctx.GetStub().GetState(id)
	if err!=nil{
		return nil,fmt.Errorf("Failed to read from world state:%v",err)
	}
	if itemJSON == nil{
		return nil,fmt.Errorf("The item %s don't exists",id)
	}

	var item Item
	err = json.Unmarshal(itemJSON, &item)
	if err!= nil{
		return nil, err
	}
	return &item,nil

}