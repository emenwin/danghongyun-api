package video

//TemplateGroup 转码模组
type TemplateGroup struct {
	UserID             int    `json:"userId"`          // 用户ID
	ID                 int    `json:"id"`              //模板组ID
	Name               string `json:"name"`            //模板组名称
	Type               int    `json:"type"`            //转码组类型 0 普通 1 VR
	DefaultGroup       bool   `json:"defaultGroup"`    //是否默认
	Trailor            bool   `json:"trailor"`         //是否要转试看片段
	TrailorDuration    int    `json:"trailorDuration"` //试看片段时长 单位秒
	LogoGroupID        int    `json:"logoGroupId"`     //台标组ID
	TranscodeTemplates string `json:"transcodeTemplates"`
}

//TranscodeTemplate 模板组内的目标参数
type TranscodeTemplate struct {
	VideoWidth        int    `json:"videoWidth"`        // 视频宽度
	VideoHeight       int    `json:"videoHeight"`       //视频高度
	Name              string `json:"name"`              //模板名称
	VideoBitrate      int    `json:"videoBitrate"`      //视频比特率，1000为1000kbps
	AudioBitrate      int    `json:"audioBitrate"`      //音频比特率，64000为64kbps
	Format            int    `json:"format"`            //转码格式，0 hls， 1 mp4， 2 flv， 3 ts
	EncryptType       int    `json:"encryptType"`       //加密类型，0 不加密，6 AES加密，7 当虹加密，其中AES加密只对hls起作用
	DefaultVideo      int    `json:"defaultVideo"`      //1 默认播放模板，0 非默
	AdvancedArguments string `json:"advancedArguments"` //高级参数,
}

//TemplateGroupListRespParam 转码模组列表返回
type TemplateGroupListRespParam struct {
	Code    int             `json:"code"`    // 整型错误码，0成功，其它值失败
	Message string          `json:"message"` // 字符串，可能包含错误或者别的一些信息
	Success bool            `json:"success"` // true 或 false
	Result  []TemplateGroup `json:"result"`  // 字符串，返回的结果
}

//TemplateGroupInfoRespParam 转码模组列表返回
type TemplateGroupInfoRespParam struct {
	Code    int           `json:"code"`    // 整型错误码，0成功，其它值失败
	Message string        `json:"message"` // 字符串，可能包含错误或者别的一些信息
	Success bool          `json:"success"` // true 或 false
	Result  TemplateGroup `json:"result"`  // 字符串，返回的结果
}
