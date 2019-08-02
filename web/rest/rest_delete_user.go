package rest

import (
	"encoding/json"
	"fmt"
	"strings"
	"net/http"
	"github.com/privateledger/blockchain/invoke"
	"github.com/privateledger/web/model"
)

func (app *RestApp) DeleteUserHandler() func(http.ResponseWriter, *http.Request) {

	return app.isAuthorized(func(w http.ResponseWriter, r *http.Request) {

		orgUser := app.Org.GetOrgUser()

		if orgUser == nil {
			respondJSON(w, map[string]string{"error": "No Session Available"})
		} else {

				var userdata model.ModelUserData
				_ = json.NewDecoder(r.Body).Decode(&userdata)
				email := userdata.Email
				role  := userdata.Role
				owner := userdata.Org
				targets := userdata.Targets
			
				fmt.Println("DeleteUserHandler : Email = " + email)
				
				orgInvoke := invoke.OrgInvoke {
					User: orgUser,
				}

				orgSetup := orgUser.Setup.ChooseORG(strings.ToLower(owner))

				err := orgUser.RemoveUser(email,orgSetup.OrgCaID, orgSetup.CaClient)

				if err != nil {
					fmt.Println("DeleteUserHandler : RemoveUser = Error : " + err.Error())
					respondJSON(w, map[string]string{"error": "Error Session User  " + err.Error()})
				} else {
					fmt.Println("Success RemoveUser ")

					// ReInitialize to Session Org

					_ = orgUser.Setup.ChooseORG(strings.ToLower(orgUser.Setup.OrgName))

					err = orgInvoke.DeleteUserFromLedger(email, targets, role)

					if err != nil {
						fmt.Println("DeleteUserHandler : Error Deleting User from ledger : " + err.Error())
						respondJSON(w, map[string]string{"error": "Error Deleting User from ledger " + err.Error()})
					}
					respondJSON(w, map[string]string{"success": "Succesfully delete the user with email - " + email})

				}			 
		}
	})
}
