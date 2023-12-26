package web

type BaseResponse[T any] struct {
	Status  int    `json:"status"`
	Data    T      `json:"data"`
	Message string `json:"message"`
}

func Null() interface{} {
	return nil
}

// Response returns new BaseResponse.
func Response[T any](status int, message string, data T) BaseResponse[T] {
	return BaseResponse[T]{
		Status:  status,
		Data:    data,
		Message: message,
	}
}
