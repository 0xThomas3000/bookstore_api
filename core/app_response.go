package core

type successResponse struct {
	StatusCode int         `json:"status_code"`
	Message    string      `json:"message"`
	Log        string      `json:"log,omitempty"`
	Paging     interface{} `json:"paging,omitempty"`
	Data       interface{} `json:"data,omitempty"`
}

func NewSuccessResponse(data, paging interface{}, log string) *successResponse {
	return &successResponse{Data: data, Paging: paging, StatusCode: 200, Message: "OK", Log: log}
}

func SimpleSuccessResponse(data interface{}) *successResponse {
	return NewSuccessResponse(data, nil, "")
}

func AddSuccessResponse(data interface{}, log string) *successResponse {
	return NewSuccessResponse(data, nil, log)
}

func UpdateSuccessResponse(log string) *successResponse {
	return NewSuccessResponse(nil, nil, log)
}
