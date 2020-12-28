package tests

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/emenwin/danghongyun-api/auth"
	"github.com/emenwin/danghongyun-api/video"
)

func TestGetUserInfo(t *testing.T) {
	accessKey := "275ee3cd-6fb"
	accessSecret := "5GcXHNYdAVVdFW0yervG"

	if TestAccessKey != "" {
		accessKey = TestAccessKey
		accessSecret = TestSecretKey
	}

	fmt.Println(accessKey)
	fmt.Println(accessSecret)

	cred := auth.New(accessKey, accessSecret)

	livemanager := video.NewLiveManager(cred)

	templateRsp, err := livemanager.GetUserInfo()

	if err != nil {
		//t.Error("sign error ")
		t.Errorf(err.Error())
	}

	fmt.Println(templateRsp)
}

func TestGetToken(t *testing.T) {
	accessKey := "275ee3cd-6fb"
	accessSecret := "5GcXHNYdAVVdFW0yervG"

	if TestAccessKey != "" {
		accessKey = TestAccessKey
		accessSecret = TestSecretKey
	}

	fmt.Println(accessKey)
	fmt.Println(accessSecret)

	cred := auth.New(accessKey, accessSecret)

	livemanager := video.NewLiveManager(cred)

	templateRsp, err := livemanager.GetToken()

	if err != nil {
		//t.Error("sign error ")
		t.Errorf(err.Error())
	}

	fmt.Println(templateRsp)
}

func TestGetVideos(t *testing.T) {
	accessKey := "275ee3cd-6fb"
	accessSecret := "5GcXHNYdAVVdFW0yervG"

	if TestAccessKey != "" {
		accessKey = TestAccessKey
		accessSecret = TestSecretKey
	}

	fmt.Println(accessKey)
	fmt.Println(accessSecret)

	cred := auth.New(accessKey, accessSecret)

	livemanager := video.NewLiveManager(cred)
	searchParams := video.SearchVideoParam{}
	searchParams.Start = "1"
	searchParams.Number = "10"
	searchParams.Sort = "title"
	searchParams.Order = "desc"

	templateRsp, err := livemanager.GetVideos(searchParams)

	if err != nil {
		//t.Error("sign error ")
		t.Errorf(err.Error())
	}

	fmt.Println(templateRsp)
}

func TestGetVideocategory(t *testing.T) {
	accessKey := "275ee3cd-6fb"
	accessSecret := "5GcXHNYdAVVdFW0yervG"

	if TestAccessKey != "" {
		accessKey = TestAccessKey
		accessSecret = TestSecretKey
	}

	fmt.Println(accessKey)
	fmt.Println(accessSecret)

	cred := auth.New(accessKey, accessSecret)

	livemanager := video.NewLiveManager(cred)

	templateRsp, err := livemanager.GetVideocategory()

	if err != nil {
		//t.Error("sign error ")
		t.Errorf(err.Error())
	}

	fmt.Println(templateRsp)
}

func TestEditVideocategory(t *testing.T) {
	accessKey := "275ee3cd-6fb"
	accessSecret := "5GcXHNYdAVVdFW0yervG"

	if TestAccessKey != "" {
		accessKey = TestAccessKey
		accessSecret = TestSecretKey
	}

	fmt.Println(accessKey)
	fmt.Println(accessSecret)

	cred := auth.New(accessKey, accessSecret)

	livemanager := video.NewLiveManager(cred)
	params := video.Videocategory{ID: 1, Category: "修改分类 -- 测试"}

	templateRsp, err := livemanager.EditVideocategory(params)

	if err != nil {
		//t.Error("sign error ")
		t.Errorf(err.Error())
	}

	fmt.Println(templateRsp)
}

func TestCreateVideocategory(t *testing.T) {
	accessKey := "275ee3cd-6fb"
	accessSecret := "5GcXHNYdAVVdFW0yervG"

	if TestAccessKey != "" {
		accessKey = TestAccessKey
		accessSecret = TestSecretKey
	}

	fmt.Println(accessKey)
	fmt.Println(accessSecret)

	cred := auth.New(accessKey, accessSecret)

	livemanager := video.NewLiveManager(cred)
	params := video.Videocategory{Category: "新增分类 -- 测试"}

	templateRsp, err := livemanager.CreateVideocategory(params)

	if err != nil {
		//t.Error("sign error ")
		t.Errorf(err.Error())
	}

	fmt.Println(templateRsp)
}
func TestDeleteVideocategory(t *testing.T) {
	accessKey := "275ee3cd-6fb"
	accessSecret := "5GcXHNYdAVVdFW0yervG"

	if TestAccessKey != "" {
		accessKey = TestAccessKey
		accessSecret = TestSecretKey
	}

	fmt.Println(accessKey)
	fmt.Println(accessSecret)

	cred := auth.New(accessKey, accessSecret)

	livemanager := video.NewLiveManager(cred)
	ids := "1,2"

	templateRsp, err := livemanager.DeleteVideocategory(ids)

	if err != nil {
		//t.Error("sign error ")
		t.Errorf(err.Error())
	}

	fmt.Println(templateRsp)
}
func TestGetTemplateGroupList(t *testing.T) {
	accessKey := "275ee3cd-6fb"
	accessSecret := "5GcXHNYdAVVdFW0yervG"

	if TestAccessKey != "" {
		accessKey = TestAccessKey
		accessSecret = TestSecretKey
	}

	fmt.Println(accessKey)
	fmt.Println(accessSecret)

	cred := auth.New(accessKey, accessSecret)

	livemanager := video.NewLiveManager(cred)
	templateGroupType := 1
	templateRsp, err := livemanager.GetTemplateGroupList(templateGroupType)

	if err != nil {
		//t.Error("sign error ")
		t.Errorf(err.Error())
	}

	fmt.Println(templateRsp)
}

func TestCreateOrEditTemplateGroup(t *testing.T) {
	accessKey := "275ee3cd-6fb"
	accessSecret := "5GcXHNYdAVVdFW0yervG"

	if TestAccessKey != "" {
		accessKey = TestAccessKey
		accessSecret = TestSecretKey
	}

	fmt.Println(accessKey)
	fmt.Println(accessSecret)

	cred := auth.New(accessKey, accessSecret)

	livemanager := video.NewLiveManager(cred)
	param := video.TemplateGroup{}
	param.Type = 0
	param.Name = "新增模板组 -- 测试"
	arr := []video.TranscodeTemplate{
		{
			Name:         "高清",
			VideoWidth:   1280,
			VideoHeight:  720,
			VideoBitrate: 1000,
			AudioBitrate: 192000,
			Format:       0,
			EncryptType:  6,
			DefaultVideo: 1,
		},
	}
	bytes, err := json.Marshal(arr)
	param.TranscodeTemplates = string(bytes)

	templateRsp, err := livemanager.CreateOrEditTemplateGroup(param)

	if err != nil {
		//t.Error("sign error ")
		t.Errorf(err.Error())
	}

	fmt.Println(templateRsp)
}

func TestDeleteTemplateGroup(t *testing.T) {
	accessKey := "275ee3cd-6fb"
	accessSecret := "5GcXHNYdAVVdFW0yervG"

	if TestAccessKey != "" {
		accessKey = TestAccessKey
		accessSecret = TestSecretKey
	}

	fmt.Println(accessKey)
	fmt.Println(accessSecret)

	cred := auth.New(accessKey, accessSecret)

	livemanager := video.NewLiveManager(cred)

	ids := "999,1000"

	templateRsp, err := livemanager.DeleteTemplateGroup(ids)

	if err != nil {
		//t.Error("sign error ")
		t.Errorf(err.Error())
	}

	fmt.Println(templateRsp)
}

func TestGetLogoGroupList(t *testing.T) {
	accessKey := "275ee3cd-6fb"
	accessSecret := "5GcXHNYdAVVdFW0yervG"

	if TestAccessKey != "" {
		accessKey = TestAccessKey
		accessSecret = TestSecretKey
	}

	fmt.Println(accessKey)
	fmt.Println(accessSecret)

	cred := auth.New(accessKey, accessSecret)

	livemanager := video.NewLiveManager(cred)

	templateRsp, err := livemanager.GetLogoGroupList()

	if err != nil {
		//t.Error("sign error ")
		t.Errorf(err.Error())
	}

	fmt.Println(templateRsp)
}

func TestCreateOrEditLogoGroup(t *testing.T) {
	accessKey := "275ee3cd-6fb"
	accessSecret := "5GcXHNYdAVVdFW0yervG"

	if TestAccessKey != "" {
		accessKey = TestAccessKey
		accessSecret = TestSecretKey
	}

	fmt.Println(accessKey)
	fmt.Println(accessSecret)

	cred := auth.New(accessKey, accessSecret)

	livemanager := video.NewLiveManager(cred)
	param := video.LogoGroup{}
	param.ID = 1
	param.Name = "修改台标组 -- 测试"
	arr := []video.Logos{
		{
			URI:     "bj-hsy-images/p2191m9e/watermark/1.png",
			Width:   100,
			Height:  50,
			LeftPos: 0,
			TopPos:  0,
			Opacity: 100,
			Resize:  100,
			Start:   0,
			End:     50,
		},
	}
	bytes, err := json.Marshal(arr)
	param.Logos = string(bytes)

	templateRsp, err := livemanager.CreateOrEditLogoGroup(param)

	if err != nil {
		//t.Error("sign error ")
		t.Errorf(err.Error())
	}

	fmt.Println(templateRsp)
}

func TestDeleteLogoGroup(t *testing.T) {
	accessKey := "275ee3cd-6fb"
	accessSecret := "5GcXHNYdAVVdFW0yervG"

	if TestAccessKey != "" {
		accessKey = TestAccessKey
		accessSecret = TestSecretKey
	}

	fmt.Println(accessKey)
	fmt.Println(accessSecret)

	cred := auth.New(accessKey, accessSecret)

	livemanager := video.NewLiveManager(cred)
	ids := "1"
	templateRsp, err := livemanager.DeleteLogoGroup(ids)

	if err != nil {
		//t.Error("sign error ")
		t.Errorf(err.Error())
	}

	fmt.Println(templateRsp)
}
