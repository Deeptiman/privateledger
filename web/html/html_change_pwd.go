package html

import (
	"fmt"
	"net/http"
)

func (app *HtmlApp) OpenChangePwdHandler() func(http.ResponseWriter, *http.Request) {

	return app.hasSavedToken(func(w http.ResponseWriter, r *http.Request, token string) {

		data := &Data {}
		orgUser := app.Org.GetOrgUser()
		
		if r.FormValue("openChangePwdSubmitted") == "true" {

			if orgUser == nil {

				data.Error = true
				data.ErrorMsg = "No session available"
				data.Response = true
				renderTemplate(w, r, "change_password.html", data)

			} else {

				data, err := data.Setup(orgUser, false)
				if err != nil {
					data.Error = true
					data.ErrorMsg = err.Error()
					data.Response = true
					renderTemplate(w, r, "change_password.html", data)
				}  

				data.Success = true
				data.Response = true
				renderTemplate(w, r, "change_password.html", data)

			}			
		} else {

				data, _ := data.Setup(orgUser, false)				
				data.Response = true
				renderTemplate(w, r, "change_password.html", data)				
		}
	})
}

func (app *HtmlApp) ChangePwdHandler() func(http.ResponseWriter, *http.Request) {

	return app.hasSavedToken(func(w http.ResponseWriter, r *http.Request, token string) {

		data := &Data{}

		data.Response = true

		orgUser := app.Org.GetOrgUser()

		if orgUser == nil {

			data.Error = true
			data.ErrorMsg = "No session available"

		} else {

			if r.FormValue("changePwdSubmitted") == "true" {

				data, err := data.Setup(orgUser, false)
				if err != nil {
					data.Error = true
					data.ErrorMsg = err.Error()
				}
			 
				emailValue := data.SessionUserData.Email
				roleValue := data.SessionUserData.Role

				oldPwdValue := hash(r.FormValue("oldPwd"))
				newPwdValue := hash(r.FormValue("newPwd"))

				err = orgUser.Setup.ChangePassword(emailValue, roleValue, oldPwdValue, newPwdValue)

				if err != nil {
					fmt.Println("Error : %s " + err.Error())
					fmt.Errorf("Unable to Change user pwd Error Msg : %s", err.Error())

					data.Error = true
					data.ErrorMsg = err.Error()					

					renderTemplate(w, r, "change_password.html", data)

				} else {
					http.Redirect(w, r, "/logout", 302)
				}
			}
		}
	})

}
