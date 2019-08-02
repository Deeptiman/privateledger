package model

import (
	"fmt"
	"strings"
	"time"	
	"github.com/golang/protobuf/ptypes/timestamp"
)
 
const (
	CO1 = "Apple"
	CO2 = "Google"
	CO3 = "Microsoft"
	CO4 = "Amazon"
)

const (

	COLLECTION_KEY = "email"
	ADMIN = "admin"
)

func GetCustomOrgName(org string) string {
	fmt.Println(" ### GetCustomOrgName = "+org)
	if strings.EqualFold(org,"org1"){
		return CO1
	} else if strings.EqualFold(org,"org2"){
		return CO2
	} else if strings.EqualFold(org,"org3"){
		return CO3
	} else if strings.EqualFold(org,"org4"){
		return CO4
	}
	return CO1
}

type User struct {

	ID    		string 						`json:"id"`
	Email 		string 						`json:"email"`	
	Name  		string 						`json:"name"`
	Mobile 		string 						`json:"mobile"`
	Age    		string 						`json:"age"`
	Salary 		string 						`json:"salary"`
	Owner		string						`json:"owner"`
	ShareAccess string						`json:"shareAccess"`
	Role		string						`json:"role"`
	Targets 	string 						`json:"targets"`
	Time		string						`json:"time"`
	Remarks		string						`json:"remarks"`	
}

type HistoryData struct {
	EmailKey			string 				`json:"emailKey"`
	TxID  				string 				`json:"txId"`
	QueryCreator			string 				`json:"creator"`
	Query				string				`json:"query"`
	TargetOrg			string				`json:"targetOrg"`
	Time				string 				`json:"time"`
	Remarks				string				`json:"remarks"`
}


func IsAdmin(role string) bool {
	return strings.EqualFold(role,ADMIN)
}

func GetTime(timestamp *timestamp.Timestamp) string {

	t := time.Unix(timestamp.GetSeconds(), 0)

	return t.Format(time.RFC1123)
}
