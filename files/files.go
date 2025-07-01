package files

import (
	"fmt"
	"os"
)

type JsonDb struct {
	filename string
}

func NewJsonDB(name string) *JsonDb {
	return &JsonDb{
		filename: name,
	}
}

func (db *JsonDb) Read(name string) ([]byte, error) {
	data, err := os.ReadFile(db.filename)

	if err != nil {
		fmt.Println("Error opening file", err)
		return nil, err
	}

	return data, nil
}

func (db *JsonDb) Write(content []byte, name string) {
	file, err := os.Create(db.filename)

	if err != nil {
		fmt.Println("Error creating file", err)
		return
	}

	defer file.Close()
	_, err = file.Write(content)

	if err != nil {
		file.Close()
		fmt.Println("Error writing to file", err)
		return
	}

	fmt.Println("File written successfully")
}
