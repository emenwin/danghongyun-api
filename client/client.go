package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptrace"
	"net/http/httputil"
	"net/url"
	"strings"

	"github.com/qiniu/api.v7/v7/auth"
	"github.com/qiniu/api.v7/v7/internal/log"
)

//UserAgent agent
var UserAgent = "Golang danghongyun/client package"

//DefaultClient Default client
var DefaultClient = Client{&http.Client{Transport: http.DefaultTransport}}

// DebugMode 用来打印调试信息
var DebugMode = false

// DeepDebugInfo 深度打印
var DeepDebugInfo = false

// --------------------------------------------------------------------

// Client 负责发送HTTP请求到七牛接口服务器
type Client struct {
	*http.Client
}

// TurnOnDebug 开启Debug模式
func TurnOnDebug() {
	DebugMode = true
}

// --------------------------------------------------------------------

func newRequest(ctx context.Context, method, reqUrl string, headers http.Header, body io.Reader) (req *http.Request, err error) {
	req, err = http.NewRequest(method, reqUrl, body)
	if err != nil {
		return
	}

	if headers == nil {
		headers = http.Header{}
	}

	req.Header = headers
	req = req.WithContext(ctx)

	//check access token
	mac, t, ok := auth.CredentialsFromContext(ctx)
	if ok {
		err = mac.AddToken(t, req)
		if err != nil {
			return
		}
	}
	if DebugMode {
		trace := &httptrace.ClientTrace{
			GotConn: func(connInfo httptrace.GotConnInfo) {
				remoteAddr := connInfo.Conn.RemoteAddr()
				log.Debug(fmt.Sprintf("Network: %s, Remote ip:%s, URL: %s", remoteAddr.Network(), remoteAddr.String(), req.URL))
			},
		}
		req = req.WithContext(httptrace.WithClientTrace(req.Context(), trace))
		bs, bErr := httputil.DumpRequest(req, DeepDebugInfo)
		if bErr != nil {
			err = bErr
			return
		}
		log.Debug(string(bs))
	}
	return
}

// DoRequest do request
func (r Client) DoRequest(ctx context.Context, method, reqUrl string, headers http.Header) (resp *http.Response, err error) {
	req, err := newRequest(ctx, method, reqUrl, headers, nil)
	if err != nil {
		return
	}
	return r.Do(ctx, req)
}

//DoRequestWith request
func (r Client) DoRequestWith(ctx context.Context, method, reqUrl string, headers http.Header, body io.Reader,
	bodyLength int) (resp *http.Response, err error) {

	req, err := newRequest(ctx, method, reqUrl, headers, body)
	if err != nil {
		return
	}
	req.ContentLength = int64(bodyLength)
	return r.Do(ctx, req)
}

//DoRequestWith64 request
func (r Client) DoRequestWith64(ctx context.Context, method, reqUrl string, headers http.Header, body io.Reader,
	bodyLength int64) (resp *http.Response, err error) {

	req, err := newRequest(ctx, method, reqUrl, headers, body)
	if err != nil {
		return
	}
	req.ContentLength = bodyLength
	return r.Do(ctx, req)
}

// DoRequestWithForm request
func (r Client) DoRequestWithForm(ctx context.Context, method, reqUrl string, headers http.Header,
	data map[string][]string) (resp *http.Response, err error) {

	if headers == nil {
		headers = http.Header{}
	}
	headers.Add("Content-Type", "application/x-www-form-urlencoded")

	requestData := url.Values(data).Encode()
	if method == "GET" || method == "HEAD" || method == "DELETE" {
		if strings.ContainsRune(reqUrl, '?') {
			reqUrl += "&"
		} else {
			reqUrl += "?"
		}
		return r.DoRequest(ctx, method, reqUrl+requestData, headers)
	}

	return r.DoRequestWith(ctx, method, reqUrl, headers, strings.NewReader(requestData), len(requestData))
}

// DoRequestWithJson request
func (r Client) DoRequestWithJson(ctx context.Context, method, reqUrl string, headers http.Header,
	data interface{}) (resp *http.Response, err error) {

	reqBody, err := json.Marshal(data)
	if err != nil {
		return
	}

	if headers == nil {
		headers = http.Header{}
	}
	headers.Add("Content-Type", "application/json")
	return r.DoRequestWith(ctx, method, reqUrl, headers, bytes.NewReader(reqBody), len(reqBody))
}

//Do do
func (r Client) Do(ctx context.Context, req *http.Request) (resp *http.Response, err error) {
	reqctx := req.Context()

	if _, ok := req.Header["User-Agent"]; !ok {
		req.Header.Set("User-Agent", UserAgent)
	}

	resp, err = r.Client.Do(req)
	return
}

// --------------------------------------------------------------------

//ErrorInfo errorinfo
type ErrorInfo struct {
	Err   string `json:"error,omitempty"`
	Key   string `json:"key,omitempty"`
	Errno int    `json:"errno,omitempty"`
	Code  int    `json:"code"`
}

//ErrorDetail error
func (r *ErrorInfo) ErrorDetail() string {

	msg, _ := json.Marshal(r)
	return string(msg)
}

//Error error
func (r *ErrorInfo) Error() string {

	return r.Err
}
