package main

import (
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"	
	pb "github.com/hyperledger/fabric/protos/peer"
	"github.com/hyperledger/fabric/protos/msp"
	"github.com/golang/protobuf/proto"
	"encoding/json"
	"github.com/privateledger/chaincode/model"
)

func (t *PrivateLedgerChaincode) createHistory(stub shim.ChaincodeStubInterface, queryCreator, targetOrg, email, query, remarks string ) pb.Response {

	serializedID, _ := stub.GetCreator()
	sId := &msp.SerializedIdentity{}
	err := proto.Unmarshal(serializedID, sId)
	if err != nil {
		return shim.Error(fmt.Sprintf("Could not deserialize a SerializedIdentity, err %s", err))
	}

	txID 		:= stub.GetTxID()	
	time, err 	:= stub.GetTxTimestamp()
	if err != nil {
		return shim.Error("Timestamp Error "+err.Error())
	}

	emailKey := email
	emailIndexKey, err := stub.CreateCompositeKey(emailKey, []string{txID})
	if err != nil {
		return shim.Error(err.Error())
	}

	fmt.Println("	################# Create History - "+email+" ###############	")
	
	fmt.Println("	EmailKey 	- "+emailIndexKey)
	fmt.Println("	TxID	 	- "+txID)
	fmt.Println("	QueryCreator	- "+queryCreator)
	fmt.Println("	Query		- "+query)
	fmt.Println("	TargetOrg	- "+targetOrg)
	fmt.Println("	Time		- "+time.String())
	fmt.Println("	Remarks		- "+remarks)	
	
	tm := model.GetTime(time)

	history := &model.HistoryData {
		EmailKey 		: emailIndexKey,
		TxID	 		: txID,
		QueryCreator	 	: queryCreator,
		Query			: query,
		TargetOrg		: targetOrg,
		Time			: tm,
		Remarks			: remarks,		
	}

	historyDataJSONasBytes, err := json.Marshal(history)

	if err != nil {
		return shim.Error(err.Error())
	}

	err = stub.PutState(emailIndexKey, historyDataJSONasBytes)

	if err != nil {
		return shim.Error(err.Error())
	}

	fmt.Println(" ############### History Created for - "+email)

	return shim.Success(nil)
}


func (t *PrivateLedgerChaincode) readHistory(stub shim.ChaincodeStubInterface, args[] string) pb.Response {

	var email, eventID string

	email = args[1]
	eventID = args[2]

	fmt.Println("	################# Read History - "+email+" ###############	")
	
	emailKey := email
	iterator, err := stub.GetStateByPartialCompositeKey(emailKey, []string{})
	if err != nil {
		return shim.Error(fmt.Sprintf("Unable to retrieve the history list of resource in the ledger: %v", err))
	}

	allHistoryData := make([]model.HistoryData, 0)

	for iterator.HasNext() {

		keyValueState, errIt := iterator.Next()
		if errIt != nil {
			return shim.Error(fmt.Sprintf("Unable to retrieve history of a user in the ledger: %v", errIt))
		}
		var historydata model.HistoryData
		err = byteToObject(keyValueState.Value, &historydata)
		if err != nil {
			return shim.Error(fmt.Sprintf("Unable to convert a history: %v", err))
		}

		allHistoryData = append(allHistoryData, historydata)
	}

	allHistoryAsByte, err := objectToByte(allHistoryData)
	if err != nil {
		return shim.Error(fmt.Sprintf("Unable to convert the history list to byte: %v", err))
	}

	err = stub.SetEvent(eventID, []byte{})
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(allHistoryAsByte)

}
