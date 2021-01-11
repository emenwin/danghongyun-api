package tests

import (
	"fmt"
	"testing"

	"github.com/emenwin/danghongyun-api/auth"
)

func TestSign(t *testing.T) {

	accessKey := "a020e193-0f1"
	accessSecret := "5GcXHNYdAVVdFW0yervG"

	action := "getUser"
	version := "2.0"
	timestamp := "1466488681033" //ms

	cred := auth.New(accessKey, accessSecret)
	q, sign := cred.Sign2(action, version, timestamp)

	result := "3d864184117e240ad4def677c48fbba509a1d0d48ea5dfb9e914c587ae3ce5bf"
	if sign != result {
		//t.Error("sign error ")
		t.Errorf("Sign error \nnow:(%s) \nshould:(%s)", sign, result)
	}

	fmt.Println(q, "---", sign)
	// 	accessKey = "a020e193-0f1"
	// accessSecret = "5GcXHNYdAVVdFW0yervG"

	queryParams := map[string]string{
		"action":    "ugcCreateChannel",
		"version":   "2.0",
		"timestamp": "1466488681033",
		"name":      "直播测试频道",
	}

	q, sign = cred.SignExt(queryParams)
	result = "69a22291e517179e76e55ad7c21269c17462e102adf0eb0526cc5012d9b4cd9a"
	if sign != result {
		//t.Error("sign error ")
		t.Errorf("Sign error \nnow:(%s) \nshould:(%s)", sign, result)
	}

	fmt.Println(q, "---", sign)

}
