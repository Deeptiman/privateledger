package html

import (
	"fmt"
	"net/http"
)

func (app *HtmlApp) LogoutHandler(w http.ResponseWriter, r *http.Request) {

	orgUser := app.Org.GetOrgUser()

	if orgUser == nil {
					
		fmt.Println("Error Logout , Session is unavailable")
	} else {

			fmt.Println("Delete Token - "+savedToken["Token"]) 

			delete(savedToken, "token")

			orgUser.Logout()

			fmt.Println("Success Logout ")

			http.Redirect(w, r, "/login.html", 302)
		
	}
}
