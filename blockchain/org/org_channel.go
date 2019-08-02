package org

import (
	"fmt"
	"strings"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/resmgmt"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/errors/retry"
	fabAPI "github.com/hyperledger/fabric-sdk-go/pkg/common/providers/fab"
	"github.com/pkg/errors"
)

func(s *OrgSetup) CreateChannelForOrg() error {

	fmt.Println("Create Channel for Org - "+s.OrgName, len(s.OrgName))

	if len(s.OrgName) == 0 {
		return errors.WithMessage(nil, " empty Org Name")
	}

	var OrgJoined bool
	var err error
	
	if !strings.EqualFold(s.OrgName, s.OrdererName){

		OrgJoined, err = s.IsJoinedChannel(s.Resmgmt, s.Peers[0])
		if err != nil {
			fmt.Println("failed to check isJoin channel")
			return errors.WithMessage(err, "  failed to check isJoin channel")
		}
	}

	if !OrgJoined {

		fmt.Println("Create Channel == ChannelID = "+s.ChannelID)
		fmt.Println("Create Channel == OrdererID = "+Orderer.OrdererID)
		fmt.Println("Create Channel == channelConfigPath = "+s.ChannelConfig)
		fmt.Println("Create Channel == SigningIdentities = ", len(s.SigningIdentities))
	
		req := resmgmt.SaveChannelRequest{
			ChannelID: s.ChannelID, 
			ChannelConfigPath: s.ChannelConfig, 
			SigningIdentities: s.SigningIdentities,
		}
		
		txID, err := s.Resmgmt.SaveChannel(
			req, resmgmt.WithOrdererEndpoint(Orderer.OrdererID))
		
		if err != nil || txID.TransactionID == "" {
			return errors.WithMessage(err, "failed to save anchor channel for - "+s.OrgName)
		}

		if err != nil {
			return errors.WithMessage(err, "failed to save channel")
		}

		var lastConfigBlock uint64
		lastConfigBlock, err = WaitForOrdererConfigUpdate(s.Resmgmt, s.ChannelID, true, lastConfigBlock)

		if err != nil {
			return errors.WithMessage(err, "failed to get Orderer config update")
		}

		fmt.Printf("Channel Orderer Config Update %lld ", lastConfigBlock)

	} else {
		fmt.Println(" Peers Already Joined channel")
	}

	return nil
}


func(s *OrgSetup) JoinChannelForOrg() error {

	fmt.Println("Join Channel for Org - "+s.OrgName, len(s.OrgName))

	if len(s.OrgName) == 0 {
		return errors.WithMessage(nil, " empty Org Name")
	}

	var OrgJoined bool
	var err error
	
	if !strings.EqualFold(s.OrgName, s.OrdererName){

		OrgJoined, err = s.IsJoinedChannel(s.Resmgmt, s.Peers[0])
		if err != nil {
			fmt.Println("failed to check isJoin channel")
			return errors.WithMessage(err, "  failed to check isJoin channel")
		}
	}

	if !OrgJoined {

		fmt.Println("Initiating Join Channel - "+s.OrgName)

			fmt.Println("  JoinChannel "+s.ChannelID)
				
			if err := s.Resmgmt.JoinChannel(s.ChannelID, resmgmt.WithRetry(retry.DefaultResMgmtOpts), resmgmt.WithOrdererEndpoint(Orderer.OrdererID)); err != nil {
				return errors.WithMessage(err, "failed to make admin join channel")
			}

		fmt.Println("  Successfully Joined the Channel "+s.ChannelID)

	} else {

		fmt.Println(" Peers Already Joined channel")
	}

	return nil
}


func(s *OrgSetup) IsJoinedChannel(orgResmgmt *resmgmt.Client,peer fabAPI.Peer) (bool, error) {

	resp, err := orgResmgmt.QueryChannels(resmgmt.WithTargets(peer))
	if err != nil {
		fmt.Println("IsJoinedChannel : failed to Query >>> "+err.Error())
		return false, err
	}
	for _, chInfo := range resp.Channels {
		fmt.Println("IsJoinedChannel : "+chInfo.ChannelId+" --- "+s.ChannelID)
		if chInfo.ChannelId == s.ChannelID {
			return true, nil
		}
	}
	return false, nil
}
