# NetSuite requests Go package

Using this package, you'll be able to call RESTLET APIs using OAuth 1.0.

### Credentials
First of all, create an integration app and, after that, create an access token. With that, use this JSON file example to put your credentials:
```json
{
  "token": {
    "key": "INSERT-YOUR-TOKEN-KEY",
    "secret": "INSERT-YOUR-TOKEN-SECRET"
  },
  "consumer": {
    "key": "INSERT-YOUR-CONSUMER-KEY",
    "secret": "INSERT-YOUR-CONSUMER-SECRET"
  }
}
```

### Examples
In [examples directory]("https://github.com/igorgbianchi/netsuite-restlet-requests-go/tree/master/examples"), there are POST and GET implementations. PUT is similar to POST, just change the HTTP method.

### Contact
For any information ask me by email: igorgbianchi@gmail.com
