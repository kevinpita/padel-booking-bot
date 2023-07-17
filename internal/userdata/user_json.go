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

func readJSONFile(data []byte) (map[string]Login, error) {
	loginMap := make(map[string]Login)
	err := json.Unmarshal(data, &loginMap)
	if err != nil {
		return nil, err
	}

	return loginMap, nil
}

func (l *loginFile) get(id int) (*Login, error) {
	f, err := os.ReadFile(l.filepath)
	if err != nil {
		return nil, fmt.Errorf("couldn't read file: %s. Error: %w", l.filepath, err)
	}

	loginMap, errJSON := readJSONFile(f)
	if errJSON != nil {
		return nil, fmt.Errorf("couldn't parse file: %s. Error: %w", l.filepath, err)
	}

	login, ok := loginMap[strconv.Itoa(id)]
	if !ok {
		return nil, fmt.Errorf("%d id could not be found", id)
	}

	return &login, nil
}
