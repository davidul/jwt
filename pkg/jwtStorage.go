package pkg

import (
	"encoding/json"
	"fmt"
	"os"
)

var jwtMap map[string]string

func init() {
	jwtMap = make(map[string]string)
}

func AddToMap(key string, value string) {
	jwtMap[key] = value
}

func GetToken(key string) string {
	return jwtMap[key]
}

func ListTokens() map[string]string {
	return jwtMap
}

func DeleteToken(key string) {
	delete(jwtMap, key)
}

func ClearTokens() {
	jwtMap = make(map[string]string)
}

func GetTokenCount() int {
	return len(jwtMap)
}

func GetTokenKeys() []string {
	keys := make([]string, 0, len(jwtMap))
	for k := range jwtMap {
		keys = append(keys, k)
	}
	return keys
}

func ToJSON() string {
	marshal, err := json.Marshal(jwtMap)
	if err != nil {
		println(err)
	}
	return string(marshal)
}

func FromJSON(jsonString string) {
	err := json.Unmarshal([]byte(jsonString), &jwtMap)
	if err != nil {
		println(err)
	}
}

func WriteToFile(filename string) error {

	var file *os.File
	_, statError := os.Stat(filename)
	if os.IsExist(statError) {
		file, statError = os.OpenFile(filename, os.O_CREATE|os.O_RDWR, 0644)
	} else {
		file, statError = os.OpenFile(filename, os.O_CREATE|os.O_RDWR, 0644)
	}

	if statError != nil {
		fmt.Print("Error: ")
		println(statError)
		return statError
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			println(err)
		}
	}(file)

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "\t")
	for i, k := range GetTokenKeys() {
		println(i, k)
	}
	err := encoder.Encode(jwtMap)
	if err != nil {
		fmt.Println("Error Encoding: ")
		fmt.Println(err)
		return err
	}

	return nil
}

func ReadFromFile(filename string) error {
	open, err := os.Open(filename)
	if err != nil {
		return err
	}

	defer func(open *os.File) {
		err := open.Close()
		if err != nil {
			println(err)
		}
	}(open)

	decoder := json.NewDecoder(open)
	err = decoder.Decode(&jwtMap)
	if err != nil {
		println(err)
		return err
	}

	return nil
}
