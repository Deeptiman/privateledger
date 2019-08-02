package org


import (
	"fmt"
	"bytes"
	"time"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/errors/retry"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/event"
	"github.com/pkg/errors"
)

func(s *OrgSetup) ExecuteChaincodeTranctionEvent(eventID, fcnName string, args [][]byte, chaincodeId string, channelClient  *channel.Client, ccEvent *event.Client) (*channel.Response, error) {

	fmt.Println(" ############# ExecuteChaincodeTranctionEvent - "+eventID+" ############## ")

	fmt.Println("  Execute Org -- "+s.OrgName)
	fmt.Println("  Execute CCId -- "+chaincodeId)

	if ccEvent == nil {
		fmt.Println(" ############### Event is Nil")
	}

    reg, notifier, err := ccEvent.RegisterChaincodeEvent(chaincodeId, eventID)
   
    if err != nil {
		return nil, fmt.Errorf("Blockchain ..... failed to register event: %v", err)
    }
    defer ccEvent.Unregister(reg)

	transientDataMap := make(map[string][]byte)
	transientDataMap["result"] = []byte("Transient data to invoke in the ledger")

	response, err := channelClient.Execute(channel.Request{
		
			ChaincodeID: chaincodeId, 
			Fcn: fcnName, 
			Args: args, 

	}, channel.WithRetry(retry.DefaultChannelOpts), channel.WithTargets(s.Peers[0],s.Peers[1]))

	if err != nil {
		fmt.Printf("failed to execute Invoke request: %s", err.Error())
		return nil, nil
	}

	if response.ChaincodeStatus == 0 {
		return nil, errors.WithMessage(nil, "Expected ChaincodeStatus")
	}

	if response.Responses[0].ChaincodeStatus != response.ChaincodeStatus {
		return nil, errors.WithMessage(nil, "Expected the chaincode status returned by successful Peer Endorsement to be same as Chaincode status for client response")
	}

	select {
		case ccEventNotify := <-notifier:
			fmt.Printf("Received CC event: %v\n", ccEventNotify)
		case <-time.After(time.Second * 60):
			return  nil, fmt.Errorf("did NOT receive CC event for eventId(%s)", eventID)
	}

	return &response, nil
}

func createKeyValuePairs(m map[string][]byte) string {
	b := new(bytes.Buffer)
	for key, value := range m {
		fmt.Fprintf(b, "%s=\"%s\"\n", key, value)
	}
	return b.String()
}

func CToGoString(c []byte) string {
    n := -1
    for i, b := range c {
        if b == 0 {
            break
        }
        n = i
    }
    return string(c[:n+1])
}