package main

import (
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

// PrivateLedgerChaincode implementation of Chaincode
type PrivateLedgerChaincode struct {
}

// Init of the chaincode
// This function is called only one when the chaincode is instantiated.
// So the goal is to prepare the ledger to handle future requests.
func (t *PrivateLedgerChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println("########### PrivateLedgerChaincode Init ###########")

	function, _ := stub.GetFunctionAndParameters()

	if function != "init" {
		return shim.Error("unknown function call")
	}

	return shim.Success(nil)
}

// Invoke
// All future requests named invoke will arrive here.
func (t *PrivateLedgerChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println("########### PrivateLedgerChaincode Invoke *********************************** ###########")

	// Get the function and arguments from the request
	_, args := stub.GetFunctionAndParameters()

	// Check whether the number of arguments is sufficient
	if len(args) < 1 {
		return shim.Error("The number of arguments is insufficient.")
	}

	fmt.Println("Called Function - "+args[0])

	if args[0] == "createUser" {

		fmt.Println("Create User Function Called")
		return t.createUser(stub, args)

	} else if args[0] == "updateUserData" {

		fmt.Println("Update User Data Function Called")
		return t.updateUserData(stub, args)

	} else if args[0] == "updateTarget" {

		fmt.Println("Update Target Data Function Called")
		return t.updateTarget(stub, args)

	} else if args[0] == "readUser" {

		fmt.Println("Read User Function Called")
		return t.readUser(stub, args)

	} else if args[0] == "readAllUser" {

		fmt.Println("Read All User Function Called")
		return t.readAllUser(stub, args)

	} else if args[0] == "readHistory" {

		fmt.Println("Read History Data Function Called")
		return t.readHistory(stub, args)
	} else if args[0] == "shareUser" {

		fmt.Println("Share Data Function Called")
		return t.shareUser(stub, args)
		
	} else if args[0] == "deleteUser" {

		fmt.Println("Delete User Function Called")
		return t.deleteUser(stub, args)
	} else if args[0] == "testInvoke" {

		eventID := args[1]

		fmt.Println(" #####  Test Event  - "+eventID)

		err := stub.SetEvent(eventID, []byte{})
		if err != nil {
			return shim.Error(err.Error())
		}

		return shim.Success(nil)
	}

	return shim.Success(nil)
}

func main() {
	// Start the chaincode and make it ready for futures requests
	err := shim.Start(new(PrivateLedgerChaincode))
	if err != nil {
		return
	}
}
