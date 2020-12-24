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

	template.Type = "1"

	params := m.Credentials.NewSignParams("liveTemplateCreate", Version, "")
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
	fmt.Println(queryparam)

	url := LiveRestAPIURL + "?" + queryparam
	var respTempalte TemplateListRespParam

	err := m.Client.CallWithJSONQuery(context.Background(),
		&respTempalte, "GET", url, nil, nil)

	if nil != err {
		return nil, err
	}

	return &respTempalte, nil
}
