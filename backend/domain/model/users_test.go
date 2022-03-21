package model

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	var email string = "suzuki@sample.com"
	var name string = "suzuki"
	var password string = "suzuki_pass"
	var roleid int = 1
	var updatedAt time.Time
	user, err := NewUser(email, name, password, roleid, updatedAt)

	if err != nil {
		t.Errorf("failed NewUser")
	}

	assert.Equal(t, email, user.Email)
	assert.Equal(t, name, user.Name)
	assert.Equal(t, password, user.Password)
	assert.Equal(t, roleid, user.RoleId)
	assert.Equal(t, updatedAt, user.UpdatedAt)

	t.Logf("user: %p", user)
	t.Logf("user.Email: %s", user.Email)
	t.Logf("user.Name: %s", user.Name)
	t.Logf("user.Password: %s", user.Password)
	t.Logf("user.RoleId: %v", user.RoleId)
	t.Logf("user.UpdatedAt: %s", user.UpdatedAt)
}
