package invoke

import (
	"fmt"
	"strings"
	"encoding/json"
	"github.com/privateledger/chaincode/model"
	"github.com/golang/protobuf/proto"
	"github.com/dgrijalva/jwt-go"
)


func(s *OrgInvoke) CreateOrgTargets(email string, accessList []int32, orgList []string) (string, error) {

	fmt.Println(" #### CreateOrgTargets ####")

	var targetOrg *model.Target
	var access int32
	targets := make(map[string][]byte)

		sessionOrg :=  s.User.Setup.OrgName
	// Session Org 
		access = model.LedgerAccess(model.ALL).Int()
		transactionHash, err := s.GetSecretMessage(sessionOrg, sessionOrg)

		if err != nil {
			return "", fmt.Errorf("Invalid Hash, unable to create hash for - "+sessionOrg, err)
		}
		targetOrg = &model.Target{
			Access 	: access,
			Remarks	: model.GetCustomOrgName(sessionOrg)+" created user - "+email,
			TransactionHash: transactionHash,
		}

		data, err := proto.Marshal(targetOrg)
		if err != nil {
			return "", fmt.Errorf("marshaling error: ", err)
		}

		targets[sessionOrg] = data
	
		for i, org := range orgList {
					
			if !strings.EqualFold(sessionOrg, org){
			
				transactionHash, err = s.GetSecretMessage(sessionOrg, org)
				if err != nil {
					return "", fmt.Errorf("Invalid Hash, unable to create hash for - "+sessionOrg, err)
				}

				var remarks string
				access = accessList[i]
				
				if access == model.LedgerAccess(model.NOACCESS).Int(){
					remarks = model.GetCustomOrgName(org)+" currently has no access to user - "+email
				} else if access == model.LedgerAccess(model.REMOVEACCESS).Int(){
					remarks = model.GetCustomOrgName(sessionOrg)+" remove all access for "+model.GetCustomOrgName(org)
				} else {
					a := model.LedgerAccess(access).String()
					remarks = model.GetCustomOrgName(sessionOrg)+" gave "+ a + " access to "+model.GetCustomOrgName(org)
				}

				targetOrg = &model.Target{
					Access 	: access, 
					Remarks	: remarks,
					TransactionHash: transactionHash,
				}

				data, err := proto.Marshal(targetOrg)
				if err != nil {
					return "", fmt.Errorf("marshaling error: ", err)
				}
			
				targets[org] = data
			
				fmt.Println(" Org ==== "+org)
			}
	}

	tgts, _ := json.Marshal(targets)

	return string(tgts), nil	
}


func(s *OrgInvoke) UpdateOrgTargets(targets string, org string, access int32) (string, string, error) {

	sessionOrg :=  strings.ToLower(s.User.Setup.OrgName)

	fmt.Println("Update Org Target - "+org+" == access = ", access)

	tgts := make(map[string][]byte)
	
	err  := json.Unmarshal([]byte(targets), &tgts)
	
	if err != nil {
		fmt.Println("UpdateOrgTargets = failed to unmarshaling error: "+err.Error())
		return "", "", fmt.Errorf("failed to unmarshaling error: ", err.Error())
	}

	remarks := ""

	if access == model.LedgerAccess(model.REMOVEACCESS).Int(){
		remarks = model.GetCustomOrgName(sessionOrg)+" remove all access for "+model.GetCustomOrgName(org)
	} else {
		a := model.LedgerAccess(access).String()
		remarks = model.GetCustomOrgName(sessionOrg)+" gave "+ a + " access to "+model.GetCustomOrgName(org)
	}
	

	transactionHash, err := s.GetSecretMessage(sessionOrg, org)

	if err != nil {
		return "", "", fmt.Errorf(" failed to generate transaction hash for "+org, err)
	}

	targetOrg := &model.Target{
		Access 	: access,
		Remarks	: remarks,	
		TransactionHash: transactionHash,	
	}

	data, err := proto.Marshal(targetOrg)
	if err != nil {
		fmt.Println("UpdateOrgTargets =  marshaling error: "+err.Error())
		return "", "", fmt.Errorf("marshaling error: ", err)
	}

	tgts[org] = data

	updateTgts, _ := json.Marshal(tgts)

	return string(updateTgts), remarks, nil
}

func(s *OrgInvoke) GetSecretMessage(sessionOrg , targetOrg string) (string, error) {

	message := "need code access for "+sessionOrg+" and "+targetOrg

	transactionHash, err := GetTransactionHash(message)

	if err != nil {
		return "", fmt.Errorf(" GetSecretMessage Error  ", err)
	}

	return transactionHash, nil
}

func GetTransactionHash(message string) (string, error){

	token := jwt.New(jwt.SigningMethodHS256)

	hash, err := token.SignedString([]byte(message))

	if err != nil {
		fmt.Println("something went wrong: %s", err.Error())
		return "", err
	}

	return hash, nil
}


func(s *OrgInvoke) ParseOrgTargets(targets string) error {

	tgts := make(map[string][]byte)
	
	err  := json.Unmarshal([]byte(targets), &tgts)
	
	if err != nil {
		fmt.Println("ParseOrgTargets = failed to unmarshaling error: "+err.Error())
		return fmt.Errorf("failed to unmarshaling error: ", err.Error())
	}

	for key, value := range tgts {

		orgName := key
		obj := []byte(value)

		fmt.Println(" ############# Unmarshal Target - "+orgName+"  ################## ")
	
		newTarget := &model.Target{}
		err = proto.Unmarshal([]byte(obj), newTarget)
		if err != nil {
			fmt.Println("ParseOrgTargets = unmarshaling error: "+err.Error())
			return fmt.Errorf("unmarshaling error: ", err.Error())
		}
	
		fmt.Println(" Target Access = ",newTarget.GetAccess())  
		fmt.Println(" Target Remarks = ",newTarget.GetRemarks()) 
		fmt.Println(" Target TransactionHash = ",newTarget.GetTransactionHash())  
	
		fmt.Println(" ################################################## ")	
	}

	return nil
}


func(s *OrgInvoke) IsAccessExists(targets string, checkAccess int32,  org string) bool{

	tgts := make(map[string][]byte)
	
	err  := json.Unmarshal([]byte(targets), &tgts)
	
	if err != nil {
		fmt.Println("GetAccessOrgList = failed to unmarshaling error: "+err.Error())
		return false
	}

	for key, value := range tgts {

		orgName := key
		access := []byte(value)

		fmt.Println(" ############# Unmarshal Target - "+orgName+"  ################## ")
	
		newTarget := &model.Target{}
		err = proto.Unmarshal([]byte(access), newTarget)
		if err != nil {
			fmt.Println("ParseOrgTargets = unmarshaling error: "+err.Error())
			return false
		}

		fmt.Println(" Target Access = ",newTarget.GetAccess())  
		fmt.Println(" Target Remarks = ",newTarget.GetRemarks())  
			
		if checkAccess == model.LedgerAccess(newTarget.GetAccess()).Int() && 
			strings.EqualFold(orgName, org){
				return true
		}

		fmt.Println(" ################################################## ")	
	}

	return false
}


func(s *OrgInvoke) OrgHasAccess(targets string, checkAccess int32, org string) bool{

	fmt.Println(" ########### Check Access ############## ")

	tgts := make(map[string][]byte)
	
	err  := json.Unmarshal([]byte(targets), &tgts)
	
	if err != nil {
		fmt.Println("GetAccessOrgList = failed to unmarshaling error: "+err.Error())
		return false
	}

	for key, value := range tgts {

		orgName := key
		access := []byte(value)

		fmt.Println(" ############# Unmarshal Target - "+orgName+"  ################## ")
	
		newTarget := &model.Target{}
		err = proto.Unmarshal([]byte(access), newTarget)
		if err != nil {
			fmt.Println("ParseOrgTargets = unmarshaling error: "+err.Error())
			return false
		}

		fmt.Println(" Target Access = ",newTarget.GetAccess())  
		fmt.Println(" Target Remarks = ",newTarget.GetRemarks())  
		
		var check = ( model.LedgerAccess(checkAccess).Int() == model.LedgerAccess(newTarget.GetAccess()).Int())	

		if (check || model.LedgerAccess(newTarget.GetAccess()).AllAccess()) && 
			strings.EqualFold(orgName, org){
				fmt.Println(org+" has access ")
				return true
		}

		fmt.Println(" ################################################## ")	
	}

	return false
}

func(s *OrgInvoke) GetAccessOrgList(targets string) ([]string, error){

	fmt.Println(" $$$$$$$$$ GetAccessOrgList $$$$$$$$$$$$$$")

	tgts := make(map[string][]byte)
	
	err  := json.Unmarshal([]byte(targets), &tgts)
	
	if err != nil {
		fmt.Println("GetAccessOrgList = failed to unmarshaling error: "+err.Error())
		return nil, fmt.Errorf("failed to unmarshaling error: ", err.Error())
	}

	var orgList [] string

	for key, value := range tgts {

		orgName := key
		access := []byte(value)

		fmt.Println(" ############# Unmarshal Target - "+orgName+"  ################## ")
	
		newTarget := &model.Target{}
		err = proto.Unmarshal([]byte(access), newTarget)
		if err != nil {
			fmt.Println("ParseOrgTargets = unmarshaling error: "+err.Error())
			return nil, fmt.Errorf("unmarshaling error: ", err.Error())
		}
			
		if !model.LedgerAccess(newTarget.GetAccess()).NoAccess(){
			orgList = append(orgList, orgName)
		} else {
			fmt.Println(" ***** No Access Orgs --- "+orgName)
		}

		fmt.Println(" Target Access = ",newTarget.GetAccess())  
		fmt.Println(" Target Remarks = ",newTarget.GetRemarks())  
	
		fmt.Println(" ################################################## ")	
	}

	fmt.Println(" $$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$ ")
	return orgList, nil

}