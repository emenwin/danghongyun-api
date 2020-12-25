package live

//TemplateRespParam 直播模板响应参数
type TemplateRespParam struct {
	Code    int      `json:"code"`    // 整型错误码，0成功，其它值失败
	Message string   `json:"message"` // 字符串，可能包含错误或者别的一些信息
	Success bool     `json:"success"` // true 或 false
	Result  Template `json:"result"`  // 字符串，返回的结果

}

//TemplateListRespParam 直播模板查询列表返回参数
type TemplateListRespParam struct {
	Code    int        `json:"code"`    // 整型错误码，0成功，其它值失败
	Message string     `json:"message"` // 字符串，可能包含错误或者别的一些信息
	Success bool       `json:"success"` // true 或 false
	Result  []Template `json:"result"`  // 字符串，返回的结果

}

//Template 直播模板
type Template struct {
	DisplayName          string `json:"displayName"`   //模版名称
	Type                 string `json:"type"`          //模板类型 1.用户模板 2.系统模板
	TranscodeType        int    `json:"transcodeType"` //所属直播类型。 0：普通直播、2：VR直播、3：进阶直播
	VideoWidth           int    `json:"videoWidth"`    //视频宽
	VideoHeight          int    `json:"videoHeight"`   //视频高
	VideoBitrate         int    `json:"videoBitrate"`  //视频码率
	AudioBitrate         int    `json:"audioBitrate"`  //音频码率
	FrameRate            int    `json:"frameRate"`     //视频帧率
	AdvancedArguments    string `json:"advancedArguments"`
	AdvancedArgumentsObj *TemplateAdvancedArguments

	ID         string `json:"id"`   //:唯一标识
	Name       string `json:"name"` //模版名称/码率/宽x高
	VideoCodec string `json:"videoCodec"`
	UserID     int64  `json:"userId"`     //用户id，如果是系统的模板，此项为空
	TemplateID string `json:"templateId"` //?
	CreateTime int64  `json:"createTime"` //创建时间
	UpdateTime int64  `json:"updateTime"` //更新时间
	IsDefault  bool   `json:"isDefault"`  //default ?
}

//TemplateAdvancedArguments 模板高级参数 //advancedArguments : 高级参数
type TemplateAdvancedArguments struct {
	Rc           string `json:"rc"`           //rc:码率控制
	GopType      int    `json:"gopType"`      //gopType: gop单位：0-帧（默认） 1-时间（毫秒）
	GopSize      int    `json:"gopSize"`      //gopSize: GOP大小
	Bframe       int    `json:"bframe"`       //bframe: B帧
	RefFrame     int    `json:"refFrame"`     //refFrame: 参考帧
	QualityLevel int    `json:"qualityLevel"` //qualityLevel ? null
	Interlace    int    `json:"interlace"`    //interlace": 帧场模式 -1-跟随源 0-帧 3-MBAFF
	SmartBorder  int    `json:"smartBorder"`  //smartBorder: 宽高变换模式：0-线性拉伸 1-智能黑边 2-自动裁剪

}

// //TemplateReqParam 直播模板请求参数
// type TemplateReqParam struct {
// 	AccessKey string `json:"accessKey"` //用户访问key
// 	Action    string `json:"action"`    //api名称，此接口为 liveTemplateCreate
// 	Version   string `json:"version"`   //api版本，目前必须为 2.0
// 	Timestamp string `json:"timestamp"` //api调用时间戳，1970年1月1日以来的毫秒数，如1466488681033
// 	Signature string `json:"signature"` //签名

// 	Template *Template
// }

// //MarshalJSON json序列化
// func (a TemplateReqParam) MarshalJSON() ([]byte, error) {

// 	b := struct {
// 		AccessKey string `json:"accessKey"` //用户访问key
// 		Action    string `json:"action"`    //api名称，此接口为 liveTemplateCreate
// 		Version   string `json:"version"`   //api版本，目前必须为 2.0
// 		Timestamp string `json:"timestamp"` //api调用时间戳，1970年1月1日以来的毫秒数，如1466488681033
// 		Signature string `json:"signature"` //签名

// 		DisplayName       string `json:"displayName"`
// 		Type              int    `json:"type"`
// 		TranscodeType     int    `json:"transcodeType"`
// 		VideoWidth        int    `json:"videoWidth"`
// 		VideoHeight       int    `json:"videoHeight"`
// 		VideoBitrate      int    `json:"videoBitrate"`
// 		AudioBitrate      int    `json:"audioBitrate"`
// 		FrameRate         int    `json:"frameRate"`
// 		AdvancedArguments string `json:"advancedArguments"`
// 	}{
// 		AccessKey: a.AccessKey,
// 		Action:    a.Action,
// 		Version:   a.Version,
// 		Timestamp: a.Timestamp,
// 		Signature: a.Signature,
// 	}

// 	if nil != a.Template {
// 		b.DisplayName = a.Template.DisplayName
// 		b.Type = a.Template.Type
// 		b.TranscodeType = a.Template.TranscodeType
// 		b.VideoWidth = a.Template.VideoWidth
// 		b.VideoHeight = a.Template.VideoHeight
// 		b.VideoBitrate = a.Template.VideoBitrate
// 		b.AudioBitrate = a.Template.AudioBitrate
// 		b.FrameRate = a.Template.FrameRate

// 		if nil != a.Template.AdvancedArgumentsObj {
// 			obj, err := json.Marshal(a.Template.AdvancedArgumentsObj)
// 			if nil == err {
// 				b.AdvancedArguments = string(obj)

// 			}
// 		}
// 	}

// 	return json.Marshal(b)

// }
