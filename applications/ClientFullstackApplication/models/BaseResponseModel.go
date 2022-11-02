package models

type BaseResponseModel[T interface{}] struct {
	Data         T      `json:"data"`
	ErrorMessage string `json:"errorMessage"`
}
