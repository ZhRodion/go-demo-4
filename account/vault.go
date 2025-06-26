package account

import (
	"demo/password/files"
	"encoding/base64"
	"encoding/json"
	"time"

	"github.com/fatih/color"
)

type Vault struct {
	Accounts  []Account `json:"accounts"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewVault() *Vault {
	file, err := files.ReadFile("vault.json")

	if err != nil {
		return &Vault{
			Accounts:  []Account{},
			UpdatedAt: time.Now(),
		}
	}

	var vault Vault
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
	files.WriteFile(string(data), "vault.json")
}

func (vault *Vault) ToBytes() ([]byte, error) {
	file, err := json.Marshal(vault)

	if err != nil {
		return nil, err
	}

	return file, nil
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
	files.WriteFile(string(jsonData), filename)
	return nil
}

// LoadBin загружает бинарные данные из JSON файла
func LoadBin(filename string) ([]byte, error) {
	file, err := files.ReadFile(filename)
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
