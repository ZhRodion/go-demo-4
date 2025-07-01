package main

import (
	"demo/password/account"
	"demo/password/files"
	"fmt"
)

// InputProvider интерфейс для получения пользовательского ввода
type InputProvider interface {
	PromptData(prompt string) string
	PromptInt(prompt string) int
}

// ConsoleInputProvider реализация для консольного ввода
type ConsoleInputProvider struct{}

func NewConsoleInputProvider() *ConsoleInputProvider {
	return &ConsoleInputProvider{}
}

func (c *ConsoleInputProvider) PromptData(prompt string) string {
	fmt.Print(prompt)
	var res string
	fmt.Scanln(&res)
	return res
}

func (c *ConsoleInputProvider) PromptInt(prompt string) int {
	fmt.Print(prompt)
	var res int
	fmt.Scanln(&res)
	return res
}

func main() {
	// Создаем единую зависимость для работы с базой данных
	db := files.NewJsonDB("vault.json")

Menu:
	for {
		if variant := menuSelector(); variant != 0 {
			switch variant {
			case 1:
				createAccount(db)
			case 2:
				searchAccount(db)
			case 3:
				deleteAccount(db)
			default:
				break Menu
			}
		}
	}
}

func createAccount(db account.Db) {
	inputProvider := NewConsoleInputProvider()
	login := inputProvider.PromptData("Login: ")
	password := inputProvider.PromptData("Password: ")
	url := inputProvider.PromptData("URL: ")

	myAccount, err := account.NewAccount(login, password, url)

	if err != nil {
		fmt.Println(err)
		return
	}

	vault := account.NewVault(db)
	vault.AddAccount(myAccount)
}

func menuSelector() int {
	inputProvider := NewConsoleInputProvider()
	fmt.Println("Choose variant: ")
	fmt.Println("1. Create account")
	fmt.Println("2. Search account")
	fmt.Println("3. Delete account")
	fmt.Println("4. Exit")

	return inputProvider.PromptInt("")
}

func searchAccount(db account.Db) {
	foundAccount, err := account.FindAccount(db)

	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	fmt.Println("Аккаунт найден:")
	foundAccount.OutputPassword()
}

func deleteAccount(db account.Db) {
	err := account.DeleteAccount(db)

	if err != nil {
		fmt.Println("Ошибка при удалении аккаунта:", err)
		return
	}

	fmt.Println("Аккаунт успешно удален!")
}
