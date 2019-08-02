package html

import (
	"fmt"
	"net/http"
)

func (app *HtmlApp) AuthenticateHandler(w http.ResponseWriter, r *http.Request) {

	
	if r.FormValue("signinSubmitted") == "true" {

		fmt.Println(" Login ")

		app.LoginHandler(w,r)

	} else if r.FormValue("signupSubmitted") == "true" {

		fmt.Println(" Register ")

		app.RegisterHandler(w,r)
	}
	renderTemplate(w, r, "register.html", nil)
}