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
	timestamp := int64(1466488681033) //ms

	cred := auth.New(accessKey, accessSecret)
	sign := cred.Sign(action, version, timestamp)

	result := "3d864184117e240ad4def677c48fbba509a1d0d48ea5dfb9e914c587ae3ce5bf"
	if sign != result {
		//t.Error("sign error ")
		t.Errorf("Sign error \nnow:(%s) \nshould:(%s)", sign, result)
	}

	fmt.Println(accessKey)
	fmt.Println(accessSecret)
	fmt.Println(action)
	fmt.Println(timestamp)
	fmt.Println(version)
	fmt.Println(sign)
}
