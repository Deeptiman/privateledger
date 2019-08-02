package org

import (
	"fmt"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/resmgmt"
	"github.com/pkg/errors"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/errors/retry"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/errors/status"
)

// WaitForOrdererConfigUpdate waits until the config block update has been committed.
// In Fabric 1.0 there is a bug that panics the orderer if more than one config update is added to the same block.
// This function may be invoked after each config update as a workaround.
func WaitForOrdererConfigUpdate(client *resmgmt.Client, channelID string, genesis bool, lastConfigBlock uint64) (uint64, error) {

	fmt.Println("**** WaitForOrdererConfigUpdate ****")

	blockNum, err := retry.NewInvoker(retry.New(retry.TestRetryOpts)).Invoke(
		func() (interface{}, error) {
			chConfig, err := client.QueryConfigFromOrderer(channelID, resmgmt.WithOrdererEndpoint(Orderer.OrdererID))
			if err != nil {
				return nil, status.New(status.TestStatus, status.GenericTransient.ToInt32(), err.Error(), nil)
			}

			currentBlock := chConfig.BlockNumber()
			if currentBlock <= lastConfigBlock && !genesis {
				return nil, status.New(status.TestStatus, status.GenericTransient.ToInt32(), fmt.Sprintf("Block number was not incremented [%d, %d]", currentBlock, lastConfigBlock), nil)
			}
			return &currentBlock, nil
		},
	)

	if err != nil {
		return *blockNum.(*uint64), errors.WithMessage(err, "failed to get Orderer Config Update")
	}

	return *blockNum.(*uint64), nil
}