package ugc

//Channel 频道
type Channel struct {
	ID             int64 `json:"id"`
	Status         int64 `json:"status"`
	CurrentSession int64 `json:"current_session"` //current_session
}

//CurrentSession 当前激活的会话信息
type CurrentSession struct {
	ID        int64       `json:"id"`
	Status    int64       `json:"status"`
	Play      string      `json:"play"`
	Push      string      `json:"push"`
	URL       interface{} `json:"url"`
	HLS       string      `json:"hls"`
	Thumbnail string      `json:"thumbnail"`
}

//ChannelInfoRespParam 频道信息返回
type ChannelInfoRespParam struct {
	Code    int     `json:"code"`    // 整型错误码，0成功，其它值失败
	Message string  `json:"message"` // 字符串，可能包含错误或者别的一些信息
	Success bool    `json:"success"` // true 或 false
	Result  Channel `json:"result"`  // 字符串，返回的结果
}

//ChannelSessionRespParam 会话信息返回
type ChannelSessionRespParam struct {
	Code    int            `json:"code"`    // 整型错误码，0成功，其它值失败
	Message string         `json:"message"` // 字符串，可能包含错误或者别的一些信息
	Success bool           `json:"success"` // true 或 false
	Result  CurrentSession `json:"result"`  // 字符串，返回的结果
}
