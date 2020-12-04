package util

import (
	"fmt"
	"reflect"
)

// map构建xml字符串
// Deprecated: Xml.ToXmlString
func BuildByMap(data map[string]interface{}) string {
	return NewXml(data).String()
}

type Xml map[string]interface{}

func (x Xml) String() string {
	return x.ToString()
}

func (x Xml) ToString() string {
	return x.ToXmlString()
}

func (x Xml) ToXmlString() string {
	s := "<xml>"
	for k, v := range x {
		s += "<" + k + ">" + v.(string) + "</" + k + ">"
	}
	s += "</xml>"
	return s
}

// map to xml
func NewXml(data map[string]interface{}) Xml {
	if data == nil {
		return Xml{}
	}
	var (
		t   reflect.Type
		res = make(Xml)
	)
	for k, v := range data {
		if v == nil {
			res[k] = ""
			continue
		}
		t = reflect.TypeOf(v)
		isString := t.Kind() == reflect.String || t.Elem().Kind() == reflect.String
		if isString {
			res[k] = v
			continue
		}
		//isStruct := t.Kind() == reflect.Struct || t.Elem().Kind() == reflect.Struct
		//if isStruct {
		//	bs, _ := json.Marshal(v)
		//	res[k] = string(bs)
		//	continue
		//}
		res[k] = fmt.Sprint(v)
	}
	return res
}
