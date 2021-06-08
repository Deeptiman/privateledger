package html

import (
	"privateledger/blockchain/invoke"
	"fmt"
	"net/http"
	"strconv"
)

func (app *HtmlApp) SharePageHandler() func(http.ResponseWriter, *http.Request) {

	return app.hasSavedToken(func(w http.ResponseWriter, r *http.Request, token string) {

		data := &Data{}

		if r.FormValue("shareSubmitted") == "true" {

			email := r.FormValue("shareEmail")
			creatorRole := r.FormValue("shareRole")

			fmt.Println(" Share - User >> " + email)

			orgUser := app.Org.GetOrgUser()

			if orgUser == nil {

				data.Error = true
				data.ErrorMsg = "No session available"

			} else {

				orgInvoke := invoke.OrgInvoke{
					User: orgUser,
					Role: creatorRole,
				}

				orgNames := orgUser.Setup.FilteredOrgNames()

				var share bool

				fmt.Println(" ******** Share User Details - " + email)

				var accessList []int32
				var orgList []string

				for _, org := range orgNames {

					access := r.FormValue(org)
					i, _ := strconv.Atoi(access)
					a := int32(i)
					accessList = append(accessList, a)
					orgList = append(orgList, org)

					fmt.Println(" Share Request - "+org+" -- "+email+" -- "+access, i, a)

				}

				targets, _ := orgInvoke.CreateOrgTargets(email, accessList, orgList)

				fmt.Println(" Targets === " + targets)

				err := orgInvoke.ShareDataToOrg(email, orgList, accessList, targets)
				if err != nil {
					fmt.Println(" failed to share " + err.Error())
				}

				share = true

				data, err := data.Setup(orgUser, false)
				if err != nil {
					data.Error = true
					data.ErrorMsg = err.Error()
				}

				data.Share = share
				data.ShareUser = email
			}
		} else {
			data.Error = true
			data.ErrorMsg = "No Share Request Received"
		}

		data.Response = true
		//renderTemplate(w, r, "index.html", data)
		http.Redirect(w, r, "/index.html", 302)
	})

}
