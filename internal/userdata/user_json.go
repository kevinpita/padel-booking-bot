package userdata

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
)

// Login email and password data used to log in into the api
// TODO: move login to its own file with the needed interface
// to be able to implement more than one 'database'
type Login struct {
	Email    string `json:"user"`
	Password string `json:"password"`
}

type loginFile struct {
	filepath string
}

func loginFromFile(file string) *loginFile {
	return &loginFile{filepath.Clean(file)}
}

func (l *loginFile) get(id int) (Login, error) {
	f, err := os.ReadFile(l.filepath)
	if err != nil {
		return Login{}, fmt.Errorf("couldn't read file: %s. Error: %w", l.filepath, err)
	}

	loginMap := make(map[string]Login)
	err = json.Unmarshal(f, &loginMap)

	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return Login{}, fmt.Errorf("couldn't parse file: %s. Error: %w", l.filepath, err)
	}

	login, ok := loginMap[strconv.Itoa(id)]
	if !ok {
		return Login{}, fmt.Errorf("%d id could not be found", id)
	}

	return login, nil
}
