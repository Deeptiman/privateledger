package rest

import (
	"privateledger/web/model"
	"encoding/json"
	"net/http"
)

func (app *RestApp) ChangePwdHandler() func(http.ResponseWriter, *http.Request) {

	return app.isAuthorized(func(w http.ResponseWriter, r *http.Request) {

		orgUser := app.Org.GetOrgUser()

		if orgUser == nil {
			respondJSON(w, map[string]string{"error": "Error Session User "})
		} else {

			var userdata model.ModelUserData

			_ = json.NewDecoder(r.Body).Decode(&userdata)

			email := userdata.Email
			role := userdata.Role
			oldPwd := hash(userdata.OldPassword)
			newPwd := hash(userdata.Password)

			verifyErr := verifyPassword(userdata.Password)

			if verifyErr != nil && len(verifyErr.Error()) > 0 {
				respondJSON(w, map[string]string{"error": verifyErr.Error(), "message": "Password must contain at least one number and one uppercase and lowercase letter, and at least 8 or more characters"})
			} else {

				err := orgUser.Setup.ChangePassword(email, role, oldPwd, newPwd)

				if err != nil {
					respondJSON(w, map[string]string{"error": "Unable to Change user pwd - " + err.Error()})
				} else {
					respondJSON(w, map[string]string{"success": "Password successfully changed for - " + email})
				}
			}
		}
	})
}
