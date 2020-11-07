package ma

type WxMaShareService interface {
	// 解密分享敏感数据.
	GetShareInfo(sessionKey, encryptedData, ivStr string) (*WxMaShareInfo, error)
}

type WxMaShareServiceImpl struct {
	service WxMaService
}

func newWxMaShareService(service WxMaService) *WxMaShareServiceImpl {
	return &WxMaShareServiceImpl{
		service: service,
	}
}

func (s *WxMaShareServiceImpl) GetShareInfo(sessionKey, encryptedData, ivStr string) (*WxMaShareInfo, error) {
	return GetShareInfo(sessionKey, encryptedData, ivStr)
}
