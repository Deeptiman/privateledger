package org

import (
	"fmt"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/providers/context"
	fabAPI "github.com/hyperledger/fabric-sdk-go/pkg/common/providers/fab"	
	contextImpl "github.com/hyperledger/fabric-sdk-go/pkg/context"
	"github.com/pkg/errors"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/errors/status"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/errors/retry"

)

func DiscoverLocalPeers(ctxProvider context.ClientProvider, expectedPeers int) ([]fabAPI.Peer, error) {

	ctx, err := contextImpl.NewLocal(ctxProvider)
	if err != nil {
		fmt.Println("    error creating local context :  %s "+err.Error())
		return nil, errors.Wrap(err, "error creating local context")
	}

	discoveredPeers, err := retry.NewInvoker(retry.New(retry.TestRetryOpts)).Invoke(

		func() (interface{}, error) {

			peers, serviceErr := ctx.LocalDiscoveryService().GetPeers()

			if serviceErr != nil {
				fmt.Println("    error getting peers for MSP :  %s "+err.Error())
				return nil, errors.Wrapf(serviceErr, "error getting peers for MSP [%s]", ctx.Identifier().MSPID)
			}
			
			fmt.Println("  MSP ID -- "+ctx.Identifier().MSPID,expectedPeers, len(peers))

			if len(peers) < expectedPeers {
				return nil, status.New(status.TestStatus, status.GenericTransient.ToInt32(), fmt.Sprintf("Expecting %d peers but got %d", expectedPeers, len(peers)), nil)
			}
			return peers, nil
		},
	)

	if err != nil {
		return nil, err
	}

	return discoveredPeers.([]fabAPI.Peer), nil
}


func LoadOrgPeers(ctxProvider context.ClientProvider) error {

	fmt.Println("    LoadOrgPeers")

	_, err := contextImpl.NewLocal(ctxProvider)

	if err != nil {		
		fmt.Println("    context creation failed : %s ", err.Error())
		return errors.WithMessage(err, "context creation failed")
	}

	 
	return nil
}