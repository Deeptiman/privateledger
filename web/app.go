package web

import (
	"privateledger/web/html"
	"privateledger/web/rest"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func ServeWeb(app *html.HtmlApp) {

	fs := http.FileServer(http.Dir("web/assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	http.HandleFunc("/token_auth", app.TokenAuthHandler())

	http.HandleFunc("/login.html", app.LoginHandler)
	http.HandleFunc("/register.html", app.RegisterHandler)
	http.HandleFunc("/index.html", app.IndexPageHandler())
	http.HandleFunc("/edit", app.EditPageHandler())
	http.HandleFunc("/share", app.SharePageHandler())
	http.HandleFunc("/delete", app.DeletePageHandler())
	http.HandleFunc("/history.html", app.GetHistoryHandler())
	http.HandleFunc("/change_password.html", app.OpenChangePwdHandler())
	http.HandleFunc("/change_pwd", app.ChangePwdHandler())
	http.HandleFunc("/logout", app.LogoutHandler)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		http.Redirect(w, r, "/login.html", http.StatusTemporaryRedirect)
	})

	fmt.Println("Listening (http://localhost:6000) ...")
	http.ListenAndServe(":6000", nil)

}

func RestServe(app *rest.RestApp) {

	r := mux.NewRouter()

	r.HandleFunc("/api/read_users", app.GetAllUsersDataHandler()).Methods("GET")
	r.HandleFunc("/api/read_user", app.GetUserDataByEmailHandler()).Methods("GET")

	r.HandleFunc("/api/user_login", app.LoginHandler).Methods("POST")
	r.HandleFunc("/api/user_register", app.RegisterHandler).Methods("POST")

	r.HandleFunc("/api/change_password", app.ChangePwdHandler()).Methods("POST")

	r.HandleFunc("/api/update_user", app.UpdateUserHandler()).Methods("PUT")

	r.HandleFunc("/api/delete_user", app.DeleteUserHandler()).Methods("DELETE")

	fmt.Println("Listening (http://localhost:4000) ...")
	log.Fatal(http.ListenAndServe(":4000", r))
}
