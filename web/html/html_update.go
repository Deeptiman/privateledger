package html

import (
	"privateledger/blockchain/invoke"
	"fmt"
	"net/http"
)

func (app *HtmlApp) EditPageHandler() func(http.ResponseWriter, *http.Request) {

	return app.hasSavedToken(func(w http.ResponseWriter, r *http.Request, token string) {

		data := &Data{}

		if r.FormValue("editSubmitted") == "true" {

			userId := r.FormValue("editUserId")
			email := r.FormValue("editEmail")
			name := r.FormValue("editName")
			age := r.FormValue("editAge")
			mobile := r.FormValue("editMobile")
			salary := r.FormValue("editSalary")
			targets := r.FormValue("editTarget")
			role := r.FormValue("editRole")
			creatorRole := r.FormValue("editCreatorRole")

			fmt.Println(" ####### Web Input for Update ####### ")

			fmt.Println(" Update UserId = " + userId)
			fmt.Println(" Update Email = " + email)
			fmt.Println(" Update Name = " + name)
			fmt.Println(" Update Age = " + age)
			fmt.Println(" Update Mobile = " + mobile)
			fmt.Println(" Update Salary = " + salary)
			fmt.Println(" Update Targets 	= " + targets)
			fmt.Println(" Update Role = " + role)

			fmt.Println(" ###################################### ")

			orgUser := app.Org.GetOrgUser()

			if orgUser == nil {

				data.Error = true
				data.ErrorMsg = "No session available"

			} else {

				orgInvoke := invoke.OrgInvoke{
					User: orgUser,
					Role: creatorRole,
				}

				err := orgInvoke.UpdateUserFromLedger(email, name, mobile, age, salary, targets, role)

				if err != nil {
					data.Error = true
					data.ErrorMsg = err.Error()
				} else {

					data, err := data.Setup(orgUser, false)
					if err != nil {
						data.Error = true
						data.ErrorMsg = err.Error()
					}

					data.UpdateUser = email
					data.Update = true
				}
			}
		} else {
			data.Error = true
			data.ErrorMsg = "No Update Request Received"
		}

		data.Response = true
		//renderTemplate(w, r, "index.html", data)
		http.Redirect(w, r, "/index.html", 302)
	})
}
