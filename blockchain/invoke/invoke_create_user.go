package invoke

import (
	"privateledger/blockchain/org"
	"privateledger/chaincode/model"
	"fmt"
	"strconv"
)

type OrgInvoke struct {
	User *org.OrgUser
	Role string
}

func (s *OrgInvoke) InvokeCreateUser(name, age, mobile, salary string) error {

	fmt.Println(" ############## Invoke Create User ################")

	var queryCreatorOrg string

	queryCreatorOrg = s.User.Setup.OrgName
	email := s.User.Username
	collection := s.User.Setup.OrgCollection
	eventID := "userInvoke"
	needHistory := strconv.FormatBool(true)

	var accessList []int32
	var orgList []string

	for _, org := range s.User.Setup.FilteredOrgNames() {

		access := model.LedgerAccess(model.NOACCESS).Int()
		accessList = append(accessList, access)
		orgList = append(orgList, org)
	}

	targets, err := s.CreateOrgTargets(email, accessList, orgList)
	if err != nil {
		targets = ""
	}

	fmt.Println(" ###### Create Data Parameters ###### ")
	fmt.Println("	Email 			= " + email)
	fmt.Println(" 	Name 			= " + name)
	fmt.Println(" 	Mobile 			= " + mobile)
	fmt.Println(" 	Age 			= " + age)
	fmt.Println(" 	Salary 			= " + salary)
	fmt.Println(" 	Owner 			= " + queryCreatorOrg)
	fmt.Println(" ################################## ")

	_, err = s.User.Setup.ExecuteChaincodeTranctionEvent(eventID, "invoke",
		[][]byte{
			[]byte("createUser"),

			[]byte(name),
			[]byte(email),
			[]byte(mobile),
			[]byte(age),
			[]byte(salary),
			[]byte(targets),
			[]byte(collection),
			[]byte(eventID),
			[]byte(queryCreatorOrg),
			[]byte(needHistory),
		}, s.User.Setup.ChaincodeId, s.User.ChannelClient, s.User.Event)

	if err != nil {
		return fmt.Errorf("Error - addUserToLedger : %s", err.Error())
	}

	fmt.Println("#### User added Successfully ####")

	return nil
}
