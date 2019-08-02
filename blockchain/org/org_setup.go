package org

import (
	"fmt"
	"os"
	"strings"
	caMsp "github.com/hyperledger/fabric-sdk-go/pkg/client/msp"
	mspclient "github.com/hyperledger/fabric-sdk-go/pkg/client/msp"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/resmgmt"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/providers/msp"
	contextAPI "github.com/hyperledger/fabric-sdk-go/pkg/common/providers/context"
	fabAPI "github.com/hyperledger/fabric-sdk-go/pkg/common/providers/fab"
	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/event"
	"github.com/hyperledger/fabric-sdk-go/pkg/fab/resource"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"github.com/hyperledger/fabric-sdk-go/third_party/github.com/hyperledger/fabric/common/cauthdsl"
	cb "github.com/hyperledger/fabric-sdk-go/third_party/github.com/hyperledger/fabric/protos/common"
	"github.com/pkg/errors"

)	

const (
	appName  	= "privateledger"
	ccPath   	= "github.com/privateledger/chaincode/"
	ccId		= "CC_ORG_V00"
	ccVersion   = "1"
	ccPolicy	= "OR ('Org1MSP.member','Org2MSP.member','Org3MSP.member', 'Org4MSP.member')"
	ordererName = "OrdererOrg"
	ordererId	= "orderer.private.ledger.com"
)

type OrgSetup struct {

	// Network parameters
	OrdererID             		string
	OrdererAdmin          		string
	OrdererName           		string
	OrdererClientContext  		contextAPI.ClientProvider
	OrdererChannelContext 		contextAPI.ChannelProvider
	OrdererResmgmt        		*resmgmt.Client
	
	// Channel parameters
	ChannelID     		  string
	ChannelConfig 		  string

	// Chaincode parameters
	ChaincodeGoPath 	  string
	ChaincodePath   	  string
	ChaincodeId 		  string
	ChainCodeVersion 	  string
	ChainCodePolicy		  string	
	
	CCPkg    		  *resource.CCPackage
	

	ConfigFile       	  	string
	OrgCaID				string
	OrgName				string
	OrgAdmin			string
	UserName			string
	OrgCollection		  	string
	OrgCollectionPolicy		string
	OrgCollectionConfig   		*cb.CollectionConfig
	
	Sdk             	  	*fabsdk.FabricSDK
	CaClient        	  	*caMsp.Client
	Resmgmt				*resmgmt.Client
	Ctx				contextAPI.ClientProvider
	MspClient			*mspclient.Client
	Peers				[]fabAPI.Peer
	ChannelContext 		  	contextAPI.ChannelProvider
	ChannelClient   	  	*channel.Client
	Event 			  	*event.Client
	SigningIdentity		  	msp.SigningIdentity	
	SigningIdentities 	  	[]msp.SigningIdentity
}


func initializeOrg(obj OrgSetup) (*OrgSetup,error) {

		fmt.Println("Initialize " + obj.OrgName + " SDK...")

		sdk, err := fabsdk.New(config.FromFile(obj.ConfigFile))
		if err != nil {
			return nil, errors.WithMessage(err, "failed to create SDK")
		}		
		fmt.Println("  SDK created for " + obj.OrgName)


		caClient, err := caMsp.New(sdk.Context())
		if err != nil {
			return nil, fmt.Errorf("failed to create new CA client: %v", err)
		}		
		fmt.Println("  CA Client created for " + obj.OrgName)


		orgCtx := sdk.Context(
			fabsdk.WithUser(obj.OrgAdmin),
			fabsdk.WithOrg(obj.OrgName))
		fmt.Println("  Context created for " + obj.OrgName)


		resMgmtClient, err := resmgmt.New(orgCtx)
		if err != nil {
			return nil, errors.WithMessage(err, "failed to create resmgmt")
		}		
		fmt.Println("  Ressource management client created for " + obj.OrgName)


		mspClient, err := mspclient.New(sdk.Context(), mspclient.WithOrg(obj.OrgName))
		if err != nil {
			return nil, errors.WithMessage(err, "failed to create MSP client")
		}		
		fmt.Println("  MSP Client created for " + obj.OrgName)


		signingIdentity, err := mspClient.GetSigningIdentity(obj.OrgAdmin)
		if err != nil {
			return nil, errors.WithMessage(err, "failed to get admin signing identity")
		}
		fmt.Println("  Signing Identity created for " + obj.OrgName)

		
		if !strings.EqualFold(obj.OrgName, ordererName){

				
				signIdentities = append(signIdentities, signingIdentity)
							
				collCfg, err = newCollectionConfig(obj.OrgCollection,obj.OrgCollectionPolicy, collCfgRequiredPeerCount,collCfgMaximumPeerCount,collCfgBlockToLive)
				if err != nil {
					fmt.Println("failed to create collection config : "+err.Error())
					return nil, errors.WithMessage(err, "failed to create collection config ")
				}
				collConfigs = append(collConfigs, collCfg)

				orgPeers, err = DiscoverLocalPeers(orgCtx, 2)
				if err != nil {
					fmt.Errorf("  failed to Discover Local Peers: %v for ",obj.OrgName, err)
					return nil,nil
				}				
				fmt.Println("  Peers Discovered for " + obj.OrgName)

				channelCtx = sdk.ChannelContext(appName, 
									fabsdk.WithUser(obj.OrgAdmin), 
									fabsdk.WithOrg(obj.OrgName))
				fmt.Println(" Channel Client create for " + obj.OrgName)
		}


		signingIdentities := []msp.SigningIdentity{signingIdentity}

		//channelClient, event, _:= obj.CreateChannelClient(sdk, obj.OrgName, obj.OrgAdmin, caClient)

		return &OrgSetup {
			ConfigFile:       	  	obj.ConfigFile,
			ChannelID:       		appName,
			ChaincodeGoPath: 		os.Getenv("GOPATH"),
			ChaincodePath:   		ccPath,
			ChaincodeId: 			ccId,	 
			ChainCodeVersion: 		ccVersion,
			ChainCodePolicy:		ccPolicy,	
			OrdererName:			ordererName,
			OrdererID:			ordererId,
			ChannelClient:			nil,
			Event:				nil,
			OrgCaID:			obj.OrgCaID,
			OrgName:			obj.OrgName,
			OrgAdmin:			obj.OrgAdmin,
			OrgCollection:		  	obj.OrgCollection,
			OrgCollectionConfig:		collCfg,
			ChannelConfig:		  	getArtifactPath()+obj.ChannelConfig,
			Sdk: 			 	sdk,
			CaClient: 		 	caClient,
			Ctx: 			 	orgCtx,
			Resmgmt: 		 	resMgmtClient,
			MspClient: 		 	mspClient,
			SigningIdentities:		signingIdentities,
			Peers: 			 	orgPeers,
			ChannelContext:  		channelCtx,			
		}, nil
}


func getArtifactPath() string {
	return os.Getenv("GOPATH") + "/src/github.com/privateledger/fixtures/artifacts/"
}

func getSigningIdentities() []msp.SigningIdentity {
	fmt.Println(" getSigningIdentities == ", len(signIdentities))
	return signIdentities
}

func getCollectionConfigs() []*cb.CollectionConfig{
	return collConfigs
}


//#################################################################################//
// 						Private Collection Config Setup						      //	
//###############################################################################//


func(s *OrgSetup) SetupCollConfig() ([]*cb.CollectionConfig,error) {
	return getCollectionConfigs() , nil
}

func newCollectionConfig(colName, policy string, reqPeerCount, maxPeerCount int32, blockToLive uint64) (*cb.CollectionConfig, error) {
	
	p, err := cauthdsl.FromString(policy)

	if err != nil {
		fmt.Println("failed to create newCollectionConfig : "+err.Error())
		return nil, err
	}
	cpc := &cb.CollectionPolicyConfig{
		Payload: &cb.CollectionPolicyConfig_SignaturePolicy{
			SignaturePolicy: p,
		},
	}
	return &cb.CollectionConfig{
		Payload: &cb.CollectionConfig_StaticCollectionConfig{
			StaticCollectionConfig: &cb.StaticCollectionConfig{
				Name:              colName,
				MemberOrgsPolicy:  cpc,
				RequiredPeerCount: reqPeerCount,
				MaximumPeerCount:  maxPeerCount,
				BlockToLive:       blockToLive,
			},
		},
	}, nil
}
