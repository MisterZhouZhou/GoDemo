package define

type BaseResult struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type CommonResult struct {
	*BaseResult
	//定义空接口，接受任何参数
	Data interface{} `json:"data"`
}

type CommonPageData struct {
	*BaseResult
	Data  interface{} `json:"data"`
	Count int 		  `json:"count"`	
}

var Success = &BaseResult{
	Code: 1,
	Msg:  "成功",
}

var Failed = &BaseResult{
	Code: -1,
	Msg:  "失败",
}


/**
  自定义错误
 */
func CustomerCommonResult(code int, msg string, data interface{}) CommonResult {
	return CommonResult {
		BaseResult: &BaseResult{
			Code: code,
			Msg:  msg,
		},
		Data:data,
	}
}


// define.CommonResult{
//	BaseResult: define.Success,
//	Data:       "注册成功",
//}