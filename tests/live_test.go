package tests

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/emenwin/danghongyun-api/auth"
	"github.com/emenwin/danghongyun-api/live"
)

func TestCreatTemplate(t *testing.T) {

	accessKey := "275ee3cd-6fb"
	accessSecret := "5GcXHNYdAVVdFW0yervG"

	if TestAccessKey != "" {
		accessKey = TestAccessKey
		accessSecret = TestSecretKey
	}

	fmt.Println(accessKey)
	fmt.Println(accessSecret)

	cred := auth.New(accessKey, accessSecret)

	livemanager := live.NewLiveManager(cred)

	template := live.Template{}
	template.DisplayName = fmt.Sprintf("TestCreatTemplate %d", rand.Intn(99999))
	template.Name = fmt.Sprintf("TestCreatTemplate Name %d", rand.Intn(99999))
	template.VideoWidth = rand.Intn(2048)
	template.VideoHeight = rand.Intn(2048)
	template.AudioBitrate = rand.Intn(1000)
	template.FrameRate = rand.Intn(1000)

	templateRsp, err := livemanager.TemplateCreate(template)

	if err != nil {
		//t.Error("sign error ")
		t.Errorf(err.Error())
	}

	fmt.Println(templateRsp)
}

func TestGetTemplates(t *testing.T) {

	accessKey := "275ee3cd-6fb"
	accessSecret := "5GcXHNYdAVVdFW0yervG"

	if TestAccessKey != "" {
		accessKey = TestAccessKey
		accessSecret = TestSecretKey
	}

	fmt.Println(accessKey)
	fmt.Println(accessSecret)

	cred := auth.New(accessKey, accessSecret)

	livemanager := live.NewLiveManager(cred)

	templateRsp, err := livemanager.GetTemplates("", "", -1)

	if err != nil {
		//t.Error("sign error ")
		t.Errorf(err.Error())
	} else if !templateRsp.Success {
		t.Errorf("Success:%v, msg:%s", templateRsp.Success, templateRsp.Message)

	}

	fmt.Println(templateRsp)
}
