package mina

// https://developers.weixin.qq.com/miniprogram/dev/api/api-login.html#wxloginobject

// Response 请求微信返回基础数据
type CommonResponse struct {
	Errcode int    `json:"errcode,omitempty"`
	Errmsg  string `json:"errmsg,omitempty"`
}

type LoginResponse struct {
	CommonResponse
	Unionid    string `json:"unionid,omitempty"`
	Openid     string `json:"openid"`
	SessionKey string `json:"session_key"`
}

// 微信小程序加密数据结构
type UserInfo struct {
	UnionId   string     `json:"unionId,omitempty"`
	OpenId    string     `json:"openId"`
	NickName  string     `json:"nickName"`
	Gender    int        `json:"gender"` // 值为1时是男性，值为2时是女性，值为0时是未知
	AvatarUrl string     `json:"avatarUrl"`
	City      string     `json:"city"`
	Province  string     `json:"province"`
	Country   string     `json:"country"`
	Language  string     `json:"language"`
	Watermark *Watermark `json:"watermark,omitempty"` //数据水印( watermark )
}

// 用于校验数据正确性
type Watermark struct {
	Appid     string `json:"appid"`
	Timestamp int64  `json:"timestamp"`
}
