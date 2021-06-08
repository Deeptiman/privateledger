package html

import (
	"privateledger/blockchain/invoke"
	"privateledger/blockchain/org"
	"privateledger/chaincode/model"
	"fmt"
	"net/http"
	"strconv"
)

type Data struct {
	Error    bool
	ErrorMsg string
	Success  bool
	Response bool

	Admin bool
	User  bool

	SessionOrg      string
	SessionUserData *model.User

	AllUsersData []model.User
	SharingOrgs  []string

	AllHistoryData []model.HistoryData
	HistoryUser    string
	History        bool

	ShareUser string
	Share     bool

	UpdateUser string
	Update     bool

	DeleteUser string
	Delete     bool

	Logout bool

	ALL    string
	WRITE  string
	DELETE string

	CustomOrg1 string
	CustomOrg2 string
	CustomOrg3 string
	CustomOrg4 string
}

func (app *HtmlApp) IndexPageHandler() func(http.ResponseWriter, *http.Request) {

	return app.hasSavedToken(func(w http.ResponseWriter, r *http.Request, token string) {

		data := &Data{}

		orgUser := app.Org.GetOrgUser()

		if orgUser == nil {

			data.Error = true
			data.ErrorMsg = "No session available"

		} else {

			data, err := data.Setup(orgUser, true)
			if err != nil && data != nil {
				data.Response = true
				data.Error = true
				data.ErrorMsg = err.Error()
			}
			data.ALL = strconv.Itoa(int(model.ALL))
			data.WRITE = strconv.Itoa(int(model.WRITE))
			data.DELETE = strconv.Itoa(int(model.DELETE))
			renderTemplate(w, r, "index.html", data)

		}

	})
}

func (data *Data) Setup(orgUser *org.OrgUser, needHistory bool) (*Data, error) {

	orgInvoke := invoke.OrgInvoke{
		User: orgUser,
	}

	/* Session User Data */

	SessionUserData, err := orgInvoke.GetUserFromLedger(orgUser.Username, needHistory)

	if err != nil {
		return nil, err
	}

	data.SessionUserData = SessionUserData
	data.SessionOrg = orgUser.Setup.OrgName

	/* Is Logged In User is Admin? */
	if model.IsAdmin(SessionUserData.Role) {

		allUsersData, err := orgInvoke.GetAllUsersFromLedger()
		if err != nil {
			return nil, err
		}

		data.Admin = true
		data.AllUsersData = allUsersData

	} else {
		data.User = true
	}

	data.SharingOrgs = orgUser.Setup.FilteredOrgNames()
	data.Success = true
	data.Response = true

	fmt.Println(" Sharing Orgs --- ", len(data.SharingOrgs))
	for _, org := range data.SharingOrgs {

		fmt.Println(" Sharing Org -- " + org)

	}

	return &Data{

		Success:         data.Success,
		Response:        data.Response,
		Admin:           data.Admin,
		User:            data.User,
		SessionOrg:      data.SessionOrg,
		SessionUserData: data.SessionUserData,
		AllUsersData:    data.AllUsersData,
		SharingOrgs:     data.SharingOrgs,
		CustomOrg1:      model.GetCustomOrgName("org1"),
		CustomOrg2:      model.GetCustomOrgName("org2"),
		CustomOrg3:      model.GetCustomOrgName("org3"),
		CustomOrg4:      model.GetCustomOrgName("org4"),
	}, nil

}
