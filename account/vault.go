package account

import (
	"demo/password/files"
	"encoding/base64"
	"encoding/json"
	"time"

	"github.com/fatih/color"
)

type Db interface {
	Read(string) ([]byte, error)
	Write([]byte, string)
}

// VaultInterface интерфейс для работы с хранилищем аккаунтов
type VaultInterface interface {
	AddAccount(*Account)
	ToBytes() ([]byte, error)
	GetAccounts() []Account
}

type Vault struct {
	Accounts  []Account `json:"accounts"`
	UpdatedAt time.Time `json:"updated_at"`
}

type VaultWithDb struct {
	Vault
	db Db
}

func NewVault(db Db) *VaultWithDb {
	file, err := db.Read("vault.json")

	if err != nil {
		return &VaultWithDb{
			Vault: Vault{
				Accounts:  []Account{},
				UpdatedAt: time.Now(),
			},
			db: db,
		}
	}

	var vault VaultWithDb
	err = json.Unmarshal(file, &vault)

	if err != nil {
		color.Red(err.Error())
	}

	return &vault
}

func (vault *Vault) AddAccount(account *Account) {
	vault.Accounts = append(vault.Accounts, *account)
	vault.UpdatedAt = time.Now()
	data, err := vault.ToBytes()
	if err != nil {
		color.Red("Error writing file", err.Error())
	}
	files.NewJsonDB("vault.json").Write(data, "vault.json")
}

func (vault *Vault) ToBytes() ([]byte, error) {
	file, err := json.Marshal(vault)

	if err != nil {
		return nil, err
	}

	return file, nil
}

func (vault *Vault) GetAccounts() []Account {
	return vault.Accounts
}

// SaveBin сохраняет бинарные данные в JSON формате с base64 кодированием
func SaveBin(data []byte, filename string) error {
	// Структура для JSON с бинарными данными
	binaryData := struct {
		Data      string    `json:"data"` // base64 encoded data
		CreatedAt time.Time `json:"created_at"`
		Size      int       `json:"size"`
	}{
		Data:      base64.StdEncoding.EncodeToString(data),
		CreatedAt: time.Now(),
		Size:      len(data),
	}

	// Маршалим в JSON
	jsonData, err := json.MarshalIndent(binaryData, "", "  ")
	if err != nil {
		return err
	}

	// Сохраняем в файл
	files.NewJsonDB(filename).Write(jsonData, filename)
	return nil
}

// LoadBin загружает бинарные данные из JSON файла
func LoadBin(filename string) ([]byte, error) {
	file, err := files.NewJsonDB(filename).Read(filename)
	if err != nil {
		return nil, err
	}

	var binaryData struct {
		Data      string    `json:"data"`
		CreatedAt time.Time `json:"created_at"`
		Size      int       `json:"size"`
	}

	err = json.Unmarshal(file, &binaryData)
	if err != nil {
		return nil, err
	}

	// Декодируем из base64
	data, err := base64.StdEncoding.DecodeString(binaryData.Data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
