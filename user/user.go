package user

import (
	"os"
	"path/filepath"
)

const baseDir string = "users"

func InitUserDir() error {
	err := os.MkdirAll(baseDir, 0755)
	if err != nil {
		return err
	}
	return nil
}

func UserExist(username string) bool {
	_, err := os.Stat(GetUserPath(username))
	if err != nil {
		return false
	}
	return true

}

func GetUserPath(username string) string {
	return filepath.Join(baseDir, username)
}
func GetUserTasks(username string) string {
	return filepath.Join(baseDir, username, "tasks.json")
}

func CreateUser(username string) error {
	err := os.Mkdir(GetUserPath(username), 0755)
	if err != nil {
		return err
	}
	return nil
}

func DeleteUser(username string) error {
	err := os.RemoveAll(GetUserPath(username))
	if err != nil {
		return err
	}
	return nil
}

// func UpdateUserName(username string) error {

// }
