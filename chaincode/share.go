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

func (t *PrivateLedgerChaincode) shareUser(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	fmt.Println(" ******** Invoke Share User ******** ")

	var queryCreator string
	var queryCreatorOrg string
	var queryCreatorRole string
	var queryAccess, queryTxnHash string

	var name, email, mobile, age, salary string
	var role string
	var remarks, shareAccess string

	var targets string
	var sharingCollection string
	var eventID string
	var sharingOrg string

	var needHistory bool

	/* User Data Parameter */
	name = args[1]
	email = args[2]
	mobile = args[3]
	age = args[4]
	salary = args[5]
	role = args[6]

	eventID = args[7]
	shareAccess = args[8]
	queryTxnHash = args[9]
	targets = args[10]
	sharingCollection = args[11]
	sharingOrg = args[12]
	queryCreatorOrg = args[13]
	queryCreatorRole = args[14]
	needHistory, _ = strconv.ParseBool(args[15])

	indexName := model.COLLECTION_KEY
	userNameIndexKey, err := stub.CreateCompositeKey(indexName, []string{email})

	if err != nil {
		return shim.Error(err.Error())
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

	fmt.Println(" SharingOrg == " + sharingOrg)
	fmt.Println(" QueryCreator == " + queryCreatorOrg)

	access, txnHash, err := parseOrgAccessList(sharingOrg, targets)

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

	if !model.LedgerAccess(access).RemoveAccess() {

		fmt.Println("##### Update Sharing Access = ", access)

		txID := stub.GetTxID()

		timestamp, err := stub.GetTxTimestamp()
		if err != nil {
			return shim.Error("Timestamp Error " + err.Error())
		}

		tm := model.GetTime(timestamp)

		user := &model.User{
			ID:          txID,
			Name:        name,
			Email:       email,
			Mobile:      mobile,
			Age:         age,
			Salary:      salary,
			ShareAccess: shareAccess,
			Owner:       queryCreatorOrg,
			Role:        role,
			Time:        tm,
			Targets:     targets,
		}

		userAsByte, err := objectToByte(user)
		if err != nil {
			return shim.Error(fmt.Sprintf("Unable convert the userData to byte: %v", err))
		}

		err = stub.PutPrivateData(sharingCollection, userNameIndexKey, userAsByte)

		if err != nil {
			return shim.Error(err.Error())
		}

	} else {

		fmt.Println("##### Delete Sharing Access for " + sharingOrg)

		err = deleteFromLedger(stub, userNameIndexKey, sharingCollection)
		if err != nil {
			return shim.Error(fmt.Sprintf("Unable to delete the user in the ledger: %v", err))
		}
	}

	err = stub.SetEvent(eventID, []byte{})
	if err != nil {
		return shim.Error(err.Error())
	}

	if needHistory {

		if strings.EqualFold(queryCreatorRole, model.ADMIN) {
			queryCreator = model.GetCustomOrgName(queryCreatorOrg) + " Admin"
		} else {
			queryCreator = email
		}

		fmt.Println(" ###### Query Access Details ###### ")
		fmt.Println(" queryCreator = " + queryCreator)
		fmt.Println(" queryCreatorRole = " + queryCreatorRole)
		fmt.Println(" ################################## ")

		query := args[0]

		if access == model.LedgerAccess(model.NOACCESS).Int() {

			remarks = sharingOrg + " currently has no access to user - " + email

		} else if access == model.LedgerAccess(model.REMOVEACCESS).Int() {

			remarks = model.GetCustomOrgName(queryCreatorOrg) + " remove all access for " + model.GetCustomOrgName(sharingOrg)

		} else {

			a := model.LedgerAccess(access).String()
			remarks = model.GetCustomOrgName(queryCreatorOrg) + " gave " + a + " access to " + model.GetCustomOrgName(sharingOrg)

		}

		t.createHistory(stub, queryCreator, sharingOrg, email, query, remarks)
	}

	fmt.Println("###############  Successfully Invoke Share User ################")

	return shim.Success(nil)
}
