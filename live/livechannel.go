package live

//Livechannel 直播频道
type Livechannel struct {
	ID                 string        `json:"id"`           //频道id
	Name               string        `json:"name"`         //频道名称
	Description        string        `json:"description"`  //频道描述
	SignalType         string        `json:"signalType"`   //源流类型 1.推流 2.拉流
	App                string        `json:"app"`          //推流端app
	Stream             string        `json:"stream"`       //流名称
	ClientID           string        `json:"clientId"`     //推流客户端id
	ThumbnailURL       string        `json:"thumbnailUrl"` //源流图片地址
	PullStream         string        `json:"pullStream"`   //拉流地址，源流类型为拉流时必填
	PullTaskID         string        `json:"pullTaskId"`
	Status             int           `json:"status"`             //频道状态 0. 初始化 1.就绪 2.直播中 3.拉流中 5启动中
	StreamStatus       int           `json:"streamStatus"`       //流状态 1-正常 2-无源
	UserID             int           `json:"userId"`             //用户id
	StreamWidth        int           `json:"streamWidth"`        // 源流宽
	StreamHeight       int           `json:"streamHeight"`       // 源流高
	StreamBitrate      int           `json:"streamBitrate"`      //源流码率
	StreamAudioBitrate int           `json:"streamAudioBitrate"` //源音频码率
	StreamFrameRate    int           `json:"streamFrameRate"`
	CreateTime         int           `json:"createTime"`       //创建时间
	UpdateTime         int           `json:"updateTime"`       //更新时间
	StartTime          int           `json:"startTime"`        //直播开始时间
	EndTime            int           `json:"endTime"`          //直播结束时间
	AutoEndTime        int           `json:"autoEndTime"`      //直播自动结束时间
	PushURL            string        `json:"pushUrl"`          //推流地址
	PublicPlayURL      string        `json:"publicPlayUrl"`    //源流rtmp播放地址
	PublicFlvPlayURL   string        `json:"publicFlvPlayUrl"` //源流flv播放地址
	LiveType           int           `json:"liveType"`         //0：普通直播、2：VR直播、3：1080P直播
	Type               string        `json:"type"`             //频道类型id
	TypeName           string        `json:"typeName"`         //频道类型名称
	DeviceID           string        `json:"deviceId"`         //设备id-与库管系统的设备进行绑定
	OutputGroupList    []OutputGroup `json:"outputGroupList"`  //输出组信息列表

}

//OutputGroup 直播输出组
type OutputGroup struct {
	ID                      string       `json:"id"`                      //直播输出组唯一标识符
	Name                    string       `json:"name"`                    //直播输出组名称
	ProtocolType            int          `json:"protocolType"`            //直播输出组协议类型 1.hls 2.rtmp
	ChannelID               string       `json:"channelId"`               //所属频道id
	TaskID                  string       `json:"taskId"`                  //任务Id
	IsLiveEncrypt           string       `json:"isLiveEncrypt"`           //是否启用加密，0.不启用 1.启用，默认0。，只有HLS协议下才可以启用加密
	CustomerID              string       `json:"customerId"`              //加密秘钥1
	ContentID               string       `json:"contentId"`               //加密秘钥2
	CreateTime              int64        `json:"createTime"`              //创建时间
	UpdateTime              int64        `json:"updateTime"`              //更新时间
	LiveLogoID              string       `json:"liveLogoId"`              //直播Logo id
	OutputTemplateList      []Template   `json:"outputTemplateList"`      //直播模板列表
	LiveOutputs             []LiveOutput `json:"liveOutputs"`             //输出信息
	LiveLogo                Livelogo     `json:"liveLogo"`                //直播logo信息
	LogoID                  string       `json:"logoId"`                  //直播Logo id
	RtmpThirdPartyOutputURL string       `json:"rtmpThirdPartyOutputUrl"` //rtmp第三方输出地址
	IsRecord                string       `json:"isRecord"`                //是否录制，0.否 1.是
	RecordTaskID            string       `json:"recordTaskId"`            //当前录制任务id
	RecordStartTime         int          `json:"recordStartTime"`         //录制开始时间
	NeedToMp4               string       `json:"needToMp4"`               //录制的文件是否转mp4，1.是 0.否
	IsTimeShift             string       `json:"isTimeShift"`             //是否时移，0.否 1.是
	TimeShiftDuration       string       `json:"timeShiftDuration"`       //时移时长，单位小时
	TimeShiftTaskID         string       `json:"timeShiftTaskId"`         //时移任务Id
	Rotation                string       `json:"rotation"`                //旋转度数
}

//LiveOutput 输出信息
type LiveOutput struct {
	ID                string      `json:"id"`
	PublicOutputURL   string      `json:"publicOutputUrl"`   //输出播放地址
	VideoWidth        int64       `json:"videoWidth"`        //分辨率：宽
	VideoHeight       int64       `json:"videoHeight"`       //分辨率：高
	VideoBitrate      int64       `json:"videoBitrate"`      //视频码率
	AudioBitrate      int64       `json:"audioBitrate"`      //音频码率
	ProtocolType      int64       `json:"protocolType"`      //直播分组协议类型 1.hls 2.rtmp
	UserID            int64       `json:"userId"`            //用户id
	ChannelID         string      `json:"channelId"`         //所属直播频道id
	LiveOutputGroupID string      `json:"liveOutputGroupId"` //所属直播输出组id
	LiveTemplateID    string      `json:"liveTemplateId"`    //直播模板Id
	CreateTime        int64       `json:"createTime"`        //创建时间
	UpdateTime        interface{} `json:"updateTime"`        //更新时间
	IsDefault         string      `json:"isDefault"`         //1. 默认 0.不默认
}

//LivechannelListRespParam 直播频道列表返回
type LivechannelListRespParam struct {
	Code    int           `json:"code"`    // 整型错误码，0成功，其它值失败
	Message string        `json:"message"` // 字符串，可能包含错误或者别的一些信息
	Success bool          `json:"success"` // true 或 false
	Result  []Livechannel `json:"result"`  // 字符串，返回的结果
}

//LivechannelRespParam 直播频道响应参数
type LivechannelRespParam struct {
	Code    int         `json:"code"`    // 整型错误码，0成功，其它值失败
	Message string      `json:"message"` // 字符串，可能包含错误或者别的一些信息
	Success bool        `json:"success"` // true 或 false
	Result  Livechannel `json:"result"`  // 字符串，返回的结果

}

//SearchLivechannelParam 频道列表查询参数
type SearchLivechannelParam struct {
	Start      int    `json:"start"`      // 当前页码，不传时默认为1
	Number     int    `json:"number"`     //每页展示数量，不传默认为15
	Sort       string `json:"sort"`       //结果按该字段排序，只能为name, type, status, streamStatus, createTime
	Order      string `json:"order"`      //asc 顺序，desc 倒序
	SearchType string `json:"searchType"` //搜索类型，只能为name，type，all，分别为按name搜索，按type搜索，按name和type一起搜索，搜索时必须同时要输入参数keyword
	Keyword    string `json:"keyword"`    //关键字
}

// LivechannelURL 频道播放地址
type LivechannelURL struct {
	ChannelID        string    `json:"channelId"`        //直播频道Id
	ChannelGroupID   string    `json:"channelGroupId"`   //输出组Id
	ChannelGroupType int64     `json:"channelGroupType"` //输出组协议类型 1.hls 2.rtmp
	PlayURL          string    `json:"playUrl"`          //父级播放地址
	PlayURLList      []PlayURL `json:"playUrlList"`      //子级播放信息列表
}

// PlayURL 频道播放地址
type PlayURL struct {
	URL          string `json:"url"`          //子级播放地址
	Width        int64  `json:"width"`        //视频宽
	Height       int64  `json:"height"`       //视频宽
	VideoBitrate int64  `json:"videoBitrate"` //视频码率
	AudioBitrate int64  `json:"audioBitrate"` //音频码率
	DefaultVideo string `json:"defaultVideo"` //是否默认 0.否 1.是
}

//CreateOrUpdatechannelParam 创建或更新频道
type CreateOrUpdatechannelParam struct {
	ID              string            `json:"id"` //频道id
	Name            string            `json:"name"`
	SignalType      string            `json:"signalType"`
	LiveType        int               `json:"liveType"`
	PullStream      string            `json:"pullStream"` //拉流地址，源流类型为拉流时必填
	Type            string            `json:"type"`       //频道类型id
	TypeName        string            `json:"typeName"`   //类型名称
	OutputGroupList []CreateGroupList `json:"outputGroupList"`
}

//CreateGroupList 创建频道组
type CreateGroupList struct {
	ID                      string               `json:"id"`                      //输出组id
	RtmpThirdPartyOutputURL string               `json:"rtmpThirdPartyOutputUrl"` //rtmp第三方输出地址
	IsRecord                string               `json:"isRecord"`                //是否录制，0.否 1.是
	ProtocolType            int                  `json:"protocolType"`
	IsTimeShift             string               `json:"isTimeShift,omitempty"`
	TimeShiftDuration       string               `json:"timeShiftDuration,omitempty"`
	IsLiveEncrypt           string               `json:"isLiveEncrypt,omitempty"`
	LogoID                  string               `json:"logoId,omitempty"`
	Rotation                string               `json:"rotation"`  //旋转度数
	NeedToMp4               string               `json:"needToMp4"` //录制的文件是否转mp4，1.是 0.否
	OutputTemplateList      []CreateTemplateList `json:"outputTemplateList"`
}

//CreateTemplateList 创建频道模板
type CreateTemplateList struct {
	TemplateID string `json:"templateId"`
	IsDefault  string `json:"isDefault"`
}
