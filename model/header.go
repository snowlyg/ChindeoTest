package model

import "strconv"

func GetHeader() map[string]string {
	return map[string]string{"X-Token": Token, "AuthType": strconv.FormatInt(int64(AuthTypeServer), 10)}
}

func GetHeaderNoToken() map[string]string {
	return map[string]string{"X-Token": "", "AuthType": strconv.FormatInt(int64(AuthTypeServer), 10)}
}

func GetHeaderAuthMini() map[string]string {
	return map[string]string{"X-Token": "", "AuthType": strconv.FormatInt(int64(AuthTypeMiniWechat), 10)}
}

func GetMiniHeader() map[string]string {
	return map[string]string{"X-Token": MiniWechatToken, "IsDev": "1", "AuthType": strconv.FormatInt(int64(AuthTypeMiniWechat), 10)}
}

func GetMiniHeaderNoToken() map[string]string {
	return map[string]string{"X-Token": "", "IsDev": "1", "AuthType": strconv.FormatInt(int64(AuthTypeMiniWechat), 10)}
}

func GetMiniHeaderNoIsDev() map[string]string {
	return map[string]string{"X-Token": "", "IsDev": "", "AuthType": strconv.FormatInt(int64(AuthTypeMiniWechat), 10)}
}

func GetMiniHeaderAuthServer() map[string]string {
	return map[string]string{"X-Token": MiniWechatToken, "IsDev": "1", "AuthType": strconv.FormatInt(int64(AuthTypeServer), 10)}
}

func GetSessionId() string {
	return PHPSESSID
}

func GetMiniSessionId() string {
	return MINIWECHATPHPSESSID
}
