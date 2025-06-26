package main

import (
	"demo/password/account"
	"demo/password/files"
	"fmt"
)

func main() {
Menu:
	for {
		if variant := menuSelector(); variant != 0 {
			switch variant {
			case 1:
				createAccount()
			case 2:
				searchAccount()
			case 3:
				deleteAccount()
			default:
				break Menu
			}
		}
	}
}

func createAccount() {
	login := promtData("Login: ")
	password := promtData("Password: ")
	url := promtData("URL: ")

	myAccount, err := account.NewAccount(login, password, url)

	if err != nil {
		fmt.Println(err)
		return
	}

	vault := account.NewVault()
	vault.AddAccount(myAccount)
	data, err := vault.ToBytes()

	if err != nil {
		fmt.Println("Error marshalling account", err)
		return
	}

	files.WriteFile(string(data), "vault.json")
}

func promtData(promt string) string {
	fmt.Print(promt)
	var res string
	fmt.Scanln(&res)
	return res
}

func menuSelector() int {
	var variant int
	fmt.Println("Choose variant: ")
	fmt.Println("1. Create account")
	fmt.Println("2. Search account")
	fmt.Println("3. Delete account")
	fmt.Println("4. Exit")

	fmt.Scanln(&variant)
	return variant
}

func searchAccount() {
	foundAccount, err := account.FindAccount()

	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	fmt.Println("Аккаунт найден:")
	foundAccount.OutputPassword()
}

func deleteAccount() {
	err := account.DeleteAccount()

	if err != nil {
		fmt.Println("Ошибка при удалении аккаунта:", err)
		return
	}

	fmt.Println("Аккаунт успешно удален!")
}
