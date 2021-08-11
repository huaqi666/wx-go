package err

type WxErr string

func (e WxErr) Error() string {
	return string(e)
}
