package video

//User 用户信息
type User struct {
	RandomWordTime         int64       `json:"randomWordTime"`         //随机码生成时间
	TrialDays              int64       `json:"trialDays"`              //试用天数
	TokenTime              interface{} `json:"tokenTime"`              //本接口所返回该字段目前无效
	RandomWord             string      `json:"randomWord"`             //随机码
	Source                 int64       `json:"source"`                 //用户来源，内部使用
	GenerateTime           int64       `json:"generateTime"`           //注册时间
	Type                   int64       `json:"type"`                   //用户类别，0 集团大用户，1 普通用户
	Trial                  int64       `json:"trial"`                  //用户用状态， 0 正式用户试，1 试用
	LastLoginIP            string      `json:"lastLoginIp"`            //上一次登录的ip
	ExpiresIn              interface{} `json:"expiresIn"`              //本接口所返回该字段目前无效
	TrialDeleteExpiredDate int64       `json:"trialDeleteExpiredDate"` //试用用户过期后是否删除数据，0 否， 1 是
	Contact                string      `json:"contact"`                //联系人名称
	Cellphone              string      `json:"cellphone"`              //手机号
	Company                string      `json:"company"`                //公司
	ID                     int64       `json:"id"`                     //用户ID
	ActivateStatus         int64       `json:"activateStatus"`         //激活状态，0 未激活，1 激活
	Qq                     string      `json:"qq"`                     //QQ号码
	RandomWordUsedCount    int64       `json:"randomWordUsedCount"`    //随机码使用次数
	TrialTimes             int64       `json:"trialTimes"`             //试用延长次数
	Abbreviation           string      `json:"abbreviation"`           // 用户简称
	ParentID               interface{} `json:"parentId"`               //父账号id
	AccessSecret           string      `json:"accessSecret"`           //访问密钥，该字段只为兼容老版本使用，请勿以此接口返回值为依据，后续可能移除
	Token                  string      `json:"token"`                  //本接口所返回该字段目前无效
	LastLoginTime          int64       `json:"lastLoginTime"`          //上一次登录的时间戳
	AccessKey              string      `json:"accessKey"`              //访问key，该字段只为兼容老版本使用，请勿以此接口返回值为依据，后续可能移除
	DomainPrefix           string      `json:"domainPrefix"`           //用户域名前缀
	Account                string      `json:"account"`                //账号
	RefreshToken           string      `json:"refreshToken"`           //本接口所返回该字段目前无效
}

//UserInfoRespParam 用户信息返回
type UserInfoRespParam struct {
	Code    int    `json:"code"`    // 整型错误码，0成功，其它值失败
	Message string `json:"message"` // 字符串，可能包含错误或者别的一些信息
	Success bool   `json:"success"` // true 或 false
	Result  *User  `json:"result"`  // 字符串，返回的结果

}
