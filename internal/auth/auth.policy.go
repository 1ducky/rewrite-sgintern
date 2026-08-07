package auth

type Role string

const (
	Admin Role = "ADMIN"
	User  Role = "USER"
)

func ParseRole(role string) (Role, error) {
	switch Role(role) {
	case Admin, User:
		return Role(role), nil
	default:
		return "", ErrRoleInvalid
	}
}

func (r Role) String() string {
	return string(r)
}

func (r Role) Compare(diff Role) bool {
	return r == diff
}
