package rest

import (
	"privateledger/blockchain/invoke"
	"privateledger/web/model"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"unicode"
)

func (app *RestApp) RegisterHandler(w http.ResponseWriter, r *http.Request) {

	var userdata model.ModelUserData

	_ = json.NewDecoder(r.Body).Decode(&userdata)

	orgName := userdata.Org
	email := userdata.Email
	name := userdata.Name
	mobile := userdata.Mobile
	age := userdata.Age
	salary := userdata.Salary
	role := userdata.Role
	password := hash(userdata.Password)
	verifyErr := verifyPassword(userdata.Password)

	fmt.Println(" ####### Rest Input for Register ####### ")

	fmt.Println(" Email = " + email)
	fmt.Println(" Password = " + password)
	fmt.Println(" Role = " + role)

	fmt.Println(" Name = " + name)
	fmt.Println(" Mobile = " + mobile)
	fmt.Println(" Age = " + age)
	fmt.Println(" Salary = " + salary)
	fmt.Println(" ###################################### ")

	if verifyErr != nil && len(verifyErr.Error()) > 0 {

		respondJSON(w, map[string]string{"error": verifyErr.Error(), "message": "Password must contain at least one number and one uppercase and lowercase letter, and at least 8 or more characters"})
	}

	Org, err := app.Org.InitializeOrg(orgName)
	if err != nil {
		respondJSON(w, map[string]string{"error": "failed to invoke user " + err.Error()})
	}

	orgUser, err := Org.RegisterUserWithCA(orgName, email, password, role)

	orgInvoke := invoke.OrgInvoke{
		User: orgUser,
	}

	if err != nil {
		respondJSON(w, map[string]string{"error": "Unable to Register : " + err.Error()})
	} else {

		token := app.processAuthentication(w, email)

		if len(token) > 0 {

			err := orgInvoke.InvokeCreateUser(name, age, mobile, salary)
			if err != nil {
				respondJSON(w, map[string]string{"error": "failed to invoke user " + err.Error()})
			} else {

				respondJSON(w, map[string]string{
					"token":  token,
					"name":   name,
					"email":  email,
					"age":    age,
					"mobile": mobile,
					"salary": salary,
				})
			}

		} else {
			respondJSON(w, map[string]string{"error": "Failed to generate token"})
		}
	}
}

func verifyPassword(password string) error {
	var uppercasePresent bool
	var lowercasePresent bool
	var numberPresent bool
	var specialCharPresent bool
	const minPassLength = 8
	const maxPassLength = 64
	var passLen int
	var errorString string

	for _, ch := range password {
		switch {
		case unicode.IsNumber(ch):
			numberPresent = true
			passLen++
		case unicode.IsUpper(ch):
			uppercasePresent = true
			passLen++
		case unicode.IsLower(ch):
			lowercasePresent = true
			passLen++
		case unicode.IsPunct(ch) || unicode.IsSymbol(ch):
			specialCharPresent = true
			passLen++
		case ch == ' ':
			passLen++
		}
	}
	appendError := func(err string) {
		if len(strings.TrimSpace(errorString)) != 0 {
			errorString += ", " + err
		} else {
			errorString = err
		}
	}
	if !lowercasePresent {
		appendError("lowercase letter missing")
	}
	if !uppercasePresent {
		appendError("uppercase letter missing")
	}
	if !numberPresent {
		appendError("atleast one numeric character required")
	}
	if !specialCharPresent {
		appendError("special character missing")
	}
	if !(minPassLength <= passLen && passLen <= maxPassLength) {
		appendError(fmt.Sprintf("password length must be between %d to %d characters long", minPassLength, maxPassLength))
	}

	if len(errorString) != 0 {
		return fmt.Errorf(errorString)
	}
	return nil
}
