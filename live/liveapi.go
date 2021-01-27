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
	params["frameRate"] = strconv.Itoa(template.FrameRate)
	params["advancedArguments"] = template.AdvancedArguments

	queryparam, _ := m.Credentials.SignExt(params)
	url := LiveRestAPIURL + "?" + queryparam

	var respTempalte TemplateRespParam
	err := m.Client.CallWithJSON(context.Background(),
		&respTempalte, "POST", url, nil, nil)

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
	params["frameRate"] = strconv.Itoa(template.FrameRate)
	params["advancedArguments"] = template.AdvancedArguments
	queryparam, _ := m.Credentials.SignExt(params)
	url := LiveRestAPIURL + "?" + queryparam
	var respTempalte TemplateRespParam
	err := m.Client.CallWithJSON(context.Background(),
		&respTempalte, "POST", url, nil, nil)

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
		&respTempalte, "POST", url, nil, nil)

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
		&respLivetype, "POST", url, nil, nil)

	if nil != err {
		return nil, err
	}

	return &respLivetype, nil
}

//GetLivelogos 查询直播logo
func (m *DHLiveManager) GetLivelogos(name string, id string) (*LivelogoListRespParam, error) {

	params := m.Credentials.NewSignParams("liveLogoList", Version, "")

	if name != "" {
		params["name"] = name
	}

	if id != "" {
		params["id"] = id
	}

	queryparam, _ := m.Credentials.SignExt(params)
	fmt.Println(queryparam)

	url := LiveRestAPIURL + "?" + queryparam
	var respLivelogo LivelogoListRespParam

	err := m.Client.CallWithJSONQuery(context.Background(),
		&respLivelogo, "GET", url, nil, nil)

	if nil != err {
		return nil, err
	}

	return &respLivelogo, nil
}

// CreateLivelogo 创建直播logo
func (m *DHLiveManager) CreateLivelogo(livelogo Livelogo) (*LivelogoRespParam, error) {

	params := m.Credentials.NewSignParams("liveLogoCreate", Version, "")
	params["name"] = livelogo.Name
	params["uri"] = livelogo.Uri
	params["width"] = strconv.Itoa(livelogo.Width)
	params["height"] = strconv.Itoa(livelogo.Height)
	params["offsetHeight"] = strconv.Itoa(livelogo.OffsetHeight)
	params["offsetWidth"] = strconv.Itoa(livelogo.OffsetWidth)
	params["opacity"] = strconv.Itoa(livelogo.Opacity)
	params["resize"] = strconv.Itoa(livelogo.Resize)
	queryparam, _ := m.Credentials.SignExt(params)
	url := LiveRestAPIURL + "?" + queryparam

	var respLivelogo LivelogoRespParam

	err := m.Client.CallWithForm(context.Background(),
		&respLivelogo, "POST", url, nil, nil)

	if nil != err {
		return nil, err
	}

	return &respLivelogo, nil
}

// EditLivelogo 修改直播logo
func (m *DHLiveManager) EditLivelogo(livelogo Livelogo) (*LivelogoRespParam, error) {

	params := m.Credentials.NewSignParams("liveLogoUpdate", Version, "")
	params["id"] = livelogo.ID
	params["name"] = livelogo.Name
	params["uri"] = livelogo.Uri
	params["width"] = strconv.Itoa(livelogo.Width)
	params["height"] = strconv.Itoa(livelogo.Height)
	params["offsetHeight"] = strconv.Itoa(livelogo.OffsetHeight)
	params["offsetWidth"] = strconv.Itoa(livelogo.OffsetWidth)
	params["opacity"] = strconv.Itoa(livelogo.Opacity)
	params["resize"] = strconv.Itoa(livelogo.Resize)
	queryparam, _ := m.Credentials.SignExt(params)
	url := LiveRestAPIURL + "?" + queryparam

	var respLivelogo LivelogoRespParam

	err := m.Client.CallWithForm(context.Background(),
		&respLivelogo, "POST", url, nil, nil)

	if nil != err {
		return nil, err
	}

	return &respLivelogo, nil
}

// DeleteLivelogo 删除 批量删除直播logo
//logo唯一标识，多个id用“,”分隔
func (m *DHLiveManager) DeleteLivelogo(livelogo Livelogo) (*LivelogoRespParam, error) {

	params := m.Credentials.NewSignParams("liveLogoDelete", Version, "")
	params["id"] = livelogo.ID
	queryparam, _ := m.Credentials.SignExt(params)
	url := LiveRestAPIURL + "?" + queryparam

	var respLivelogo LivelogoRespParam

	err := m.Client.CallWithForm(context.Background(),
		&respLivelogo, "POST", url, nil, nil)

	if nil != err {
		return nil, err
	}

	return &respLivelogo, nil
}

//GetPageLivechannels 分页查询直播频道
func (m *DHLiveManager) GetPageLivechannels(searchLivechannel SearchLivechannelParam) (*LivechannelListRespParam, error) {

	params := m.Credentials.NewSignParams("liveGetChannelsByPage", Version, "")

	if searchLivechannel.Start != 0 {
		params["start"] = strconv.Itoa(searchLivechannel.Start)
	}
	if searchLivechannel.Number != 0 {
		params["number"] = strconv.Itoa(searchLivechannel.Number)
	}

	if searchLivechannel.Sort != "" {
		params["sort"] = searchLivechannel.Sort
	}

	if searchLivechannel.Order != "" {
		params["order"] = searchLivechannel.Order
	}

	if searchLivechannel.SearchType != "" {
		params["searchType"] = searchLivechannel.SearchType
	}

	if searchLivechannel.Keyword != "" {
		params["keyword"] = searchLivechannel.Keyword
	}

	queryparam, _ := m.Credentials.SignExt(params)

	url := LiveRestAPIURL + "?" + queryparam
	var respLivechannel LivechannelListRespParam

	err := m.Client.CallWithJSONQuery(context.Background(),
		&respLivechannel, "GET", url, nil, nil)

	if nil != err {
		return nil, err
	}

	return &respLivechannel, nil
}

//GetLivechannelByID 根据ID查询数据
func (m *DHLiveManager) GetLivechannelByID(id string) (*LivechannelRespParam, error) {

	params := m.Credentials.NewSignParams("liveGetChannelById", Version, "")

	params["id"] = id

	queryparam, _ := m.Credentials.SignExt(params)

	url := LiveRestAPIURL + "?" + queryparam
	var respLivechannel LivechannelRespParam

	err := m.Client.CallWithJSONQuery(context.Background(),
		&respLivechannel, "GET", url, nil, nil)

	if nil != err {
		return nil, err
	}

	return &respLivechannel, nil
}

//GetLivechannelsPlayURL 获取直播频道地址
func (m *DHLiveManager) GetLivechannelsPlayURL(channelID string) (*LivechannelRespParam, error) {

	params := m.Credentials.NewSignParams("liveGetChannelsPlayUrl", Version, "")

	params["channelId"] = channelID

	queryparam, _ := m.Credentials.SignExt(params)

	url := LiveRestAPIURL + "?" + queryparam
	var respLivechannel LivechannelRespParam

	err := m.Client.CallWithJSONQuery(context.Background(),
		&respLivechannel, "GET", url, nil, nil)

	if nil != err {
		return nil, err
	}

	return &respLivechannel, nil
}

// CreateLivechannel 创建直播频道
func (m *DHLiveManager) CreateLivechannel(param CreateOrUpdatechannelParam) (*LivechannelRespParam, error) {

	params := m.Credentials.NewSignParams("liveCreateChannel", Version, "")

	queryparam, _ := m.Credentials.SignExt(params)
	url := LiveRestAPIURL + "?" + queryparam

	var respLivechannel LivechannelRespParam
	liveChannel := CreateOrUpdatechannelParam{}
	liveChannel.Name = param.Name
	liveChannel.SignalType = param.SignalType
	liveChannel.PullStream = param.PullStream
	liveChannel.Type = param.Type
	liveChannel.TypeName = param.TypeName
	liveChannel.LiveType = param.LiveType
	// if liveChannel.FillInfo.FileType != -1 {
	// 	liveChannel.FillInfo = param.FillInfo
	// }

	liveChannel.OutputGroupList = param.OutputGroupList

	err := m.Client.CallWithJSON(context.Background(),
		&respLivechannel, "POST", url, nil, liveChannel)

	if nil != err {
		return nil, err
	}

	return &respLivechannel, nil
}

// EditLivechannel 修改直播频道
func (m *DHLiveManager) EditLivechannel(param CreateOrUpdatechannelParam) (*LivechannelRespParam, error) {

	params := m.Credentials.NewSignParams("liveUpdate", Version, "")

	queryparam, _ := m.Credentials.SignExt(params)
	url := LiveRestAPIURL + "?" + queryparam

	var respLivechannel LivechannelRespParam
	liveChannel := CreateOrUpdatechannelParam{}
	liveChannel.ID = param.ID
	liveChannel.Name = param.Name
	liveChannel.SignalType = param.SignalType
	liveChannel.PullStream = param.PullStream
	liveChannel.Type = param.Type
	liveChannel.TypeName = param.TypeName
	liveChannel.LiveType = param.LiveType
	// liveChannel.FillInfo = param.FillInfo
	liveChannel.OutputGroupList = param.OutputGroupList

	err := m.Client.CallWithJSON(context.Background(),
		&respLivechannel, "POST", url, nil, liveChannel)

	if nil != err {
		return nil, err
	}

	return &respLivechannel, nil
}

// DeleteLivechannel 删除 批量删除直播频道
//唯一标识，多个id用“,”分隔
func (m *DHLiveManager) DeleteLivechannel(ids string) (*LivechannelRespParam, error) {

	params := m.Credentials.NewSignParams("liveDelete", Version, "")
	params["ids"] = ids
	queryparam, _ := m.Credentials.SignExt(params)
	url := LiveRestAPIURL + "?" + queryparam

	var respLivelogo LivechannelRespParam

	err := m.Client.CallWithForm(context.Background(),
		&respLivelogo, "POST", url, nil, nil)

	if nil != err {
		return nil, err
	}

	return &respLivelogo, nil
}

// SetLiveCallback 设置的直播频道状态变化回调地址
func (m *DHLiveManager) SetLiveCallback(value string) (*LivechannelRespParam, error) {
	params := m.Credentials.NewSignParams("setLiveCallback", Version, "")
	params["value"] = value
	queryparam, _ := m.Credentials.SignExt(params)
	url := LiveRestAPIURL + "?" + queryparam

	var respLivelogo LivechannelRespParam

	err := m.Client.CallWithForm(context.Background(),
		&respLivelogo, "POST", url, nil, nil)

	if nil != err {
		return nil, err
	}

	return &respLivelogo, nil
}

//StartLiveChannel 开始直播
func (m *DHLiveManager) StartLiveChannel(id string) (*LivechannelRespParam, error) {
	params := m.Credentials.NewSignParams("liveStartChannel", Version, "")
	params["id"] = id
	queryparam, _ := m.Credentials.SignExt(params)
	url := LiveRestAPIURL + "?" + queryparam

	var respLivelogo LivechannelRespParam

	err := m.Client.CallWithForm(context.Background(),
		&respLivelogo, "POST", url, nil, nil)

	if nil != err {
		return nil, err
	}

	return &respLivelogo, nil
}

//StopLiveChannel 结束直播
func (m *DHLiveManager) StopLiveChannel(id string) (*LivechannelRespParam, error) {
	params := m.Credentials.NewSignParams("liveStopChannel", Version, "")
	params["id"] = id
	queryparam, _ := m.Credentials.SignExt(params)
	url := LiveRestAPIURL + "?" + queryparam

	var respLivelogo LivechannelRespParam

	err := m.Client.CallWithForm(context.Background(),
		&respLivelogo, "POST", url, nil, nil)

	if nil != err {
		return nil, err
	}

	return &respLivelogo, nil
}

//StartPullLiveChannel 开始直播拉流
func (m *DHLiveManager) StartPullLiveChannel(id string) (*LivechannelRespParam, error) {
	params := m.Credentials.NewSignParams("liveStartChannelPull", Version, "")
	params["id"] = id
	queryparam, _ := m.Credentials.SignExt(params)
	url := LiveRestAPIURL + "?" + queryparam

	var respLivelogo LivechannelRespParam

	err := m.Client.CallWithForm(context.Background(),
		&respLivelogo, "POST", url, nil, nil)

	if nil != err {
		return nil, err
	}

	return &respLivelogo, nil
}

//StopPullLiveChannel 结束直播拉流
func (m *DHLiveManager) StopPullLiveChannel(id string) (*LivechannelRespParam, error) {
	params := m.Credentials.NewSignParams("liveStopChannelPull", Version, "")
	params["id"] = id
	queryparam, _ := m.Credentials.SignExt(params)
	url := LiveRestAPIURL + "?" + queryparam

	var respLivelogo LivechannelRespParam

	err := m.Client.CallWithForm(context.Background(),
		&respLivelogo, "POST", url, nil, nil)

	if nil != err {
		return nil, err
	}

	return &respLivelogo, nil
}

//StartRecordLiveChannel 开始录制
///channelId	直播频道id	String	必选
//outputGroupId	直播频道输出组id 必选
func (m *DHLiveManager) StartRecordLiveChannel(channelID, outputGroupID string) (*LivechannelRespParam, error) {
	params := m.Credentials.NewSignParams("liveStartChannelRecord", Version, "")
	params["channelId"] = channelID
	params["outputGroupId"] = outputGroupID
	queryparam, _ := m.Credentials.SignExt(params)
	url := LiveRestAPIURL + "?" + queryparam

	var respLivelogo LivechannelRespParam

	err := m.Client.CallWithForm(context.Background(),
		&respLivelogo, "POST", url, nil, nil)

	if nil != err {
		return nil, err
	}

	return &respLivelogo, nil
}

//StopRecordLiveChannel 结束录制
func (m *DHLiveManager) StopRecordLiveChannel(channelID, outputGroupID string) (*LivechannelRespParam, error) {
	params := m.Credentials.NewSignParams("liveStopChannelRecord", Version, "")
	params["channelId"] = channelID
	params["outputGroupId"] = outputGroupID
	queryparam, _ := m.Credentials.SignExt(params)
	url := LiveRestAPIURL + "?" + queryparam

	var respLivelogo LivechannelRespParam

	err := m.Client.CallWithForm(context.Background(),
		&respLivelogo, "POST", url, nil, nil)

	if nil != err {
		return nil, err
	}

	return &respLivelogo, nil
}

//EditLiveTimeshift 直播时移
//要时移的时间点与当前时间的差值，负整数，单位秒
func (m *DHLiveManager) EditLiveTimeshift(timeshift int, playURL string) (*LivechannelRespParam, error) {
	url := playURL + "?timeshift=" + strconv.Itoa(timeshift)

	var respLivelogo LivechannelRespParam

	err := m.Client.CallWithForm(context.Background(),
		&respLivelogo, "POST", url, nil, nil)

	if nil != err {
		return nil, err
	}

	return &respLivelogo, nil
}

//BackLiveTime 回看
//回看2018-01-01 12:00:00 ~ 2018-01-01 12:30:00 半小时内的视频内容
//beginTime=20180101040000&endTime=20180101043000
func (m *DHLiveManager) BackLiveTime(beginTime, endTime, playURL string) (*LivechannelRespParam, error) {
	url := playURL + "?beginTime=" + beginTime + "&endTime=" + endTime

	var respLivelogo LivechannelRespParam

	err := m.Client.CallWithForm(context.Background(),
		&respLivelogo, "POST", url, nil, nil)

	if nil != err {
		return nil, err
	}

	return &respLivelogo, nil
}
