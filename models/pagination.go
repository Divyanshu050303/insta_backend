package models

type PaginationModel struct {
	Page  *int `json:"page"`
	Limit *int `json:"limit"`
}
