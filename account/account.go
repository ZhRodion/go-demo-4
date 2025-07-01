package account

import (
	"demo/password/files"
	"encoding/json"
	"errors"
	"fmt"
	"math/rand/v2"
	"net/url"
	"time"

	"github.com/fatih/color"
)

var lettersRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

type Account struct {
	Login     string    `json:"login"`
	Password  string    `json:"password"`
	URL       string    `json:"url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (account *Account) OutputPassword() {
	color.Cyan(account.Login)
	color.Green(account.Password)
	color.Blue(account.URL)
}

func (account *Account) GeneratePassword(length int) {
	result := make([]rune, length)
	for i := range length {
		result[i] = lettersRunes[rand.IntN(len(lettersRunes))]
	}

	account.Password = string(result)
}

func NewAccount(login, password, urlString string) (*Account, error) {

	if login == "" {
		return nil, errors.New("login is required")
	}

	_, err := url.ParseRequestURI(urlString)

	if err != nil {
		return nil, errors.New("invalid url")
	}

	newAccount := &Account{
		Login:    login,
		Password: password,
		URL:      urlString,
	}

	if len(password) == 0 {
		newAccount.GeneratePassword(12)
	}

	return &Account{
		Login:     newAccount.Login,
		Password:  newAccount.Password,
		URL:       newAccount.URL,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}

func (account *Account) ToBytes() ([]byte, error) {
	file, err := json.Marshal(account)

	if err != nil {
		fmt.Println("Error marshalling account", err)
		return nil, err
	}

	return file, nil
}

func (account *Account) SearchAccount() (*Account, error) {

	fmt.Println("Enter login: ")
	var login string
	fmt.Scanln(&login)

	if login == "" {
		return nil, errors.New("login is required")
	}

	accountLogin := account.Login

	if accountLogin == login {
		return account, nil
	}

	return nil, errors.New("account not found")
}

func FindAccount() (*Account, error) {
	fmt.Println("Enter login: ")
	var inputLogin string
	fmt.Scanln(&inputLogin)

	if inputLogin == "" {
		return nil, errors.New("login is required")
	}

	file, err := files.NewJsonDB("vault.json").Read("vault.json")

	if err != nil {
		color.Red("Error reading file", err.Error())
		return nil, err
	}

	var vault Vault
	err = json.Unmarshal(file, &vault)

	if err != nil {
		color.Red("Error unmarshalling file", err.Error())
		return nil, err
	}

	for _, account := range vault.Accounts {
		if account.Login == inputLogin {
			return &account, nil
		}
	}

	return nil, errors.New("account not found")
}

func DeleteAccount() error {
	fmt.Println("Enter login to delete: ")
	var inputLogin string
	fmt.Scanln(&inputLogin)

	if inputLogin == "" {
		return errors.New("login is required")
	}

	// Читаем текущий vault
	file, err := files.NewJsonDB("vault.json").Read("vault.json")
	if err != nil {
		color.Red("Error reading file", err.Error())
		return err
	}

	var vault Vault
	err = json.Unmarshal(file, &vault)
	if err != nil {
		color.Red("Error unmarshalling file", err.Error())
		return err
	}

	// Ищем и удаляем аккаунт
	found := false
	for i, account := range vault.Accounts {
		if account.Login == inputLogin {
			// Удаляем аккаунт из слайса
			vault.Accounts = append(vault.Accounts[:i], vault.Accounts[i+1:]...)
			found = true
			break
		}
	}

	if !found {
		return errors.New("account not found")
	}

	// Сохраняем обновленный vault обратно в файл
	data, err := vault.ToBytes()
	if err != nil {
		color.Red("Error marshalling vault", err.Error())
		return err
	}

	files.NewJsonDB("vault.json").Write(data, "vault.json")
	return nil
}
