package invoke

import (
	"fmt"
	"sync"
	"strings"
	"strconv"
	"github.com/privateledger/chaincode/model"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/event"
	"github.com/privateledger/blockchain/org"
)


func(s *OrgInvoke) UpdateUserFromLedger(email , name, mobile, age, salary, targets, role string) error {

	fmt.Println(" ############## Invoke Update Data ################")
	
	if !s.OrgHasAccess(targets, 
		model.LedgerAccess(model.WRITE).Int(),
		s.User.Setup.OrgName){
		return fmt.Errorf("org has no write access")
	}

	orgList, _ := s.GetAccessOrgList(targets)
	user, _ := s.GetUserFromLedger(email, false)

	updateUser := &model.User {
		Name: 			name, 
		Email: 			email, 
		Mobile: 		mobile, 
		Age: 			age, 
		Salary: 		salary,
	}

	/*	
	########################  Access Override #########################################
		Query Initiator Org will override the access privillege for other participated
		Organizations, and in the remark the query initator org will claim performing 
		the query. 

		This scenario will appear, if participated Org during a transaction don't have 
		access to perform the query
		
		A transaction hash stored in both Org collection will verfiy the invoked query
	*/

		queryAccess := strconv.Itoa(int(model.WRITE))

	// ###############################################################################//

		fmt.Println(" ######## Update User Data ##########")	

		threads := len(orgList)	
		var wg sync.WaitGroup

		respond := make(chan string, threads)	
		wg.Add(threads)

		for _, orgName := range orgList {

			fmt.Println(" Particiapted Org - "+orgName)

			orgSetup := s.User.Setup.ChooseORG(strings.ToLower(orgName))
			orgSdk			:=  orgSetup.Sdk
			orgAdmin		:=  orgSetup.OrgAdmin
			caClient 		:= 	orgSetup.CaClient		
			channelClient, event, _:=  orgSetup.CreateChannelClient(orgSdk, orgName, orgAdmin, caClient)

			queryTxnHash, err := s.GetSecretMessage(user.Owner, orgName)

			if err != nil {
				fmt.Errorf("Invalid Hash, unable to invoke update query for - "+orgName, err)
			} else {
			
				go s.updateUserData(respond, &wg, orgSetup, updateUser, queryAccess, queryTxnHash, channelClient, event)					
			}
		}

		wg.Wait()
		close(respond)

		for queryResp := range respond {
			fmt.Println("Update Response: "+ queryResp)
		}
		

	fmt.Println(" ###################################################### ")

	return nil
}


func(s *OrgInvoke) updateUserData(respond chan<- string, wg *sync.WaitGroup, orgSetup *org.OrgSetup, user *model.User, queryAccess, queryTxnHash string, channelClient *channel.Client, ccEvent *event.Client) {

	name := user.Name
	email := user.Email
	mobile := user.Mobile
	age := user.Age
	salary := user.Salary
	
	eventID := "updateUserDataInvoke"		
	queryCreatorOrg := s.User.Setup.OrgName
	queryCreatorRole := s.Role
	targetOrg := orgSetup.OrgName
	needHistory := strconv.FormatBool(true)


	fmt.Println(" ########### Invoke Update Query Details ########### ")
	fmt.Println("	queryCreatorOrg - "+queryCreatorOrg)
	fmt.Println("   queryCreatorRole - "+queryCreatorRole)
	fmt.Println("	queryAccess - "+queryAccess)
	fmt.Println("	queryTxnHash - "+queryTxnHash)
	fmt.Println("	Collection - "+orgSetup.OrgCollection)
	fmt.Println(" ##################################### ")

	fmt.Println(" ###### Update Data Parameters ###### ")
	fmt.Println(" Email			= "+email)
	fmt.Println(" Name 			= "+name)
	fmt.Println(" Mobile 		= "+mobile)
	fmt.Println(" Age			= "+age)
	fmt.Println(" Salary 		= "+salary)
	fmt.Println(" ################################## ")


	_, err := orgSetup.ExecuteChaincodeTranctionEvent(eventID, "invoke",
		[][]byte{

			[]byte("updateUserData"),
			
			[]byte(name),
			[]byte(email),
			[]byte(mobile),
			[]byte(age),
			[]byte(salary),

			[]byte(orgSetup.OrgCollection),
			[]byte(queryAccess),
			[]byte(queryTxnHash),

			[]byte(eventID),
			[]byte(queryCreatorOrg),
			[]byte(queryCreatorRole),
			[]byte(targetOrg),
			[]byte(needHistory),

		}, orgSetup.ChaincodeId, orgSetup.ChannelClient, orgSetup.Event)

	if err != nil {
		fmt.Errorf(" Error - Update User Data From Ledger : %s ",err.Error())
	} 

	defer wg.Done()

	respond <- fmt.Sprintf("%s responded to update query: %s", queryCreatorOrg, orgSetup.OrgCollection)

}