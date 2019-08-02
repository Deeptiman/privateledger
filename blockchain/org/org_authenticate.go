package org

import (
	"fmt"
	"strings"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/event"
	caMsp "github.com/hyperledger/fabric-sdk-go/pkg/client/msp"
)

var sessionOrgName = make(map[string]string)
var sessionUser = make(map[string]string)
var secretKey = make(map[string]string)

type OrgUser struct {
	Username        	string
	ChannelClient   	*channel.Client
	Event				*event.Client
	Setup	 			OrgSetup
}

var sessionOrgUser *OrgUser

func(s *OrgSetup) GetOrgUser() *OrgUser{
	return sessionOrgUser
}

func (s *OrgUser) Logout() {
	fmt.Println("Logout , delete the session") 
	delete(sessionUser, "name")
	sessionOrgUser = nil
}

func (s *OrgUser) RevokeUser(email string) error {

	_, err := s.Setup.CaClient.Revoke(&caMsp.RevocationRequest{
		Name: email,
	})

	if err != nil {
		return fmt.Errorf("failed to revoke signing identity for '%s': %v", email, err)
	}

	return nil
}

func (s *OrgUser) RemoveUserFromCA(email string, caID string, caClient *caMsp.Client) error {

	_, err := caClient.RemoveIdentity(&caMsp.RemoveIdentityRequest{

		ID:     email,
		Force:  true,
		CAName: caID,
	})

	if err != nil {
		return fmt.Errorf("failed to remove signing identity for '%s': %v", email, err)
	}
	return nil
}

func (s *OrgUser) RemoveUser(email string, caID string, caClient *caMsp.Client) error {
	 
	err := s.RemoveUserFromCA(email, caID, caClient)

	if err != nil {
		return fmt.Errorf("failed to remove signing identity for '%s': %v", email, err)
	} else {

		/*err = s.DeleteUserFromLedger()
		if err != nil {
			return fmt.Errorf("unable to delete user from blockchain:>>> %v", err)
		}*/
	}
	return nil
}


func (s *OrgSetup) ChangePassword(email, role, oldPwd, newPwd string) error {

	//fmt.Println("Change PWD : Email = " + email + " , OLD PWD = " + oldPwd + " , Saved PWD = " + secretKey["secret"] + " ,  New PWD = " + newPwd)

	if !strings.EqualFold(oldPwd, secretKey["secret"]) {
		return fmt.Errorf("Old password don't matched, can't change pwd for the email: '%s'",email)
	}

	if strings.EqualFold(oldPwd, newPwd) {
		return fmt.Errorf("failed old password, new password should not be same: '%s'",email)
	}

	orgUser := s.GetOrgUser()

	err := orgUser.RemoveUserFromCA(email, s.OrgCaID, s.CaClient)

	if err != nil {
		return fmt.Errorf("failed to remove identity '%s': %v", email, err)
	}

	fmt.Println("User Removed")

	org := s.OrgName

	_, err = s.RegisterUserWithCA(org, email, newPwd, role)

	if err != nil {
		return fmt.Errorf("failed to register with CA '%s': %v", email, err)
	}

	fmt.Println("User Re-Registered")

	_, err = s.LoginUserWithCA(email, newPwd)

	if err != nil {
		return fmt.Errorf("failed to enroll user '%s': %v", email, err)
	}

	fmt.Println("User Enrolled")

	err = s.ReEnrollUser(email)

	if err != nil {
		return fmt.Errorf("failed to re-enroll user '%s': %v", email, err)
	}

	fmt.Println("User Re-Enrolled")

	return nil
}

func (s *OrgSetup) ReEnrollUser(email string) error {

	err := s.CaClient.Reenroll(email)

	if err != nil {
		return fmt.Errorf("failed to re-enroll user '%s': %v", email, err)
	}
	return nil
}

func(s *OrgSetup) RegisterUserWithCA(org, email, password, role string) (*OrgUser,error) {

	var caid string

	fmt.Println(" ****** Register User ****** ")

	orgSetup := s.ChooseORG(org)

	caid  	  =		orgSetup.OrgCaID 
	caClient :=  	orgSetup.CaClient

	fmt.Println("CA Register Org      === " + org)
	fmt.Println("CA Register CaID     === " + caid)
	fmt.Println("CA Register Email 	  === " + email)
	fmt.Println("CA Register Password === " + password)
	fmt.Println("CA Register Role 	  === " + role)
	
	affl := strings.ToLower(org) + ".department1"

	_, err := caClient.Register(&caMsp.RegistrationRequest{
		Name:           email,
		Secret:         password,
		Type:           "peer",
		MaxEnrollments: -1,
		Affiliation:    affl,
		Attributes: []caMsp.Attribute{
			{
				Name:  "role",
				Value: role,
				ECert: true,
			},
		},
		CAName: caid,
	})

	if err != nil {
		return nil, fmt.Errorf("unable to register user with CA '%s': %v", email, err)
	}

	sessionOrgName["orgName"] = org
	sessionUser["name"] = email
	secretKey["secret"] = password

	fmt.Println("Successfully register user "+email)

	orgUser, err := s.LoginUserWithCA(email, password)
	if err != nil {
		return nil, fmt.Errorf("unable to login '%s': %v", email, err)
	}

	fmt.Println("Org Register Name === " + orgUser.Setup.OrgName)


	return orgUser, nil
}


func(s *OrgSetup) LoginUserWithCA(email, password string) (*OrgUser, error) {

	fmt.Println(" ****** Login User ****** ")

	//orgSetup := s.ChooseORG(org)
	caClient := s.CaClient

	err := caClient.Enroll(email, caMsp.WithSecret(password))
	if err != nil {
		return nil, fmt.Errorf("failed to enroll identity '%s': %v", email, err)
	}

	sessionOrgName["orgName"] = s.OrgName
	sessionUser["name"] = email
	secretKey["secret"] = password

	fmt.Println("Org - "+s.OrgName)

	channelClient, event, err := s.CreateChannelClient(s.Sdk, s.OrgName, email, caClient)

	if err != nil {
		return nil, fmt.Errorf("unable to create channel client '%s': %v", email, err)
	}

	fmt.Println("Org Enroll Name === " + s.OrgName)

	sessionOrgUser = &OrgUser{
		Username: email,
		ChannelClient: channelClient,
		Event:	event,
		Setup: *s,
	}

	return sessionOrgUser, nil
}


func(s *OrgSetup) CreateChannelClient(sdk *fabsdk.FabricSDK, org string, email string, caClient *caMsp.Client) (*channel.Client, *event.Client, error) {

	SigningIdentity, err := caClient.GetSigningIdentity(email)
	if err != nil {
		fmt.Println(" failed to get signing identity "+email+" ---- "+err.Error())
		return nil, nil, fmt.Errorf("failed to get signing identity for '%s': %v", email, err)
	}	
	fmt.Println("Signing Identity Created for "+email)

	clientContext := sdk.ChannelContext(
		s.ChannelID, 
		fabsdk.WithUser(email),
		fabsdk.WithOrg(org),
		fabsdk.WithIdentity(SigningIdentity))

	ChannelClient, err := channel.New(clientContext)
	if err != nil {
	   return nil, nil, fmt.Errorf("failed to create new channel client for '%s': %v", email, err)
	}	
	s.ChannelClient = ChannelClient
	fmt.Println("Channel client created")

	// Creation of the client which will enables access to our channel events
	Event, err := event.New(clientContext)
	if err != nil {
	   return nil, nil, fmt.Errorf("failed to create new event client %v", err)
	}
	s.Event = Event
	fmt.Println("Event client created")	
	 
	return s.ChannelClient, s.Event, nil
}


func(s *OrgSetup) AddAffiliationOrg() error {

	org := s.OrgName
	caid := s.OrgCaID
	caClient := s.CaClient

	affl := strings.ToLower(org) + ".department1"

	fmt.Println("Initializing Affiliation for " + affl)

	afRes, err := caClient.GetAffiliation(affl)

	if afRes != nil && err != nil {

		fmt.Println("Affiliation Exists")

		AfInfo := afRes.AffiliationInfo
		CAName := afRes.CAName

		fmt.Println("AfInfo : " + AfInfo.Name)
		fmt.Println("CAName : " + CAName)
	} else {

		fmt.Println("Add Affiliation " + affl)

		_, err = caClient.AddAffiliation(&caMsp.AffiliationRequest{

			Name:   affl,
			Force:  true,
			CAName: caid,
		})

		if err != nil {
			return fmt.Errorf("Failed to add affiliation for CA '%s' : %v ", caid, err)
		}
	}

	return nil
}