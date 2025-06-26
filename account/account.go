package account

import (
	"errors"
	"math/rand/v2"
	"net/url"
	"time"

	"github.com/fatih/color"
)

var lettersRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

type Account struct {
	login    string
	password string
	url      string
}

type AccountWithTimesStamp struct {
	Account   Account
	createdAt time.Time
	updatedAt time.Time
}

func (account *Account) OutputPassword() {
	color.Cyan(account.login)
	color.Green(account.password)
	color.Blue(account.url)
}

func (account *Account) GeneratePassword(length int) {
	result := make([]rune, length)
	for i := range length {
		result[i] = lettersRunes[rand.IntN(len(lettersRunes))]
	}

	account.password = string(result)
}

func NewAccountWithTimesStamp(login, password, urlString string) (*AccountWithTimesStamp, error) {

	if login == "" {
		return nil, errors.New("login is required")
	}

	_, err := url.ParseRequestURI(urlString)

	if err != nil {
		return nil, errors.New("invalid url")
	}

	newAccount := &Account{
		login:    login,
		password: password,
		url:      urlString,
	}

	if len(password) == 0 {
		newAccount.GeneratePassword(12)
	}

	return &AccountWithTimesStamp{
		Account:   *newAccount,
		createdAt: time.Now(),
		updatedAt: time.Now(),
	}, nil
}
