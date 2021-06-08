package html

import (
	"privateledger/blockchain/invoke"
	"fmt"
	"net/http"
	"strings"
)

func (app *HtmlApp) DeletePageHandler() func(http.ResponseWriter, *http.Request) {

	return app.hasSavedToken(func(w http.ResponseWriter, r *http.Request, token string) {

		data := &Data{}

		if r.FormValue("deleteSubmitted") == "true" {

			email := r.FormValue("deleteEmail")
			role := r.FormValue("deleteRole")
			owner := r.FormValue("deleteOwner")
			creatorRole := r.FormValue("deleteCreatorRole")
			targets := r.FormValue("deleteTargets")

			fmt.Println(" Delete Email = " + email)
			fmt.Println(" Delete Role = " + role)
			fmt.Println(" Delete Owner = " + owner)

			orgUser := app.Org.GetOrgUser()

			if orgUser == nil {

				data.Error = true
				data.ErrorMsg = "No session available"

			} else {

				orgInvoke := invoke.OrgInvoke{
					User: orgUser,
					Role: creatorRole,
				}

				orgSetup := orgUser.Setup.ChooseORG(strings.ToLower(owner))

				err := orgUser.RemoveUser(email, orgSetup.OrgCaID, orgSetup.CaClient)

				if err != nil {
					fmt.Println("Error Removing User - " + email + " :: " + err.Error())
				}

				// ReInitialize to Session Org

				_ = orgUser.Setup.ChooseORG(strings.ToLower(orgUser.Setup.OrgName))

				err = orgInvoke.DeleteUserFromLedger(email, targets, role)

				if err != nil {

					data.Error = true
					data.ErrorMsg = err.Error()

					fmt.Println("Error Deleting User from ledger - " + email + " :: " + err.Error())

				} else {

					if strings.EqualFold(email, orgUser.Username) {
						data.Logout = true
						http.Redirect(w, r, "/logout", 302)
					} else {

						data, err := data.Setup(orgUser, false)
						if err != nil {
							data.Error = true
							data.ErrorMsg = err.Error()
						}

						data.DeleteUser = email
						data.Delete = true
					}
				}
			}

		} else {
			data.Error = true
			data.ErrorMsg = "No Delete Request Received"
		}

		data.Response = true
		//renderTemplate(w, r, "index.html", data)
		http.Redirect(w, r, "/index.html", 302)
	})
}
