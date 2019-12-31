package authz

import "testing"

func TestCasbinAuthz(t *testing.T) {
	s := New("../../deploy/authz/model.conf", "../../deploy/authz/policy.csv")
	if ok := s.allow("user", "post/1111", "like"); !ok {
		t.Errorf("user should be able to like")
	}
	if ok := s.allow("guest", "post/1111", "like"); ok {
		t.Errorf("guest should not be able to like")
	}

	if ok := s.allow("user", "post/1111", "delete"); ok {
		t.Errorf("user should not be able to delete")
	}
	if ok := s.allow("owner", "post/1111", "delete"); !ok {
		t.Errorf("owner should be able to delete")
	}

}
