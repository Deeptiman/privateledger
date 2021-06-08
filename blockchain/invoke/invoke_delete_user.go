package invoke

import (
	"privateledger/blockchain/org"
	"privateledger/chaincode/model"
	"fmt"
	"strconv"
	"strings"
	"sync"

	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/event"
)

func (s *OrgInvoke) DeleteUserFromLedger(email, targets, role string) error {

	fmt.Println(" ############## Invoke Delete User ################")

	if !s.OrgHasAccess(targets,
		model.LedgerAccess(model.DELETE).Int(),
		s.User.Setup.OrgName) {
		return fmt.Errorf("org has no delete access")
	}

	orgList, _ := s.GetAccessOrgList(targets)
	user, _ := s.GetUserFromLedger(email, false)

	/*
		########################  Access Override #########################################
			Query Initiator Org will override the access privillege for other participated
			Organizations, and in the remark the query initator org will claim performing
			the query.

			This scenario will appear, if participated Org during a transaction don't have
			access to perform the query

			A transaction hash stored in both Org collection will verfiy the invoked query
	*/
	queryAccess := strconv.Itoa(int(model.DELETE))

	// ###############################################################################//

	fmt.Println(" ######## Delete User Data ##########")

	threads := len(orgList)
	var wg sync.WaitGroup

	respond := make(chan string, threads)
	wg.Add(threads)

	for _, orgName := range orgList {

		fmt.Println(" Particiapted Org - " + orgName)

		orgSetup := s.User.Setup.ChooseORG(strings.ToLower(orgName))
		orgSdk := orgSetup.Sdk
		orgAdmin := orgSetup.OrgAdmin
		caClient := orgSetup.CaClient
		channelClient, event, _ := orgSetup.CreateChannelClient(orgSdk, orgName, orgAdmin, caClient)

		queryTxnHash, err := s.GetSecretMessage(user.Owner, orgName)

		if err != nil {
			return fmt.Errorf("Invalid Hash, unable to invoke delete query for - "+orgName, err)
		}

		go s.deleteUserData(respond, &wg, orgSetup, email, role, queryAccess, queryTxnHash, channelClient, event)
	}

	wg.Wait()
	close(respond)

	for queryResp := range respond {
		fmt.Println("Update Response: " + queryResp)
	}

	return nil
}

func (s *OrgInvoke) deleteUserData(respond chan<- string, wg *sync.WaitGroup, orgSetup *org.OrgSetup, email, role, queryAccess, queryTxnHash string, channelClient *channel.Client, ccEvent *event.Client) {

	eventID := "deleteUserDataInvoke"
	queryCreatorOrg := s.User.Setup.OrgName
	queryCreatorRole := s.Role
	targetOrg := orgSetup.OrgName
	needHistory := strconv.FormatBool(true)

	fmt.Println(" ########### Invoke Update Query Details ########### ")
	fmt.Println("	queryCreatorOrg - " + queryCreatorOrg)
	fmt.Println("   queryCreatorRole - " + queryCreatorRole)
	fmt.Println("	queryAccess - " + queryAccess)
	fmt.Println("	queryTxnHash - " + queryTxnHash)
	fmt.Println("	Collection - " + orgSetup.OrgCollection)
	fmt.Println(" ##################################### ")

	_, err := orgSetup.ExecuteChaincodeTranctionEvent(eventID, "invoke",
		[][]byte{
			[]byte("deleteUser"),
			[]byte(email),
			[]byte(orgSetup.OrgCollection),
			[]byte(eventID),

			[]byte(queryAccess),
			[]byte(queryTxnHash),

			[]byte(role),
			[]byte(queryCreatorOrg),
			[]byte(queryCreatorRole),
			[]byte(targetOrg),
			[]byte(needHistory),
		}, orgSetup.ChaincodeId, orgSetup.ChannelClient, orgSetup.Event)

	if err != nil {
		fmt.Errorf("Error - DeleteUserFromLedger : %s", err.Error())
	}

	defer wg.Done()

	respond <- fmt.Sprintf("%s responded to delete query: %s", queryCreatorOrg, orgSetup.OrgCollection)
}
