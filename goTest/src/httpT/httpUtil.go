package httpT

import "encoding/json"
import "../define"

//将回馈信息转换成json，byte
func FeeBookCommonResultToByte(commonResult define.CommonResult) []byte {
	s, _ := json.Marshal(commonResult)
	return s
}


//将回馈信息转换成json，string
func FeeBookCommonResultToString(commonResult define.CommonResult) string {
	s, _ := json.Marshal(commonResult)
	return string(s)
}


