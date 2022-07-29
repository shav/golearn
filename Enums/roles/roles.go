package roles

import (
	"errors"
	"strings"
)

type Role struct {
	name string
}

var (
	Unknown   = Role{""}
	Guest     = Role{"guest"}
	Member    = Role{"member"}
	Moderator = Role{"moderator"}
	Admin     = Role{"admin"}
)

var roles = map[string]Role{
	Guest.name:     Guest,
	Member.name:    Member,
	Moderator.name: Moderator,
	Admin.name:     Admin,
}

func (r Role) String() string {
	return r.name
}

func FromString(str string) (Role, error) {
	if role, ok := roles[strings.ToLower(str)]; ok {
		return role, nil
	}
	return Unknown, errors.New("unknown role")
}
