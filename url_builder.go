package mina

import "net/url"

const (
	// BaseURL 微信请求基础URL
	baseURL   = "https://api.weixin.qq.com"
	codeAPI   = "/sns/jscode2session"
	grantType = "authorization_code"
)

func CodeToURL(appID, secret, code string) (string, error) {
	u, err := url.Parse(baseURL + codeAPI)
	if err != nil {
		return "", err
	}

	query := u.Query()

	query.Set("appid", appID)
	query.Set("secret", secret)
	query.Set("js_code", code)
	query.Set("grant_type", grantType)

	u.RawQuery = query.Encode()

	return u.String(), nil
}
