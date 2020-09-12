package netsuite

import ( 
    b64 "encoding/base64"
    "fmt"
    "math/rand"
    "strconv"
    "crypto/hmac"
	"crypto/sha1"
    "net/url"
    "strings"
    "time"
    "sort"
)


func generateNonce() string {
    var nonce string
    rand.Seed(time.Now().UnixNano())

    for i := 0; i < 8; i++ {
        nonce += strconv.Itoa(rand.Intn(10))
    }
   
    return nonce
}

func splitURL(urlPath string) (string, string){
    splitted := strings.Split(urlPath, "?")

    return splitted[0], splitted[1]
}


func escapeURL(str *string) {
	*str = url.QueryEscape(*str)
    *str = strings.ReplaceAll(*str, "-", "%2B")
	*str = strings.ReplaceAll(*str, "_", "%2F")
}


func generateBodyHash(body string) string{
	encodedBytes := hmac.New(sha1.New, []byte(body))
	bodyHash := b64.URLEncoding.EncodeToString(encodedBytes.Sum(nil))
	escapeURL(&bodyHash)
	
	return bodyHash
}


func generateSignature(baseString string, cred credentials) string {
	key := cred.Consumer.Secret + "&" + cred.Token.Secret
    signature := hmac.New(sha1.New, []byte(key))
    signature.Write([]byte(baseString))
	sha1Signature := b64.URLEncoding.EncodeToString(signature.Sum(nil))
    escapeURL(&sha1Signature)
	
	return sha1Signature
}

func insertURLParameters(OAuthHeader map[string]string, urlParameters string){
    var aux []string
    splittedParameters := strings.Split(urlParameters, "&")
    for i := 0; i < len(splittedParameters); i++{
        aux = strings.Split(splittedParameters[i], "=")
        OAuthHeader[aux[0]] = aux[1]
    }
}

func sortMap(mapObject map[string]string) []string{
    keys := make([]string, 0, len(mapObject))
    for k := range mapObject {
        keys = append(keys, k)
    }

    sort.Strings(keys)

    return keys
}


func makeBaseString(httpMethod string, basePath string, OAuthHeader map[string]string, sortedKeys[]string) string {
    var baseString string
    
    for _, key := range sortedKeys {
        baseString += key + "=" + OAuthHeader[key] + "&"
    }
    stringLength := len(baseString)
    baseString = baseString[:stringLength-1]
    baseString = url.QueryEscape(baseString)
    baseString = httpMethod + "&" + basePath + "&" + baseString

    return baseString
}


func makeHeader(realm string, httpMethod string, urlPath string, body string, cred credentials) string {
    OAuthHeader := make(map[string]string)
	basePath, urlParameters := splitURL(urlPath)
    insertURLParameters(OAuthHeader, urlParameters)
    basePath = url.QueryEscape(basePath)
	OAuthHeader["oauth_timestamp"] = fmt.Sprintf("%d", time.Now().Unix())
	OAuthHeader["oauth_nonce"] = generateNonce()
    OAuthHeader["oauth_body_hash"] = generateBodyHash(body)
    OAuthHeader["oauth_token"] = cred.Token.Key
    OAuthHeader["oauth_consumer_key"] = cred.Consumer.Key
    OAuthHeader["oauth_signature_method"] = "HMAC-SHA1"
    OAuthHeader["oauth_version"] = "1.0"
    sortedKeys := sortMap(OAuthHeader)
    baseString := makeBaseString(httpMethod, basePath, OAuthHeader, sortedKeys)
	
    sha1Signature := generateSignature(baseString, cred)

	authorizationHeader := fmt.Sprintf(`OAuth realm="%s", oauth_version="1.0", oauth_nonce="%s", oauth_timestamp="%s", oauth_token="%s", oauth_consumer_key="%s", `+
                                       `oauth_body_hash="%s", oauth_signature_method="HMAC-SHA1", oauth_signature="%s"`, 
						   realm, OAuthHeader["oauth_nonce"], OAuthHeader["oauth_timestamp"], OAuthHeader["oauth_token"], OAuthHeader["oauth_consumer_key"], OAuthHeader["oauth_body_hash"], sha1Signature)
	
	return authorizationHeader
}