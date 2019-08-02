package org

import (
	"fmt"
)

func(s *OrgSetup) ChooseORG(org string) *OrgSetup {

	fmt.Println("Input  Org "+org)

	orgSetup, _ := s.InitializeOrg(org)

	fmt.Println(" ########### Org Details ################# ")
	fmt.Println(" OrgName - "+orgSetup.OrgName)

	return &orgSetup
}



