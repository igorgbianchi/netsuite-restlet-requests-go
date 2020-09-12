package main

import (
	"github.com/igorgbianchi/netsuite-restlet-requests-go/netsuite"
	"fmt"
)

func main() {
	urlPath := "https://123456.restlets.api.netsuite.com/app/site/hosting/restlet.nl?script=1&deploy=1"
	httpMethod := "GET"
	realm := "123456"
	credentials, err := netsuite.NewCredentials("examples/get/netsuite-credentials.json")
	if(err != nil){
		fmt.Println(err)
	}
	fmt.Println(credentials)
	response, err := netsuite.Request(httpMethod, urlPath, credentials, "", realm)
	if(err == nil){
		fmt.Println(response)
	}else{
		fmt.Println(err)
	}
}
