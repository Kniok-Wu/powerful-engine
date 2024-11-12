package Response

type Response map[string]interface{}

func NewResponse(message string) Response {
	resposne := make(map[string]interface{})
	resposne["msg"] = message
	return resposne
}

// Error 出现了未知的异常
func (r Response) Error() Response {
	r["code"] = 10025
	return r
}

// Invalid 错误的请求内容
func (r Response) Invalid() Response {
	r["code"] = 10024
	return r
}

// Success 正常的请求返回
func (r Response) Success() Response {
	r["code"] = 10000
	return r
}

// Data 设置一个返回的数据
func (r Response) Data(key string, value interface{}) Response {
	if r["data"] == nil {
		r["data"] = make(map[string]interface{})
	}

	r["data"].(map[string]interface{})[key] = value
	return r
}
