package model

type CommonParams struct {
	Pkg string `json:"pkg" form:"pkg"`
	Vn string `json:"vn" form:"vn"`
	Ts int64 `json:"ts" form:"ts"`
}
