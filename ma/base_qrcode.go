package ma

// 二维码数据
type BaseCode struct {
	Width     int           `json:"width"`
	AutoColor bool          `json:"auto_color"`
	IsHyaline bool          `json:"is_hyaline"`
	LineColor CodeLineColor `json:"line_color"`
}

// 二维码数据
type WxaCode struct {
	BaseCode
	Path string `json:"path"`
}

// 二维码数据
type WxaCodeUnlimited struct {
	BaseCode
	Scene string `json:"scene"`
	Page  string `json:"page"`
}

// 二维码线颜色
type CodeLineColor struct {
	R string `json:"r"`
	B string `json:"b"`
	G string `json:"g"`
}

// 默认白色
func DefaultCodeLineColor() CodeLineColor {
	return CodeLineColor{R: "0", B: "0", G: "0"}
}
