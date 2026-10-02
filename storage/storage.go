package storage

import (
	"encoding/json"
	"os"
)

func SaveToFile(filepath string, data any) error {
	bytes, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(filepath, bytes, 0644)
}

func ReadFromFile(filepath string, target any) error {
	bytes, err := os.ReadFile(filepath)
	if err != nil {
		return err
	}
	err = json.Unmarshal(bytes, target)
	if err != nil {
		return err
	}
	return nil
}
