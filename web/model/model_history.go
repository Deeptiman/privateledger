package model

import (
	"github.com/golang/protobuf/ptypes/timestamp"
)

type HistoryData struct {
	EmailKey			string 					`json:"emailKey"`
	TxID  				string 					`json:"txId"`
	QueryCreator		string 					`json:"creator"`
	Query				string					`json:"query"`
	QueryCreatorOrg		string					`json:"queryCreatorOrg"`
	Time				*timestamp.Timestamp 	`json:"time"`
	Remarks				string					`json:"remarks"`
}