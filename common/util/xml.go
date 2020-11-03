package util

func BuildByMap(data map[string]interface{}) string {
	s := "<xml>"
	for k, v := range data {
		s += "<" + k + ">" + v.(string) + "</" + k + ">"
	}
	s += "</xml>"
	return s
}
