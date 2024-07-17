package types

import "database/sql"

type OrderBy struct {
	Order string `json:"order"`
	Key   string `json:"key"`
}

type GetRequest struct {
	PageSize  int     `json:"pageSize"`
	PageIndex int     `json:"pageIndex"`
	Search    string  `json:"search"`
	OrderBy   OrderBy `json:"orderBy"`
}

type Category struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	AddedBy     string `json:"added_by"`
	IconID      int    `json:"icon_id"`
	FrgColor    string `json:"frg_color"`
	BgColor     string `json:"bg_color"`
	URL         string `json:"url"`
	Active      int    `json:"active"`
	Description sql.NullString `json:"description"`
	IsApproved  int    `json:"is_approved"`
}