package main

import (
	"github.com/igorgbianchi/netsuite-restlet-requests-go/netsuite"
	"fmt"
)

type Body struct{
	Data []NFS `json:"data"`
}

type NFS struct{
	NfsID string `json:"nfsId"`
	NfsURL string `json:"nfsUrl"`
}

func main() {
	urlPath := "https://12346.restlets.api.netsuite.com/app/site/hosting/restlet.nl?script=1&deploy=1"
	httpMethod := "POST"
	realm := "123456"
	credentials, _ := netsuite.NewCredentials("examples/post/netsuite-credentials.json")
	body := Body {
		Data: []NFS {
			NFS{
				NfsID: "22121", NfsURL: "https://google.com",
			},
			NFS{
				NfsID: "22122", NfsURL: "https://bing.com",
			},
		},
	}

	bodyStr, _ := netsuite.StructToStr(body)
	response, err := netsuite.Request(httpMethod, urlPath, credentials, bodyStr, realm)
	if(err == nil){
		fmt.Println(response)
	}else{
		fmt.Println(err)
	}
}
