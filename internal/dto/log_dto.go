package dto

type GetLogDTO struct {
	Path string `json:"path"`
	Order int    `json:"order"`
}