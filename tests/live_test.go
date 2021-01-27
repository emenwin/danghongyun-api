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

	templateRsp, err := livemanager.GetTemplates("", "", 0)

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

	param := live.Template{ID: "166e264f20a9434e9dff8771f3433966"}

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

	templateRsp, err := livemanager.GetLivelogos("北体传媒", "")

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

func TestGetPageLivechannels(t *testing.T) {
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
	searchLivechannel := live.SearchLivechannelParam{}
	searchLivechannel.Number = 0
	searchLivechannel.Start = 0

	templateRsp, err := livemanager.GetPageLivechannels(searchLivechannel)

	if err != nil {
		//t.Error("sign error ")
		t.Errorf(err.Error())
	} else if !templateRsp.Success {
		t.Errorf("Success:%v, msg:%s", templateRsp.Success, templateRsp.Message)

	}

	fmt.Println(templateRsp)

}

func TestGetLivechannelByID(t *testing.T) {
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
	id := "338f277869454810bdea06c060fa5d53"

	templateRsp, err := livemanager.GetLivechannelByID(id)

	if err != nil {
		//t.Error("sign error ")
		t.Errorf(err.Error())
	} else if !templateRsp.Success {
		t.Errorf("Success:%v, msg:%s", templateRsp.Success, templateRsp.Message)

	}

	fmt.Println(templateRsp)

}

func TestGetLivechannelPlayURL(t *testing.T) {
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
	id := "39cc0761445b4de0ad9a937075f0b04e"

	templateRsp, err := livemanager.GetLivechannelsPlayURL(id)

	if err != nil {
		//t.Error("sign error ")
		t.Errorf(err.Error())
	} else if !templateRsp.Success {
		t.Errorf("Success:%v, msg:%s", templateRsp.Success, templateRsp.Message)

	}

	fmt.Println(templateRsp)

}

func TestCreateLivechannel(t *testing.T) {

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

	livechannel := live.CreateOrUpdatechannelParam{}
	// livechannel.FillInfo.FileType = 1
	rd := rand.New(rand.NewSource(time.Now().UnixNano()))
	livechannel.Name = fmt.Sprintf("添加直播频道20210127 ---  测试 %d", rd.Intn(99999))
	livechannel.SignalType = "1"
	livechannel.LiveType = 0
	livechannel.TypeName = "普通/300kbps/854x480"
	templateList := []live.CreateTemplateList{
		{
			TemplateID: "addByAuto-dcz-20181220-2",
			IsDefault:  "0",
		},
	}
	grouList := []live.CreateGroupList{
		{
			ProtocolType:            1,
			RtmpThirdPartyOutputURL: "",
			Rotation:                "0",
			IsRecord:                "1",
			NeedToMp4:               "1",
			IsTimeShift:             "0",
			TimeShiftDuration:       "",
			IsLiveEncrypt:           "0",
			LogoID:                  "5b99a7d13afd4e7eb4a52e15210982f6",
			OutputTemplateList:      templateList,
			IsFastEdit:              1,
		},
	}
	livechannel.OutputGroupList = grouList
	// livechannel.FillInfo.FileType = -1
	// livechannel.FillInfo.FileID = ""
	// livechannel.FillInfo.FileHTTPURL = ""
	// livechannel.FillInfo.FileLocalURL = ""

	templateRsp, err := livemanager.CreateLivechannel(livechannel)

	if err != nil {
		//t.Error("sign error ")
		t.Errorf(err.Error())
	}

	fmt.Println(templateRsp)
}

func TestEditLivechannel(t *testing.T) {

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

	livechannel := live.CreateOrUpdatechannelParam{}
	livechannel.ID = "6fc646b57e444068a2237f3f5870880f"
	rd := rand.New(rand.NewSource(time.Now().UnixNano()))
	livechannel.Name = fmt.Sprintf("修改直播频道20200107 ---  测试 %d", rd.Intn(99999))
	livechannel.SignalType = "1"
	livechannel.PullStream = ""
	livechannel.LiveType = 0
	livechannel.TypeName = "类型名称--修改--测试"
	livechannel.Type = "3daa3949c294456ab40273ab2d8afba1"
	templateList := []live.CreateTemplateList{
		{
			TemplateID: "4a8235f0ad80496b940c44e1436dff87",
			IsDefault:  "0",
		},
	}
	grouList := []live.CreateGroupList{
		{
			ProtocolType:            1,
			RtmpThirdPartyOutputURL: "",
			Rotation:                "90",
			IsRecord:                "1",
			NeedToMp4:               "1",
			IsTimeShift:             "0",
			TimeShiftDuration:       "",
			IsLiveEncrypt:           "0",
			LogoID:                  "5b99a7d13afd4e7eb4a52e15210982f6",
			OutputTemplateList:      templateList,
		},
	}
	livechannel.OutputGroupList = grouList

	templateRsp, err := livemanager.EditLivechannel(livechannel)

	if err != nil {
		//t.Error("sign error ")
		t.Errorf(err.Error())
	}

	fmt.Println(templateRsp)
}

func TestDeleteLivechannel(t *testing.T) {

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

	ids := "b2ec7fb099694915a209b5b65d016419,d5a4f6ccbb4a4d8ea1e1b2e0433d2501"
	templateRsp, err := livemanager.DeleteLivechannel(ids)

	if err != nil {
		//t.Error("sign error ")
		t.Errorf(err.Error())
	}

	fmt.Println(templateRsp)
}

func TestSetLiveCallback(t *testing.T) {

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

	value := "https://igit.xempower.cn/team-china/danghongyun-api-golang/tree/main"
	templateRsp, err := livemanager.SetLiveCallback(value)

	if err != nil {
		//t.Error("sign error ")
		t.Errorf(err.Error())
	}

	fmt.Println(templateRsp)
}

func TestStartLiveChannel(t *testing.T) {

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

	id := "39cc0761445b4de0ad9a937075f0b04e"
	templateRsp, err := livemanager.StartLiveChannel(id)

	if err != nil {
		//t.Error("sign error ")
		t.Errorf(err.Error())
	}

	fmt.Println(templateRsp)
}

func TestStopLiveChannel(t *testing.T) {

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

	id := "0f4f319ef71a4d9ea2a54a91a9d27f19"
	templateRsp, err := livemanager.StopLiveChannel(id)

	if err != nil {
		//t.Error("sign error ")
		t.Errorf(err.Error())
	}

	fmt.Println(templateRsp)
}
func TestStartPullLiveChannel(t *testing.T) {

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

	id := "e08d4c124699490c92442441b360b7bb"
	templateRsp, err := livemanager.StartPullLiveChannel(id)

	if err != nil {
		//t.Error("sign error ")
		t.Errorf(err.Error())
	}

	fmt.Println(templateRsp)
}
func TestStopPullLiveChannel(t *testing.T) {

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

	id := "0f4f319ef71a4d9ea2a54a91a9d27f19"
	templateRsp, err := livemanager.StopPullLiveChannel(id)

	if err != nil {
		//t.Error("sign error ")
		t.Errorf(err.Error())
	}

	fmt.Println(templateRsp)
}

func TestStopRecordLiveChannel(t *testing.T) {

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

	channelID := "0f4f319ef71a4d9ea2a54a91a9d27f19"
	outputGroupID := "0f4f319ef71a4d9ea2a54a91a9d27f19"
	templateRsp, err := livemanager.StopRecordLiveChannel(channelID, outputGroupID)

	if err != nil {
		//t.Error("sign error ")
		t.Errorf(err.Error())
	}

	fmt.Println(templateRsp)
}
func TestStartRecordLiveChannel(t *testing.T) {

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

	channelID := "0f4f319ef71a4d9ea2a54a91a9d27f19"
	outputGroupID := "0f4f319ef71a4d9ea2a54a91a9d27f19"
	templateRsp, err := livemanager.StartRecordLiveChannel(channelID, outputGroupID)

	if err != nil {
		//t.Error("sign error ")
		t.Errorf(err.Error())
	}

	fmt.Println(templateRsp)
}
