package model

import "strconv"

func GetHeader() map[string]string {
	return map[string]string{"X-Token": Token, "AuthType": strconv.FormatInt(int64(AuthTypeServer), 10)}
}

func GetHeaderNoToken() map[string]string {
	return map[string]string{"X-Token": "", "AuthType": strconv.FormatInt(int64(AuthTypeServer), 10)}
}

func GetMiniHeader(uuid string) map[string]string {
	if len(uuid) == 0 {
		uuid = "5205857593c2eacc6f6c1da376b32ca3"
	}
	return map[string]string{"X-Token": MiniWechatToken, "IsDev": "1", "AuthType": strconv.FormatInt(int64(AuthTypeMiniWechat), 10), "UUID": uuid}
}

func GetMiniHeaderNoToken(uuid string) map[string]string {
	if len(uuid) == 0 {
		uuid = "5205857593c2eacc6f6c1da376b32ca3"
	}
	return map[string]string{"X-Token": "", "IsDev": "1", "AuthType": strconv.FormatInt(int64(AuthTypeMiniWechat), 10), "UUID": uuid}
}

func GetMiniHeaderNoIsDev(uuid string) map[string]string {
	if len(uuid) == 0 {
		uuid = "5205857593c2eacc6f6c1da376b32ca3"
	}
	return map[string]string{"X-Token": "", "IsDev": "", "AuthType": strconv.FormatInt(int64(AuthTypeMiniWechat), 10), "UUID": uuid}
}

func GetMiniHeaderAuthServer(uuid string) map[string]string {
	if len(uuid) == 0 {
		uuid = "5205857593c2eacc6f6c1da376b32ca3"
	}
	return map[string]string{"X-Token": MiniWechatToken, "IsDev": "1", "AuthType": strconv.FormatInt(int64(AuthTypeServer), 10), "UUID": uuid}
}

func GetSessionId() string {
	return PHPSESSID
}

func GetMiniSessionId() string {
	return MINIWECHATPHPSESSID
}
