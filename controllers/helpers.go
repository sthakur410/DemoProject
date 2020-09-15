package controllers

type Response struct {
	Code  int         `json:"code"`
	Msg   string      `json:"msg"`
	Model interface{} `json:"model"`
}

func ReturnResponseFromCodeMessage(code int, message string, model interface{}) (response Response) {
	response.Code = code
	response.Msg = message
	response.Model = model
	return
}
