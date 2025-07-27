package functions

import (
	"fmt"
	"os"
)

func (f *Function) getenv(envName string, defaultValue interface{}) (interface{}, error) {
	val, exists := os.LookupEnv(envName)
	if !exists {
		if defaultValue != nil {
			return fmt.Sprint(defaultValue), nil
		}
		return "", nil
	}
	return val, nil
}
