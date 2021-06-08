package html

import (
	"privateledger/chaincode/model"
	"fmt"
	"net/http"
)

func (app *HtmlApp) LoginHandler(w http.ResponseWriter, r *http.Request) {

	data := &struct {
		TransactionId string
		ErrorMsg      string
		Error         bool
		Register      bool
		Login         bool
		Success       bool
		Response      bool
		Username      string
		UserOrg       string
		CustomOrg1    string
		CustomOrg2    string
		CustomOrg3    string
		CustomOrg4    string
	}{
		TransactionId: "",
		ErrorMsg:      "",
		Error:         false,
		Register:      false,
		Login:         false,
		Success:       false,
		Response:      false,
		Username:      "",
		UserOrg:       "",
		CustomOrg1:    model.GetCustomOrgName("org1"),
		CustomOrg2:    model.GetCustomOrgName("org2"),
		CustomOrg3:    model.GetCustomOrgName("org3"),
		CustomOrg4:    model.GetCustomOrgName("org4"),
	}

	if r.FormValue("signinSubmitted") == "true" {

		data.Login = true

		orgName := r.FormValue("company")
		email := r.FormValue("email")
		password := hash(r.FormValue("password"))

		fmt.Println("Sign In ---> org = " + orgName + " , email = " + email + " , pwd = " + password)

		Org, err := app.Org.InitializeOrg(orgName)
		if err != nil {
			fmt.Errorf("Web ----- unable to login user : %v", err)

			data.Error = true
			data.ErrorMsg = "failed to initialize org " + err.Error()
		} else {

			_, err := Org.LoginUserWithCA(email, password)

			if err != nil {
				fmt.Println("Web Error ----->>> Unable to Login Error Msg : %v", err)

				fmt.Errorf("Web ----- Unable to Login Error Msg : %s", err.Error())

				data.Error = true
				data.ErrorMsg = err.Error()
			} else {

				token := app.processAuthentication(w, email)

				if len(token) > 0 {

					data.Username = email
					data.Success = true

					http.Redirect(w, r, "/index.html", 302)
				} else {

					data.Error = true
					data.ErrorMsg = "Token generation failed"
				}
			}
		}
		data.Response = true
	}

	renderTemplate(w, r, "login.html", data)
}
