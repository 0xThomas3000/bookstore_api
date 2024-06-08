package core

type successResponse struct {
	Data   interface{} `json:"data"`
	Paging interface{} `json:"paging,omitempty"`
}

func NewSuccessResponse(data, paging interface{}) *successResponse {
	return &successResponse{Data: data, Paging: paging}
}

func SimpleSuccessResponse(data interface{}) *successResponse {
	return NewSuccessResponse(data, nil)
}
