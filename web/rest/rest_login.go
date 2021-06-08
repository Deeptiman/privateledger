package rest

import (
	"privateledger/blockchain/invoke"
	"privateledger/web/model"
	"encoding/json"
	"fmt"
	"net/http"
)

func (app *RestApp) LoginHandler(w http.ResponseWriter, r *http.Request) {

	var userdata model.ModelUserData

	_ = json.NewDecoder(r.Body).Decode(&userdata)

	orgName := userdata.Org
	email := userdata.Email
	password := hash(userdata.Password)

	Org, err := app.Org.InitializeOrg(orgName)
	if err != nil {
		respondJSON(w, map[string]string{"error": "failed to invoke user " + err.Error()})
	}

	fmt.Println("Sign In --->  emailValue = " + email)

	orgUser, err := Org.LoginUserWithCA(email, password)

	orgInvoke := invoke.OrgInvoke{
		User: orgUser,
	}

	if err != nil {
		respondJSON(w, map[string]string{"error": "Unable to Login : " + err.Error()})
	} else {

		fmt.Println("Logged In User : " + orgUser.Username)

		token := app.processAuthentication(w, email)

		if len(token) > 0 {

			UserData, err := orgInvoke.GetUserFromLedger(email, true)

			if err != nil {
				respondJSON(w, map[string]string{"error": "No User Data Found"})
			} else {
				respondJSON(w, map[string]string{

					"token":  token,
					"name":   UserData.Name,
					"email":  UserData.Email,
					"age":    UserData.Age,
					"mobile": UserData.Mobile,
					"salary": UserData.Salary,
				})
			}

		} else {
			respondJSON(w, map[string]string{"error": "Failed to generate token"})
		}
	}
}
