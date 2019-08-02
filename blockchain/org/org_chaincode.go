package org

import (
	"fmt"
	"strings"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/resmgmt"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/errors/retry"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/errors/status"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/providers/fab"
	"github.com/hyperledger/fabric-sdk-go/third_party/github.com/hyperledger/fabric/common/cauthdsl"
	"github.com/pkg/errors"
	"github.com/hyperledger/fabric-sdk-go/pkg/fab/resource"
	packager "github.com/hyperledger/fabric-sdk-go/pkg/fab/ccpackager/gopackager"
	
)


func(s *OrgSetup) CreateCCPackage() (*resource.CCPackage,error) {
	fmt.Println(" Creating Chaincode Package ")
	fmt.Println(" 	- ChaincodePath "+s.ChaincodePath)
	fmt.Println(" 	- ChaincodeGoPath "+s.ChaincodeGoPath)
	ccPkg, err := packager.NewCCPackage(s.ChaincodePath, s.ChaincodeGoPath)
	if err != nil {
		return nil,errors.WithMessage(err, "failed to create chaincode package")
	}	 
	s.CCPkg = ccPkg
	return ccPkg, nil
}

//#################################################################################//
// 							Install Chaincode for Orgs						  //	
//###############################################################################//

func(s *OrgSetup) InstallCCForOrg(ccPkg *resource.CCPackage) error {

	if strings.EqualFold(s.OrgName,s.OrdererName) || len(s.OrgName) == 0{
		return nil	
	}

	fmt.Println("Initiating Install Chaincode For - "+s.OrgName)

	// Ensure that Gossip has propagated it's view of local peers before invoking
	// install since some peers may be missed if we call InstallCC too early
	orgPeers, err := DiscoverLocalPeers(s.Ctx, 2)
	if err != nil {
		return errors.WithMessage(err, "failed to Discover Local Peers for "+s.OrgName)
	}
	s.Peers = orgPeers
	fmt.Println("  Peers Discovered for " + s.OrgName)

	fmt.Println("\n  Installing Chaincode for "+s.OrgName)

	fmt.Println("Install CC ")
	fmt.Println("CC Name -  "+s.ChaincodeId)
	fmt.Println("CC Version -  ",s.ChainCodeVersion)
	fmt.Println("CC ChaincodePath - "+s.ChaincodePath)

	req := resmgmt.InstallCCRequest{
		Name:    s.ChaincodeId,
		Path:    s.ChaincodePath,
		Version: s.ChainCodeVersion,
		Package: ccPkg,
	}

	_, err = s.Resmgmt.InstallCC(req,
		resmgmt.WithRetry(retry.DefaultResMgmtOpts))

	if err != nil {
		fmt.Println("failed to install chaincode : "+err.Error())
		return errors.WithMessage(err, "  failed to install chaincode")
	}

		if err != nil {
			fmt.Println("failed to install chaincode : "+err.Error())
			return errors.WithMessage(err, "  failed to install chaincode for "+s.OrgName)
		}
	fmt.Println("  Chaincode installed for "+s.OrgName)

	return nil
}


//#################################################################################//
// 							Instantiate Chaincode for Orgs						  //	
//###############################################################################//


func(s *OrgSetup) InstantiateCCForOrg(orgPeers []fab.Peer) error {

	if strings.EqualFold(s.OrgName,s.OrdererName) || len(s.OrgName) == 0{
		return nil	
	}

	fmt.Println("\n  Instantiating Chaincode for - "+s.OrgName)

	fmt.Println("InstantiateCC  ")
	fmt.Println("InstantiateCC CC Policy - "+s.ChainCodePolicy)
	fmt.Println("InstantiateCC CC Name -  "+s.ChaincodeId)
	fmt.Println("InstantiateCC CC ChaincodePath - "+s.ChaincodePath)

	ccPolicy, err := cauthdsl.FromString(s.ChainCodePolicy)

	if err != nil{
		fmt.Println("failed policy : "+err.Error())
		return errors.WithMessage(err,"failed policy")
	}
	
	cfg, _ := s.SetupCollConfig()
	
	if cfg == nil {
		return errors.WithMessage(err, " Collection is nil")
	}
 
	resp, err := s.Resmgmt.InstantiateCC(
		s.ChannelID,
		resmgmt.InstantiateCCRequest{
			
			Name: 	 	s.ChaincodeId,
			Path:    	s.ChaincodePath,
			Version: 	s.ChainCodeVersion,
			Args:    	[][]byte{[]byte("init")},
			Policy:  	ccPolicy,
			CollConfig: cfg,

	},resmgmt.WithRetry(retry.DefaultResMgmtOpts), resmgmt.WithTargets(orgPeers[0], orgPeers[1]))


	if err != nil || resp.TransactionID == "" {
		fmt.Println("failed to instantiate the chaincode : "+err.Error())
		return errors.WithMessage(err, "  failed to instantiate the chaincode")
	}
	if err != nil {
		return errors.WithMessage(err, "  failed to instantiate the chaincode for "+s.OrgName)		
	}

	fmt.Println("   Chaincode is instantiated for Org1 & Org2 successfully")

	return nil
}


//#################################################################################//
// 							Upgrade Chaincode for Orgs						      //	
//###############################################################################//


func(s *OrgSetup) UpgradeCCForOrg(orgPeers []fab.Peer) error {

	fmt.Println("UpgradeCC  ")
	fmt.Println("Upgrade CC Policy - "+s.ChainCodePolicy)
	fmt.Println("Upgrade CC Name -  "+s.ChaincodeId)
	fmt.Println("Upgrade CC Version -  "+s.ChainCodeVersion)
	fmt.Println("Upgrade CC ChaincodePath - "+s.ChaincodePath)

	ccPolicy, err := cauthdsl.FromString(s.ChainCodePolicy)

	if err != nil{
		return errors.WithMessage(err,"failed policy")
	}

	cfg, _ := s.SetupCollConfig()
	
	if cfg == nil {
		return errors.WithMessage(err, " Collection is nil")
	}
	
	req := resmgmt.UpgradeCCRequest{
		Name: 		s.ChaincodeId, 
		Version: 	s.ChainCodeVersion, 
		Path: 		s.ChaincodePath, 
		Args:  		[][]byte{[]byte("init")},
		Policy: 	ccPolicy,
		CollConfig: cfg,		
	}

	resp, err := s.Resmgmt.UpgradeCC(s.ChannelID, req, resmgmt.WithRetry(retry.DefaultResMgmtOpts),resmgmt.WithTargets(orgPeers[0], orgPeers[1]))

	if err != nil {
		return errors.WithMessage(err, " >>>> failed to upgrade chaincode")
	}

	if resp.TransactionID == "" {
		return errors.WithMessage(nil, " ***** failed to upgrade chaincode, no transaction ID")
	}

	return nil
}





//#################################################################################//
// 				Query Installed & Instantiated Chaincode						  //	
//###############################################################################//

func QueryInstalledCC(orgID string, resMgmt *resmgmt.Client,
	ccName, ccVersion string, peers []fab.Peer) (bool, error) {

	installed, err := retry.NewInvoker(retry.New(retry.TestRetryOpts)).Invoke(

		func() (interface{}, error) {

			ok, _ := isCCInstalled(orgID, resMgmt, ccName, ccVersion, peers)
			if !ok {
				return &ok, status.New(status.TestStatus,
					status.GenericTransient.ToInt32(),
					fmt.Sprintf("  Chaincode [%s:%s] is not installed on all peers for [%s]", ccName, ccVersion, orgID), nil)
			}
			return &ok, nil
		},
	)

	if err != nil {
		return false, errors.WithMessage(err, "  Got error checking if chaincode was installed")
	}

	return *(installed).(*bool), nil
}


func isCCInstalled(orgID string, resMgmt *resmgmt.Client,
	ccName, ccVersion string, peers []fab.Peer) (bool, error) {

	fmt.Println("\n  Querying "+orgID+" peers to see if chaincode was installed")

	installedOnAllPeers := true

	for _, peer := range peers {

		fmt.Println("\n   Querying ..."+ peer.URL())
		resp, err := resMgmt.QueryInstalledChaincodes(resmgmt.WithTargets(peer))

		if err != nil {
			return false, errors.WithMessage(err, "  QueryInstalledChaincodes for peer [%s] failed : "+peer.URL())
		}
		found := false

		for _, ccInfo := range resp.Chaincodes {
			fmt.Println("   "+orgID+" found chaincode "+ccInfo.Name+" --- "+ccName+ " with version "+ ccInfo.Version+" -- "+ccVersion)
			if ccInfo.Name == ccName && ccInfo.Version == ccVersion {
				found = true
				break
			}
		}

		if !found {
			fmt.Println("   "+orgID+" chaincode is not installed on peer "+ peer.URL())
			installedOnAllPeers = false
		}
	}

	return installedOnAllPeers, nil
}


func(s *OrgSetup) QueryInstalledCCForOrg() error {

	if strings.EqualFold(s.OrgName,s.OrdererName) || len(s.OrgName) == 0{
		return nil	
	}
	
	fmt.Println(" Checking CC Installed in Org - "+s.OrgName)

	orgPeers, err := DiscoverLocalPeers(s.Ctx, 2)
	if err != nil {
		return errors.WithMessage(err, "failed to Discover Local Peers for "+s.OrgName)
	}
	s.Peers = orgPeers
	fmt.Println("  Peers Discovered for " + s.OrgName)

	OrgCCInstalled, err := QueryInstalledCC(
		s.OrgName, 
		s.Resmgmt, 
		s.ChaincodeId, 
		s.ChainCodeVersion, 
		s.Peers)

	if err != nil {
		errors.WithMessage(err, "  Got error checking if chaincode was installed for "+s.OrgName)
	}	

	if OrgCCInstalled {
		fmt.Println("\n  Chaincode is installed on all peers in "+s.OrgName)
	} else {
		return errors.WithMessage(nil, "  Chaincode is not installed for "+s.OrgName)
	}

	return nil
}


func(s *OrgSetup) QueryInstantiatedCCForOrg() error {

	if strings.EqualFold(s.OrgName,s.OrdererName) || len(s.OrgName) == 0{
		return nil	
	}
	
	fmt.Println("Checking Chaincode Instantiated in Org - "+s.OrgName+" , "+s.ChainCodeVersion)

	OrgCCfound, err := QueryInstantiatedCC(s.OrgName, 
		s.Resmgmt, s.ChannelID,
		s.ChaincodeId, s.ChainCodeVersion, s.Peers)

	if err != nil {
		return errors.WithMessage(err, "  failed to queryInstantiatedCC for "+s.OrgName)
	}

	if !OrgCCfound {
		return errors.WithMessage(err, "  failed to find instantiated chaincode [%s:%s] in at least one peer in "+s.OrgName+" on channel [%s] "+s.ChainCodeVersion+" , "+s.ChannelID)
	}

	fmt.Println("  Queried instantiated Chaincode for "+s.OrgName+" successfully")

	return nil
}

func QueryInstantiatedCC(orgID string, resMgmt *resmgmt.Client, channelID, ccName, ccVersion string, peers []fab.Peer) (bool,error){

	if len(peers) < 0 {
		return false, errors.WithMessage(nil, "  Expecting one or more peers")
	}

	fmt.Println("   Querying "+orgID+" peers to see if chaincode was instantiated on channel - "+channelID)

	instantiated, err := retry.NewInvoker(retry.New(retry.TestRetryOpts)).Invoke(
		func() (interface{}, error) {
			ok, err := isCCInstantiated(resMgmt, channelID, ccName, ccVersion, peers)
			if !ok {
				return &ok, errors.WithMessage(err, "  Did NOT find instantiated chaincode [%s:%s] on one or more peers in [%s].")
			}
			return &ok, nil
		},
	)

	if err != nil {
		return false, errors.WithMessage(err, "  Got error checking if chaincode was instantiated")
	}

	return *(instantiated).(*bool), nil
}

func isCCInstantiated(resMgmt *resmgmt.Client, channelID, ccName, ccVersion string, peers []fab.Peer) (bool, error) {

	installedOnAllPeers := true

	for _, peer := range peers {

		fmt.Println("\n   Querying peer "+peer.URL()+" for instantiated chaincode")

		chaincodeQueryResponse, err := resMgmt.QueryInstantiatedChaincodes(channelID, resmgmt.WithRetry(retry.DefaultResMgmtOpts), resmgmt.WithTargets(peer))

		if err != nil {
			return false, errors.WithMessage(err, "  QueryInstantiatedChaincodes return error")
		}
		fmt.Println("\n   Found instantiated chaincodes on peer "+peer.URL())

		found := false

		for _, chaincode := range chaincodeQueryResponse.Chaincodes {
			fmt.Println("   Found instantiated chaincode Name: "+chaincode.Name+", Version: "+chaincode.Version+", Path: "+chaincode.Path+" on peer "+peer.URL())
			if chaincode.Name == ccName && chaincode.Version == ccVersion {
				found = true
				break
			}
		}
		if !found {
			fmt.Println("  "+ccName+" chaincode is not instantiated on peer "+ peer.URL())
			installedOnAllPeers = false
		}
	}
	return installedOnAllPeers, nil
}


func(s *OrgSetup) TestInvoke(org string) error {

	fmt.Println(" ********** Test Invoke - "+org+" **********")

	eventID := "testInvoke - "+org

	orgSetup := s.ChooseORG(org)
	orgName 			:= 	orgSetup.OrgName
	orgSdk				:=  orgSetup.Sdk
	orgAdmin			:=  orgSetup.OrgAdmin
	caClient 			:= 	orgSetup.CaClient
	channelClient, event,_ := s.CreateChannelClient(orgSdk, orgName, orgAdmin, caClient)

	_, err := s.ExecuteChaincodeTranctionEvent(eventID, "invoke",[][]byte{
		[]byte("testInvoke"),
		[]byte(eventID),

	}, s.ChaincodeId, channelClient,event)

	if err != nil {
		return fmt.Errorf("Error - Test Invoke failed for "+org+" : %s", err.Error())
	}

	fmt.Println(" ********** "+org+" Test Invoke Successful ********** ")

	return nil
}