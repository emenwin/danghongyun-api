package ugc

import (
	"context"

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

//GetChannelInfo 频道信息
func (ctl *DHLiveManager) GetChannelInfo(ID string) (*ChannelInfoRespParam, error) {
	params := ctl.Credentials.NewSignParams("ugcGetChannel", Version, "")
	params["id"] = ID

	queryparam, _ := ctl.Credentials.SignExt(params)
	url := LiveRestAPIURL + "?" + queryparam

	var respTempalte ChannelInfoRespParam
	err := ctl.Client.CallWithJSONQuery(context.Background(),
		&respTempalte, "GET", url, nil, nil)

	if nil != err {
		return nil, err
	}

	return &respTempalte, nil
}

//CreateChannel 新增频道
func (ctl *DHLiveManager) CreateChannel(name string) (*ChannelInfoRespParam, error) {

	params := ctl.Credentials.NewSignParams("ugcCreateChannel", Version, "")
	params["name"] = name

	queryparam, _ := ctl.Credentials.SignExt(params)
	url := LiveRestAPIURL + "?" + queryparam
	var respTempalte ChannelInfoRespParam
	err := ctl.Client.CallWithJSON(context.Background(),
		&respTempalte, "POST", url, nil, nil)

	if nil != err {
		return nil, err
	}

	return &respTempalte, nil
}

//EditChannel 修改频道
func (ctl *DHLiveManager) EditChannel(ID, name string) (*ChannelInfoRespParam, error) {

	params := ctl.Credentials.NewSignParams("ugcUpdateChannel", Version, "")
	params["id"] = ID
	params["name"] = name

	queryparam, _ := ctl.Credentials.SignExt(params)
	url := LiveRestAPIURL + "?" + queryparam
	var respTempalte ChannelInfoRespParam
	err := ctl.Client.CallWithJSON(context.Background(),
		&respTempalte, "POST", url, nil, nil)

	if nil != err {
		return nil, err
	}

	return &respTempalte, nil
}

//DeleteChannel 删除频道
func (ctl *DHLiveManager) DeleteChannel(ID string) (*ChannelInfoRespParam, error) {

	params := ctl.Credentials.NewSignParams("ugcDeleteChannel", Version, "")
	params["id"] = ID

	queryparam, _ := ctl.Credentials.SignExt(params)
	url := LiveRestAPIURL + "?" + queryparam
	var respTempalte ChannelInfoRespParam
	err := ctl.Client.CallWithJSON(context.Background(),
		&respTempalte, "POST", url, nil, nil)

	if nil != err {
		return nil, err
	}

	return &respTempalte, nil
}

//BlockChannel 屏蔽频道
func (ctl *DHLiveManager) BlockChannel(ID string) (*ChannelInfoRespParam, error) {

	params := ctl.Credentials.NewSignParams("ugcBlockChannel", Version, "")
	params["id"] = ID

	queryparam, _ := ctl.Credentials.SignExt(params)
	url := LiveRestAPIURL + "?" + queryparam
	var respTempalte ChannelInfoRespParam
	err := ctl.Client.CallWithJSON(context.Background(),
		&respTempalte, "POST", url, nil, nil)

	if nil != err {
		return nil, err
	}

	return &respTempalte, nil
}

//RestoreChannel 恢复频道
func (ctl *DHLiveManager) RestoreChannel(ID string) (*ChannelInfoRespParam, error) {

	params := ctl.Credentials.NewSignParams("ugcRestoreChannel", Version, "")
	params["id"] = ID

	queryparam, _ := ctl.Credentials.SignExt(params)
	url := LiveRestAPIURL + "?" + queryparam
	var respTempalte ChannelInfoRespParam
	err := ctl.Client.CallWithJSON(context.Background(),
		&respTempalte, "POST", url, nil, nil)

	if nil != err {
		return nil, err
	}

	return &respTempalte, nil
}

//GetChannelSessionInfo 获取会话信息
func (ctl *DHLiveManager) GetChannelSessionInfo(ID string) (*ChannelSessionRespParam, error) {
	params := ctl.Credentials.NewSignParams("ugcGetChannelSession", Version, "")
	params["id"] = ID

	queryparam, _ := ctl.Credentials.SignExt(params)
	url := LiveRestAPIURL + "?" + queryparam

	var respTempalte ChannelSessionRespParam
	err := ctl.Client.CallWithJSONQuery(context.Background(),
		&respTempalte, "GET", url, nil, nil)

	if nil != err {
		return nil, err
	}

	return &respTempalte, nil
}

//CreateChannelSession 创建会话信息
func (ctl *DHLiveManager) CreateChannelSession(ID string) (*ChannelSessionRespParam, error) {

	params := ctl.Credentials.NewSignParams("ugcCreateChannelSession", Version, "")
	params["id"] = ID // 频道id

	queryparam, _ := ctl.Credentials.SignExt(params)
	url := LiveRestAPIURL + "?" + queryparam
	var respTempalte ChannelSessionRespParam
	err := ctl.Client.CallWithJSON(context.Background(),
		&respTempalte, "POST", url, nil, nil)

	if nil != err {
		return nil, err
	}

	return &respTempalte, nil
}

//StopChannelSession 停止会话信息
func (ctl *DHLiveManager) StopChannelSession(ID string) (*ChannelSessionRespParam, error) {

	params := ctl.Credentials.NewSignParams("ugcStopChannelSession", Version, "")
	params["id"] = ID

	queryparam, _ := ctl.Credentials.SignExt(params)
	url := LiveRestAPIURL + "?" + queryparam
	var respTempalte ChannelSessionRespParam
	err := ctl.Client.CallWithJSON(context.Background(),
		&respTempalte, "POST", url, nil, nil)

	if nil != err {
		return nil, err
	}

	return &respTempalte, nil
}

//DeleteChannelSession 删除会话信息
func (ctl *DHLiveManager) DeleteChannelSession(ID string) (*ChannelSessionRespParam, error) {

	params := ctl.Credentials.NewSignParams("ugcDeleteChannelSession", Version, "")
	params["id"] = ID

	queryparam, _ := ctl.Credentials.SignExt(params)
	url := LiveRestAPIURL + "?" + queryparam
	var respTempalte ChannelSessionRespParam
	err := ctl.Client.CallWithJSON(context.Background(),
		&respTempalte, "POST", url, nil, nil)

	if nil != err {
		return nil, err
	}

	return &respTempalte, nil
}
