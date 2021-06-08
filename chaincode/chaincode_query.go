package main

import (
	"privateledger/chaincode/model"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/golang/protobuf/proto"
	"github.com/hyperledger/fabric-chaincode-go/shim"
)

func getPrivateDataFromLedger(stub shim.ChaincodeStubInterface, key, collection string, result interface{}) error {

	resultAsByte, err := stub.GetPrivateData(collection, key)
	if err != nil {
		return fmt.Errorf("Get User Data Error " + err.Error())
	}

	if err != nil {
		return fmt.Errorf("unable to retrieve the object in the ledger: %v", err)
	}
	if resultAsByte == nil {
		return fmt.Errorf("the object doesn't exist in the ledger")
	}
	err = byteToObject(resultAsByte, result)
	if err != nil {
		return fmt.Errorf("unable to convert the result to object: %v", err)
	}
	return nil
}

// deleteFromLedger delete an object in the ledger
func deleteFromLedger(stub shim.ChaincodeStubInterface, key, collection string) error {

	err := stub.DelPrivateData(collection, key)
	if err != nil {
		return fmt.Errorf("unable to delete the object in the ledger: %v", err)
	}
	return nil
}

func objectToByte(object interface{}) ([]byte, error) {
	objectAsByte, err := json.Marshal(object)
	if err != nil {
		return nil, fmt.Errorf("unable convert the object to byte: %v", err)
	}
	return objectAsByte, nil
}

func byteToObject(objectAsByte []byte, result interface{}) error {
	err := json.Unmarshal(objectAsByte, result)
	if err != nil {
		return fmt.Errorf("unable to convert the result to object: %v", err)
	}
	return nil
}

func parseOrgAccessList(org string, targets string) (int32, string, error) {

	tgts := make(map[string][]byte)

	err := json.Unmarshal([]byte(targets), &tgts)

	if err != nil {
		fmt.Println("ParseOrgAccessList = failed to unmarshaling error: " + err.Error())
		return -1, "", fmt.Errorf("failed to unmarshaling error: ", err.Error())
	}

	for key, value := range tgts {

		orgName := key
		obj := []byte(value)

		fmt.Println(" ############# Unmarshal Target - " + orgName + "  ################## ")

		newTarget := &model.Target{}
		err = proto.Unmarshal([]byte(obj), newTarget)
		if err != nil {
			fmt.Println("ParseOrgTargets = unmarshaling error: " + err.Error())
			return -1, "", fmt.Errorf("unmarshaling error: ", err.Error())
		}

		if strings.EqualFold(org, orgName) {
			return newTarget.GetAccess(), newTarget.GetTransactionHash(), nil
		}

		fmt.Println(" ###### Org - " + orgName + " Access ###### ")
		fmt.Println(" Target Access = ", newTarget.GetAccess())
		fmt.Println(" Target Remarks = ", newTarget.GetRemarks())
		fmt.Println(" Target TransactionHash = ", newTarget.GetTransactionHash())

		fmt.Println(" ################################################## ")
	}

	return -1, "", nil
}
