package live

//SignalType String，源流类型 1.推流 2.拉流
type SignalType string

const (
	//SignalTypePush 1 push
	SignalTypePush = "1"
	//SignalTypePull 2 pull
	SignalTypePull = "2"
)

//LiveType Integer，直播类型，0.普通直播 2.VR直播 3.进阶直播，不传默认跟随模版类型
type LiveType int

const (
	//LiveTypeNormal 0 普通
	LiveTypeNormal = 0
	//LiveTypeVR 2 vr
	LiveTypeVR = 2
	//LiveTypeAdvance 进阶直播，不传默认跟随模版类型
	LiveTypeAdvance = 3
)

//TemplateType 模板类型
type TemplateType int
