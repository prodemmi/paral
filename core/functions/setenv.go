package functions

import (
	"fmt"
	"os"
)

func (f *Function) setenv(key string, value interface{}) (interface{}, error) {
	strVal := fmt.Sprint(value)
	err := os.Setenv(key, strVal)
	return nil, err
}
