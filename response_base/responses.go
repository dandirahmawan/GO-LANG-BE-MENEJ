package response_base

type Resp struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type RespMsg struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func ResponseData(isSuccess bool, msg string, data interface{}) Resp {
	var r Resp
	r.Success = isSuccess
	r.Message = msg
	r.Data = data

	return r
}

func ResponseMsg(isSuccess bool, msg string) RespMsg {
	var r RespMsg
	r.Success = isSuccess
	r.Message = msg

	return r
}
