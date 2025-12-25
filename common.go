package mysunxapi

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"net/url"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"time"

	jsoniter "github.com/json-iterator/go"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
)

const (
	BIT_BASE_10 = 10
	BIT_SIZE_64 = 64
	BIT_SIZE_32 = 32
)

type RequestType string

const (
	GET    = "GET"
	POST   = "POST"
	DELETE = "DELETE"
	PUT    = "PUT"
)

var NIL_REQBODY = []byte{}

var json = jsoniter.ConfigCompatibleWithStandardLibrary

var log = logrus.New()

func SetLogger(logger *logrus.Logger) {
	log = logger
}

func GetPointer[T any](v T) *T {
	return &v
}

func HmacSha256(secret, data string) []byte {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(data))
	return h.Sum(nil)
}

type MySunx struct {
}

const (
	SUNX_API_HTTP = "api.sunx.io"
	IS_GZIP       = true
)

type Client struct {
	AccessKey string
	SecretKey string
}

type RestClient struct {
	c *Client
}

func (*MySunx) NewPrivateRestClient(accessKey, secretKey string) *PrivateRestClient {
	return &PrivateRestClient{
		c: &Client{
			AccessKey: accessKey,
			SecretKey: secretKey,
		},
	}
}

func (*MySunx) NewPublicRestClient() *PublicRestClient {
	return &PublicRestClient{
		c: &Client{},
	}
}

type PublicRestClient RestClient

func (c *RestClient) PublicRestClient() *PublicRestClient {
	return &PublicRestClient{
		c: c.c,
	}
}

type PrivateRestClient RestClient

func (c *RestClient) PrivateRestClient() *PrivateRestClient {
	return &PrivateRestClient{
		c: c.c,
	}
}

func sunxCallApi[T any](url url.URL, reqBody []byte, method string) (*SunxRestRes[T], error) {

	headerMap := map[string]string{"Content-Type": "application/json"}

	body, err := RequestWithHeader(url.String(), reqBody, method, headerMap, IS_GZIP)
	if err != nil {
		return nil, err
	}

	res, err := handlerCommonRes[T](body)
	if err != nil {
		return nil, err
	}
	return res, res.handlerError()
}

type APIType int

const (
	REST APIType = iota
	WS
)

func sunxGetHost(apiType APIType) string {
	switch apiType {
	case REST:
		return SUNX_API_HTTP
	}
	return ""
}

// URL标准封装 带路径参数 不带签名
func sunxHandlerRequestAPIWithoutSignature[T any](apiType APIType, request *T, name string) url.URL {
	reqMap := sunxHandlerReq(request)
	queryStr := ""
	for k, v := range reqMap {
		queryStr += k + "=" + v + "&"
	}
	queryStr = strings.TrimRight(queryStr, "&")
	return url.URL{
		Scheme:   "https",
		Host:     sunxGetHost(apiType),
		Path:     name,
		RawQuery: queryStr,
	}
}

// URL标准封装 带路径参数 带签名
func sunxHandlerRequestAPIWithSignature[T any](client *Client, apiType APIType, method string, request *T, name string) url.URL {
	reqMap := sunxHandlerReq(request)
	reqMap["AccessKeyId"] = client.AccessKey
	reqMap["SignatureMethod"] = "HmacSHA256"
	reqMap["SignatureVersion"] = "2"
	timestamp := time.Now().UTC().Format("2006-01-02T15:04:05")
	reqMap["Timestamp"] = url.QueryEscape(timestamp)

	sortQueryFunc := func(m map[string]string) string {
		keys := []string{}
		for k := range m {
			keys = append(keys, k)
		}

		sort.Strings(keys)

		query := ""
		for i, k := range keys {
			if i < len(keys)-1 {
				query += k + "=" + m[k] + "&"
			} else {
				query += k + "=" + m[k]
			}
		}
		return query
	}
	query := sortQueryFunc(reqMap)
	hmacSha256Data := method + "\n" + sunxGetHost(apiType) + "\n" + name + "\n" + query
	sign := HmacSha256(client.SecretKey, hmacSha256Data)
	query += "&Signature=" + url.QueryEscape(base64.StdEncoding.EncodeToString(sign))

	// log.Warn(hmacSha256Data)
	// log.Warn(base64.StdEncoding.EncodeToString(sign))
	// log.Warn(method)
	// log.Warn(sunxGetHost(apiType))
	// log.Warn(name)
	// log.Warn(query)

	return url.URL{
		Scheme:   "https",
		Host:     sunxGetHost(apiType),
		Path:     name,
		RawQuery: query,
	}
}

func sunxHandlerReq[T any](req *T) map[string]string {
	reqMap := make(map[string]string)
	t := reflect.TypeOf(req)
	v := reflect.ValueOf(req)
	if v.IsNil() {
		return reqMap
	}
	t = t.Elem()
	v = v.Elem()
	count := v.NumField()
	for i := 0; i < count; i++ {
		argName := t.Field(i).Tag.Get("json")
		switch v.Field(i).Elem().Kind() {
		case reflect.String:
			reqMap[argName] = url.QueryEscape(v.Field(i).Elem().String())
		case reflect.Int, reflect.Int64:
			reqMap[argName] = strconv.FormatInt(v.Field(i).Elem().Int(), BIT_BASE_10)
		case reflect.Float32, reflect.Float64:
			reqMap[argName] = decimal.NewFromFloat(v.Field(i).Elem().Float()).String()
		case reflect.Bool:
			reqMap[argName] = strconv.FormatBool(v.Field(i).Elem().Bool())
		case reflect.Struct:
			sv := reflect.ValueOf(v.Field(i).Interface())
			ToStringMethod := sv.MethodByName("String")
			args := make([]reflect.Value, 0)
			result := ToStringMethod.Call(args)
			reqMap[argName] = url.QueryEscape(result[0].String())
		case reflect.Slice:
			s := v.Field(i).Interface()
			d, _ := json.Marshal(s)
			reqMap[argName] = url.QueryEscape(string(d))
		case reflect.Invalid:
		default:
			log.Errorf("req type error %s:%s", argName, v.Field(i).Elem().Kind())
		}
	}
	return reqMap
}

// func sunxHandlerReq[T any](req *T) string {
// 	var argBuffer bytes.Buffer

// 	t := reflect.TypeOf(req)
// 	v := reflect.ValueOf(req)
// 	if v.IsNil() {
// 		return ""
// 	}
// 	t = t.Elem()
// 	v = v.Elem()
// 	count := v.NumField()
// 	for i := 0; i < count; i++ {
// 		argName := t.Field(i).Tag.Get("json")
// 		switch v.Field(i).Elem().Kind() {
// 		case reflect.String:
// 			argBuffer.WriteString(argName + "=" + v.Field(i).Elem().String() + "&")
// 		case reflect.Int, reflect.Int64:
// 			argBuffer.WriteString(argName + "=" + strconv.FormatInt(v.Field(i).Elem().Int(), BIT_BASE_10) + "&")
// 		case reflect.Float32, reflect.Float64:
// 			argBuffer.WriteString(argName + "=" + decimal.NewFromFloat(v.Field(i).Elem().Float()).String() + "&")
// 		case reflect.Bool:
// 			argBuffer.WriteString(argName + "=" + strconv.FormatBool(v.Field(i).Elem().Bool()) + "&")
// 		case reflect.Struct:
// 			sv := reflect.ValueOf(v.Field(i).Interface())
// 			ToStringMethod := sv.MethodByName("String")
// 			args := make([]reflect.Value, 0)
// 			result := ToStringMethod.Call(args)
// 			argBuffer.WriteString(argName + "=" + result[0].String() + "&")
// 		case reflect.Slice:
// 			s := v.Field(i).Interface()
// 			d, _ := json.Marshal(s)
// 			argBuffer.WriteString(argName + "=" + url.QueryEscape(string(d)) + "&")
// 		case reflect.Invalid:
// 		default:
// 			log.Errorf("req type error %s:%s", argName, v.Field(i).Elem().Kind())
// 		}
// 	}
// 	return strings.TrimRight(argBuffer.String(), "&")
// }
