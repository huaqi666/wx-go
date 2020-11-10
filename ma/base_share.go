package ma

import "github.com/cliod/wx-go/common"

// 分享的敏感信息
type WxMaShareInfo struct {
	common.Err

	OpenGId string `json:"open_g_id"`
}
