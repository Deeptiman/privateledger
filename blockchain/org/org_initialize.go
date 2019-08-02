package org

import (
	"fmt"
	"strings"
	fabAPI "github.com/hyperledger/fabric-sdk-go/pkg/common/providers/fab"
	contextAPI "github.com/hyperledger/fabric-sdk-go/pkg/common/providers/context"
	cb "github.com/hyperledger/fabric-sdk-go/third_party/github.com/hyperledger/fabric/protos/common"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/providers/msp"
	"github.com/pkg/errors"
)

const (
	collCfgBlockToLive       = 1000
	collCfgRequiredPeerCount = 0
	collCfgMaximumPeerCount  = 3
)

var Orderer OrgSetup
var OrgList []OrgSetup
var OrgNames []string
var totalOrg = 5

var orgPeers []fabAPI.Peer
var channelCtx contextAPI.ChannelProvider
var collCfg *cb.CollectionConfig

var signIdentities = make([]msp.SigningIdentity, 0, totalOrg-1)
var collConfigs = make([]*cb.CollectionConfig, 0, totalOrg-1)


func(s *OrgSetup) Init(processAll bool) error {

	OrgList = make([]OrgSetup, 0, totalOrg-1)
	OrgNames = []string{"org1","org2","org3","org4"}

	if processAll {
		s.InitializeAllOrgs()
	}

	return nil
}

func(s *OrgSetup) GetOrgNames() []string {
	return OrgNames
}

func(s *OrgSetup) FilteredOrgNames() []string {
	var filteredOrg []string
	for _, org := range OrgNames {

		if !strings.EqualFold(s.OrgName, org){
			filteredOrg = append(filteredOrg, org)
		}
	}
	return filteredOrg
}

func(s *OrgSetup) InitializeOrg(org string) (OrgSetup,error) {

	var obj OrgSetup

	switch name := org; name {
			
		case "org1":

			obj = OrgSetup {
				OrgAdmin: 				"Admin",
				OrgName:  				"org1",
				ConfigFile: 			"config-org1.yaml",
				OrgCaID: 				"ca.org1.private.ledger.com",
				ChannelConfig: 			"Org1MSPanchors.tx",
				OrgCollection: 			"collectionOrg1",
				OrgCollectionPolicy: 	"OR ('Org1MSP.member')",
			}
		
			break

		case "org2":

			obj = OrgSetup {
				OrgAdmin: 				"Admin",
				OrgName:  				"org2",
				ConfigFile: 			"config-org2.yaml",
				OrgCaID: 				"ca.org2.private.ledger.com",
				ChannelConfig: 			"Org2MSPanchors.tx",
				OrgCollection: 			"collectionOrg2",
				OrgCollectionPolicy: 	"OR ('Org2MSP.member')",
			}
			 
			break

		case "org3":

			obj = OrgSetup {
				OrgAdmin: 				"Admin",
				OrgName:  				"org3",
				ConfigFile: 			"config-org3.yaml",
				OrgCaID: 				"ca.org3.private.ledger.com",
				ChannelConfig: 			"Org3MSPanchors.tx",
				OrgCollection: 			"collectionOrg3",
				OrgCollectionPolicy: 	"OR ('Org3MSP.member')",
			}
			 
			break

		case "org4":

			obj = OrgSetup {
				OrgAdmin: 				"Admin",
				OrgName:  				"org4",
				ConfigFile: 			"config-org4.yaml",
				OrgCaID: 				"ca.org4.private.ledger.com",
				ChannelConfig: 			"Org4MSPanchors.tx",
				OrgCollection: 			"collectionOrg4",
				OrgCollectionPolicy: 	"OR ('Org4MSP.member')",
			} 

			break	 
	}
	
	orgSetup, err := InitiateOrg(obj)
	if err != nil {
		return OrgSetup{}, errors.WithMessage(err, " failed to initiate Org")
	}
				
	return orgSetup, nil

}

func InitiateOrderer() (OrgSetup,error) {

	obj := OrgSetup {
		OrgAdmin: 				"Admin",
		OrgName:  				"OrdererOrg",
		ConfigFile: 			"config-org1.yaml",
		ChannelConfig: 			"privateledger.channel.tx",
	}

	orderer, err  := initializeOrg(obj)

	if err != nil {
		return OrgSetup{}, fmt.Errorf("  failed to setup Org - "+obj.OrgName+" - "+err.Error())
	}

	if orderer == nil {
		return OrgSetup{}, fmt.Errorf("  failed to setup Org - "+obj.OrgName)
	}

	Orderer = *orderer

	Orderer.OrdererID = "orderer.private.ledger.com"

	fmt.Println(" **** Setup Created for "+Orderer.OrgName+" **** ")

	return Orderer, nil
}


func(s *OrgSetup) InitializeAllOrgs() error {

	ordererSetup, err := InitiateOrderer()
	if err != nil {
		return errors.WithMessage(err, " failed to initiate Orderer")
	}
	OrgList = append(OrgList, ordererSetup)

	for _, org := range OrgNames {

		orgSetup, err := s.InitializeOrg(org)
		if err != nil {
			return errors.WithMessage(err, " failed to initiate Org - "+org)
		}
		OrgList = append(OrgList, orgSetup)
	}

	Orderer.SigningIdentities = getSigningIdentities()
	OrgList[0] = Orderer

	return nil
}

func InitiateOrg(obj OrgSetup) (OrgSetup,error) {

	org, err := initializeOrg(obj)
	if err != nil {
		return OrgSetup{}, fmt.Errorf("  failed to setup Org - "+obj.OrgName+" - "+err.Error())
	}

	if org == nil {
		return OrgSetup{}, fmt.Errorf("  failed to setup Org - "+obj.OrgName)
	}
	
	fmt.Println(" **** Setup Created for "+org.OrgName+" **** ")

	return *org, nil
}
 