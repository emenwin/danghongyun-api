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
	template.VideoWidth = rand.Intn(1000)
	template.VideoHeight = rand.Intn(1000)
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

	templateRsp, err := livemanager.GetTemplates("VR高清", "2", 2)

	if err != nil {
		//t.Error("sign error ")
		t.Errorf(err.Error())
	} else if !templateRsp.Success {
		t.Errorf("Success:%v, msg:%s", templateRsp.Success, templateRsp.Message)

	}

	fmt.Println(templateRsp)
}

func TestEditTemplate(t *testing.T) {
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
	template.ID = "26c3ebf49cf64153a270ace17a074e58"
	template.DisplayName = "修改直播模板名称"

	templateRsp, err := livemanager.EditTemplate(template)

	if err != nil {
		//t.Error("sign error ")
		t.Errorf(err.Error())
	}

	fmt.Println(templateRsp)
}

func TestDeleteTemplate(t *testing.T) {
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

	param := live.Template{ID: "26c3ebf49cf64153a270ace17a074e58"}

	templateRsp, err := livemanager.DeleteTemplates(param)

	if err != nil {
		//t.Error("sign error ")
		t.Errorf(err.Error())
	}

	fmt.Println(templateRsp)

}

func TestGetLivetypes(t *testing.T) {
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

	templateRsp, err := livemanager.GetLivetypes("", "3daa3949c294456ab40273ab2d8afba1")

	if err != nil {
		//t.Error("sign error ")
		t.Errorf(err.Error())
	} else if !templateRsp.Success {
		t.Errorf("Success:%v, msg:%s", templateRsp.Success, templateRsp.Message)

	}

	fmt.Println(templateRsp)
}

func TestCreatLivetype(t *testing.T) {

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

	livetype := live.Livetype{}
	livetype.Name = fmt.Sprintf("直播类型 %d", rand.Intn(99999))
	templateRsp, err := livemanager.CreateLivetype(livetype)

	if err != nil {
		//t.Error("sign error ")
		t.Errorf(err.Error())
	}

	fmt.Println(templateRsp)
}

func TestEditLivetype(t *testing.T) {

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

	livetype := live.Livetype{}
	livetype.ID = "331f756633ce4ec887d1df021cc51b35"
	livetype.Name = "修改直播类型--测试"
	templateRsp, err := livemanager.EditLivetype(livetype)

	if err != nil {
		//t.Error("sign error ")
		t.Errorf(err.Error())
	}

	fmt.Println(templateRsp)
}

func TestDeleteLivetype(t *testing.T) {

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

	livetype := live.Livetype{ID: "e402a628a14a4ce29de447741c3ba46d"}

	templateRsp, err := livemanager.DeleteLivetype(livetype)

	if err != nil {
		//t.Error("sign error ")
		t.Errorf(err.Error())
	}

	fmt.Println(templateRsp)
}
