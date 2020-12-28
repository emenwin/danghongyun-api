package video

//LogoGroup 台标组
type LogoGroup struct {
	UserID string `json:"userId"` // 用户ID
	ID     int    `json:"id"`     //台标组ID
	Name   string `json:"name"`   //台标组名称
	Logos  string `json:"logos"`  //台标组内的目标参数，
}

//LogoGroupListRespParam 转码模组列表返回
type LogoGroupListRespParam struct {
	Code    int         `json:"code"`    // 整型错误码，0成功，其它值失败
	Message string      `json:"message"` // 字符串，可能包含错误或者别的一些信息
	Success bool        `json:"success"` // true 或 false
	Result  []LogoGroup `json:"result"`  // 字符串，返回的结果
}

//LogoGroupInfoRespParam 转码模组列表返回
type LogoGroupInfoRespParam struct {
	Code    int       `json:"code"`    // 整型错误码，0成功，其它值失败
	Message string    `json:"message"` // 字符串，可能包含错误或者别的一些信息
	Success bool      `json:"success"` // true 或 false
	Result  LogoGroup `json:"result"`  // 字符串，返回的结果
}

//Logos 台标组内的目标参数，
type Logos struct {
	URI     string `json:"uri"`     //	uri	台标相对路径	String	必须
	Width   int    `json:"width"`   //	width	台标宽度	Integer	必须
	Height  int    `json:"height"`  //height	台标高度	Integer	必须
	LeftPos int    `json:"leftPos"` //leftPos	台标左上角位置宽度百分比，如90%，则leftPos=90；以下参数百分比类似	Integer	必须
	TopPos  int    `json:"topPos"`  //topPos	台标左上角位置高度百分比	Integer	必须
	Opacity int    `json:"opacity"` //opacity	台标透明度百分比	Integer	必须
	Resize  int    `json:"resize"`  //resize	台标缩放值百分比	Integer	必须
	Start   int    `json:"start"`   //	start	台标开始时间点，百分比	Integer	必须
	End     int    `json:"end"`     //end	台标结束时间点，百分比	Integer	必
}
