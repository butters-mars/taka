package authz

import (
	"context"

	"github.com/butters-mars/taka/def"
)

// Service authrization service interface
type Service interface {
	Allow(ctx context.Context, user *def.E, obj *def.E, action string, args []interface{}) (bool, error)
}
