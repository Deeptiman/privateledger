package rest

import (
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"net/http"

	"privateledger/blockchain/org"
)

var PORT = "4000"

type RestApp struct {
	Org *org.OrgSetup
}

func hash(s string) string {
	h := sha1.New()
	h.Write([]byte(s))
	sha1_hash := hex.EncodeToString(h.Sum(nil))

	return sha1_hash
}

func respondJSON(w http.ResponseWriter, payload interface{}) {

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(payload)
}
