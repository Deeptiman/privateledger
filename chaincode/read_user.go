package main

import (
	"privateledger/chaincode/model"
	"fmt"
	"strconv"
	"strings"

	"github.com/hyperledger/fabric-chaincode-go/shim"
	pb "github.com/hyperledger/fabric-protos-go/peer"
)

func (t *PrivateLedgerChaincode) readUser(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	fmt.Println(" ******** Invoke Read User ******** ")

	var user model.User
	var email, role, eventID string
	var queryCreatorOrg string
	var queryCreator string
	var needHistory bool
	var collection string

	email = args[1]
	collection = args[2]
	eventID = args[3]
	queryCreatorOrg = args[4]
	needHistory, _ = strconv.ParseBool(args[5])

	role, err := t.getRole(stub)
	if err != nil {
		return shim.Error(fmt.Sprintf("Unable to get roles from the account: %v", err))
	}

	fmt.Println(" Read User - Role === " + role)
	fmt.Println(" Read User - Collection === " + collection)

	fmt.Println("##### Read " + email + " User #####")

	indexName := model.COLLECTION_KEY
	userNameIndexKey, err := stub.CreateCompositeKey(indexName, []string{email})

	err = getPrivateDataFromLedger(stub, userNameIndexKey, collection, &user)
	if err != nil {
		return shim.Error(fmt.Sprintf("Unable to retrieve userData in the ledger: %v", err))
	}

	userAsByte, err := objectToByte(user)
	if err != nil {
		return shim.Error(fmt.Sprintf("Unable convert the userData to byte: %v", err))
	}

	err = stub.SetEvent(eventID, []byte{})
	if err != nil {
		return shim.Error(err.Error())
	}

	/*	Created History for Read by email Transaction */

	if needHistory {
		if strings.EqualFold(role, model.ADMIN) {
			queryCreator = model.GetCustomOrgName(queryCreatorOrg) + " Admin"
		} else {
			queryCreator = email
		}

		query := args[0]
		remarks := queryCreator + " read " + email + " 's user details"
		t.createHistory(stub, queryCreator, queryCreatorOrg, email, query, remarks)
	}

	return shim.Success(userAsByte)
}

func (t *PrivateLedgerChaincode) readAllUser(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	fmt.Println("##### Read All User #####")

	role, err := t.getRole(stub)
	if err != nil {
		return shim.Error(fmt.Sprintf("Unable to get roles from the account: %v", err))
	}

	if !strings.EqualFold(role, model.ADMIN) {
		return shim.Error(fmt.Sprintf("Only admin can read all the user data from the ledger: %v", err))
	}

	var eventID, collection string

	eventID = args[1]
	collection = args[2]

	fmt.Println(" Read All User - Role : " + role)
	fmt.Println(" Read All User - Collection : " + collection)

	indexName := model.COLLECTION_KEY

	iterator, err := stub.GetPrivateDataByPartialCompositeKey(collection, indexName, []string{})
	if err != nil {
		return shim.Error(fmt.Sprintf("Unable to retrieve the list of resource in the ledger: %v", err))
	}

	allUsers := make([]model.User, 0)

	for iterator.HasNext() {
		keyValueState, errIt := iterator.Next()
		if errIt != nil {
			return shim.Error(fmt.Sprintf("Unable to retrieve a user in the ledger: %v", errIt))
		}
		var user model.User
		err = byteToObject(keyValueState.Value, &user)
		if err != nil {
			return shim.Error(fmt.Sprintf("Unable to convert a user: %v", err))
		}

		fmt.Println("Read User : " + user.Name + " -- " + user.Email)

		allUsers = append(allUsers, user)
	}

	allUsersAsByte, err := objectToByte(allUsers)
	if err != nil {
		return shim.Error(fmt.Sprintf("Unable to convert the users list to byte: %v", err))
	}

	err = stub.SetEvent(eventID, []byte{})
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(allUsersAsByte)
}
