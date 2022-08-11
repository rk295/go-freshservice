package freshservice

import (
	"context"
)

// DepartmentsService is an interface for interacting with the department
// endpoints of the Freshservice API
type DepartmentsService interface {
	List(context.Context, QueryFilter) ([]DepartmentDetails, string, error)
	Get(context.Context, int) (*DepartmentDetails, error)
}
