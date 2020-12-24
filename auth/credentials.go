package auth

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"strconv"

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

//Sign2 sign
func (ath *Credentials) Sign2(action string, version string) (queryparams, timestamp, signature string) {
	h := hmac.New(sha256.New, []byte(ath.SecretKey))
	now := utils.NowMillis()
	timestamp = strconv.FormatInt(now, 10)
	data := fmt.Sprintf("%saccessKey=%saction=%stimestamp=%sversion=%s", ath.SecretKey, ath.AccessKey, action, timestamp, version)
	h.Write([]byte(data))
	signature = hex.EncodeToString(h.Sum(nil))

	queryparams = fmt.Sprintf("accessKey=%s&action=%s&version=%s&timestamp=%s&signature=%s",
		ath.AccessKey,
		action,
		version,
		timestamp,
		signature)
	return
}

type (
	xHeaderItem struct {
		HeaderName  string
		HeaderValue string
	}
	xHeaders []xHeaderItem
)

func (headers xHeaders) Len() int {
	return len(headers)
}

func (headers xHeaders) Less(i, j int) bool {
	if headers[i].HeaderName < headers[j].HeaderName {
		return true
	} else if headers[i].HeaderName > headers[j].HeaderName {
		return false
	} else {
		return headers[i].HeaderValue < headers[j].HeaderValue
	}
}

func (headers xHeaders) Swap(i, j int) {
	headers[i], headers[j] = headers[j], headers[i]
}
