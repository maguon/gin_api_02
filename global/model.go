package global

import (
	"time"
)

type EXTEND_MODEL struct {
	ID        int64 `json:"id" form:"id" gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type EXTEND_SEARCH struct {
	PageNumber   int       `json:"pageNumber" form:"pageNumber"` // 页码
	PageSize     int       `json:"pageSize" form:"pageSize"`     // 每页大小
	CreatedStart time.Time `json:"createdStart" form:"createdStart"`
	CreatedEnd   time.Time `json:"createdEnd" form:"createdEnd"`
	UpdatedStart time.Time `json:"updatedStart" form:"updatedStart"`
	UpdatedEnd   time.Time `json:"updatedEnd" form:"updatedEnd"`
}
