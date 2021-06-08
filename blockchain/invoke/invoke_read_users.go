package invoke

import (
	"privateledger/chaincode/model"
	"encoding/json"
	"fmt"
	"strconv"
)

func (s *OrgInvoke) GetAllUsersFromLedger() ([]model.User, error) {

	fmt.Println(" ############## Invoke Read All User ################")

	eventID := "getAllUsersInvoke"
	collection := s.User.Setup.OrgCollection

	response, err := s.User.Setup.ExecuteChaincodeTranctionEvent(eventID, "invoke",
		[][]byte{
			[]byte("readAllUser"),
			[]byte(eventID),
			[]byte(collection),
		}, s.User.Setup.ChaincodeId, s.User.ChannelClient, s.User.Event)

	if err != nil {
		return nil, fmt.Errorf("Error - addUserToLedger : %s", err.Error())
	}

	fmt.Println("Response Received")

	allUsers := make([]model.User, 0)

	if response != nil && response.Payload == nil {
		return nil, fmt.Errorf("unable to get response for the query: %v", err)
	}

	if response != nil {
		err = json.Unmarshal(response.Payload, &allUsers)
		if err != nil {
			return nil, fmt.Errorf("unable to convert response to the object given for the query: %v", err)
		}
	}

	if len(allUsers) < 1 {
		return nil, fmt.Errorf("No records found")
	}

	/*fmt.Println("#### All Users Records Found ####")
	for _, user := range allUsers {

		fmt.Println("Read User : "+user.Name+" -- "+user.Email)

		targets := user.Targets
		err := s.ParseOrgTargets(targets)

		if err != nil {
			fmt.Errorf("Failed to parse Org Targets - %s ",err.Error())
		}
	}*/

	return allUsers, nil
}

func (s *OrgInvoke) GetUserFromLedger(email string, needHistory bool) (*model.User, error) {

	fmt.Println(" ############## Invoke Read User From Ledger ################")

	eventID := "getUserInvoke"
	queryCreatorOrg := s.User.Setup.OrgName
	collection := s.User.Setup.OrgCollection

	fmt.Println(" Email = " + email)

	fmt.Println(" Need History - " + strconv.FormatBool(needHistory))

	response, err := s.User.Setup.ExecuteChaincodeTranctionEvent(eventID, "invoke",
		[][]byte{
			[]byte("readUser"),
			[]byte(email),
			[]byte(collection),
			[]byte(eventID),
			[]byte(queryCreatorOrg),
			[]byte(strconv.FormatBool(needHistory)),
		}, s.User.Setup.ChaincodeId, s.User.ChannelClient, s.User.Event)

	if err != nil {
		return nil, fmt.Errorf("Error - Get User From Ledger : %s", err.Error())
	}

	if response == nil {
		return nil, fmt.Errorf("Error - No User found ")
	}

	var user *model.User

	err = json.Unmarshal(response.Payload, &user)
	if err != nil {
		return nil, fmt.Errorf("unable to convert response to the object given for the query: %v", err)
	}

	fmt.Println("#### User Found #### ")
	fmt.Println(" Email 	= " + user.Email)
	fmt.Println(" Mobile  	= " + user.Mobile)
	fmt.Println(" Age 		= " + user.Age)
	fmt.Println(" Salary 	= " + user.Salary)
	fmt.Println(" ################################## ")

	return user, nil
}
