package file_ops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

// folder should be same name as package, file can be any name
// use uppercase to expose to other packages
func GetFloatFromFile(fileName string) (float64, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return 1000, errors.New("failed to find file")
	}
	valueText := string(data)
	value, err := strconv.ParseFloat(valueText, 64)
	if err != nil {
		return 1000, errors.New("failed to parse stored value")
	}
	return value, err
}

func WriteFloatToFile(fileName string, value float64) {
	valueTxt := fmt.Sprint(value)
	os.WriteFile(fileName, []byte(valueTxt), 0644)
}
