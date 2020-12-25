package live

import (
	"context"
	"fmt"
	"strconv"

	"github.com/emenwin/danghongyun-api/auth"
	"github.com/emenwin/danghongyun-api/client"
)

const (
	//LiveRestAPIURL rest url
	LiveRestAPIURL = "http://api.danghongyun.com/rest"
	//Version version
	Version = "2.0"
)

// DHLiveManager 提供了对直播进行管理的操作
type DHLiveManager struct {
	Client      *client.Client
	Credentials *auth.Credentials
}

// NewLiveManager 用来构建一个新的直播管理对象
func NewLiveManager(cred *auth.Credentials) *DHLiveManager {

	return NewLiveManagerEx(cred, nil)
}

// NewLiveManagerEx 用来构建一个新的直播管理对象
func NewLiveManagerEx(cred *auth.Credentials, clt *client.Client) *DHLiveManager {

	if clt == nil {
		clt = &client.DefaultClient
	}

	return &DHLiveManager{
		Client:      &client.DefaultClient,
		Credentials: cred,
	}
}

// TemplateCreate 用来修改文件状态, 禁用和启用文件的可访问性
// http://api.danghongyun.com/rest
func (m *DHLiveManager) TemplateCreate(template Template) (*TemplateRespParam, error) {

	params := m.Credentials.NewSignParams("liveTemplateCreate", Version, "")
	params["type"] = "1"
	params["transcodeType"] = strconv.Itoa(template.TranscodeType)
	params["videoWidth"] = strconv.Itoa(template.VideoWidth)
	params["videoHeight"] = strconv.Itoa(template.VideoHeight)
	params["videoBitrate"] = strconv.Itoa(template.VideoBitrate)
	params["audioBitrate"] = strconv.Itoa(template.AudioBitrate)
	params["displayName"] = template.DisplayName

	queryparam, _ := m.Credentials.SignExt(params)
	url := LiveRestAPIURL + "?" + queryparam

	var respTempalte TemplateRespParam
	err := m.Client.CallWithJSON(context.Background(),
		&respTempalte, "POST", url, nil, template)

	if nil != err {
		return nil, err
	}

	return &respTempalte, nil
}

// GetTemplates 查询直播模板
// name: 模板名称  可选 （”“）
// ttype : string 直播模版类型。1：用户模版，2：系统模版	可选（""）
// transcodeType	所属直播类型。 0：普通直播、2：VR直播、3：进阶直播 可选(-)
// http://api.danghongyun.com/rest
func (m *DHLiveManager) GetTemplates(name string, ttype string, transcodeType int) (*TemplateListRespParam, error) {

	params := m.Credentials.NewSignParams("liveGetTemplates", Version, "")

	if name != "" {
		params["name"] = name
	}

	if ttype != "" {
		params["type"] = ttype
	}

	if transcodeType != -1 {
		s := strconv.FormatInt(int64(transcodeType), 10)
		params["transcodeType"] = s
	}

	queryparam, _ := m.Credentials.SignExt(params)

	url := LiveRestAPIURL + "?" + queryparam
	var respTempalte TemplateListRespParam

	err := m.Client.CallWithJSONQuery(context.Background(),
		&respTempalte, "GET", url, nil, nil)

	if nil != err {
		return nil, err
	}

	return &respTempalte, nil
}

//EditTemplate 修改直播模板
func (m *DHLiveManager) EditTemplate(template Template) (*TemplateRespParam, error) {

	params := m.Credentials.NewSignParams("liveTemplateUpdate", Version, "")
	params["id"] = template.ID
	params["type"] = "1"
	params["transcodeType"] = strconv.Itoa(template.TranscodeType)
	params["videoWidth"] = strconv.Itoa(template.VideoWidth)
	params["videoHeight"] = strconv.Itoa(template.VideoHeight)
	params["videoBitrate"] = strconv.Itoa(template.VideoBitrate)
	params["audioBitrate"] = strconv.Itoa(template.AudioBitrate)
	params["displayName"] = template.DisplayName
	queryparam, _ := m.Credentials.SignExt(params)
	url := LiveRestAPIURL + "?" + queryparam
	var respTempalte TemplateRespParam
	err := m.Client.CallWithJSON(context.Background(),
		&respTempalte, "POST", url, nil, template)

	if nil != err {
		return nil, err
	}

	return &respTempalte, nil
}

//DeleteTemplates 删除 批量删除直播模板
//id: 直播模版唯一标识,多个ID用“,”分隔
func (m *DHLiveManager) DeleteTemplates(param Template) (*TemplateRespParam, error) {

	params := m.Credentials.NewSignParams("liveTemplateDelete", Version, "")
	params["id"] = param.ID
	queryparam, _ := m.Credentials.SignExt(params)
	url := LiveRestAPIURL + "?" + queryparam
	var respTempalte TemplateRespParam
	err := m.Client.CallWithJSON(context.Background(),
		&respTempalte, "POST", url, nil, param)

	if nil != err {
		return nil, err
	}

	return &respTempalte, nil
}

//GetLivetypes  查询直播类型
func (m *DHLiveManager) GetLivetypes(name string, id string) (*LivetypeListRespParam, error) {

	params := m.Credentials.NewSignParams("liveChannelTypeList", Version, "")

	if name != "" {
		params["name"] = name
	}

	if id != "" {
		params["id"] = id
	}

	queryparam, _ := m.Credentials.SignExt(params)
	fmt.Println(queryparam)

	url := LiveRestAPIURL + "?" + queryparam
	var respLivetype LivetypeListRespParam

	err := m.Client.CallWithJSONQuery(context.Background(),
		&respLivetype, "GET", url, nil, nil)

	if nil != err {
		return nil, err
	}

	return &respLivetype, nil
}

// CreateLivetype 创建直播类型
func (m *DHLiveManager) CreateLivetype(livetype Livetype) (*LivetypeRespParam, error) {

	params := m.Credentials.NewSignParams("liveChannelTypeCreate", Version, "")
	params["name"] = livetype.Name
	queryparam, _ := m.Credentials.SignExt(params)
	url := LiveRestAPIURL + "?" + queryparam

	var respLivetype LivetypeRespParam

	err := m.Client.CallWithForm(context.Background(),
		&respLivetype, "POST", url, nil, nil)

	if nil != err {
		return nil, err
	}

	return &respLivetype, nil
}

//EditLivetype 修改直播类型
func (m *DHLiveManager) EditLivetype(livetype Livetype) (*LivetypeRespParam, error) {

	params := m.Credentials.NewSignParams("liveChannelTypeUpdate", Version, "")
	params["id"] = livetype.ID
	params["name"] = livetype.Name
	queryparam, _ := m.Credentials.SignExt(params)
	url := LiveRestAPIURL + "?" + queryparam

	var respLivetype LivetypeRespParam
	err := m.Client.CallWithJSON(context.Background(),
		&respLivetype, "POST", url, nil, nil)

	if nil != err {
		return nil, err
	}

	return &respLivetype, nil
}

//DeleteLivetype 删除直播类型
func (m *DHLiveManager) DeleteLivetype(livetype Livetype) (*LivetypeRespParam, error) {

	params := m.Credentials.NewSignParams("liveChannelTypeDelete", Version, "")
	params["id"] = livetype.ID
	queryparam, _ := m.Credentials.SignExt(params)
	url := LiveRestAPIURL + "?" + queryparam

	var respLivetype LivetypeRespParam
	err := m.Client.CallWithJSON(context.Background(),
		&respLivetype, "POST", url, nil, livetype)

	if nil != err {
		return nil, err
	}

	return &respLivetype, nil
}
