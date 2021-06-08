package main

import (
	"privateledger/chaincode/model"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/hyperledger/fabric-chaincode-go/shim"
	pb "github.com/hyperledger/fabric-protos-go/peer"
	"github.com/pkg/errors"
)

func (t *PrivateLedgerChaincode) updateUserData(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	fmt.Println(" ******** Invoke Update User ******** ")

	var user model.User
	var name, email, mobile, age, salary, eventID string

	var collection string
	var targetOrg string
	var queryCreatorOrg string
	var queryCreatorRole string
	var queryCreator string
	var queryAccess, queryTxnHash string
	var needHistory bool

	/* User Data Parameter */
	name = args[1]
	email = args[2]
	mobile = args[3]
	age = args[4]
	salary = args[5]

	collection = args[6]
	queryAccess = args[7]
	queryTxnHash = args[8]

	eventID = args[9]
	queryCreatorOrg = args[10]
	queryCreatorRole = args[11]
	targetOrg = args[12]
	needHistory, _ = strconv.ParseBool(args[13])

	indexName := model.COLLECTION_KEY
	userNameIndexKey, err := stub.CreateCompositeKey(indexName, []string{email})

	if err != nil {
		return shim.Error(err.Error())
	}

	err = getPrivateDataFromLedger(stub, userNameIndexKey, collection, &user)
	if err != nil {
		return shim.Error(fmt.Sprintf("Unable to retrieve userData in the ledger: %v", err))
	}

	/*
			In this Section : The Transaction Hash for both org needs to be verify to get access to the
		ledger

			Transaction Hash : combination of  ( Owner Org + Sharing Org )

				Ex : Owner Org - org1 ,  Sharing Org - org2

				Transaction Hash =   Jwt.SignedString([]byte(org1 + org2))

			if the input hash and stored hash in the sharing collection matched then , it will
			allow to perform the transaction.

			It requires, when the Owner Org wants to perform few queries in the sharing collection Org
			but the Org don't have certain access query, then Owner Org will override the access and pass
			the transaction hash to perform the transaction on behalf of sharing collection Org.

	*/

	fmt.Println(" ********************* Validate Transaction Hash ******************** ")

	fmt.Println("       ############### Parse Access List ################             ")

	_, txnHash, err := parseOrgAccessList(targetOrg, user.Targets)

	if err != nil {
		return shim.Error(fmt.Sprintf("Unable to retrieve access details in the ledger: %v", err))
	}

	fmt.Println(" ################################################### ")

	fmt.Println(" ###### Input  -  queryTxnHash = " + queryTxnHash)
	fmt.Println(" ###### Stored -  txnHash = " + txnHash)
	fmt.Println(" ######  QueryAccess = " + queryAccess)

	if !strings.EqualFold(queryTxnHash, txnHash) {
		return shim.Error(fmt.Sprintf("Invalid transaction hash", errors.New("didn't match hash")))
	}

	fmt.Println(" ********************************************************************* ")

	///////////////////////////////////////////////////////////////////////////////////////////////////////

	fmt.Println(" ********************* Perfom the Update Transaction ******************** ")

	i, _ := strconv.Atoi(queryAccess)
	access := int32(i)

	/* Check if the access has Write or All access */

	if model.LedgerAccess(access).WriteAccess() ||
		model.LedgerAccess(access).AllAccess() {

		userdata := &model.User{
			ID:          user.ID,
			Name:        name,
			Email:       user.Email,
			Mobile:      mobile,
			Age:         age,
			Salary:      salary,
			Owner:       user.Owner,
			ShareAccess: user.ShareAccess,
			Role:        user.Role,
			Time:        user.Time,
			Targets:     user.Targets,
		}

		userDataJSONasBytes, err := json.Marshal(userdata)
		if err != nil {
			return shim.Error(err.Error())
		}

		err = stub.PutPrivateData(collection, userNameIndexKey, userDataJSONasBytes)

		if err != nil {
			return shim.Error(err.Error())
		}

		err = stub.SetEvent(eventID, []byte{})
		if err != nil {
			return shim.Error(err.Error())
		}

		fmt.Println(" ###### Update Data Parameters ###### ")
		fmt.Println(" Email		= " + email)
		fmt.Println(" Name 		= " + name)
		fmt.Println(" Mobile 		= " + mobile)
		fmt.Println(" Age		= " + age)
		fmt.Println(" Salary 		= " + salary)
		fmt.Println(" ################################## ")

		/*	Created History for Read by email Transaction */

		if needHistory {
			if strings.EqualFold(queryCreatorRole, model.ADMIN) {
				queryCreator = model.GetCustomOrgName(queryCreatorOrg) + " Admin"
			} else {
				queryCreator = email
			}

			fmt.Println(" ###### Query Access Details ###### ")
			fmt.Println(" queryCreator = " + queryCreator)
			fmt.Println(" queryCreatorRole = " + queryCreatorRole)
			fmt.Println(" queryAccess  = " + queryAccess)
			fmt.Println(" queryTxnHash = " + queryTxnHash)
			fmt.Println(" ################################## ")

			var change []string

			if !strings.EqualFold(name, user.Name) {
				change = append(change, " Name to "+name+" , ")
			}

			if !strings.EqualFold(mobile, user.Mobile) {
				change = append(change, " Mobile number to "+mobile+" , ")
			}

			if !strings.EqualFold(age, user.Age) {
				change = append(change, " Age to "+age+" , ")
			}

			if !strings.EqualFold(salary, user.Salary) {
				change = append(change, " Salary to "+salary+" , ")
			}

			query := args[0]
			remarks := queryCreator + " has done following changes \n " + " [ " + strings.Join(change[:], "\n") + " ] "
			t.createHistory(stub, queryCreator, targetOrg, email, query, remarks)
		}

	} else {
		return shim.Error("No WRITE access for " + queryCreatorOrg + " to upate " + email + "'s data")
	}

	return shim.Success(nil)
}

func (t *PrivateLedgerChaincode) updateTarget(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	fmt.Println("###############  Update Target Invoke ################")

	var user model.User
	var email string
	var targets string
	var collection string
	var eventID string

	email = args[1]
	targets = args[2]
	collection = args[3]
	eventID = args[4]

	indexName := model.COLLECTION_KEY
	userNameIndexKey, err := stub.CreateCompositeKey(indexName, []string{email})

	if err != nil {
		return shim.Error(err.Error())
	}

	err = getPrivateDataFromLedger(stub, userNameIndexKey, collection, &user)
	if err != nil {
		return shim.Error(fmt.Sprintf("Unable to retrieve userData in the ledger: %v", err))
	}

	updateUser := &model.User{
		ID:          user.ID,
		Name:        user.Name,
		Email:       user.Email,
		Mobile:      user.Mobile,
		Age:         user.Age,
		Salary:      user.Salary,
		Owner:       user.Owner,
		ShareAccess: user.ShareAccess,
		Role:        user.Role,
		Time:        user.Time,
		Targets:     targets,
	}

	userDataJSONasBytes, err := json.Marshal(updateUser)

	if err != nil {
		return shim.Error(err.Error())
	}

	err = stub.PutPrivateData(collection, userNameIndexKey, userDataJSONasBytes)

	if err != nil {
		return shim.Error(err.Error())
	}

	err = stub.SetEvent(eventID, []byte{})
	if err != nil {
		return shim.Error(err.Error())
	}

	fmt.Println("###############  Successfully Invoke Update Target ################")

	return shim.Success(nil)
}
