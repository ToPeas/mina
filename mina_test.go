package mina

import "testing"

func TestCheckSignature(t *testing.T) {
	signature := "9723a3c9be77cf53abab266547322daa2547e509"
	sessionKey := "Hj84aljX3iEU6+DAChvPjg=="
	rawData := `{"nickName":"solarhell","gender":1,"language":"zh_CN","city":"Nanjing","province":"Jiangsu","country":"China","avatarUrl":"https://wx.qlogo.cn/mmopen/vi_32/Q0j4TwGTfTI5J2lPp4NaUXibU2U6icHZ0lZf2j1QJhHIh89tPP2j1pKMB5Ax91RslrevFm1jMS5v1lDS4zmCSE3A/132"}`
	ok := CheckSignature(signature, sessionKey, rawData)
	if !ok {
		t.Error("出错了")
	}
	t.Log("success CheckSignature")
}

func TestDecryptUserInfo(t *testing.T) {
	sessionKey := "tiihtNczf5v6AKRyjwEUhQ=="
	encryptedData := "CiyLU1Aw2KjvrjMdj8YKliAjtP4gsMZMQmRzooG2xrDcvSnxIMXFufNstNGTyaGS9uT5geRa0W4oTOb1WT7fJlAC+oNPdbB+3hVbJSRgv+4lGOETKUQz6OYStslQ142dNCuabNPGBzlooOmB231qMM85d2/fV6ChevvXvQP8Hkue1poOFtnEtpyxVLW1zAo6/1Xx1COxFvrc2d7UL/lmHInNlxuacJXwu0fjpXfz/YqYzBIBzD6WUfTIF9GRHpOn/Hz7saL8xz+W//FRAUid1OksQaQx4CMs8LOddcQhULW4ucetDf96JcR3g0gfRK4PC7E/r7Z6xNrXd2UIeorGj5Ef7b1pJAYB6Y5anaHqZ9J6nKEBvB4DnNLIVWSgARns/8wR2SiRS7MNACwTyrGvt9ts8p12PKFdlqYTopNHR1Vf7XjfhQlVsAJdNiKdYmYVoKlaRv85IfVunYzO0IKXsyl7JCUjCpoG20f0a04COwfneQAGGwd5oa+T8yO5hzuyDb/XcxxmK01EpqOyuxINew=="
	iv := "r7BXXKkLb8qrSNn05n0qiA=="


	_, err := DecryptUserInfo(sessionKey, encryptedData, iv)
	if err != nil {
		t.Error(err)
	}

	t.Log("success DecryptUserInfo")
}
