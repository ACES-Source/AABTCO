// Code generated by "scopegen"; DO NOT EDIT.
package accountapplication

var Scopes = map[string]string{
	"https://auth.bnk.to/accountapplication.read": "View accountapplication data",
	"https://auth.bnk.to/accountapplication.write": "Manage accountapplication data",
}

var AuthScopes = map[string][]string{
	"/accountapplication.AccountApplicationService/CreateAccountApplication": []string{"https://auth.bnk.to/accountapplication.write"},
	"/accountapplication.AccountApplicationService/GetAccountApplication": []string{"https://auth.bnk.to/accountapplication.read"},
	"/accountapplication.AccountApplicationService/GetAccountApplications": []string{"https://auth.bnk.to/accountapplication.read"},
	"/accountapplication.AccountApplicationService/UpdateAccountApplicationStatus": []string{"https://auth.bnk.to/accountapplication.write"},
}