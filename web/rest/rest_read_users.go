package rest

import (
	"privateledger/blockchain/invoke"
	"privateledger/web/model"
	"encoding/json"
	"fmt"
	"net/http"
)

func (app *RestApp) GetUserDataByEmailHandler() func(http.ResponseWriter, *http.Request) {

	return app.isAuthorized(func(w http.ResponseWriter, r *http.Request) {

		orgUser := app.Org.GetOrgUser()

		if orgUser == nil {
			respondJSON(w, map[string]string{"error": "No Session Available"})
		} else {

			var userdata model.ModelUserData
			_ = json.NewDecoder(r.Body).Decode(&userdata)
			email := userdata.Email

			fmt.Println(" Session User - " + orgUser.Username)
			fmt.Println(" Session OrgName - " + orgUser.Setup.OrgName)

			orgInvoke := invoke.OrgInvoke{
				User: orgUser,
			}

			UserData, err := orgInvoke.GetUserFromLedger(email, true)

			if err != nil {
				respondJSON(w, map[string]string{"error": "No User Data Found"})
			} else {
				respondJSON(w, UserData)
			}
		}
	})
}

func (app *RestApp) GetAllUsersDataHandler() func(http.ResponseWriter, *http.Request) {

	return app.isAuthorized(func(w http.ResponseWriter, r *http.Request) {

		orgUser := app.Org.GetOrgUser()

		if orgUser == nil {
			respondJSON(w, map[string]string{"error": "No Session Available"})
		} else {

			orgInvoke := invoke.OrgInvoke{
				User: orgUser,
			}

			allUsersData, err := orgInvoke.GetAllUsersFromLedger()

			if err != nil {
				respondJSON(w, map[string]string{"error": "No User Data Found"})
			} else {
				respondJSON(w, allUsersData)
			}

		}
	})
}
