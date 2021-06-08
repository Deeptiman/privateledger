package rest

import (
	"privateledger/blockchain/invoke"
	"privateledger/web/model"
	"encoding/json"
	"fmt"
	"net/http"
)

func (app *RestApp) UpdateUserHandler() func(http.ResponseWriter, *http.Request) {

	return app.isAuthorized(func(w http.ResponseWriter, r *http.Request) {

		orgUser := app.Org.GetOrgUser()

		if orgUser == nil {
			respondJSON(w, map[string]string{"error": "No Session Available"})
		} else {

			var userdata model.ModelUserData
			_ = json.NewDecoder(r.Body).Decode(&userdata)

			name := userdata.Name
			email := userdata.Email
			mobile := userdata.Mobile
			age := userdata.Age
			salary := userdata.Salary
			role := userdata.Role
			targets := userdata.Targets

			fmt.Println(" ####### Rest Input for Update ####### ")

			fmt.Println(" Update Email	 	= " + email)
			fmt.Println(" Update Name 		= " + name)
			fmt.Println(" Update Mobile 	= " + mobile)
			fmt.Println(" Update Age 		= " + age)
			fmt.Println(" Update Salary 	= " + salary)
			fmt.Println(" Update Targets 	= " + targets)
			fmt.Println(" Update Role 		= " + role)
			fmt.Println(" ###################################### ")

			orgInvoke := invoke.OrgInvoke{
				User: orgUser,
			}

			user, _ := orgInvoke.GetUserFromLedger(email, false)

			if user != nil {

				err := orgInvoke.UpdateUserFromLedger(email, name, mobile, age, salary, targets, role)

				if err != nil {
					respondJSON(w, map[string]string{"error": "Error Update User Data = " + err.Error()})
				} else {

					UserData, err := orgInvoke.GetUserFromLedger(email, false)

					if err != nil {
						respondJSON(w, map[string]string{"error": "No User Data Found"})
					} else {
						respondJSON(w, UserData)
					}
				}
			} else {
				respondJSON(w, map[string]string{"error": "No User Data Found"})
			}
		}
	})
}
