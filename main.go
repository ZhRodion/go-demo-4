package main

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"net/url"
	"time"
)

type account struct {
	login    string
	password string
	url      string
}

type accountWithTimesStamp struct {
	account   account
	createdAt time.Time
	updatedAt time.Time
}

func (account *account) outputPassword() {
	fmt.Println(account.login, account.password, account.url)
}

func (account *account) generatePassword(length int) {
	result := make([]rune, length)
	for i := range length {
		result[i] = lettersRunes[rand.IntN(len(lettersRunes))]
	}

	account.password = string(result)
}

func newAccountWithTimesStamp(login, password, urlString string) (*accountWithTimesStamp, error) {

	if login == "" {
		return nil, errors.New("login is required")
	}

	_, err := url.ParseRequestURI(urlString)

	if err != nil {
		return nil, errors.New("invalid url")
	}

	newAccount := &account{
		login:    login,
		password: password,
		url:      urlString,
	}

	if password == "" {
		newAccount.generatePassword(12)
	}

	return &accountWithTimesStamp{
		account:   *newAccount,
		createdAt: time.Now(),
		updatedAt: time.Now(),
	}, nil
}

var lettersRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func main() {

	login := promtData("Login: ")
	password := promtData("Password: ")
	url := promtData("URL: ")

	myAccount, err := newAccountWithTimesStamp(login, password, url)

	if err != nil {
		fmt.Println(err)
		return
	}

	myAccount.account.generatePassword(12)
	myAccount.account.outputPassword()

}

func promtData(promt string) string {
	fmt.Print(promt)
	var res string
	fmt.Scanln(&res)
	return res
}
