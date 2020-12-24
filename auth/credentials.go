package auth

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"net/url"
	"sort"
	"strings"

	"fmt"

	"github.com/emenwin/danghongyun-api/utils"
)

// Credentials 鉴权类，用于生成请求签名签名
// https://developer.danghongyun.com/rest_api_new.html 文档
type Credentials struct {
	AccessKey string
	SecretKey string
}

//New 构建一个Credentials对象
func New(accessKey, secretKey string) *Credentials {
	return &Credentials{accessKey, secretKey}
}

// Sign 生成签名
// action: action
// version: version
// timestamp: timestamp 毫秒
// signature:请求签名，签名生成规则:对所有参数按照名称排序(CASE_INSENSITIVE_ORDER)生成一个字符串，并在字符串头上加上accessSecret，然后对该字符串调用[HmacSHA256](http://en.wikipedia.org/wiki/Hash-based_message_authentication_code) 算法生成signature。
// 每一个API，根据输入参数的不同，生成的签名都是不同的。
func (ath *Credentials) Sign(action string, version string, timestamp int64) (signature string) {
	h := hmac.New(sha256.New, []byte(ath.SecretKey))
	data := fmt.Sprintf("%saccessKey=%saction=%stimestamp=%dversion=%s", ath.SecretKey, ath.AccessKey, action, timestamp, version)
	h.Write([]byte(data))
	signature = hex.EncodeToString(h.Sum(nil))

	return
}

//NewSignParams new signparmas
func (ath *Credentials) NewSignParams(action string, version, timestamp string) map[string]string {

	queryParams := map[string]string{
		"action":    action,
		"version":   version,
		"timestamp": timestamp,
	}
	if timestamp == "" {
		queryParams["timestamp"] = utils.NowMillisStr()
	}
	return queryParams
}

//Sign2 sign
func (ath *Credentials) Sign2(action string, version, timestamp string) (queryString, signature string) {

	queryParams := map[string]string{
		"action":    action,
		"version":   version,
		"timestamp": timestamp,
	}
	return ath.SignExt(queryParams)
}

//SignExt sign
func (ath *Credentials) SignExt(queryParams map[string]string) (queryString, signature string) {
	h := hmac.New(sha256.New, []byte(ath.SecretKey))

	queryParams["accessKey"] = ath.AccessKey

	items := xHeaders{}
	for k, v := range queryParams {
		item := xItem{k, v}
		items = append(items, item)
	}
	sort.Sort(items)

	toSign := ath.SecretKey
	params := []string{}

	for _, item := range items {
		//注意：
		//url 参数需要escape
		//参与签名不需要
		kv := item.Name + "=" + item.Value
		toSign += kv
		kv2 := item.Name + "=" + url.QueryEscape(item.Value)

		params = append(params, kv2)
	}

	// fmt.Println("\n\n")
	// fmt.Println("queryParams:", queryParams)
	// fmt.Println("params:", params)
	// fmt.Println("toSign:", toSign)
	// fmt.Println("\n\n")

	h.Write([]byte(toSign))

	signature = hex.EncodeToString(h.Sum(nil))

	params = append(params, "signature="+signature)
	queryString = strings.Join(params, "&")

	//fmt.Println("queryString:", queryString)
	return
}

type (
	xItem struct {
		Name  string
		Value string
	}
	xHeaders []xItem
)

func (headers xHeaders) Len() int {
	return len(headers)
}

func (headers xHeaders) Less(i, j int) bool {
	if headers[i].Name < headers[j].Name {
		return true
	} else if headers[i].Name > headers[j].Name {
		return false
	} else {
		return headers[i].Value < headers[j].Value
	}
}

func (headers xHeaders) Swap(i, j int) {
	headers[i], headers[j] = headers[j], headers[i]
}
