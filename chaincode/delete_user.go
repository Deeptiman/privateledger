package main

import (
	"privateledger/chaincode/model"
	"fmt"
	"strconv"
	"strings"

	"github.com/hyperledger/fabric-chaincode-go/shim"
	pb "github.com/hyperledger/fabric-protos-go/peer"
	"github.com/pkg/errors"
)

func (t *PrivateLedgerChaincode) deleteUser(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	fmt.Println(" ******** Invoke Delete User ******** ")

	var user model.User

	var email, role, eventID string

	var queryCreatorOrg string
	var queryCreatorRole string
	var targetOrg string
	var queryCreator string

	var queryAccess, queryTxnHash string
	var collection string
	var needHistory bool

	email = args[1]
	collection = args[2]
	eventID = args[3]

	queryAccess = args[4]
	queryTxnHash = args[5]

	role = args[6]
	queryCreatorOrg = args[7]
	queryCreatorRole = args[8]
	targetOrg = args[9]
	needHistory, _ = strconv.ParseBool(args[10])

	fmt.Println(" ###### Delete Data Parameters ###### ")
	fmt.Println(" Email		= " + email)
	fmt.Println(" Role		= " + role)
	fmt.Println(" EventID	= " + eventID)
	fmt.Println(" ################################## ")

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

	fmt.Println(" ********************* Perfom the Delete Transaction ******************** ")

	i, _ := strconv.Atoi(queryAccess)
	access := int32(i)

	/* Check if the access has Delete or All access */

	if model.LedgerAccess(access).DeleteAccess() ||
		model.LedgerAccess(access).AllAccess() {

		err = deleteFromLedger(stub, userNameIndexKey, collection)
		if err != nil {
			return shim.Error(fmt.Sprintf("Unable to delete the user in the ledger: %v", err))
		}

		err = stub.SetEvent(eventID, []byte{})
		if err != nil {
			return shim.Error(err.Error())
		}

		/*	Created History for Delete by email Transaction */

		if needHistory {
			var remarks string
			if strings.EqualFold(queryCreatorRole, model.ADMIN) {
				queryCreator = model.GetCustomOrgName(queryCreatorOrg) + " Admin"
				remarks = queryCreator + " has deleted user - " + email
			} else {
				queryCreator = email
				remarks = queryCreator + " has deleted the account"
			}

			fmt.Println(" ###### Query Access Details ###### ")
			fmt.Println(" queryCreator = " + queryCreator)
			fmt.Println(" queryCreatorRole = " + queryCreatorRole)
			fmt.Println(" queryAccess  = " + queryAccess)
			fmt.Println(" queryTxnHash = " + queryTxnHash)
			fmt.Println(" ################################## ")

			query := args[0]
			t.createHistory(stub, queryCreator, targetOrg, email, query, remarks)
		}

	} else {
		return shim.Error("No Delete access for " + queryCreatorOrg + " to delete " + email)
	}

	return shim.Success(nil)
}
