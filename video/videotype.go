package video

//Videocategory 视频分类
type Videocategory struct {
	ID       int    `json:"id"`       //ID
	Category string `json:"category"` //分类名称
}

//VideotypeListRespParam 视频分类列表返回
type VideotypeListRespParam struct {
	Code    int             `json:"code"`    // 整型错误码，0成功，其它值失败
	Message string          `json:"message"` // 字符串，可能包含错误或者别的一些信息
	Success bool            `json:"success"` // true 或 false
	Result  []Videocategory `json:"result"`  // 字符串，返回的结果
}

//VideotypeInfoRespParam 视频分类列表返回
type VideotypeInfoRespParam struct {
	Code    int           `json:"code"`    // 整型错误码，0成功，其它值失败
	Message string        `json:"message"` // 字符串，可能包含错误或者别的一些信息
	Success bool          `json:"success"` // true 或 false
	Result  Videocategory `json:"result"`  // 字符串，返回的结果
}
