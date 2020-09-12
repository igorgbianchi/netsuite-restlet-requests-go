package netsuite

import (
    "net/http"
    "time"
    "io/ioutil"
    "strings"
)

func Request(httpMethod string, urlPath string, cred credentials, body string, realm string) (string, error){
    var req *http.Request
    var err error
    authorizationHeader := makeHeader(realm, httpMethod, urlPath, "", cred)

    if(body != ""){
        req, err = http.NewRequest(httpMethod, urlPath, strings.NewReader(body))
    }else{
        req, err = http.NewRequest(httpMethod, urlPath, nil)
    }
    
    if(err != nil){
        return "", err
    }

    req.Header.Set("Content-Type", "application/json")
    req.Header.Set("Authorization", authorizationHeader)
    client := &http.Client{Timeout: time.Second * 240}

    resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
 
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
    
    bodyString := string(respBody)

	return bodyString, nil
}
