package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptrace"
	"net/http/httputil"
	"net/url"
	"strings"

	"log"
)

func init() {
	log.SetPrefix("TRACE: ")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
}

//UserAgent agent
var UserAgent = "Golang danghongyun/client package"

//DefaultClient Default client
var DefaultClient = Client{&http.Client{Transport: http.DefaultTransport}}

// DebugMode 用来打印调试信息
var DebugMode = true

// DeepDebugInfo 深度打印
var DeepDebugInfo = true

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

func newRequest(ctx context.Context, method, reqURL string, headers http.Header, body io.Reader) (req *http.Request, err error) {
	req, err = http.NewRequest(method, reqURL, body)
	if err != nil {
		return
	}

	if headers == nil {
		headers = http.Header{}
	}

	req.Header = headers
	req = req.WithContext(ctx)

	if DebugMode {
		trace := &httptrace.ClientTrace{
			GotConn: func(connInfo httptrace.GotConnInfo) {
				remoteAddr := connInfo.Conn.RemoteAddr()
				log.Println(fmt.Sprintf("\n\nNetwork: %s, Remote ip:%s, URL: %s", remoteAddr.Network(), remoteAddr.String(), req.URL))
			},
		}
		req = req.WithContext(httptrace.WithClientTrace(req.Context(), trace))
		bs, bErr := httputil.DumpRequest(req, DeepDebugInfo)
		if bErr != nil {
			err = bErr

			return
		}
		log.Println("\n", string(bs))
	}
	return
}

// DoRequest do request
func (r Client) DoRequest(ctx context.Context, method, reqURL string, headers http.Header) (resp *http.Response, err error) {
	req, err := newRequest(ctx, method, reqURL, headers, nil)
	if err != nil {
		return
	}
	return r.Do(ctx, req)
}

//DoRequestWith request
func (r Client) DoRequestWith(ctx context.Context, method, reqURL string, headers http.Header, body io.Reader,
	bodyLength int) (resp *http.Response, err error) {

	req, err := newRequest(ctx, method, reqURL, headers, body)
	if err != nil {
		return
	}
	req.ContentLength = int64(bodyLength)
	return r.Do(ctx, req)
}

//DoRequestWith64 request
func (r Client) DoRequestWith64(ctx context.Context, method, reqURL string, headers http.Header, body io.Reader,
	bodyLength int64) (resp *http.Response, err error) {

	req, err := newRequest(ctx, method, reqURL, headers, body)
	if err != nil {
		return
	}
	req.ContentLength = bodyLength
	return r.Do(ctx, req)
}

// DoRequestWithForm request
func (r Client) DoRequestWithForm(ctx context.Context, method, reqURL string, headers http.Header,
	data map[string][]string) (resp *http.Response, err error) {

	if headers == nil {
		headers = http.Header{}
	}
	headers.Add("Content-Type", "application/x-www-form-urlencoded")

	requestData := url.Values(data).Encode()
	if method == "GET" || method == "HEAD" || method == "DELETE" {
		if strings.ContainsRune(reqURL, '?') {
			reqURL += "&"
		} else {
			reqURL += "?"
		}
		return r.DoRequest(ctx, method, reqURL+requestData, headers)
	}

	return r.DoRequestWith(ctx, method, reqURL, headers, strings.NewReader(requestData), len(requestData))
}

// DoRequestWithJSON request
func (r Client) DoRequestWithJSON(ctx context.Context, method, reqURL string, headers http.Header,
	data interface{}) (resp *http.Response, err error) {

	reqBody, err := json.Marshal(data)
	if err != nil {
		return
	}

	if DebugMode {
		log.Println("reqBody:\n", string(reqBody))
	}
	if headers == nil {
		headers = http.Header{}
	}

	headers.Add("Content-Type", "application/json")
	return r.DoRequestWith(ctx, method, reqURL, headers, bytes.NewReader(reqBody), len(reqBody))
}

// DoRequestQueryWithJSON request
func (r Client) DoRequestQueryWithJSON(ctx context.Context, method, reqURL string, headers http.Header,
	data map[string][]string) (resp *http.Response, err error) {

	if headers == nil {
		headers = http.Header{}
	}
	headers.Add("Content-Type", "application/json")

	requestData := url.Values(data).Encode()
	//if method == "GET" || method == "HEAD" || method == "DELETE" {
	// if strings.ContainsRune(reqURL, '?') {
	// 	reqURL += "&"
	// } else {
	// 	reqURL += "?"
	// }
	return r.DoRequest(ctx, method, reqURL+requestData, headers)
	//}

}

//Do do
func (r Client) Do(ctx context.Context, req *http.Request) (resp *http.Response, err error) {

	if _, ok := req.Header["User-Agent"]; !ok {
		req.Header.Set("User-Agent", UserAgent)
	}

	resp, err = r.Client.Do(req)
	return
}

// --------------------------------------------------------------------

//ErrorInfo errorinfo
type ErrorInfo struct {
	HTTPCode int    `json:"http_code"` //http code
	Code     int    `json:"code"`      // 整型错误码，0成功，其它值失败
	Message  string `json:"message"`   // 字符串，可能包含错误或者别的一些信息
	Result   string `json:"result"`    // 字符串，返回的结果
	Success  bool   `json:"success"`   // true 或 false

}

//ErrorDetail error
func (r *ErrorInfo) ErrorDetail() string {

	msg, _ := json.Marshal(r)
	return string(msg)
}

//Error error
func (r *ErrorInfo) Error() string {

	return r.Message
}

//ResponseError 返回错误
func ResponseError(resp *http.Response) (err error) {

	e := &ErrorInfo{
		HTTPCode: resp.StatusCode,
		Success:  false,
	}

	return e
}

//Call get call
func (r Client) Call(ctx context.Context, ret interface{}, method, reqURL string, headers http.Header) (err error) {

	resp, err := r.DoRequestWith(ctx, method, reqURL, headers, nil, 0)
	if err != nil {
		return err
	}
	return CallRet(ctx, ret, resp)
}

//CallWithJSON api call
func (r Client) CallWithJSON(ctx context.Context, ret interface{}, method, reqURL string, headers http.Header,
	param interface{}) (err error) {

	resp, err := r.DoRequestWithJSON(ctx, method, reqURL, headers, param)
	if err != nil {
		return err
	}
	return CallRet(ctx, ret, resp)
}

//CallWithJSONQuery api call
func (r Client) CallWithJSONQuery(ctx context.Context, ret interface{}, method, reqURL string, headers http.Header,
	param map[string][]string) (err error) {

	resp, err := r.DoRequestQueryWithJSON(ctx, method, reqURL, headers, param)
	if err != nil {
		return err
	}
	return CallRet(ctx, ret, resp)
}

//CallWithForm do requeset
func (r Client) CallWithForm(ctx context.Context, ret interface{}, method, reqURL string, headers http.Header,
	param map[string][]string) (err error) {

	resp, err := r.DoRequestWithForm(ctx, method, reqURL, headers, param)
	if err != nil {
		return err
	}
	return CallRet(ctx, ret, resp)
}

//CallRet 返回对象转换
func CallRet(ctx context.Context, ret interface{}, resp *http.Response) (err error) {

	defer func() {
		io.Copy(ioutil.Discard, resp.Body)
		resp.Body.Close()
	}()

	if DebugMode {
		bs, dErr := httputil.DumpResponse(resp, DeepDebugInfo)
		if dErr != nil {
			err = dErr
			log.Println(dErr.Error())

			return
		}
		log.Println("\n\n", string(bs))
	}

	if resp.StatusCode/100 == 2 {
		if ret != nil && resp.ContentLength != 0 {
			err = json.NewDecoder(resp.Body).Decode(ret)
			if err != nil {
				return
			}
		}
		return nil
	}
	return ResponseError(resp)
}
