package live

import (
	"context"
	"log"

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
func (m *DHLiveManager) TemplateCreate(parma TemplateReqParam) (*TemplateRespParam, error) {

	parma.AccessKey = m.Credentials.AccessKey
	parma.Version = Version
	parma.Action = "liveTemplateCreate"
	timestamp, signature := m.Credentials.Sign2(parma.Action, Version)
	parma.Timestamp = timestamp
	parma.Signature = signature

	var respTempalte TemplateRespParam
	err := m.Client.CallWithJSON(context.Background(),
		&respTempalte, "POST", LiveRestAPIURL, nil, parma)

	log.Println(err)
	if nil != err {
		return nil, err
	}

	return &respTempalte, nil
}
