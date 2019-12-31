package authz

import (
	"context"
	"fmt"

	"github.com/casbin/casbin"
	"github.com/casbin/casbin/util"
	"github.com/butters-mars/tiki/logging"
	"github.com/butters-mars/taka/def"
)

const (
	// RoleGuest -
	RoleGuest = "guest"
	// RoleUser -
	RoleUser = "user"
	// RoleOwner -
	RoleOwner = "owner"
	// RoleAdmin -
	RoleAdmin = "admin"
)

// CasbinAuthz authz service based on casbin
type CasbinAuthz struct {
	enforcer *casbin.Enforcer
}

// NewCasbinAuthz creates a CasbinAuthz
func NewCasbinAuthz(modelPath, policyPath string) *CasbinAuthz {
	enforcer := casbin.NewEnforcer(modelPath, policyPath)
	enforcer.AddFunction("keyMatch2", util.KeyMatch2Func)

	return &CasbinAuthz{enforcer: enforcer}
}

// Allow impl authz.Service methods
func (s CasbinAuthz) Allow(ctx context.Context, user *def.E, obj *def.E, action string, args []interface{}) (bool, error) {
	if obj == nil {
		return false, fmt.Errorf("obj is nul")
	}

	var uid string
	if user != nil {
		uid = user.ID
	}
	role := s.determineRole(user, obj, args)
	objPath := fmt.Sprintf("%s/%s", obj.Type, obj.ID)
	ok := s.enforcer.Enforce(role, objPath, action)
	logging.WDebug("check access_control", "uid", uid, "role", role, "action", action, "obj", objPath, "result", ok)
	return ok, nil
}

/*
// Allow impl authz.Service methods
func (s CasbinAuthz) Allow(user *def.E, action string, args []interface{}) (bool, error) {
	if obj == nil {
		return false, fmt.Errorf("obj is nul")
	}
	role := s.determineRole(user, obj, args)
	ok := s.enforcer.Enforce(role, fmt.Sprintf("%s/%s", obj.Type, obj.ID), action)
	return ok, nil
}
*/

func (s CasbinAuthz) determineRole(user, obj *def.E, args []interface{}) string {
	if user == nil {
		return RoleGuest
	}

	var oid string
	if obj != nil {
		oid = obj.ID1
	}
	if oid == "" {
		return RoleUser
	}

	if user.ID == oid {
		return RoleOwner
	}

	return RoleUser
}
func (s CasbinAuthz) allow(userRole, objPath, action string) bool {
	ok, err := s.enforcer.EnforceSafe(userRole, objPath, action)
	if err != nil {
		return false
	}

	return ok
}
