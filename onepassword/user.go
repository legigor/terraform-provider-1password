package onepassword

import (
	"encoding/json"
)

const (
	// UserResource is 1Password's internal designator for Groups
	UserResource = "user"
)

// Group represents a 1Password Group resource
type User struct {
	UUID  string
	Name  string
	Email string
	State string
}

// ReadUser gets an existing 1Password User
func (o *OnePassClient) ReadUser(id string) (*User, error) {
	user := &User{}
	res, err := o.runCmd(opPasswordGet, UserResource, id)
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(res, user); err != nil {
		return nil, err
	}
	return user, nil
}

// CreateUser creates a new 1Password User
func (o *OnePassClient) CreateUser(v *User) (*User, error) {
	args := []string{opPasswordCreate, UserResource, v.Email, v.Name}
	res, err := o.runCmd(args...)
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(res, v); err != nil {
		return nil, err
	}
	return v, nil
}

// DeleteUser deletes a 1Password User
func (o *OnePassClient) DeleteUser(id string) error {
	return o.Delete(UserResource, id)
}
