package live

//LivetypeRespParam 直播类型响应参数
type LivetypeRespParam struct {
	Code    int      `json:"code"`    // 整型错误码，0成功，其它值失败
	Message string   `json:"message"` // 字符串，可能包含错误或者别的一些信息
	Success bool     `json:"success"` // true 或 false
	Result  Livetype `json:"result"`  // 字符串，返回的结果

}

//LivetypeListRespParam 直播类型查询列表返回参数
type LivetypeListRespParam struct {
	Code    int        `json:"code"`    // 整型错误码，0成功，其它值失败
	Message string     `json:"message"` // 字符串，可能包含错误或者别的一些信息
	Success bool       `json:"success"` // true 或 false
	Result  []Livetype `json:"result"`  // 字符串，返回的结果
}

//Livetype 直播类型
type Livetype struct {
	ID         string `json:"id,omitempty" `        //:唯一标识
	Name       string `json:"name"`                 //直播类型名称
	UserID     int64  `json:"userId,omitempty"`     //用户id
	CreateTime int64  `json:"createTime,omitempty"` //创建时间
	UpdateTime int64  `json:"updateTime,omitempty"` //更新时间
}
