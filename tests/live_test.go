package tests

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

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
	rd := rand.New(rand.NewSource(time.Now().UnixNano()))

	template.DisplayName = fmt.Sprintf("模板添加 --测试 %d", rd.Intn(99999))
	template.VideoWidth = rd.Intn(800)
	template.VideoHeight = rd.Intn(800)
	template.AudioBitrate = rd.Intn(1000)
	template.VideoBitrate = rd.Intn(1000)
	template.FrameRate = rd.Intn(50)
	template.TranscodeType = 0
	template.AdvancedArguments = "{\"rc\":\"CBR\",\"gopSize\":4000,\"bFrame\":0,\"refFrame\":1,\"qualityLevel\":null,\"interlace\":-1,\"smartBorder\":1}"

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
	rd := rand.New(rand.NewSource(time.Now().UnixNano()))
	template.ID = "605f8dbe0cec4773b9df64cc88ce89ef"
	template.DisplayName = fmt.Sprintf("模板修改 --测试 %d", rd.Intn(99999))
	template.VideoWidth = rd.Intn(800)
	template.VideoHeight = rd.Intn(800)
	template.AudioBitrate = rd.Intn(1000)
	template.VideoBitrate = rd.Intn(1000)
	template.FrameRate = rd.Intn(50)
	template.TranscodeType = 0
	template.AdvancedArguments = "{\"rc\":\"CBR\",\"gopSize\":4000,\"bFrame\":0,\"refFrame\":1,\"qualityLevel\":null,\"interlace\":-1,\"smartBorder\":1}"

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

	param := live.Template{ID: "605f8dbe0cec4773b9df64cc88ce89ef,dc00bb3f3d8045449a12d9186b90df3b"}

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
	rd := rand.New(rand.NewSource(time.Now().UnixNano()))

	livetype.Name = fmt.Sprintf("直播类型 %d", rd.Intn(99999))
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

func TestGetLivelogos(t *testing.T) {
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

	templateRsp, err := livemanager.GetLivelogos("", "")

	if err != nil {
		//t.Error("sign error ")
		t.Errorf(err.Error())
	} else if !templateRsp.Success {
		t.Errorf("Success:%v, msg:%s", templateRsp.Success, templateRsp.Message)

	}

	fmt.Println(templateRsp)
}

func TestCreatLivelogo(t *testing.T) {

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

	livelogo := live.Livelogo{}
	rd := rand.New(rand.NewSource(time.Now().UnixNano()))
	livelogo.Name = fmt.Sprintf("添加直播logo ---  测试 %d", rd.Intn(99999))
	livelogo.Uri = "bj-hsy-images/kii721n6/watermarkLive/65112b7ae7394313a679020c3543cae9.jpg"
	livelogo.Height = 200
	livelogo.Width = 200
	livelogo.OffsetHeight = 1
	livelogo.OffsetWidth = 1
	livelogo.Resize = 100
	livelogo.Opacity = 100
	templateRsp, err := livemanager.CreateLivelogo(livelogo)

	if err != nil {
		//t.Error("sign error ")
		t.Errorf(err.Error())
	}

	fmt.Println(templateRsp)
}

func TestEditLivelogo(t *testing.T) {

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

	livelogo := live.Livelogo{}
	rd := rand.New(rand.NewSource(time.Now().UnixNano()))
	livelogo.ID = "57039b8eacce498b8dff5cdf47f4f326"
	livelogo.Name = fmt.Sprintf("修改直播logo --测试 %d", rd.Intn(99999))
	livelogo.Uri = "bj-hsy-images/kii721n6/watermarkLive/65112b7ae7394313a679020c3543cae9.jpg"
	livelogo.Height = 200
	livelogo.Width = 200
	livelogo.OffsetHeight = 1
	livelogo.OffsetWidth = 1
	livelogo.Resize = 100
	livelogo.Opacity = 100
	templateRsp, err := livemanager.EditLivelogo(livelogo)

	if err != nil {
		//t.Error("sign error ")
		t.Errorf(err.Error())
	}

	fmt.Println(templateRsp)
}

func TestDeleteLivelogo(t *testing.T) {

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

	livelogo := live.Livelogo{ID: "57039b8eacce498b8dff5cdf47f4f326,e99149aefe814bed96d2151e3d36652e"}
	templateRsp, err := livemanager.DeleteLivelogo(livelogo)

	if err != nil {
		//t.Error("sign error ")
		t.Errorf(err.Error())
	}

	fmt.Println(templateRsp)
}
