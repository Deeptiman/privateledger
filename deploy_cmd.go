package main

import (
	"fmt"
	"strings"
	"github.com/privateledger/blockchain/org"
	"github.com/pkg/errors"
)

func DeployCMD(setup *org.OrgSetup,cmd string) error {

	fmt.Println("Deploy CMD "+cmd)

	err := setup.Init(true)
	if err != nil {
		return errors.WithMessage(err, "  failed to initialize  : "+err.Error())
	} 
	fmt.Println(" initialized successfully\n")

	switch id := cmd; id {

		case "1":

			/*  Channel Create for Org */

			for _, s := range org.OrgList {
					
				err = s.CreateChannelForOrg()
				if err != nil {
					return errors.WithMessage(err, "  failed to Create Channel "+err.Error())
				}
			}

			break

		case "2":
			
			for _, s := range org.OrgList {

				if !strings.EqualFold(s.OrgName, s.OrdererName){

					err = s.JoinChannelForOrg()
					if err != nil {
						return errors.WithMessage(err, "  failed to Join Channel "+err.Error())
					}

				}
			}

			break

		case "3":

			/* Install Chaincode */

			ccPkg, err := org.OrgList[1].CreateCCPackage()
			if err != nil {
				fmt.Println("  failed to create chaincode Pkg : "+err.Error())
				return errors.WithMessage(err, "  failed to create chaincode Pkg : "+err.Error())
			}
			 
			Org1 := org.OrgList[1]
			err = Org1.InstallCCForOrg(ccPkg)
			if err != nil {
				return errors.WithMessage(err, "  failed to Install Chaincode "+" : "+err.Error())
			}
			fmt.Println("Install CC successfull for - "+Org1.OrgName)


			Org2 := org.OrgList[2]
			err = Org2.InstallCCForOrg(ccPkg)
			if err != nil {
				return errors.WithMessage(err, "  failed to Install Chaincode "+" : "+err.Error())
			}
			fmt.Println("Install CC successfull for - "+Org2.OrgName)

			Org3 := org.OrgList[3]
			err = Org3.InstallCCForOrg(ccPkg)
			if err != nil {
				return errors.WithMessage(err, "  failed to Install Chaincode "+" : "+err.Error())
			}
			fmt.Println("Install CC successfull for - "+Org3.OrgName)


			Org4 := org.OrgList[4]
			err = Org4.InstallCCForOrg(ccPkg)
			if err != nil {
				return errors.WithMessage(err, "  failed to Install Chaincode "+" : "+err.Error())
			}
			fmt.Println("Install CC successfull for - "+Org4.OrgName)


			/*for _, s := range org.OrgList {

				if !strings.EqualFold(s.OrgName, s.OrdererName) {

					err = s.InstallCCForOrg(ccPkg)
					if err != nil {
						return errors.WithMessage(err, "  failed to Install Chaincode "+" : "+err.Error())
					}
					fmt.Println("Install CC successfull for - "+s.OrgName)
				}
			}*/
			break

		case "4":

			/* Instantiate Chaincode */

			Org1 := org.OrgList[1]
			org1Peers := Org1.Peers

			err = Org1.InstantiateCCForOrg(org1Peers)
			if err != nil {
				return errors.WithMessage(err, "  failed to Instantiate Chaincode "+" : "+err.Error())
			}
			fmt.Println("Instantiate CC successfull ")

			for _, s := range org.OrgList {
				
				if !strings.EqualFold(s.OrgName, s.OrdererName){
				
					err := s.TestInvoke(s.OrgName)

					if err !=nil {
						fmt.Println(" Invoke failed for - "+s.OrgName+" : "+err.Error())
					}
					
				}
			}

			break

		case "5":

			var testOrg string
			fmt.Println(" Enter the Org name- ( org1, org2, org3, org4) ")
			fmt.Scanln(&testOrg)

			s := org.OrgList[4]

			err := s.TestInvoke(testOrg)

			if err !=nil {
				fmt.Println(" Invoke failed for - "+testOrg+" : "+err.Error())
			}

			break;

		case "6":
						 
		 
			Org1 := org.OrgList[1]
			Org2 := org.OrgList[2]
			Org3 := org.OrgList[3]
			Org4 := org.OrgList[4]
				

			ccPkg, err := Org1.CreateCCPackage()
			if err != nil {
				fmt.Println("  failed to create chaincode Pkg : "+err.Error())
				return errors.WithMessage(err, "  failed to create chaincode Pkg : "+err.Error())
			}
		
			err = Org1.InstallCCForOrg(ccPkg)
			if err != nil {
				return errors.WithMessage(err, "  failed to Install Chaincode "+" : "+err.Error())
			}
			fmt.Println("Install CC successfull for - "+Org1.OrgName)

			err = Org2.InstallCCForOrg(ccPkg)
			if err != nil {
				return errors.WithMessage(err, "  failed to Install Chaincode "+" : "+err.Error())
			}
			fmt.Println("Install CC successfull for - "+Org2.OrgName)

			err = Org3.InstallCCForOrg(ccPkg)
			if err != nil {
				return errors.WithMessage(err, "  failed to Install Chaincode "+" : "+err.Error())
			}
			fmt.Println("Install CC successfull for - "+Org3.OrgName)


			err = Org4.InstallCCForOrg(ccPkg)
			if err != nil {
				return errors.WithMessage(err, "  failed to Install Chaincode "+" : "+err.Error())
			}
			fmt.Println("Install CC successfull for - "+Org4.OrgName)


			org1Peers := Org1.Peers
			
			err = Org1.UpgradeCCForOrg(org1Peers)
			if err != nil {
				return errors.WithMessage(err, "  failed to Upgrade Chaincode "+" : "+err.Error())
			}
			fmt.Println("Upgrade CC successfull ")

			for _, s := range org.OrgList {
				
				if !strings.EqualFold(s.OrgName, s.OrdererName){
				
					err := s.TestInvoke(s.OrgName)

					if err !=nil {
						fmt.Println(" Invoke failed for - "+s.OrgName+" : "+err.Error())
					}
					
				}
			}

			break

		case "7":

			for _, s := range org.OrgList {

				err = s.QueryInstalledCCForOrg()
				if err != nil {
					return errors.WithMessage(err, "  failed to Query Installed Chaincode "+" : "+err.Error())
				}	
				fmt.Println("Query CC successfull ")
			}
			break

		case "8":

			for _, s := range org.OrgList {

				err = s.QueryInstantiatedCCForOrg()
				if err != nil {
					return errors.WithMessage(err, "  failed to Query Instantiated Chaincode "+" : "+err.Error())
				}
				fmt.Println("Query CC successfull ")
			}
			break

		case "9":

			Org3 := org.OrgList[3]
			
			err := Org3.AddAffiliationOrg()

			if err != nil {
				return errors.WithMessage(err, "  failed to affiliate org "+" : "+err.Error())	
			}

			Org4 := org.OrgList[4]

			err = Org4.AddAffiliationOrg()

			if err != nil {
				return errors.WithMessage(err, "  failed to affiliate org "+" : "+err.Error())	
			}

			break
			 
	}

	return nil
}