package util

import (
	jsonsys "encoding/json"
	jsoniter "github.com/json-iterator/go"
)

var defaultJson = jsoniter.ConfigCompatibleWithStandardLibrary

func Marshal(v interface{}) ([]byte, error) {
	return defaultJson.Marshal(v)
}

func MarshalToString(v interface{}) (string, error) {
	return defaultJson.MarshalToString(v)
}

func MarshalToStringWithoutErr(v interface{}) string {
	result, _ := defaultJson.MarshalToString(v)
	return result
}

func MarshalToStringPointWithoutErr(v interface{}) *string {
	result, _ := defaultJson.MarshalToString(v)
	return &result
}

func MarshalIndentStringWithoutError(v interface{}, prefix, indent string) string {
	bs, _ := defaultJson.MarshalIndent(v, prefix, indent)
	return string(bs)
}

func Valid(data []byte) bool {
	return defaultJson.Valid(data)
}

func UnmarshalFromString(data string, v interface{}) error {
	return defaultJson.UnmarshalFromString(data, v)
}

func Unmarshal(data []byte, v interface{}) error {
	return defaultJson.Unmarshal(data, v)
}

func RawMessage(data []byte) jsonsys.RawMessage {
	return jsonsys.RawMessage(data)
}
