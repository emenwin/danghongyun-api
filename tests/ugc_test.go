package tests

import (
	"fmt"
	"testing"

	"github.com/emenwin/danghongyun-api/auth"
	"github.com/emenwin/danghongyun-api/ugc"
)

func TestCreateChannel(t *testing.T) {
	accessKey := "275ee3cd-6fb"
	accessSecret := "5GcXHNYdAVVdFW0yervG"

	if TestAccessKey != "" {
		accessKey = TestAccessKey
		accessSecret = TestSecretKey
	}

	fmt.Println(accessKey)
	fmt.Println(accessSecret)

	cred := auth.New(accessKey, accessSecret)

	livemanager := ugc.NewLiveManager(cred)
	name := "新增频道1"
	templateRsp, err := livemanager.CreateChannel(name)

	if err != nil {
		//t.Error("sign error ")
		t.Errorf(err.Error())
	}

	fmt.Println(templateRsp)
}
