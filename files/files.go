package files

import (
	"encoding/json"
	"fmt"
	"os"
)

func ReadFile(name string) ([]byte, error) {
	data, err := os.ReadFile(name)

	if err != nil {
		fmt.Println("Error opening file", err)
		return nil, err
	}

	return data, nil
}

func IsValidJSON(filename string) bool {
	data, err := ReadFile(filename)
	if err != nil {
		return false
	}

	var js interface{}
	return json.Unmarshal(data, &js) == nil
}

func WriteFile(content, name string) {
	file, err := os.Create(name)

	if err != nil {
		fmt.Println("Error creating file", err)
		return
	}

	defer file.Close()
	_, err = file.WriteString(content)

	if err != nil {
		file.Close()
		fmt.Println("Error writing to file", err)
		return
	}

	fmt.Println("File written successfully")
}
