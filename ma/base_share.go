package ma

import "github.com/cliod/wx-go/common"

type WxMaShareInfo struct {
	common.Err

	OpenGId string `json:"open_g_id"`
}
