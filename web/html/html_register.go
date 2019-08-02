package html

import (
	"fmt"
	"net/http"
	"github.com/privateledger/blockchain/invoke"
	"github.com/privateledger/chaincode/model"
)

func (app *HtmlApp) RegisterHandler(w http.ResponseWriter, r *http.Request) {

	data := &struct {
		TransactionId string
		ErrorMsg      string
		Error         bool
		Register   	  bool
		Login		  bool
		Success       bool
		Response      bool
		Username      string
		UserOrg		  string
		CustomOrg1	  string
		CustomOrg2	  string
		CustomOrg3	  string
		CustomOrg4	  string
	}{
		TransactionId: "",
		ErrorMsg:      "",
		Error:         false,
		Register:      false,
		Login:		   false,
		Success:       false,
		Response:      false,
		Username:      "",
		UserOrg:	   "",
		CustomOrg1:	   model.GetCustomOrgName("org1"),
		CustomOrg2:	   model.GetCustomOrgName("org2"),
		CustomOrg3:	   model.GetCustomOrgName("org3"),
		CustomOrg4:	   model.GetCustomOrgName("org4"),
	}

	fmt.Println("RegisterHandler : " + r.FormValue("signupSubmitted"))

	if r.FormValue("signupSubmitted") == "true" {

		fmt.Println("Register : signupSubmitted")

		data.Register = true

		orgName 	:= r.FormValue("company")
		name 		:= r.FormValue("name")
		mobile 		:= r.FormValue("mobile")
		age 		:= r.FormValue("age")
		salary 		:= r.FormValue("salary")

		email 		:= r.FormValue("email")		
		password 	:= hash(r.FormValue("password"))
		role 		:= r.FormValue("role")

		fmt.Println(" ####### Web Input for Register ####### ")

		fmt.Println(" Email = "+email)
		fmt.Println(" Password = "+password)	
		fmt.Println(" Role = "+role)

		fmt.Println(" Name = "+name)
		fmt.Println(" Age = "+age)
		fmt.Println(" Mobile = "+mobile)
		fmt.Println(" Salary = "+salary)
		fmt.Println(" ###################################### ")

		Org, err := app.Org.InitializeOrg(orgName)
		if err != nil {

			fmt.Errorf("Web ----- Unable to initialize org : %s", err.Error())

			data.Error = true
			data.ErrorMsg = err.Error()

		} else {

				orgUser, err := Org.RegisterUserWithCA(orgName, email, password, role)

				orgInvoke := invoke.OrgInvoke {
					User: orgUser,
				}

				if err != nil {

					fmt.Println("Web Error ----->>> Unable to Register Error Msg : %v", err)
				
					data.Error = true
					data.ErrorMsg = err.Error()

				} else {

					data.Username = email
					data.Success = true

					token := app.processAuthentication(w, email)

					if len(token) > 0 {
						err := orgInvoke.InvokeCreateUser(name, age, mobile, salary)

						if err != nil {
							fmt.Errorf("failed to invoke user "+err.Error())
						} else {

							http.Redirect(w, r, "/index.html", 302)
						}
					}
				}
		}
		data.Response = true
	}

	renderTemplate(w, r, "register.html", data)
}
