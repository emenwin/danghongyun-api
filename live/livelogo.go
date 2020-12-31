package live

//Livelogo 直播logo
type Livelogo struct {
	ID           string `json:"id"`           //:唯一标识
	Name         string `json:"name"`         //logo名称
	UserID       int64  `json:"userId"`       //用户id
	Width        int    `json:"width"`        //台标宽度
	Height       int    `json:"height"`       //台标宽度
	OffsetWidth  int    `json:"offsetWidth"`  //台标左上角距离视频左边的距离(宽度以640为基准，实际输出会按照倍率计算)
	OffsetHeight int    `json:"offsetHeight"` //台标左上角距离视频顶部的距离(高度以360为基准，实际输出会按照倍率计算)
	Opacity      int    `json:"opacity"`      //透明度
	Resize       int    `json:"resize"`       //基于原图的放大倍数，例：100为保持原图大小 50为缩小一倍
	HttpUrl      string `json:"httpUrl"`      //台标访问地址
	Uri          string `json:"uri"`          //台标路径
}

//LivelogoListRespParam 直播logo查询列表返回参数
type LivelogoListRespParam struct {
	Code    int        `json:"code"`    // 整型错误码，0成功，其它值失败
	Message string     `json:"message"` // 字符串，可能包含错误或者别的一些信息
	Success bool       `json:"success"` // true 或 false
	Result  []Livelogo `json:"result"`  // 字符串，返回的结果
}

//LivelogoRespParam 直播logo响应参数
type LivelogoRespParam struct {
	Code    int      `json:"code"`    // 整型错误码，0成功，其它值失败
	Message string   `json:"message"` // 字符串，可能包含错误或者别的一些信息
	Success bool     `json:"success"` // true 或 false
	Result  Livelogo `json:"result"`  // 字符串，返回的结果

}
