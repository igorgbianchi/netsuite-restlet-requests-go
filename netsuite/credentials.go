package netsuite

import ( 
  "encoding/json"
  "io/ioutil"
  "os"
)

type credentials struct {
	Token fields `json:"token"`
	Consumer fields `json:"consumer"`
}

type fields struct {
  Key string `json:"key"`
  Secret string `json:"secret"`
}

type error interface {
	Error() string
}

func NewCredentials(path string) (credentials, error){
	var cred credentials
	jsonFile, err := os.Open(path)

	if err != nil {
		return cred, err
	}

	byteValue, err := ioutil.ReadAll(jsonFile)

	if err != nil {
		return cred, err
	}

	err = json.Unmarshal(byteValue, &cred)

	if err != nil {
		return cred, err
	}

	return cred, nil
}
