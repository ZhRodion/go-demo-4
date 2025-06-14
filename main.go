package main

import (
	"demo/password/account"
	"demo/password/files"
	"fmt"
)

func main() {

	login := promtData("Login: ")
	password := promtData("Password: ")
	url := promtData("URL: ")

	myAccount, err := account.NewAccountWithTimesStamp(login, password, url)

	if err != nil {
		fmt.Println(err)
		return
	}

	myAccount.Account.OutputPassword()
	files.WriteFile()
}

func promtData(promt string) string {
	fmt.Print(promt)
	var res string
	fmt.Scanln(&res)
	return res
}
