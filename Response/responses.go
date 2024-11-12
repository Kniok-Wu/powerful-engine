package Response

func NewStandardError() Response {
	return NewResponse("服务繁忙，请稍后再试").Error()
}

func NewStandardSuccess() Response {
	return NewResponse("请求成功！").Success()
}
