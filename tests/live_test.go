package tests

import (
	"fmt"
	"testing"

	"github.com/emenwin/danghongyun-api/auth"
	"github.com/emenwin/danghongyun-api/live"
)

func TestCreatTemplate(t *testing.T) {

	accessKey := "a020e193-0f1"
	accessSecret := "5GcXHNYdAVVdFW0yervG"
	cred := auth.New(accessKey, accessSecret)

	livemanager := live.NewLiveManager(cred)

	param := live.TemplateReqParam{}
	tempalte, err := livemanager.TemplateCreate(param)

	if err != nil {
		//t.Error("sign error ")
		t.Errorf(err.Error())
	}

	fmt.Println(tempalte)
}
