package html

import (
	"privateledger/blockchain/invoke"
	"fmt"
	"net/http"
)

func (app *HtmlApp) GetHistoryHandler() func(http.ResponseWriter, *http.Request) {

	return app.hasSavedToken(func(w http.ResponseWriter, r *http.Request, token string) {

		data := &Data{}

		if r.FormValue("historySubmitted") == "true" {

			email := r.FormValue("historyEmail")

			fmt.Println(" History Email = " + email)

			orgUser := app.Org.GetOrgUser()

			if orgUser == nil {

				data.Error = true
				data.ErrorMsg = "No session available"
				data.Response = true
				renderTemplate(w, r, "history.html", data)

			} else {

				orgInvoke := invoke.OrgInvoke{
					User: orgUser,
				}

				allHistoryData, err := orgInvoke.GetHistoryFromLedger(email)

				if err != nil {

					data.Error = true
					data.ErrorMsg = "Error getting history Data for user : " + email + " , msg = " + err.Error()
					data.Response = true
					renderTemplate(w, r, "history.html", data)

				} else {

					data, err := data.Setup(orgUser, false)
					if err != nil {
						data.Error = true
						data.ErrorMsg = err.Error()
					}
					fmt.Println(" HTML History Res = ", len(allHistoryData))
					data.History = len(allHistoryData) > 0
					data.HistoryUser = email
					data.AllHistoryData = allHistoryData

					data.Response = true
					renderTemplate(w, r, "history.html", data)
				}
			}
		}
	})

}
