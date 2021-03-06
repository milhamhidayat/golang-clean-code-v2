package domain

import (
	"context"
	"time"
)

// DepartmentFilter represent query filter
type DepartmentFilter struct {
	IDs     []string
	Keyword string
	Num     int
	Cursor  string
}

// Department represent department data
type Department struct {
	ID          string    `json:"id"`
	Name        string    `json:"name" validate:"required"`
	Description string    `json:"description"`
	CreatedTime time.Time `json:"created_time"`
	UpdatedTime time.Time `json:"updated_time"`
}

// DepartmentService represent service contract for department
type DepartmentService interface {
	Create(ctx context.Context, d *Department) (err error)
	Fetch(ctx context.Context, filter DepartmentFilter) (departments []Department, nextCursor string, err error)
	Get(ctx context.Context, departmentID string) (department Department, err error)
	Update(ctx context.Context, d Department) (department Department, err error)
	Delete(ctx context.Context, departmentID string) (err error)
}

// DepartmentRepository represent repository contract for department
type DepartmentRepository interface {
	Create(ctx context.Context, d *Department) (err error)
	Fetch(ctx context.Context, filter DepartmentFilter) (departments []Department, nextCursor string, err error)
	Get(ctx context.Context, departmentID string) (department Department, err error)
	Update(ctx context.Context, d Department) (department Department, err error)
	Delete(ctx context.Context, departmentID string) (err error)
}
