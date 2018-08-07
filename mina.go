package mina

import (
	"encoding/json"
	"errors"
)

func (c *Client) Login(appID, secret, code string) (lr LoginResponse, err error) {
	if appID == "" || secret == "" || code == "" {
		return lr, errors.New("param can not be null")
	}

	api, err := CodeToURL(appID, secret, code)
	if err != nil {
		return lr, err
	}

	res, err := c.client.Get(api)

	if err != nil {
		return lr, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return lr, errors.New("连接到微信服务器出错")
	}

	err = json.NewDecoder(res.Body).Decode(&lr)
	if err != nil {
		return lr, err
	}

	if lr.Errcode != 0 {
		return lr, errors.New(lr.Errmsg)
	}

	return lr, nil
}

func CheckSignature(signature, sessionKey, rawData string) bool {
	return signature == Sha1(rawData+sessionKey)
}

func DecryptUserInfo(sessionKey, encryptedData, iv string) (ui UserInfo, err error) {
	plaintext, err := AesCBCDecrypt(sessionKey, encryptedData, iv)
	if err != nil {
		return ui, err
	}

	err = json.Unmarshal(plaintext, &ui)
	if err != nil {
		return ui, err
	}

	return ui, nil
}
