package invoke

import (
	"privateledger/chaincode/model"
	"encoding/json"
	"fmt"
)

func (s *OrgInvoke) GetHistoryFromLedger(email string) ([]model.HistoryData, error) {

	fmt.Println(" ############## Invoke Get History ################")

	eventID := "getHistoryByEmail"

	historyData, err := s.User.Setup.ExecuteChaincodeTranctionEvent(eventID, "invoke",
		[][]byte{
			[]byte("readHistory"),
			[]byte(email),
			[]byte(eventID),
		}, s.User.Setup.ChaincodeId, s.User.ChannelClient, s.User.Event)

	if err != nil {
		return nil, fmt.Errorf("Error - Get History Data From Ledger : %s", err.Error())
	}

	fmt.Println(" ********** History Response Received ************** ")

	allHistoryData := make([]model.HistoryData, 0)

	if historyData != nil && historyData.Payload == nil {
		return nil, fmt.Errorf("unable to get response for the query: %v", err)
	}

	if historyData != nil {
		err = json.Unmarshal(historyData.Payload, &allHistoryData)
		if err != nil {
			return nil, fmt.Errorf("unable to convert response to the object given for the query: %v", err)
		}
	}

	if len(allHistoryData) < 1 {
		return nil, fmt.Errorf("No history records found")
	}

	fmt.Println("Total History for "+email, len(allHistoryData))

	/*for _, history := range allHistoryData {
		fmt.Println("History - "+history.QueryCreator+" -- "+history.QueryCreatorOrg+" -- "+history.Query+" -- "+history.Remarks)
	}*/

	return allHistoryData, nil
}
