package web

type Response struct {
	Code  int         `json:"code"`
	Data  interface{} `json:"data,omitempty"`
	Error string      `json:"error,omitempty"`
}

func NewResponse(code int, data interface{}, err error) *Response {
	if code > 400 {
		return &Response{
			Code:  code,
			Error: err.Error(),
		}
	}

	return &Response{
		Code: code,
		Data: data,
	}
}
