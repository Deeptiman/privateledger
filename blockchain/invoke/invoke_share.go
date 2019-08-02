package invoke

import (
	"fmt"
	"sync"
	"strconv"
	"github.com/privateledger/chaincode/model"
	"github.com/privateledger/blockchain/org"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/event"

)

func(s *OrgInvoke) ShareDataToOrg(email string, orgList []string, accessList []int32, targets string) error {

	fmt.Println(" ############## Invoke Share Data ################")

	userData, err := s.GetUserFromLedger(email, true)

	if err != nil {
		return fmt.Errorf("unable to get user for the query: %v", err)
	}
	
	threads := len(orgList)	
	var wg sync.WaitGroup

	respond := make(chan string, threads)	
	wg.Add(threads)

	for i, orgName := range orgList {

		fmt.Println(" *********** Sharing Orgs === "+orgName)

		orgSetup := s.User.Setup.ChooseORG(orgName)
		orgName 			:= 	orgSetup.OrgName
		orgSdk				:=  orgSetup.Sdk
		orgAdmin			:=  orgSetup.OrgAdmin
		caClient 			:= 	orgSetup.CaClient
		channelClient, event,_ := s.User.Setup.CreateChannelClient(orgSdk, orgName, orgAdmin, caClient)
	
		s.shareData(respond, &wg, orgSetup, email, orgName, userData, accessList[i], targets, channelClient, event)
	}
		
	wg.Wait()
	close(respond)

	for queryResp := range respond {
		fmt.Println("Share Response: "+ queryResp)
	}


	fmt.Println(" ####### Update in Owner Collection - "+userData.Owner+" ######### ")	

	appFcnName := "updateTarget"
	eventID := "updateTargetInvoke"
	channelClient	:=  s.User.Setup.ChannelClient
	event			:=  s.User.Setup.Event

	fmt.Println(" ownerCollection - "+s.User.Setup.OrgCollection)

	_, err = s.User.Setup.ExecuteChaincodeTranctionEvent(eventID, "invoke",
		[][]byte{
			[]byte(appFcnName),
			[]byte(email),
			[]byte(targets),
			[]byte(s.User.Setup.OrgCollection),
			[]byte(eventID),
		}, s.User.Setup.ChaincodeId, channelClient,event)

	if err != nil {
	 	fmt.Errorf("Error - Share User Data From Ledger : %s", err.Error())
	}	

	fmt.Println(" ###################################################### ")


	return nil
}


func(s *OrgInvoke) shareData(respond chan<- string, wg *sync.WaitGroup,  orgSetup *org.OrgSetup, email, sharingOrg string, userData *model.User, access int32, targets string, channelClient *channel.Client, ccEvent *event.Client) {

	owner := s.User.Setup.OrgName	
	queryCreatorOrg := s.User.Setup.OrgName
	queryCreatorRole := s.Role
	shareAccess := fmt.Sprint(access)
	needHistory := strconv.FormatBool(true)

	queryTxnHash, err := s.GetSecretMessage(owner, sharingOrg)

	if err != nil {
		fmt.Errorf("Invalid Hash, unable to invoke share query for - "+orgSetup.OrgName, err)
	}

	fmt.Println(" ###### Sharing Data from "+userData.Owner+" to "+orgSetup.OrgName+" ####### ")

	fmt.Println("	Collection : "+orgSetup.OrgCollection)

		appFcnName := "shareUser"
		eventID := "shareUserInvoke"
				
		_, err = orgSetup.ExecuteChaincodeTranctionEvent(eventID, "invoke",

			[][]byte{
				
				[]byte(appFcnName),

				[]byte(userData.Name),
				[]byte(userData.Email),					
				[]byte(userData.Mobile),
				[]byte(userData.Age),
				[]byte(userData.Salary),				
				[]byte(userData.Role),	
				[]byte(eventID),				

				[]byte(shareAccess),
				[]byte(queryTxnHash),

				[]byte(targets),
				[]byte(orgSetup.OrgCollection),
				[]byte(sharingOrg),
				[]byte(queryCreatorOrg),
				[]byte(queryCreatorRole),
				[]byte(needHistory),

			}, orgSetup.ChaincodeId, channelClient, ccEvent)

		if err != nil {
			fmt.Errorf("Error - Share User Data From Ledger : %s", err.Error())
		}
	
	defer wg.Done()

	respond <- fmt.Sprintf("%s responded to share query: %s", orgSetup.OrgName, orgSetup.OrgCollection)

	fmt.Println(" ###################################################### ")
}