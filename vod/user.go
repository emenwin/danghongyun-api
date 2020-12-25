package vod

import "encoding/json"

//User 用户信息
type User struct {
	RandomWordTime         int64       `json:"randomWordTime"`
	TrialDays              int64       `json:"trialDays"`
	TokenTime              interface{} `json:"tokenTime"`
	RandomWord             string      `json:"randomWord"`
	Source                 int64       `json:"source"`
	GenerateTime           int64       `json:"generateTime"`
	Type                   int64       `json:"type"`
	Trial                  int64       `json:"trial"`
	LastLoginIP            string      `json:"lastLoginIp"`
	ExpiresIn              interface{} `json:"expiresIn"`
	TrialDeleteExpiredDate int64       `json:"trialDeleteExpiredDate"`
	Contact                string      `json:"contact"`
	Cellphone              string      `json:"cellphone"`
	Company                string      `json:"company"`
	ID                     int64       `json:"id"`
	ActivateStatus         int64       `json:"activateStatus"`
	Qq                     string      `json:"qq"`
	RandomWordUsedCount    int64       `json:"randomWordUsedCount"`
	TrialTimes             int64       `json:"trialTimes"`
	Abbreviation           string      `json:"abbreviation"`
	ParentID               interface{} `json:"parentId"`
	AccessSecret           string      `json:"accessSecret"`
	Token                  string      `json:"token"`
	LastLoginTime          int64       `json:"lastLoginTime"`
	AccessKey              string      `json:"accessKey"`
	DomainPrefix           string      `json:"domainPrefix"`
	Account                string      `json:"account"`
	RefreshToken           string      `json:"refreshToken"`
}

//UnmarshalUser data to user
func UnmarshalUser(data []byte) (User, error) {
	var r User
	err := json.Unmarshal(data, &r)
	return r, err
}

//Marshal marshal
func (r *User) Marshal() ([]byte, error) {
	return json.Marshal(r)
}
