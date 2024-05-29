package controllers

import "github.com/sergey4qb/ltp/internal/core/domain/valueobject"

type LTPResponse struct {
	LTP []valueobject.LTP `json:"ltp"`
}
