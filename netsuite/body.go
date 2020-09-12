package netsuite

import (
	"encoding/json"
)

func StructToStr(s interface{}) (string, error) {
	output, err := json.Marshal(s)
	if(err != nil){
		return "", err
	}

	return string(output), nil
}